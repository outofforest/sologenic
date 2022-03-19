package supply

import (
	"github.com/outofforest/sologenic/cache/supply/providers/grouping"
	"github.com/outofforest/sologenic/cache/supply/providers/individual"
)

var (
	USD = Token{Symbol: "USD", DecimalPlaces: 2}
	EUR = Token{Symbol: "EUR", DecimalPlaces: 2}
	BTC = Token{Symbol: "BTC", DecimalPlaces: 8}
	ETH = Token{Symbol: "ETH", DecimalPlaces: 18}
)

var (
	EURUSD = TokenPair{BaseToken: EUR, QuoteToken: USD}
	USDBTC = TokenPair{BaseToken: USD, QuoteToken: BTC}
	USDETH = TokenPair{BaseToken: USD, QuoteToken: ETH}
	EURBTC = TokenPair{BaseToken: EUR, QuoteToken: BTC}
	EURETH = TokenPair{BaseToken: EUR, QuoteToken: ETH}
	BTCETH = TokenPair{BaseToken: BTC, QuoteToken: ETH}
)

var (
	individualProvider = individual.New()
	groupingProvider   = grouping.New()
)

type SourceMap map[TokenPair]Provider

var Sources = SourceMap{
	EURUSD: individualProvider,
	USDBTC: groupingProvider,
	USDETH: groupingProvider,
	EURBTC: individualProvider,
	EURETH: individualProvider,
	BTCETH: groupingProvider,
}
