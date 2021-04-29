package exchange

import (
	"errors"
	MStableContract "hex/amm/v1/MStableContract"

	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type MStableProvider struct {
	Fee             int64
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
	provider.Fee = 3

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

func (provider *MStableProvider) getAmountOutImpl(from string, to string, amountIn *big.Int) (amountOut *big.Int, fee *big.Int, err error) {
	upFrom := strings.ToUpper(from)
	upTo := strings.ToUpper(to)

	realAmountIn := new(big.Int).Mul(amountIn, tokenDecimalMap[upFrom])
	fee = new(big.Int).Mul(amountIn, big.NewInt(provider.Fee))

	if upFrom == "MUSD" || upTo == "MUSD" {
		amountOut, err = _getAmountFromAsset(provider.Pools["MUSD"], upFrom, upTo, realAmountIn)
	} else {
		fee = new(big.Int).Mul(fee, big.NewInt(2))
		poolName := getPoolName(upFrom, upTo)
		amountOut, err = provider.Pools[poolName].GetSwapOutput(nil, nameAddress[upFrom], nameAddress[upTo], realAmountIn)
	}

	// amountOut = new(big.Int).Div(tmpOut, tokenDecimalMap[upTo])
	return
}

// todo: if it's BUSD, we have extra 12 zeros
func (provider *MStableProvider) GetAmountOut(from string, to string, amountIn *big.Int) (*big.Rat, *big.Rat, error) {
	amountOut, fee, err := provider.getAmountOutImpl(from, to, amountIn)
	if err != nil {
		// log.Fatal("cannot ")
		return nil, nil, err
	}
	amountOut = new(big.Int).Div(amountOut, tokenDecimalMap[strings.ToUpper(to)])

	return new(big.Rat).SetInt(amountOut), new(big.Rat).SetFrac(fee, big.NewInt(1e4)), err
}

func (provider *MStableProvider) GetPriceAfterAmount(from string, to string, amountIn *big.Int) (price *big.Rat, err error) {

	dx1 := big.NewInt(10000)
	dx2 := big.NewInt(20000)
	dy1, _, _ := provider.getAmountOutImpl(from, to, dx1)
	dy2, _, _ := provider.getAmountOutImpl(from, to, dx2)

	x, y := calculateXY(dx1, dy1, dx2, dy2)

	amountOut, _, _ := provider.getAmountOutImpl(from, to, amountIn)

	newY := new(big.Int).Sub(y, amountOut)
	if newY.Cmp(big.NewInt(0)) <= 0 {
		err = errors.New("Wrong Amount")
	}
	newY = new(big.Int).Div(newY, tokenDecimalMap[strings.ToUpper(to)])
	price = new(big.Rat).SetFrac(new(big.Int).Add(x, amountIn), newY)
	return
}
