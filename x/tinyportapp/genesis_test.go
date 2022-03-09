package tinyportapp_test

import (
	"testing"

	keepertest "github.com/notional-labs/tinyportapp/testutil/keeper"
	"github.com/notional-labs/tinyportapp/testutil/nullify"
	"github.com/notional-labs/tinyportapp/x/tinyportapp"
	"github.com/notional-labs/tinyportapp/x/tinyportapp/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by tinyport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TinyportappKeeper(t)
	tinyportapp.InitGenesis(ctx, *k, genesisState)
	got := tinyportapp.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by tinyport scaffolding # genesis/test/assert
}
