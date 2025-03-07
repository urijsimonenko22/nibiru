package keeper_test

import (
	"fmt"
	"testing"

	"github.com/NibiruChain/nibiru/simapp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/NibiruChain/nibiru/x/common"
	"github.com/NibiruChain/nibiru/x/stablecoin/types"
)

func TestGetParams(t *testing.T) {
	nibiruApp, ctx := simapp.NewTestNibiruAppAndContext(true)
	stableKeeper := &nibiruApp.StablecoinKeeper

	params := types.DefaultParams()

	stableKeeper.SetParams(ctx, params)

	require.EqualValues(t, params, stableKeeper.GetParams(ctx))
}

func TestNewParams_Errors(t *testing.T) {
	tests := []struct {
		name          string
		params        types.Params
		expectedError error
	}{
		{
			"collateral ratio bigger than 1",
			types.NewParams(
				sdk.MustNewDecFromStr("2"),
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("0.002"),
				"15 min",
				sdk.MustNewDecFromStr("0.0025"),
				sdk.MustNewDecFromStr("0.9999"),
				sdk.MustNewDecFromStr("1.0001"),
				true,
			),
			fmt.Errorf(
				"collateral ratio is above max value(1e6): %s",
				sdk.MustNewDecFromStr("2").Mul(sdk.NewDec(1*common.Precision)).TruncateInt()),
		},
		{
			"fee ratio bigger than 1",
			types.NewParams(
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("2"),
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("0.002"),
				"15 min",
				sdk.MustNewDecFromStr("0.0025"),
				sdk.MustNewDecFromStr("0.9999"),
				sdk.MustNewDecFromStr("1.0001"),
				true,
			),
			fmt.Errorf(
				"fee ratio is above max value(1e6): %s",
				sdk.MustNewDecFromStr("2").Mul(sdk.NewDec(1*common.Precision)).TruncateInt()),
		},
		{
			"stable EF fee ratio bigger than 1",
			types.NewParams(
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("1"),
				sdk.MustNewDecFromStr("2"),
				sdk.MustNewDecFromStr("0.002"),
				"15 min",
				sdk.MustNewDecFromStr("0.0025"),
				sdk.MustNewDecFromStr("0.9999"),
				sdk.MustNewDecFromStr("1.0001"),
				false,
			),
			fmt.Errorf(
				"stable EF fee ratio is above max value(1e6): %s",
				sdk.MustNewDecFromStr("2").Mul(sdk.NewDec(1*common.Precision)).TruncateInt()),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.Validate()
			require.EqualError(t, err, tc.expectedError.Error())
		})
	}
}
