package exchange

import (
	"math/big"
)

type ExchangeProvider interface {
	GetAmountOut(from string, to string, amountIn *big.Int) (*big.Rat, error)
}

const (
	USDCString = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	USDTString = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	DAIString  = "0x6b175474e89094c44da98b954eedeac495271d0f"
	BUSDString = "0x4fabb145d64652a948d72533023f6e7a623c7c53"
	GUSDString = "0x056fd409e1d7a124bd7017459dfea2f387b6d5cd"
	SUSDString = "0x57ab1ec28d129707052df4df418d58a2d46d5f51"
	MUSDString = "0xe2f2a5C287993345a840Db3B0845fbC70f5935a5"
)
