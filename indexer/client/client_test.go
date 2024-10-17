package client

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	rpcUrl    = "https://autumn-magical-orb.ethereum-holesky.quiknode.pro/9982c719a7beb2e0ba7b076175fbb12e07c586e8"
	beaconUrl = ""
)

func TestGetLatestBlock(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(ctx, nil, rpcUrl, beaconUrl)
	require.NoError(t, err)
	blockNumber, err := client.GetLatestBlockNumber(ctx)
	require.NoError(t, err)

	fmt.Println(blockNumber)
}

func TestTraceTransaction(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(ctx, nil, rpcUrl, beaconUrl)
	require.NoError(t, err)

	res, err := client.TraceTransaction(ctx, "0xea758cffadf4821d8b1fecd6360b6ad8ae88597aae8b8df3b0d79a7df2564945")
	require.NoError(t, err)

	resJson, err := json.MarshalIndent(res, "", "\t")
	require.NoError(t, err)
	fmt.Println(string(resJson))
}

func TestTraceBlock(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(ctx, nil, rpcUrl, beaconUrl)
	require.NoError(t, err)

	res, err := client.TraceBlock(ctx, 2550745)
	require.NoError(t, err)

	resJson, err := json.MarshalIndent(res, "", "\t")
	require.NoError(t, err)
	fmt.Println(string(resJson))
}
