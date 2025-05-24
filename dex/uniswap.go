package dex

import (
    "context"
    "fmt"
    "math/big"
	"os"


    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"

    uni "github.com/aditysoni/flashloan-bot/contracts/uniswapV2Router"
	types "github.com/aditysoni/flashloan-bot/types"
)



func UniswapOutput(client bind.ContractBackend, amountIn *big.Int, path []common.Address) (*types.DexOutput, error) {
    ctx := context.Background()

	uniswapRouterV2Str := os.Getenv("UNISWAP_ROUTER_V2")
	if uniswapRouterV2Str == "" {
		return nil, fmt.Errorf("UNISWAP_ROUTER_V2 environment variable is not set")
	}


    uniswapRouterV2Addr := common.HexToAddress(uniswapRouterV2Str)

    uniswap, err := uni.NewUniswapv2router(uniswapRouterV2Addr, client)
    if err != nil {
        return nil, fmt.Errorf("failed to load Uniswap router: %w", err)
    }

    outs, err := uniswap.GetAmountsOut(&bind.CallOpts{Context: ctx}, amountIn, path)
    if err != nil {
        return nil, fmt.Errorf("Uniswap GetAmountsOut failed: %w", err)
    }

    return &types.DexOutput{
        DEX:    "Uniswap",
        Output: outs[len(outs)-1],
        Router: uniswapRouterV2Addr,
    }, nil
}
