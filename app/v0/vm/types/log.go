package types

import (
	"github.com/netcloth/netcloth-chain/hexutil"
	sdk "github.com/netcloth/netcloth-chain/types"
)

// Log represents a contract log event. These events are generated by the LOG opcode and
// stored/indexed by the node.

type Log struct {
	// address of the contract that generated the event
	Address sdk.AccAddress `json:"address" yaml:"address"`
	// list of topics provided by the contract.
	Topics []sdk.Hash `json:"topics" yaml:"topics"`
	// supplied by the contract, usually ABI-encoded
	Data hexutil.Bytes `json:"data" yaml:"data`

	BlockNumber uint64   `json:"blockNumber" yaml:"blockNumber"`
	TxHash      sdk.Hash `json:"transactionHash" yaml:"transactionHash"`
	TxIndex     uint     `json:"transactionIndex" yaml:"transactionIndex"`
	BlockHash   sdk.Hash `json:"blockHash" yaml:"blockHash"`
	Index       uint64   `json:"logIndex" yaml:"logIndex"`

	Removed bool `json:"removed" yaml:"removed"`
}
