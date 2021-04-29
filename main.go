package main

import (
	"fmt"
	"log"
	"math/big"

	"hex/amm/v1/exchange"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/a30f13ea65fe406a86783fa912982906")
	if err != nil {
		log.Fatal(err)
	}

	// ammPtr := flag.String("amm", "uniswap", "string")
	// fromPtr := flag.String("from", "DAI", "string")
	// toPtr := flag.String("to", "USDC", "string")

	// var ammMap = make(map[string](*exchange.ExchangeProvider))
	// ammMap["uniswap"] = exchange.NewUniswapProvider(client)

	// provider, err := exchange.NewUniswapProvider(client)
	provider, err := exchange.NewMStableProvider(client)
	// provider, err := exchange.NewCurveProvider(client)

	if err != nil {
		log.Fatal(err)
	}

	// amountIn := big.NewInt(1e6)
	tokenA := "DAI"
	tokenB := "USDC"

	// amountOut, fee, err := provider.GetAmountOut(tokenA, tokenB, amountIn)
	// fmt.Printf("swap %s %s for %s %s with fee: %s \n", tokenA, amountIn.String(), tokenB, amountOut.FloatString(2), fee.FloatString(2))

	var i int64
	for i = 1; i <= 1e9; i = i * 10 {
		amtIn := big.NewInt(i)
		amountOut, fee, err := provider.GetAmountOut(tokenA, tokenB, amtIn)
		if err != nil {
			log.Fatal("cannot do this swap", err)
			continue
		}

		price := new(big.Rat).Mul(new(big.Rat).SetInt64(i), new(big.Rat).Inv(amountOut))
		fmt.Printf("swap %s %s for %s %s with price %s and fee: %s \n", tokenA, amtIn.String(), tokenB, amountOut.FloatString(2), price.FloatString(5), fee.FloatString(2))
	}
	// price, err := provider.GetPriceAfterAmount("DAI", "USDC", million)
	// fmt.Printf("new price after swap: %s \n", price.FloatString(8))

	if err != nil {
		log.Fatal(err)
	}
}
