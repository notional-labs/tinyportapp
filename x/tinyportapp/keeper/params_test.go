package keeper_test

import (
	"testing"

	testkeeper "github.com/notional-labs/tinyportapp/testutil/keeper"
	"github.com/notional-labs/tinyportapp/x/tinyportapp/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TinyportappKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
