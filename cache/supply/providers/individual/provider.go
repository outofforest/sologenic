package individual

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
	res := make([]supply.FetchFn, 0, 2*len(pairs))
	for _, p := range pairs {
		res = append(res, fetchFn(p, supply.BidPriceType), fetchFn(p, supply.AskPriceType))
	}
	return res
}

func fetchFn(pair supply.TokenPair, priceType supply.PriceType) supply.FetchFn {
	return func(ctx context.Context) ([]supply.ExchangeRate, error) {
		if err := helpers.Wait(ctx); err != nil {
			return nil, err
		}
		return []supply.ExchangeRate{
			{
				Pair:  pair,
				Type:  priceType,
				Price: rand.Uint64(),
			},
		}, nil
	}
}
