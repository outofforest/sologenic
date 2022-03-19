package grouping

import (
	"context"
	"math/rand"

	"github.com/outofforest/sologenic/cache/helpers"
	"github.com/outofforest/sologenic/cache/supply"
)

func New() supply.Provider {
	return &provider{}
}

type provider struct {
}

func (p *provider) Fetchers(pairs []supply.TokenPair) []supply.FetchFn {
	bid := map[supply.Token][]supply.TokenPair{}
	ask := map[supply.Token][]supply.TokenPair{}

	for _, p := range pairs {
		bid[p.BaseToken] = append(bid[p.BaseToken], p)
		ask[p.QuoteToken] = append(ask[p.QuoteToken], p)
	}

	res := make([]supply.FetchFn, 0, len(bid)+len(ask))
	for _, pairs := range bid {
		res = append(res, fetchFn(pairs, supply.BidPriceType))
	}
	for _, pairs := range ask {
		res = append(res, fetchFn(pairs, supply.AskPriceType))
	}
	return res
}

func fetchFn(pairs []supply.TokenPair, priceType supply.PriceType) supply.FetchFn {
	return func(ctx context.Context) ([]supply.ExchangeRate, error) {
		if err := helpers.Wait(ctx); err != nil {
			return nil, err
		}
		res := make([]supply.ExchangeRate, 0, len(pairs))
		for _, p := range pairs {
			res = append(res, supply.ExchangeRate{
				Pair:  p,
				Type:  priceType,
				Price: rand.Uint64(),
			})
		}
		return res, nil
	}
}
