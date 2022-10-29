package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "mdc/testutil/keeper"
	"mdc/x/mdc/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MdcKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
