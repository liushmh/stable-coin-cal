package exchange

import (
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

func (provider *MStableProvider) GetAmountOut(from string, to string, amountIn *big.Int) (*big.Rat, error) {
	upFrom := strings.ToUpper(from)
	upTo := strings.ToUpper(to)

	if upFrom == "MUSD" || upTo == "MUSD" {
		amountOut, err := _getAmountFromAsset(provider.Pools["MUSD"], upFrom, upTo, amountIn)
		return new(big.Rat).SetInt(amountOut), err
	}

	poolName := getPoolName(upFrom, upTo)
	amountOut, err := provider.Pools[poolName].GetSwapOutput(nil, nameAddress[upFrom], nameAddress[upTo], amountIn)

	return new(big.Rat).SetInt(amountOut), err
}
