package types

import (
	sdk "github.com/netcloth/netcloth-chain/types"
)

const (
	QueryIPALNodeList = "list"
	QueryIPALNode     = "node"
	QueryIPALNodes    = "nodes"
	QueryParameters   = "params"
)

type QueryIPALNodeParams struct {
	AccAddr sdk.AccAddress
}

type QueryIPALNodesParams struct {
	AccAddrs []sdk.AccAddress `json:"acc_addrs"`
}

func NewQueryIPALNodeParams(AccAddr sdk.AccAddress) QueryIPALNodeParams {
	return QueryIPALNodeParams{
		AccAddr: AccAddr,
	}
}

func NewQueryIPALNodesParams(AccAddrs []sdk.AccAddress) QueryIPALNodesParams {
	return QueryIPALNodesParams{
		AccAddrs: AccAddrs,
	}
}
