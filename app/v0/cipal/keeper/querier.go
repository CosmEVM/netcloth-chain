package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/netcloth/netcloth-chain/app/v0/cipal/types"
	"github.com/netcloth/netcloth-chain/codec"
	sdk "github.com/netcloth/netcloth-chain/types"
	sdkerrors "github.com/netcloth/netcloth-chain/types/errors"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QueryCIPAL:
			return queryCIPAL(ctx, req, k)
		case types.QueryCIPALs:
			return queryCIPALs(ctx, req, k)
		case types.QueryCIPALCount:
			return queryCIPALCount(ctx, req, k)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query path: %s", path[0])
		}
	}
}

func queryCIPAL(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var queryParams types.QueryCIPALParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &queryParams)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	cipal, found := k.GetCIPALObject(ctx, queryParams.AccAddr)
	if found {
		bz, err := codec.MarshalJSONIndent(types.ModuleCdc, cipal)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		return bz, nil
	}

	return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "not found")
}

type CIPALCount struct {
	Count int `json:"count" yaml:"count"`
}

func NewCIPALCount(c int) CIPALCount {
	return CIPALCount{
		Count: c,
	}
}

func queryCIPALCount(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	count := k.GetCIPALObjectCount(ctx)
	res := NewCIPALCount(count)

	bz, err := codec.MarshalJSONIndent(types.ModuleCdc, res)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func queryCIPALs(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryCIPALsParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	cipals := types.CIPALObjects{}
	for _, addr := range params.AccAddrs {
		cipal, found := k.GetCIPALObject(ctx, addr)
		if found {
			cipals = append(cipals, cipal)
		}
	}

	bz, err := codec.MarshalJSONIndent(types.ModuleCdc, cipals)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
