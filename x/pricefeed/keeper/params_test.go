package keeper_test

import (
	"testing"

	"github.com/NibiruChain/nibiru/x/pricefeed/types"
	"github.com/NibiruChain/nibiru/x/testutil"
	"github.com/NibiruChain/nibiru/x/testutil/sample"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	testCases := []struct {
		name string
		test func()
	}{
		{
			name: "calling GetParams without setting returns default",
			test: func() {
				nibiruApp, ctx := testutil.NewNibiruApp(true)
				k := nibiruApp.PriceKeeper
				require.EqualValues(t, types.DefaultParams(), k.GetParams(ctx))
			},
		},
		{
			name: "params match after manual set and include default",
			test: func() {
				nibiruApp, ctx := testutil.NewNibiruApp(true)
				k := nibiruApp.PriceKeeper
				params := types.Params{
					Pairs: []types.Pair{
						{Token1: "btc", Token0: "usd", Oracles: nil, Active: true},
						{Token1: "xrp", Token0: "usd", Oracles: nil, Active: true},
					},
				}
				k.SetParams(ctx, params)
				require.EqualValues(t, params, k.GetParams(ctx))

				params.Pairs = append(params.Pairs, types.DefaultPairs...)
				k.SetParams(ctx, params)
				require.EqualValues(t, params, k.GetParams(ctx))
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			tc.test()
		})
	}
}

func TestWhitelistOracles(t *testing.T) {
	var noOracles []sdk.AccAddress

	testCases := []struct {
		name string
		test func()
	}{
		{
			name: "genesis - no oracle provided",
			test: func() {
				nibiruApp, ctx := testutil.NewNibiruApp(true)
				pk := &nibiruApp.PriceKeeper

				oracle := sample.AccAddress()
				for _, pair := range pk.GetPairs(ctx) {
					require.NotContains(t, pair.Oracles, oracle)
					require.EqualValues(t, pair.Oracles, noOracles)
				}
				require.EqualValues(t,
					pk.GetAuthorizedAddresses(ctx), noOracles)
			},
		},
		{
			name: "multiple oracles whitelisted at different times ",
			test: func() {
				nibiruApp, ctx := testutil.NewNibiruApp(true)
				pk := &nibiruApp.PriceKeeper

				for _, pair := range pk.GetPairs(ctx) {
					require.EqualValues(t, pair.Oracles, noOracles)
				}
				require.EqualValues(t,
					pk.GetAuthorizedAddresses(ctx), noOracles)

				oracleA := sample.AccAddress()
				oracleB := sample.AccAddress()

				wantOracles := []sdk.AccAddress{oracleA}
				pk.WhitelistOracles(ctx, wantOracles)
				gotOracles := pk.GetAuthorizedAddresses(ctx)
				require.EqualValues(t, wantOracles, gotOracles)
				require.NotContains(t, gotOracles, oracleB)

				wantOracles = []sdk.AccAddress{oracleA, oracleB}
				pk.WhitelistOracles(ctx, wantOracles)
				gotOracles = pk.GetAuthorizedAddresses(ctx)
				require.EqualValues(t, wantOracles, gotOracles)
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			tc.test()
		},
		)
	}
}
