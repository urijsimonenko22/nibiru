package keeper_test

import (
	"fmt"
	"testing"

	"github.com/NibiruChain/nibiru/x/testutil"

	"github.com/NibiruChain/nibiru/x/testutil/testapp"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/NibiruChain/nibiru/x/common"
	"github.com/NibiruChain/nibiru/x/dex/keeper"
	"github.com/NibiruChain/nibiru/x/dex/types"
	"github.com/NibiruChain/nibiru/x/testutil/mock"
)

func TestCreatePool(t *testing.T) {
	tests := []struct {
		name               string
		creatorAddr        sdk.AccAddress
		poolParams         types.PoolParams
		poolAssets         []types.PoolAsset
		senderInitialFunds sdk.Coins
		expectedErr        error
	}{
		{
			name:        "invalid creator addr",
			creatorAddr: []byte{},
			poolParams:  types.PoolParams{PoolType: types.PoolType_BALANCER},
			poolAssets:  []types.PoolAsset{},
			expectedErr: fmt.Errorf("empty address string is not allowed"),
		},
		{
			name:        "not enough assets",
			poolParams:  types.PoolParams{PoolType: types.PoolType_BALANCER},
			poolAssets:  []types.PoolAsset{},
			expectedErr: types.ErrTooFewPoolAssets,
		},
		{
			name:       "too many assets",
			poolParams: types.PoolParams{PoolType: types.PoolType_BALANCER},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin("aaa", 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin("bbb", 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin("ccc", 1),
					Weight: sdk.OneInt(),
				},
			},
			expectedErr: types.ErrTooManyPoolAssets,
		},
		{
			name:       "asset not whitelisted 1",
			poolParams: types.PoolParams{PoolType: types.PoolType_BALANCER},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin(common.DenomUSDC, 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin("aaaa", 1),
					Weight: sdk.OneInt(),
				},
			},
			expectedErr: types.ErrTokenNotAllowed,
		},
		{
			name:       "asset not whitelisted 2",
			poolParams: types.PoolParams{PoolType: types.PoolType_BALANCER},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin("aaa", 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin(common.DenomUSDC, 1),
					Weight: sdk.OneInt(),
				},
			},
			expectedErr: types.ErrTokenNotAllowed,
		},
		{
			name:       "insufficient pool creation fee",
			poolParams: types.PoolParams{PoolType: types.PoolType_BALANCER},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin(common.DenomUSDC, 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin(common.DenomNUSD, 1),
					Weight: sdk.OneInt(),
				},
			},
			senderInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 1e9-1),
				sdk.NewInt64Coin("aaa", 1),
				sdk.NewInt64Coin("bbb", 1),
			),
			expectedErr: fmt.Errorf("999999999unibi is smaller than 1000000000unibi: insufficient funds"),
		},
		{
			name:       "insufficient initial deposit",
			poolParams: types.PoolParams{PoolType: types.PoolType_BALANCER},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin(common.DenomNIBI, 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin(common.DenomUSDC, 1),
					Weight: sdk.OneInt(),
				},
			},
			senderInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 1e9),
			),
			expectedErr: fmt.Errorf("0unibi is smaller than 1unibi: insufficient funds"),
		},
		{
			name:       "successful pool creation",
			poolParams: types.PoolParams{PoolType: types.PoolType_BALANCER},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin(common.DenomNUSD, 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin(common.DenomUSDC, 1),
					Weight: sdk.OneInt(),
				},
			},
			senderInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin(common.DenomNIBI, 1e9),
				sdk.NewInt64Coin(common.DenomNUSD, 1),
				sdk.NewInt64Coin(common.DenomUSDC, 1),
			),
			expectedErr: nil,
		},
		{
			name:        "invalid creator addr - Stableswap",
			creatorAddr: []byte{},
			poolParams:  types.PoolParams{PoolType: types.PoolType_STABLESWAP, A: sdk.OneInt()},
			poolAssets:  []types.PoolAsset{},
			expectedErr: fmt.Errorf("empty address string is not allowed"),
		},
		{
			name:        "not enough assets - Stableswap",
			poolParams:  types.PoolParams{PoolType: types.PoolType_STABLESWAP, A: sdk.OneInt()},
			poolAssets:  []types.PoolAsset{},
			expectedErr: types.ErrTooFewPoolAssets,
		},
		{
			name:       "too many assets - Stableswap",
			poolParams: types.PoolParams{PoolType: types.PoolType_STABLESWAP, A: sdk.OneInt()},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin("aaa", 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin("bbb", 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin("ccc", 1),
					Weight: sdk.OneInt(),
				},
			},
			expectedErr: types.ErrTooManyPoolAssets,
		},
		{
			name:       "asset not whitelisted 1 - Stableswap",
			poolParams: types.PoolParams{PoolType: types.PoolType_STABLESWAP, A: sdk.OneInt()},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin(common.DenomUSDC, 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin("aaaa", 1),
					Weight: sdk.OneInt(),
				},
			},
			expectedErr: types.ErrTokenNotAllowed,
		},
		{
			name:       "asset not whitelisted 2 - Stableswap",
			poolParams: types.PoolParams{PoolType: types.PoolType_STABLESWAP, A: sdk.OneInt()},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin("aaa", 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin(common.DenomUSDC, 1),
					Weight: sdk.OneInt(),
				},
			},
			expectedErr: types.ErrTokenNotAllowed,
		},
		{
			name:       "insufficient pool creation fee - Stableswap",
			poolParams: types.PoolParams{PoolType: types.PoolType_STABLESWAP, A: sdk.OneInt()},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin(common.DenomUSDC, 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin(common.DenomNUSD, 1),
					Weight: sdk.OneInt(),
				},
			},
			senderInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 1e9-1),
				sdk.NewInt64Coin("aaa", 1),
				sdk.NewInt64Coin("bbb", 1),
			),
			expectedErr: fmt.Errorf("999999999unibi is smaller than 1000000000unibi: insufficient funds"),
		},
		{
			name:       "insufficient initial deposit - Stableswap",
			poolParams: types.PoolParams{PoolType: types.PoolType_STABLESWAP, A: sdk.OneInt()},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin(common.DenomNIBI, 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin(common.DenomUSDC, 1),
					Weight: sdk.OneInt(),
				},
			},
			senderInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 1e9),
			),
			expectedErr: fmt.Errorf("0unibi is smaller than 1unibi: insufficient funds"),
		},
		{
			name:       "successful pool creation - Stableswap",
			poolParams: types.PoolParams{PoolType: types.PoolType_STABLESWAP, A: sdk.OneInt()},
			poolAssets: []types.PoolAsset{
				{
					Token:  sdk.NewInt64Coin(common.DenomNUSD, 1),
					Weight: sdk.OneInt(),
				},
				{
					Token:  sdk.NewInt64Coin(common.DenomUSDC, 1),
					Weight: sdk.OneInt(),
				},
			},
			senderInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin(common.DenomNIBI, 1e9),
				sdk.NewInt64Coin(common.DenomNUSD, 1),
				sdk.NewInt64Coin(common.DenomUSDC, 1),
			),
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testapp.NewTestNibiruAppAndContext(true)
			msgServer := keeper.NewMsgServerImpl(app.DexKeeper)

			if tc.creatorAddr == nil {
				tc.creatorAddr = ed25519.GenPrivKey().PubKey().Address().Bytes()
			}
			if tc.senderInitialFunds != nil {
				require.NoError(t, simapp.FundAccount(app.BankKeeper, ctx, tc.creatorAddr, tc.senderInitialFunds))
			}

			msgCreatePool := types.MsgCreatePool{
				Creator:    tc.creatorAddr.String(),
				PoolParams: &tc.poolParams,
				PoolAssets: tc.poolAssets,
			}

			_, err := msgServer.CreatePool(sdk.WrapSDKContext(ctx), &msgCreatePool)
			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
				testutil.RequireNotHasTypedEvent(t, ctx, &types.EventPoolCreated{
					Creator: tc.creatorAddr.String(),
					PoolId:  1,
				})
			} else {
				require.NoError(t, err)
				testutil.RequireHasTypedEvent(t, ctx, &types.EventPoolCreated{
					Creator: tc.creatorAddr.String(),
					PoolId:  1,
				})
			}
		})
	}
}

func TestMsgServerJoinPool(t *testing.T) {
	const shareDenom = "nibiru/pool/1"
	tests := []struct {
		name                     string
		joinerInitialFunds       sdk.Coins
		initialPool              types.Pool
		tokensIn                 sdk.Coins
		expectedNumSharesOut     sdk.Coin
		expectedRemCoins         sdk.Coins
		expectedJoinerFinalFunds sdk.Coins
		expectedFinalPool        types.Pool
	}{
		{
			name: "join with all assets",
			joinerInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			initialPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100),
			tokensIn: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			expectedNumSharesOut:     sdk.NewInt64Coin(shareDenom, 100),
			expectedRemCoins:         sdk.NewCoins(),
			expectedJoinerFinalFunds: sdk.NewCoins(sdk.NewInt64Coin(shareDenom, 100)),
			expectedFinalPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 200),
					sdk.NewInt64Coin(common.DenomNUSD, 200),
				),
				/*shares=*/ 200),
		},
		{
			name: "join with some assets, none remaining",
			joinerInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			initialPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100),
			tokensIn: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 50),
				sdk.NewInt64Coin(common.DenomNUSD, 50),
			),
			expectedNumSharesOut: sdk.NewInt64Coin(shareDenom, 50),
			expectedRemCoins:     sdk.NewCoins(),
			expectedJoinerFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin(shareDenom, 50),
				sdk.NewInt64Coin("unibi", 50),
				sdk.NewInt64Coin(common.DenomNUSD, 50),
			),
			expectedFinalPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 150),
					sdk.NewInt64Coin(common.DenomNUSD, 150),
				),
				/*shares=*/ 150),
		},
		{
			name: "join with some assets, some remaining",
			joinerInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			initialPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100),
			tokensIn: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 50),
				sdk.NewInt64Coin(common.DenomNUSD, 75),
			),
			expectedNumSharesOut: sdk.NewInt64Coin(shareDenom, 50),
			expectedRemCoins: sdk.NewCoins(
				sdk.NewInt64Coin(common.DenomNUSD, 25),
			),
			expectedJoinerFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin(shareDenom, 50),
				sdk.NewInt64Coin("unibi", 50),
				sdk.NewInt64Coin(common.DenomNUSD, 50),
			),
			expectedFinalPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 150),
					sdk.NewInt64Coin(common.DenomNUSD, 150),
				),
				/*shares=*/ 150),
		},
		{
			name: "join with all assets - Stablepool",
			joinerInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			initialPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100),
			tokensIn: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			expectedNumSharesOut:     sdk.NewInt64Coin(shareDenom, 100),
			expectedRemCoins:         sdk.NewCoins(),
			expectedJoinerFinalFunds: sdk.NewCoins(sdk.NewInt64Coin(shareDenom, 100)),
			expectedFinalPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 200),
					sdk.NewInt64Coin(common.DenomNUSD, 200),
				),
				/*shares=*/ 200),
		},
		{
			name: "join with some assets, none remaining - Stablepool",
			joinerInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			initialPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100),
			tokensIn: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 50),
				sdk.NewInt64Coin(common.DenomNUSD, 50),
			),
			expectedNumSharesOut: sdk.NewInt64Coin(shareDenom, 50),
			expectedRemCoins:     []sdk.Coin{},
			expectedJoinerFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin(shareDenom, 50),
				sdk.NewInt64Coin("unibi", 50),
				sdk.NewInt64Coin(common.DenomNUSD, 50),
			),
			expectedFinalPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 150),
					sdk.NewInt64Coin(common.DenomNUSD, 150),
				),
				/*shares=*/ 150),
		},
		{
			name: "join with some assets, some remaining - Stablepool",
			joinerInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			initialPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100),
			tokensIn: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 50),
				sdk.NewInt64Coin(common.DenomNUSD, 75),
			),
			expectedNumSharesOut: sdk.NewInt64Coin(shareDenom, 62),
			expectedRemCoins:     []sdk.Coin{},
			expectedJoinerFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin(shareDenom, 62),
				sdk.NewInt64Coin("unibi", 50),
				sdk.NewInt64Coin(common.DenomNUSD, 25),
			),
			expectedFinalPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 150),
					sdk.NewInt64Coin(common.DenomNUSD, 175),
				),
				/*shares=*/ 162),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testapp.NewTestNibiruAppAndContext(true)

			poolAddr := testutil.AccAddress()
			tc.initialPool.Address = poolAddr.String()
			tc.expectedFinalPool.Address = poolAddr.String()
			app.DexKeeper.SetPool(ctx, tc.initialPool)

			joinerAddr := testutil.AccAddress()
			require.NoError(t, simapp.FundAccount(app.BankKeeper, ctx, joinerAddr, tc.joinerInitialFunds))

			msgServer := keeper.NewMsgServerImpl(app.DexKeeper)
			resp, err := msgServer.JoinPool(
				sdk.WrapSDKContext(ctx),
				types.NewMsgJoinPool(joinerAddr.String(), tc.initialPool.Id, tc.tokensIn, false),
			)

			require.NoError(t, err)
			require.Equal(t, types.MsgJoinPoolResponse{
				Pool:             &tc.expectedFinalPool,
				NumPoolSharesOut: tc.expectedNumSharesOut,
				RemainingCoins:   tc.expectedRemCoins,
			}, *resp)
			require.Equal(t, tc.expectedJoinerFinalFunds, app.BankKeeper.GetAllBalances(ctx, joinerAddr))
			expectedEvent := &types.EventPoolJoined{
				Address:       joinerAddr.String(),
				PoolId:        1,
				TokensIn:      tc.tokensIn,
				PoolSharesOut: resp.NumPoolSharesOut,
				RemCoins:      resp.RemainingCoins,
			}
			testutil.RequireHasTypedEvent(t, ctx, expectedEvent)
		})
	}
}

func TestMsgServerExitPool(t *testing.T) {
	const shareDenom = "nibiru/pool/1"
	tests := []struct {
		name                     string
		joinerInitialFunds       sdk.Coins
		initialPoolFunds         sdk.Coins
		initialPool              types.Pool
		poolSharesIn             sdk.Coin
		expectedTokensOut        sdk.Coins
		expectedJoinerFinalFunds sdk.Coins
		expectedFinalPool        types.Pool
	}{
		{
			name: "exit all pool shares",
			joinerInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
				sdk.NewInt64Coin(shareDenom, 100),
			),
			initialPoolFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			initialPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			poolSharesIn: sdk.NewInt64Coin(shareDenom, 100),
			expectedTokensOut: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 99),
				sdk.NewInt64Coin(common.DenomNUSD, 99),
			),
			expectedJoinerFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 199),
				sdk.NewInt64Coin(common.DenomNUSD, 199),
			),
			expectedFinalPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 1),
					sdk.NewInt64Coin(common.DenomNUSD, 1),
				),
				/*shares=*/ 0,
			),
		},
		{
			name: "exit half pool shares",
			joinerInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
				sdk.NewInt64Coin(shareDenom, 100),
			),
			initialPoolFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			initialPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			poolSharesIn: sdk.NewInt64Coin(shareDenom, 50),
			expectedTokensOut: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 49),
				sdk.NewInt64Coin(common.DenomNUSD, 49),
			),
			expectedJoinerFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 149),
				sdk.NewInt64Coin(common.DenomNUSD, 149),
				sdk.NewInt64Coin(shareDenom, 50),
			),
			expectedFinalPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 51),
					sdk.NewInt64Coin(common.DenomNUSD, 51),
				),
				/*shares=*/ 50,
			),
		},
		{
			name: "exit all pool shares - StablePool",
			joinerInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
				sdk.NewInt64Coin(shareDenom, 100),
			),
			initialPoolFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			initialPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			poolSharesIn: sdk.NewInt64Coin(shareDenom, 100),
			expectedTokensOut: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 99),
				sdk.NewInt64Coin(common.DenomNUSD, 99),
			),
			expectedJoinerFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 199),
				sdk.NewInt64Coin(common.DenomNUSD, 199),
			),
			expectedFinalPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 1),
					sdk.NewInt64Coin(common.DenomNUSD, 1),
				),
				/*shares=*/ 0,
			),
		},
		{
			name: "exit half pool shares - StablePool",
			joinerInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
				sdk.NewInt64Coin(shareDenom, 100),
			),
			initialPoolFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
				sdk.NewInt64Coin(common.DenomNUSD, 100),
			),
			initialPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			poolSharesIn: sdk.NewInt64Coin(shareDenom, 50),
			expectedTokensOut: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 49),
				sdk.NewInt64Coin(common.DenomNUSD, 49),
			),
			expectedJoinerFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 149),
				sdk.NewInt64Coin(common.DenomNUSD, 149),
				sdk.NewInt64Coin(shareDenom, 50),
			),
			expectedFinalPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 51),
					sdk.NewInt64Coin(common.DenomNUSD, 51),
				),
				/*shares=*/ 50,
			),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testapp.NewTestNibiruAppAndContext(true)

			poolAddr := testutil.AccAddress()
			tc.initialPool.Address = poolAddr.String()
			tc.expectedFinalPool.Address = poolAddr.String()
			app.DexKeeper.SetPool(ctx, tc.initialPool)

			sender := testutil.AccAddress()
			require.NoError(t, simapp.FundAccount(app.BankKeeper, ctx, sender, tc.joinerInitialFunds))
			require.NoError(t, simapp.FundAccount(app.BankKeeper, ctx, tc.initialPool.GetAddress(), tc.initialPoolFunds))

			msgServer := keeper.NewMsgServerImpl(app.DexKeeper)
			resp, err := msgServer.ExitPool(
				sdk.WrapSDKContext(ctx),
				types.NewMsgExitPool(sender.String(), tc.initialPool.Id, tc.poolSharesIn),
			)
			require.NoError(t, err)
			require.Equal(t,
				types.MsgExitPoolResponse{
					TokensOut: tc.expectedTokensOut,
				},
				*resp,
			)
			require.Equal(t,
				tc.expectedJoinerFinalFunds,
				app.BankKeeper.GetAllBalances(ctx, sender),
			)

			expectedEvent := &types.EventPoolExited{
				Address:      sender.String(),
				PoolId:       1,
				PoolSharesIn: tc.poolSharesIn,
				TokensOut:    resp.TokensOut,
			}

			testutil.RequireHasTypedEvent(t, ctx, expectedEvent)
		})
	}
}

func TestMsgServerSwapAssets(t *testing.T) {
	tests := []struct {
		name string

		// test setup
		userInitialFunds sdk.Coins
		initialPool      types.Pool
		tokenIn          sdk.Coin
		tokenOutDenom    string

		// expected results
		expectedError          error
		expectedTokenOut       sdk.Coin
		expectedUserFinalFunds sdk.Coins
		expectedFinalPool      types.Pool
	}{
		{
			name: "regular swap",
			userInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
			),
			initialPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			tokenIn:          sdk.NewInt64Coin("unibi", 100),
			tokenOutDenom:    common.DenomNUSD,
			expectedTokenOut: sdk.NewInt64Coin(common.DenomNUSD, 50),
			expectedUserFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin(common.DenomNUSD, 50),
			),
			expectedFinalPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 200),
					sdk.NewInt64Coin(common.DenomNUSD, 50),
				),
				/*shares=*/ 100,
			),
			expectedError: nil,
		},
		{
			name: "not enough user funds",
			userInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 1),
			),
			initialPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			tokenIn:       sdk.NewInt64Coin("unibi", 100),
			tokenOutDenom: common.DenomNUSD,
			expectedUserFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 1),
			),
			expectedFinalPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			expectedError: sdkerrors.ErrInsufficientFunds,
		},
		{
			name: "invalid token in denom",
			userInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("foo", 100),
			),
			initialPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			tokenIn:       sdk.NewInt64Coin("foo", 100),
			tokenOutDenom: common.DenomNUSD,
			expectedUserFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("foo", 100),
			),
			expectedFinalPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			expectedError: types.ErrTokenDenomNotFound,
		},
		{
			name: "invalid token out denom",
			userInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
			),
			initialPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			tokenIn:       sdk.NewInt64Coin("unibi", 100),
			tokenOutDenom: "foo",
			expectedUserFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
			),
			expectedFinalPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			expectedError: types.ErrTokenDenomNotFound,
		},
		{
			name: "same token in and token out denom",
			userInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
			),
			initialPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			tokenIn:       sdk.NewInt64Coin("unibi", 100),
			tokenOutDenom: "unibi",
			expectedUserFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
			),
			expectedFinalPool: mock.DexPool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			expectedError: types.ErrSameTokenDenom,
		},
		{
			name: "regular swap - StableSwap",
			userInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
			),
			initialPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			tokenIn:          sdk.NewInt64Coin("unibi", 100),
			tokenOutDenom:    common.DenomNUSD,
			expectedTokenOut: sdk.NewInt64Coin(common.DenomNUSD, 30),
			expectedUserFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin(common.DenomNUSD, 30),
			),
			expectedFinalPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 200),
					sdk.NewInt64Coin(common.DenomNUSD, 70),
				),
				/*shares=*/ 100,
			),
			expectedError: nil,
		},
		{
			name: "not enough user funds - StableSwap",
			userInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 1),
			),
			initialPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			tokenIn:       sdk.NewInt64Coin("unibi", 100),
			tokenOutDenom: common.DenomNUSD,
			expectedUserFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 1),
			),
			expectedFinalPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			expectedError: sdkerrors.ErrInsufficientFunds,
		},
		{
			name: "invalid token in denom - StableSwap",
			userInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("foo", 100),
			),
			initialPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			tokenIn:       sdk.NewInt64Coin("foo", 100),
			tokenOutDenom: common.DenomNUSD,
			expectedUserFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("foo", 100),
			),
			expectedFinalPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			expectedError: types.ErrTokenDenomNotFound,
		},
		{
			name: "invalid token out denom - StableSwap",
			userInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
			),
			initialPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			tokenIn:       sdk.NewInt64Coin("unibi", 100),
			tokenOutDenom: "foo",
			expectedUserFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
			),
			expectedFinalPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			expectedError: types.ErrTokenDenomNotFound,
		},
		{
			name: "same token in and token out denom - StableSwap",
			userInitialFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
			),
			initialPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			tokenIn:       sdk.NewInt64Coin("unibi", 100),
			tokenOutDenom: "unibi",
			expectedUserFinalFunds: sdk.NewCoins(
				sdk.NewInt64Coin("unibi", 100),
			),
			expectedFinalPool: mock.DexStablePool(
				/*poolId=*/ 1,
				/*assets=*/ sdk.NewCoins(
					sdk.NewInt64Coin("unibi", 100),
					sdk.NewInt64Coin(common.DenomNUSD, 100),
				),
				/*shares=*/ 100,
			),
			expectedError: types.ErrSameTokenDenom,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			app, ctx := testapp.NewTestNibiruAppAndContext(true)
			msgServer := keeper.NewMsgServerImpl(app.DexKeeper)

			// fund pool account
			poolAddr := testutil.AccAddress()
			tc.initialPool.Address = poolAddr.String()
			tc.expectedFinalPool.Address = poolAddr.String()
			require.NoError(t,
				simapp.FundAccount(
					app.BankKeeper,
					ctx,
					poolAddr,
					tc.initialPool.PoolBalances(),
				),
			)
			app.DexKeeper.SetPool(ctx, tc.initialPool)

			// fund user account
			sender := testutil.AccAddress()
			require.NoError(t, simapp.FundAccount(app.BankKeeper, ctx, sender, tc.userInitialFunds))

			// swap assets
			resp, err := msgServer.SwapAssets(
				sdk.WrapSDKContext(ctx),
				types.NewMsgSwapAssets(sender.String(), tc.initialPool.Id, tc.tokenIn, tc.tokenOutDenom),
			)

			if tc.expectedError != nil {
				require.ErrorIs(t, err, tc.expectedError)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					types.MsgSwapAssetsResponse{
						TokenOut: tc.expectedTokenOut,
					},
					*resp,
				)

				// check events
				testutil.RequireHasTypedEvent(t, ctx, &types.EventAssetsSwapped{
					Address:  sender.String(),
					PoolId:   1,
					TokenIn:  tc.tokenIn,
					TokenOut: tc.expectedTokenOut,
				})
			}

			// check user's final funds
			require.Equal(t,
				tc.expectedUserFinalFunds,
				app.BankKeeper.GetAllBalances(ctx, sender),
			)

			// check final pool state
			finalPool, err := app.DexKeeper.FetchPool(ctx, tc.initialPool.Id)
			require.NoError(t, err)
			require.Equal(t, tc.expectedFinalPool, finalPool)
		})
	}
}
