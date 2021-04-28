package exchange

import (
	"errors"
	"fmt"
	MStableContract "hex/amm/v1/MStableContract"

	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MStableProvider struct {
	TransactionFee  int64
	ContractAddress common.Address
	Pools           map[string](*MStableContract.MStable)
}

// var mAssetAddress = common.HexToAddress("0xe2f2a5C287993345a840Db3B0845fbC70f5935a5")

var poolAssetAddress = map[string]common.Address{
	"MUSD": common.HexToAddress("0xe2f2a5C287993345a840Db3B0845fbC70f5935a5"),
	"BUSD": common.HexToAddress("0xfE842e95f8911dcc21c943a1dAA4bd641a1381c6"),
	"GUSD": common.HexToAddress("0x4fB30C5A3aC8e85bC32785518633303C4590752d"),
}

func NewMStableProvider(client *ethclient.Client) (*MStableProvider, error) {
	provider := MStableProvider{}
	provider.Pools = make(map[string](*MStableContract.MStable))

	err := error(nil)

	for k, p := range poolAssetAddress {
		instance, err := MStableContract.NewMStable(p, client)
		if err != nil {
			continue
		}
		provider.Pools[k] = instance
	}

	return &provider, err
}

func getPoolName(from string, to string) string {
	if from == "BUSD" || to == "BUSD" {
		return "BUSD"
	}

	if from == "GUSD" || to == "GUSD" {
		return "GUSD"
	}
	return "MUSD"
}

func _getAmountFromAsset(instance *MStableContract.MStable, from string, to string, amountIn *big.Int) (*big.Int, error) {
	if from == "MUSD" {
		return instance.GetRedeemOutput(nil, nameAddress[to], amountIn)
	} else {
		return instance.GetMintOutput(nil, nameAddress[from], amountIn)
	}
}

func (provider *MStableProvider) getAmountOutImpl(from string, to string, amountIn *big.Int) (*big.Int, error) {
	upFrom := strings.ToUpper(from)
	upTo := strings.ToUpper(to)

	if upFrom == "MUSD" || upTo == "MUSD" {
		return _getAmountFromAsset(provider.Pools["MUSD"], upFrom, upTo, amountIn)
	}

	poolName := getPoolName(upFrom, upTo)
	return provider.Pools[poolName].GetSwapOutput(nil, nameAddress[upFrom], nameAddress[upTo], amountIn)
}

// todo: if it's BUSD, we have extra 12 zeros
func (provider *MStableProvider) GetAmountOut(from string, to string, amountIn *big.Int) (*big.Rat, error) {
	amountOut, err := provider.getAmountOutImpl(from, to, amountIn)
	return new(big.Rat).SetInt(amountOut), err
}

func (provider *MStableProvider) GetPriceAfterAmount(from string, to string, amountIn *big.Int) (price *big.Rat, err error) {
	dx1 := big.NewInt(10000000000)
	dx2 := big.NewInt(20000000000)
	dy1, _ := provider.getAmountOutImpl(from, to, dx1)
	dy2, _ := provider.getAmountOutImpl(from, to, dx2)

	fmt.Printf("dy1 is %d, dy2 is %d \n", dy1, dy2)

	x, y := calculateXY(dx1, dy1, dx2, dy2)

	fmt.Printf("x is %d, y is %d \n", x, y)
	amountOut, _ := provider.getAmountOutImpl(from, to, amountIn)

	newY := new(big.Int).Sub(y, amountOut)
	if newY.Cmp(big.NewInt(0)) <= 0 {
		err = errors.New("Wrong Amount")
	}

	price = new(big.Rat).SetFrac(new(big.Int).Add(x, amountIn), newY)
	return
}
