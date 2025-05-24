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



func SushuswapOutput(client bind.ContractBackend, amountIn *big.Int, path []common.Address) (*types.DexOutput, error) {
    ctx := context.Background()

	sushiswapRotuerAddr := os.Getenv("SUSHISWAP_ROUTER_V2")
	if sushiswapRotuerAddr == "" {
		return nil, fmt.Errorf("UNISWAP_ROUTER_V2 environment variable is not set")
	}


    sushiswapRouterV2 := common.HexToAddress(sushiswapRotuerAddr)

    sushiswap, err := uni.NewUniswapv2router(sushiswapRouterV2, client)
    if err != nil {
        return nil, fmt.Errorf("failed to load Sushiswap router: %w", err)
    }

    outs, err := sushiswap.GetAmountsOut(&bind.CallOpts{Context: ctx}, amountIn, path)
    if err != nil {
        return nil, fmt.Errorf("Sushiswap GetAmountsOut failed: %w", err)
    }

    return &types.DexOutput{
        DEX:    "SushiSwap",
        Output: outs[len(outs)-1],
        Router: sushiswapRouterV2,
    }, nil
}
