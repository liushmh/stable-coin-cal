package exchange

import (
	"errors"
	UniswapContract "hex/amm/v1/UniswapContract"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapProvider struct {
	TransactionFee  int64
	ContractAddress common.Address
	Pairs           map[string](*UniswapContract.UniswapV2Pair)
}

// var nameAddress = map[string]common.Address{
// 	"USDC": common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
// 	"USDT": common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
// 	"DAI":  common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
// 	"BUSD": common.HexToAddress("0x4fabb145d64652a948d72533023f6e7a623c7c53"),
// 	"GUSD": common.HexToAddress("0x056fd409e1d7a124bd7017459dfea2f387b6d5cd"),
// }

var pairAddress = map[string]common.Address{
	"BUSD-USDC": common.HexToAddress("0x524847c615639e76fe7d0fe0b16be8c4eac9cf3c"),
	"DAI-USDC":  common.HexToAddress("0xae461ca67b15dc8dc81ce7615e0320da1a9ab8d5"),
	"DAI-USDT":  common.HexToAddress("0xb20bd5d04be54f870d5c0d3ca85d82b34b836405"),
	"USDC-USDT": common.HexToAddress("0x3041cbd36888becc7bbcbc0045e3b1f144466f5f"),
}

func NewUniswapProvider(client *ethclient.Client) (*UniswapProvider, error) {
	provider := UniswapProvider{}
	provider.Pairs = make(map[string](*UniswapContract.UniswapV2Pair))

	err := error(nil)

	for k, p := range pairAddress {
		instance, err := UniswapContract.NewUniswapV2Pair(p, client)
		if err != nil {
			continue
		}
		provider.Pairs[k] = instance
	}

	return &provider, err
}

func getPairName(from string, to string) string {
	upFrom := strings.ToUpper(from)
	upTo := strings.ToUpper(to)

	if upFrom <= upTo {
		return upFrom + "-" + upTo
	}
	return upTo + "-" + upFrom
}

func calculateY(total_x, total_y, x *big.Int) *big.Int {
	cp := new(big.Int).Mul(total_x, total_y)
	new_x := new(big.Int).Add(total_x, x)
	new_y := new(big.Int).Div(cp, new_x)

	dy := new(big.Int).Sub(total_y, new_y)
	// fmt.Printf("get : %s y \n", dy.FloatString(8))
	// price := new(big.Rat).Mul(new(big.Rat).SetInt(x), new(big.Rat).Inv(dy))
	return dy
}

func (provider *UniswapProvider) GetAmountOut(from string, to string, amountIn *big.Int) (*big.Rat, error) {

	pairName := getPairName(from, to)

	reserves, err := provider.Pairs[pairName].GetReserves(nil)
	// if reserves != nil {
	// 	return nil, errors.New("no such pair in uniswap")
	// }
	if err != nil {
		log.Fatal(err)
	}
	amountIn = new(big.Int).Mul(amountIn, tokenDecimalMap[strings.ToUpper(from)])
	tradingAmount := new(big.Int).Div(new(big.Int).Mul(amountIn, big.NewInt(997)), big.NewInt(1000))

	var dy *big.Int
	if from < to {
		dy = calculateY(reserves.Reserve0, reserves.Reserve1, tradingAmount)
	} else {
		dy = calculateY(reserves.Reserve1, reserves.Reserve0, tradingAmount)
	}

	dy = new(big.Int).Div(dy, tokenDecimalMap[strings.ToUpper(to)])

	return new(big.Rat).SetInt(dy), nil
}

func (provider *UniswapProvider) GetPriceAfterAmount(from string, to string, amountIn *big.Int) (price *big.Rat, err error) {

	pairName := getPairName(from, to)

	reserves, err := provider.Pairs[pairName].GetReserves(nil)
	if err != nil {
		log.Fatal(err)
	}

	amountIn = new(big.Int).Mul(amountIn, tokenDecimalMap[strings.ToUpper(from)])
	dx := new(big.Int).Div(new(big.Int).Mul(amountIn, big.NewInt(997)), big.NewInt(1000))

	var x, y, dy *big.Int

	if from < to {
		dy = calculateY(reserves.Reserve0, reserves.Reserve1, dx)
		x = reserves.Reserve0
		y = reserves.Reserve1
	} else {
		dy = calculateY(reserves.Reserve1, reserves.Reserve0, dx)
		x = reserves.Reserve1
		y = reserves.Reserve0
	}

	x = new(big.Int).Add(x, dx)
	y = new(big.Int).Sub(y, dy)

	if y.Cmp(big.NewInt(0)) <= 0 {
		err = errors.New("Wrong Amount")
	}
	// y = new(big.Int).Div(y, tokenDecimalMap[strings.ToUpper(to)])
	x = new(big.Int).Mul(x, tokenDecimalMap[strings.ToUpper(to)])
	price = new(big.Rat).SetFrac(x, y)
	return
}
