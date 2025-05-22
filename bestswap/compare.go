package bestswap

import (
    "context"
    "fmt"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"

    uni "github.com/aditysoni/flashloan-bot/contracts/uniswapV2Router"
)

type DexOutput struct {
    DEX      string
    Output   *big.Int
    Router   common.Address
}

func GetBestDexOutput(client bind.ContractBackend) (*DexOutput, error) {
    ctx := context.Background()

    fmt.Print("inside")

    amountIn := big.NewInt(1e18) // 1 WETH
	path := []common.Address{
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), // WETH
		common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"), // DAI
	}


    // Uniswap Router
    uniswapAddr := common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
    uniswap, err := uni.NewUniswapv2router(uniswapAddr, client)
    if err != nil {
        return nil, fmt.Errorf("failed to load uniswap router: %w", err)
    }

    fmt.Println("amountIn:", amountIn.String())


    uniOuts, err := uniswap.GetAmountsOut(&bind.CallOpts{Context: ctx}, amountIn, path)
    if err != nil {
        return nil, fmt.Errorf("uniswap GetAmountsOut failed: %w", err)
    }

    fmt.Print("uniswap1")


    // SushiSwap Router
    sushiAddr := common.HexToAddress("0xd9e1CE17f2641f24aE83637ab66a2cca9C378B9F")
    sushiswap, err := uni.NewUniswapv2router(sushiAddr, client)
    if err != nil {
        return nil, fmt.Errorf("failed to load sushiswap router: %w", err)
    }

    fmt.Print("uniswap-2")

    sushiOuts, err := sushiswap.GetAmountsOut(&bind.CallOpts{Context: ctx}, amountIn, path)
    if err != nil {
        return nil, fmt.Errorf("sushiswap GetAmountsOut failed: %w", err)
    }

    // Compare outputs
    fmt.Print("uniswap-2")

    uniOutput := uniOuts[len(uniOuts)-1]
    sushiOutput := sushiOuts[len(sushiOuts)-1]

    if uniOutput.Cmp(sushiOutput) > 0 {
        return &DexOutput{DEX: "Uniswap", Output: uniOutput, Router: uniswapAddr}, nil
    }

    return &DexOutput{DEX: "SushiSwap", Output: sushiOutput, Router: sushiAddr}, nil
}
