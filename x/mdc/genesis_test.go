package mdc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "mdc/testutil/keeper"
	"mdc/testutil/nullify"
	"mdc/x/mdc"
	"mdc/x/mdc/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MdcKeeper(t)
	mdc.InitGenesis(ctx, *k, genesisState)
	got := mdc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
