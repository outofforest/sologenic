package types

import "context"

// Token represents an exchangeable token
type Token struct {
	Symbol        string
	DecimalPlaces uint32
}

func (t Token) String() string {
	return t.Symbol
}

// TokenPair represents a tradeable pair of tokens
type TokenPair struct {
	BaseToken  Token
	QuoteToken Token
}

func (tp TokenPair) String() string {
	return tp.BaseToken.Symbol + tp.QuoteToken.Symbol
}

// PriceType is the type of the price returned from fetcher
type PriceType int

const (
	// BidPriceType means that rate object contains bid price
	BidPriceType PriceType = iota

	// AskPriceType means that rate object contains ask price
	AskPriceType
)

// ExchangeRate is the rate returned from external service
type ExchangeRate struct {
	Pair  TokenPair
	Type  PriceType
	Price uint64
}

// FetchFn defines schema of function fetching exchange rate from external service
type FetchFn func(ctx context.Context) ([]ExchangeRate, error)

// Provider implements logic for fetching exchange rate from external service
type Provider interface {
	Fetchers(pairs []TokenPair) []FetchFn
}
