package exchange

import (
	CurveContract "hex/amm/v1/CurveContract"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/juliangruber/go-intersect"
)

type CurveProvider struct {
	Fee           int64 //0.04%,  4000000
	PoolInstances map[string](*CurveContract.Curve)
}

const (
	compoundPool = "0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56"
	threePool    = "0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7"
	susdPool     = "0xA5407eAE9Ba41422680e2e00537571bcC53efBfD"
	gusdPool     = "0x4f062658EaAF2C1ccf8C8e36D6824CDf41167956"
	busdv2Pool   = "0x4807862AA8b2bF68830e4C8dc86D0e9A998e085a"
	aavePool     = "0xDeBF20617708857ebe4F679508E7b7863a8A8EeE"
	yPool        = "0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51"
)

var poolList = []string{"0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56",
	"0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7",
	"0xA5407eAE9Ba41422680e2e00537571bcC53efBfD",
	"0x4f062658EaAF2C1ccf8C8e36D6824CDf41167956",
	"0x4807862AA8b2bF68830e4C8dc86D0e9A998e085a",
	"0xDeBF20617708857ebe4F679508E7b7863a8A8EeE",
	"0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51"}

// create all the pool instance
var namePoolMap = map[string]([]string){
	"USDC": []string{threePool, susdPool, compoundPool, aavePool, yPool, gusdPool, busdv2Pool},
	"USDT": []string{threePool, susdPool, aavePool, yPool, gusdPool, busdv2Pool},
	"DAI":  []string{threePool, susdPool, compoundPool, aavePool, yPool, gusdPool, busdv2Pool},
	"SUSD": []string{susdPool},
	"BUSD": []string{busdv2Pool},
	"GUSD": []string{gusdPool},
}

var nameIndexMap = map[string](*big.Int){
	"DAI":  big.NewInt(0),
	"USDC": big.NewInt(1),
	"USDT": big.NewInt(2),

	"SUSD": big.NewInt(3),
	"BUSD": big.NewInt(0),
	"GUSD": big.NewInt(0),
}

func NewCurveProvider(client *ethclient.Client) (*CurveProvider, error) {
	provider := CurveProvider{}
	provider.PoolInstances = make(map[string](*CurveContract.Curve))
	provider.Fee = 4
	err := error(nil)

	for _, p := range poolList {
		instance, err := CurveContract.NewCurve(common.HexToAddress(p), client)
		if err != nil {
			continue
		}
		provider.PoolInstances[p] = instance
	}

	return &provider, err
}

func getPoolList(from string, to string) []interface{} {
	list1 := namePoolMap[from]
	list2 := namePoolMap[to]
	return intersect.Simple(list1, list2).([]interface{})
}

func getCoinIndex(coin string, pool string) *big.Int {
	switch pool {
	case busdv2Pool, gusdPool:
		if coin != "BUSD" && coin != "GUSD" {
			return new(big.Int).Add(nameIndexMap[coin], big.NewInt(1))
		}
	}
	return nameIndexMap[coin]
}

func (provider *CurveProvider) getAmountOutImpl(from string, to string, amountIn *big.Int, excludedPool *string) (*big.Int, string, error) {
	poolList := getPoolList(from, to)

	amountIn = new(big.Int).Mul(amountIn, tokenDecimalMap[strings.ToUpper(from)])
	// todo: add check if swapping pare is invalid

	max := big.NewInt(0)
	var selectedPool string
	for _, p := range poolList {
		ps := p.(string)
		if excludedPool != nil && *excludedPool == ps {
			continue
		}
		fromIdx := getCoinIndex(from, ps)
		toIdx := getCoinIndex(to, ps)

		amountOut, _ := provider.PoolInstances[ps].GetDyUnderlying(nil, fromIdx, toIdx, amountIn)
		amountOut = new(big.Int).Div(amountOut, tokenDecimalMap[strings.ToUpper(to)])
		// fmt.Printf("from: %d , to : %d pool: %s, amount out: %d \n ", fromIdx, toIdx, ps, amountOut)

		if amountOut.Cmp(max) > 0 {
			max = amountOut
			selectedPool = ps
		}
	}
	return max, selectedPool, nil
}

func (provider *CurveProvider) GetAmountOut(from string, to string, amountIn *big.Int) (*big.Rat, *big.Rat, error) {
	amountOut, _, _ := provider.getAmountOutImpl(from, to, amountIn, nil)

	retMax := new(big.Rat).SetInt(amountOut)
	// tmp := new(big.Int).Mul(big.NewInt(provider.Fee), amountIn)
	fee := new(big.Rat).SetFrac(new(big.Int).Mul(big.NewInt(provider.Fee), amountIn), big.NewInt(1e4))
	return retMax, fee, nil
}

func (provider *CurveProvider) GetPriceAfterAmount(from string, to string, amountIn *big.Int) (price *big.Rat, err error) {
	amountOut, pool, _ := provider.getAmountOutImpl(from, to, amountIn, nil)

	fromIdx := getCoinIndex(from, pool)
	toIdx := getCoinIndex(to, pool)
	fromBalance, _ := provider.PoolInstances[pool].Balances(nil, fromIdx)
	toBalance, _ := provider.PoolInstances[pool].Balances(nil, toIdx)

	// fmt.Printf("Choose pool: %s, with balance: %d \n", pool, fromBalance)

	fromBalance = new(big.Int).Add(fromBalance, amountIn)
	// fromBalance = new(big.Int).Mul(fromBalance, tokenDecimalMap[strings.ToUpper(to)])

	price = new(big.Rat).SetFrac(fromBalance, new(big.Int).Sub(toBalance, amountOut))

	amountOut, _, _ = provider.getAmountOutImpl(from, to, amountIn, &pool) // get other pool amount
	// amountOut = new(big.Int).Div(amountOut, tokenDecimalMap[strings.ToUpper(to)])
	otherPoolPrice := new(big.Rat).SetFrac(amountOut, amountIn) // calculate other pools' best price

	if price.Cmp(otherPoolPrice) < 0 {
		price = otherPoolPrice
	}
	// price = new(big.Rat).SetFrac(price, tokenDecimalMap[strings.ToUpper(to)])
	return
}

// https://etherscan.io/address/0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56#code
