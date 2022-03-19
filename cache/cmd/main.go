package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/outofforest/logger"
	"github.com/outofforest/run"
	"github.com/outofforest/sologenic/cache/supply"
	"github.com/outofforest/sologenic/cache/supply/types"
	"github.com/ridge/must"
	"github.com/ridge/parallel"
	"go.uber.org/zap"
)

const workers = 2

type rate struct {
	Bid uint64
	Ask uint64
}

func main() {
	run.Service("cache", nil, func(ctx context.Context) error {
		pairsByProvider := map[types.Provider][]types.TokenPair{}
		for pair, provider := range supply.Sources {
			pairsByProvider[provider] = append(pairsByProvider[provider], pair)
		}

		fetchers := []types.FetchFn{}
		for provider, pairs := range pairsByProvider {
			fetchers = append(fetchers, provider.Fetchers(pairs)...)
		}

		ctxReady, readyClose := context.WithCancel(ctx)
		defer readyClose()

		return parallel.Run(ctx, func(ctx context.Context, spawn parallel.SpawnFn) error {
			queue := make(chan types.FetchFn, len(fetchers))
			for _, f := range fetchers {
				queue <- f
			}

			results := make(chan types.ExchangeRate)
			for i := 0; i < workers; i++ {
				spawn(fmt.Sprintf("worker-%d", i), parallel.Fail, func(ctx context.Context) error {
					log := logger.Get(ctx)
					for {
						select {
						case <-ctx.Done():
							return ctx.Err()
						case f := <-queue:
							err := func() error {
								ctxFetch, cancel := context.WithTimeout(ctx, 5*time.Second)
								defer cancel()

								rates, err := f(ctxFetch)
								if err != nil {
									if ctx.Err() != nil {
										return ctx.Err()
									}
									// Error is just logged here. It is not returned to let service continue its job.
									// But this error should be signaled using prometheus metric or sth like that.
									log.Error("Fetching exchange rate failed", zap.Error(err))
									return nil
								}
								for _, r := range rates {
									select {
									case <-ctx.Done():
										return ctx.Err()
									case results <- r:
									}
								}
								return nil
							}()
							if err != nil {
								return err
							}
							queue <- f
						}
					}
				})
			}

			rates := map[string]rate{}
			var mu sync.RWMutex
			spawn("integrator", parallel.Fail, func(ctx context.Context) error {
				readyPairs := map[string]bool{}
				for {
					select {
					case <-ctx.Done():
						return ctx.Err()
					case r := <-results:
						pair := r.Pair.String()
						mu.Lock()
						rate := rates[pair]
						switch r.Type {
						case types.BidPriceType:
							rate.Bid = r.Price
						case types.AskPriceType:
							rate.Ask = r.Price
						default:
							panic("unsupported price type")
						}
						rates[pair] = rate
						mu.Unlock()

						if rate.Bid > 0 && rate.Ask > 0 {
							readyPairs[pair] = true
							if len(readyPairs) == len(supply.Sources) {
								readyClose()
							}
						}
					}
				}
			})
			spawn("query", parallel.Fail, func(ctx context.Context) error {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-ctxReady.Done():
				}

				for {
					mu.RLock()
					rates := must.Bytes(json.MarshalIndent(rates, "", "  "))
					mu.RUnlock()
					fmt.Println(string(rates))

					select {
					case <-ctx.Done():
						return ctx.Err()
					case <-time.After(10 * time.Second):
					}
				}
			})
			return nil
		})
	})
}
