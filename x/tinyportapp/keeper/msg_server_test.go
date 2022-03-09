package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/notional-labs/tinyportapp/testutil/keeper"
	"github.com/notional-labs/tinyportapp/x/tinyportapp/keeper"
	"github.com/notional-labs/tinyportapp/x/tinyportapp/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TinyportappKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
