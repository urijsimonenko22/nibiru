package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/NibiruChain/nibiru/x/common"
)

func TestSnapshotValidate(t *testing.T) {
	tests := []struct {
		name      string
		snapshot  ReserveSnapshot
		shouldErr bool
	}{
		{
			name: "happy path",
			snapshot: ReserveSnapshot{
				Pair:              common.Pair_ETH_NUSD,
				BaseAssetReserve:  sdk.OneDec(),
				QuoteAssetReserve: sdk.OneDec(),
				TimestampMs:       time.Now().UnixMilli(),
			},
			shouldErr: false,
		},
		{
			name: "invalid pair",
			snapshot: ReserveSnapshot{
				Pair:              common.AssetPair{Token0: "$invalid", Token1: "valid"},
				BaseAssetReserve:  sdk.OneDec(),
				QuoteAssetReserve: sdk.OneDec(),
				TimestampMs:       time.Now().UnixMilli(),
			},
			shouldErr: true,
		},
		{
			name: "base asset negative",
			snapshot: ReserveSnapshot{
				Pair:              common.Pair_ETH_NUSD,
				BaseAssetReserve:  sdk.NewDec(-1),
				QuoteAssetReserve: sdk.OneDec(),
				TimestampMs:       time.Now().UnixMilli(),
			},
			shouldErr: true,
		},
		{
			name: "quote asset negative",
			snapshot: ReserveSnapshot{
				Pair:              common.Pair_ETH_NUSD,
				BaseAssetReserve:  sdk.ZeroDec(),
				QuoteAssetReserve: sdk.NewDec(-1),
				TimestampMs:       time.Now().UnixMilli(),
			},
			shouldErr: true,
		},
		{
			name: "timestamp lower than smallest UTC ('0001-01-01 00:00:00 +0000 UTC')",
			// see time.UnixMilli(-62135596800000).UTC())
			snapshot: ReserveSnapshot{
				Pair:              common.Pair_ETH_NUSD,
				BaseAssetReserve:  sdk.ZeroDec(),
				QuoteAssetReserve: sdk.ZeroDec(),
				TimestampMs:       -62135596800000 - 1,
			},
			shouldErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.snapshot.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.Nil(t, err)
			}
		})
	}
}
