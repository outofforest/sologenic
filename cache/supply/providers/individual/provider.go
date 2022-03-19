package individual

import (
	"context"
	"math/rand"

	"github.com/outofforest/sologenic/cache/helpers"
	"github.com/outofforest/sologenic/cache/supply/types"
)

// Individual provider is a fake implementation of external service delivering exchange rates.
// The logic here mocks the service which is able to deliver only one rate per query.
// Assuming we want to query for these pairs:
// - BTCETH
// - USDETH
// - USDBTC
// separate query for each one has to be sent to get bid prices. To get ask prices we have to query for these pairs:
// - ETHBTC
// - ETHUSD
// - BTCUSD
// meaning that 6 queries have to be sent in total.

// New returns new individual provider
func New() types.Provider {
	return &provider{}
}

type provider struct {
}

func (p *provider) Fetchers(pairs []types.TokenPair) []types.FetchFn {
	res := make([]types.FetchFn, 0, 2*len(pairs))
	for _, p := range pairs {
		res = append(res, fetchFn(p, types.BidPriceType), fetchFn(p, types.AskPriceType))
	}
	return res
}

func fetchFn(pair types.TokenPair, priceType types.PriceType) types.FetchFn {
	return func(ctx context.Context) ([]types.ExchangeRate, error) {
		if err := helpers.Wait(ctx); err != nil {
			return nil, err
		}
		return []types.ExchangeRate{
			{
				Pair:  pair,
				Type:  priceType,
				Price: rand.Uint64(),
			},
		}, nil
	}
}
