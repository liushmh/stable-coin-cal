package main

import (
	"fmt"
	"log"
	"math/big"

	"hex/amm/v1/exchange"

	"github.com/ethereum/go-ethereum/ethclient"
)

func calculateY(total_x, total_y, x *big.Int) big.Rat {
	cp := new(big.Int).Mul(total_x, total_y)
	new_x := new(big.Int).Add(total_x, x)
	new_y := new(big.Rat).SetFrac(cp, new_x)

	dy := new(big.Rat).Sub(new(big.Rat).SetInt(total_y), new_y)

	fmt.Printf("get : %s y \n", dy.FloatString(8))
	price := new(big.Rat).Mul(new(big.Rat).SetInt(x), new(big.Rat).Inv(dy))
	return *price
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/a30f13ea65fe406a86783fa912982906")
	if err != nil {
		log.Fatal(err)
	}

	// provider, err := exchange.NewUniswapProvider(client)
	provider, err := exchange.NewMStableProvider(client)

	if err != nil {
		log.Fatal(err)
	}

	million := new(big.Int)
	million.SetString("1000000000000", 10)

	info, err := provider.GetAmountOut("USDT", "BUSD", million)
	fmt.Printf("how much we get: %s \n", info.FloatString(8))

	price, err := provider.GetPriceAfterAmount("USDT", "BUSD", million)
	fmt.Printf("new price after swap: %s \n", price.FloatString(8))

	if err != nil {
		log.Fatal(err)
	}
}
