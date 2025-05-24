package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
"os"

	"github.com/aditysoni/flashloan-bot/bestswap"
	"github.com/joho/godotenv"

)

func main() {
	// Connect to Ethereum node
	err := godotenv.Load()
	if err != nil {
        log.Fatal("Error loading .env file")
    }
    rpc := os.Getenv("RPC") 
	if rpc == "" {
		log.Fatal("RPC not provided")

	}
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal("Failed to connect to Ethereum:", err)
	}

	best, err := bestswap.GetBestDexOutput(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Best DEX: %s, Output: %s\n", best.DEX, best.Output.String())

}
