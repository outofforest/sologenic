package grouping

import (
	"context"
	"math/rand"

	"github.com/outofforest/sologenic/cache/helpers"
	"github.com/outofforest/sologenic/cache/supply/types"
)

// New returns new grouping provider
func New() types.Provider {
	return &provider{}
}

type provider struct {
}

func (p *provider) Fetchers(pairs []types.TokenPair) []types.FetchFn {
	bid := map[types.Token][]types.TokenPair{}
	ask := map[types.Token][]types.TokenPair{}

	for _, p := range pairs {
		bid[p.BaseToken] = append(bid[p.BaseToken], p)
		ask[p.QuoteToken] = append(ask[p.QuoteToken], p)
	}

	res := make([]types.FetchFn, 0, len(bid)+len(ask))
	for _, pairs := range bid {
		res = append(res, fetchFn(pairs, types.BidPriceType))
	}
	for _, pairs := range ask {
		res = append(res, fetchFn(pairs, types.AskPriceType))
	}
	return res
}

func fetchFn(pairs []types.TokenPair, priceType types.PriceType) types.FetchFn {
	return func(ctx context.Context) ([]types.ExchangeRate, error) {
		if err := helpers.Wait(ctx); err != nil {
			return nil, err
		}
		res := make([]types.ExchangeRate, 0, len(pairs))
		for _, p := range pairs {
			res = append(res, types.ExchangeRate{
				Pair:  p,
				Type:  priceType,
				Price: rand.Uint64(),
			})
		}
		return res, nil
	}
}
