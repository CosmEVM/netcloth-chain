package vm

import (
	"math/big"

	sdk "github.com/netcloth/netcloth-chain/types"
)

type StateDB interface {
	SubBalance(sdk.AccAddress, *big.Int)
	AddBalance(sdk.AccAddress, *big.Int)
	GetBalance(sdk.AccAddress) *big.Int

	AddPreimage(sdk.Hash, []byte)
}
