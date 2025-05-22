package main

import (
	"fmt"
	"log"
	// "math/big"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	// factory "github.com/aditysoni/flashloan-bot/contracts/uniswapV2Factory"
	// pair "github.com/aditysoni/flashloan-bot/contracts/uniswapV2Pair"
	"github.com/aditysoni/flashloan-bot/bestswap"
)

func main() {
	// Connect to Ethereum node
	client, err := ethclient.Dial("https://go.getblock.io/e9141006c08a4042a6272cc370f4310a")
	if err != nil {
		log.Fatal("Failed to connect to Ethereum:", err)
	}

	// Uniswap V2 Factory address
	// factoryAddress := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f") // Mainnet
	// factoryInstance, err := factory.NewUniswapv2factory(factoryAddress, client)
	// if err != nil {
	//     log.Fatal("Failed to load factory contract:", err)
	// }

	// // Get number of pairs
	// count, err := factoryInstance.AllPairsLength(nil)
	// if err != nil {
	//     log.Fatal("Failed to get pairs length:", err)
	// }

	// fmt.Println("Total Pairs:", count)

	// // Iterate through all pairs
	// for i := int64(100); i < count.Int64(); i++ {
	//     pairAddr, err := factoryInstance.AllPairs(nil, big.NewInt(i))
	//     if err != nil {
	//         log.Printf("Failed to get pair at index %d: %v", i, err)
	//         continue
	//     }

	//     pairInstance, err := pair.NewUniswapv2pair(pairAddr, client)
	//     if err != nil {
	//         log.Printf("Failed to bind pair contract: %v", err)
	//         continue
	//     }

	//     token0, _ := pairInstance.Token0(nil)
	//     token1, _ := pairInstance.Token1(nil)
	//     reserves, _ := pairInstance.GetReserves(nil)

	//     fmt.Printf("[%d] Pair: %s | Token0: %s | Token1: %s | Reserves: %s / %s\n",
	//         i,
	//         pairAddr.Hex(),
	//         token0.Hex(),
	//         token1.Hex(),
	//         reserves.Reserve0.String(),
	//         reserves.Reserve1.String(),
	//     )
	// }

	// amountIn := big.NewInt(1e18) // 1 WETH
	// path := []common.Address{
	// 	common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), // WETH
	// 	common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"), // DAI
	// }

	best, err := bestswap.GetBestDexOutput(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Best DEX: %s, Output: %s\n", best.DEX, best.Output.String())

}
