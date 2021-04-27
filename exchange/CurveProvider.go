package exchange

import (
	"fmt"
	CurveContract "hex/amm/v1/CurveContract"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/juliangruber/go-intersect"
)

type CurveProvider struct {
	TransactionFee int64
	PoolInstances  map[string](*CurveContract.Curve)
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
	"BUSD": big.NewInt(3),
	"GUSD": big.NewInt(3),
}

func NewCurveProvider(client *ethclient.Client) (*CurveProvider, error) {
	provider := CurveProvider{}
	provider.PoolInstances = make(map[string](*CurveContract.Curve))

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

// func getCoinIndex(coin string, pool string) []string {
// 	list1 := namePoolMap[from]
// 	list2 := namePoolMap[from]
// 	return intersect.Simple(list1, list2)
// }

func (provider *CurveProvider) GetAmountOut(from string, to string, amountIn *big.Int) (*big.Rat, error) {
	poolList := getPoolList(from, to)

	fromIdx := nameIndexMap[from]
	toIdx := nameIndexMap[to]
	// add check if swapping pare is invalid

	max := big.NewInt(0)
	for _, p := range poolList {
		ps := p.(string)

		amountOut, _ := provider.PoolInstances[ps].GetDyUnderlying(nil, fromIdx, toIdx, amountIn)
		if amountOut.Cmp(max) > 0 {
			max = amountOut
			fmt.Printf("pool is: %s, %s \n", ps, max.String())
		}
	}
	retMax := new(big.Rat).SetInt(max)
	return retMax, nil
}
