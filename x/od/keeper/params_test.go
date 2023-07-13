package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "opendao/testutil/keeper"
	"opendao/x/od/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OdKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
