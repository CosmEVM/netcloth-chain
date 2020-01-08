package vm

import (
	"github.com/netcloth/netcloth-chain/modules/vm/common"
	"github.com/netcloth/netcloth-chain/modules/vm/keeper"
	"github.com/netcloth/netcloth-chain/modules/vm/types"
)

const (
	ModuleName        = types.ModuleName
	StoreKey          = types.StoreKey
	CodeKey           = types.CodeKey
	RouterKey         = types.RouterKey
	QuerierRoute      = types.QuerierRoute
	DefaultCodespace  = types.DefaultCodespace
	DefaultParamspace = keeper.DefaultParamspace
)

type (
	Keeper        = keeper.Keeper
	MsgContract   = types.MsgContract
	CommitStateDB = types.CommitStateDB
	Log           = types.Log
)

var (
	NewKeeper        = keeper.NewKeeper
	NewCommitStateDB = types.NewCommitStateDB

	CreateAddress  = common.CreateAddress
	CreateAddress2 = common.CreateAddress2

	ErrOutOfGas                 = types.ErrOutOfGas
	ErrCodeStoreOutOfGas        = types.ErrCodeStoreOutOfGas
	ErrDepth                    = types.ErrDepth
	ErrTraceLimitReached        = types.ErrTraceLimitReached
	ErrInsufficientBalance      = types.ErrInsufficientBalance
	ErrContractAddressCollision = types.ErrContractAddressCollision
	ErrNoCompatibleInterpreter  = types.ErrNoCompatibleInterpreter
	ErrEmptyInputs              = types.ErrEmptyInputs
	ErrNoCodeExist              = types.ErrNoCodeExist
	ErrMaxCodeSizeExceeded      = types.ErrMaxCodeSizeExceeded
	ErrWriteProtection          = types.ErrWriteProtection
	ErrReturnDataOutOfBounds    = types.ErrReturnDataOutOfBounds
	ErrExecutionReverted        = types.ErrExecutionReverted
	ErrInvalidJump              = types.ErrInvalidJump
	ErrGasUintOverflow          = types.ErrGasUintOverflow
)
