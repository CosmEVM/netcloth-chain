package app

import (
	"github.com/netcloth/netcloth-chain/app/v0"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/netcloth/netcloth-chain/codec"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestExport(t *testing.T) {
	db := db.NewMemDB()
	app := NewNCHApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0)
	setGenesis(app)

	// Making a new app object with the db, so that initchain hasn't been called
	newApp := NewNCHApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0)
	_, _, err := newApp.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

func setGenesis(app *NCHApp) error {
	genesisState := v0.NewDefaultGenesisState()

	stateBytes, err := codec.MarshalJSONIndent(app.Engine.GetCurrentProtocol().GetCodec(), genesisState)
	if err != nil {
		return err
	}

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	return nil
}
