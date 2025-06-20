package bestswap

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/aditysoni/flashloan-bot/dex"
	types "github.com/aditysoni/flashloan-bot/types"
)

func GetBestDexOutput(client bind.ContractBackend, tokenInitial string, tokenFinal string ) (*types.DexOutput, error) {
	amountIn := big.NewInt(1e18) // 1 WETH

	// WETH -> DAI path
	path := []common.Address{
		common.HexToAddress(tokenInitial), // WETH
		common.HexToAddress(tokenFinal), // DAI
	}

	// Fetch from Uniswap
	uniOut, err := dex.UniswapOutput(client, amountIn, path)
	if err != nil {
		return nil, fmt.Errorf("uniswap error: %w", err)
	}

	// Fetch from SushiSwap
	sushiOut, err := dex.SushuswapOutput(client, amountIn, path)
	if err != nil {
		return nil, fmt.Errorf("sushiswap error: %w", err)
	}

	// Compare and return best
	if uniOut.Output.Cmp(sushiOut.Output) > 0 {
		return uniOut, nil
	}
	return sushiOut, nil
}
