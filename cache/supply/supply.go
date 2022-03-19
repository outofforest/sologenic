package supply

import "context"

type Token struct {
	Symbol        string
	DecimalPlaces uint32
}

type TokenPair struct {
	BaseToken  Token
	QuoteToken Token
}

type PriceType int

const (
	BidPriceType PriceType = iota
	AskPriceType
)

type ExchangeRate struct {
	Pair  TokenPair
	Type  PriceType
	Price uint64
}

type FetchFn func(ctx context.Context) ([]ExchangeRate, error)

type Provider interface {
	Fetchers(pairs []TokenPair) []FetchFn
}
