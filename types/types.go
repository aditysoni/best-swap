package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type DexOutput struct {
    DEX    string
    Output *big.Int
    Router common.Address
}