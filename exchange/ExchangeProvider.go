package exchange

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ExchangeProvider interface {
	GetAmountOut(from string, to string, amountIn *big.Int) (*big.Rat, error)
	GetPriceAfterAmount(from string, to string, amountIn *big.Int) (price *big.Rat, err error)
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

var nameAddress = map[string]common.Address{
	"USDC": common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
	"USDT": common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
	"DAI":  common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
	"BUSD": common.HexToAddress("0x4fabb145d64652a948d72533023f6e7a623c7c53"),
	"GUSD": common.HexToAddress("0x056fd409e1d7a124bd7017459dfea2f387b6d5cd"),
	"SUSD": common.HexToAddress("0x57ab1ec28d129707052df4df418d58a2d46d5f51"),
	"MUSD": common.HexToAddress("0xe2f2a5C287993345a840Db3B0845fbC70f5935a5"),
}

var tokenDecimalMap = map[string](*big.Int){
	"USDC": big.NewInt(1e6),
	"USDT": big.NewInt(1e6),
	"DAI":  big.NewInt(1e18),
	"BUSD": big.NewInt(1e18),
	"GUSD": big.NewInt(1e2),
	"SUSD": big.NewInt(1e18),
	"MUSD": big.NewInt(1e18),
}

func calculateXY(dx1 *big.Int, dy1 *big.Int, dx2 *big.Int, dy2 *big.Int) (x *big.Int, y *big.Int) {
	n := new(big.Int).Sub(new(big.Int).Mul(dx2, dy2), new(big.Int).Mul(big.NewInt(2), new(big.Int).Mul(dx1, dy1)))
	d := new(big.Int).Sub(new(big.Int).Mul(big.NewInt(2), dy1), dy2)

	x = new(big.Int).Div(n, d)
	n2 := new(big.Int).Add(new(big.Int).Mul(x, dy1), new(big.Int).Mul(dx1, dy1))
	y = new(big.Int).Div(n2, dx1)
	return
}
