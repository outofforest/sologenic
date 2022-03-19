package supply

import (
	"github.com/outofforest/sologenic/cache/supply/providers/grouping"
	"github.com/outofforest/sologenic/cache/supply/providers/individual"
	"github.com/outofforest/sologenic/cache/supply/types"
)

// Tokens
var (
	USD = types.Token{Symbol: "USD", DecimalPlaces: 2}
	EUR = types.Token{Symbol: "EUR", DecimalPlaces: 2}
	BTC = types.Token{Symbol: "BTC", DecimalPlaces: 8}
	ETH = types.Token{Symbol: "ETH", DecimalPlaces: 18}
)

// Tradeable token pairs
var (
	EURUSD = types.TokenPair{BaseToken: EUR, QuoteToken: USD}
	USDBTC = types.TokenPair{BaseToken: USD, QuoteToken: BTC}
	USDETH = types.TokenPair{BaseToken: USD, QuoteToken: ETH}
	EURBTC = types.TokenPair{BaseToken: EUR, QuoteToken: BTC}
	EURETH = types.TokenPair{BaseToken: EUR, QuoteToken: ETH}
	BTCETH = types.TokenPair{BaseToken: BTC, QuoteToken: ETH}
)

// Providers
var (
	individualProvider = individual.New()
	groupingProvider   = grouping.New()
)

// SourceMap connects token pair with provider supplying exchange rates for that pair
type SourceMap map[types.TokenPair]types.Provider

// Sources maps token pairs to providers
var Sources = SourceMap{
	EURUSD: individualProvider,
	USDBTC: groupingProvider,
	USDETH: groupingProvider,
	EURBTC: individualProvider,
	EURETH: individualProvider,
	BTCETH: groupingProvider,
}
