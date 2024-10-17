package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (c *Client) GetLatestBlockNumber(ctx context.Context) (uint64, error) {
	return c.execution.BlockNumber(ctx)
}

func (c *Client) GetAccountBalance(ctx context.Context, account string, blockNumber uint64) (uint64, error) {
	acc := common.HexToAddress(account)
	num := big.NewInt(int64(blockNumber))

	balance, err := c.execution.BalanceAt(ctx, acc, num)
	if err != nil {
		return 0, err
	}
	return balance.Uint64(), nil
}

func (c *Client) TraceTransaction(ctx context.Context, txHash string) (interface{}, error) {
	var result interface{}
	err := c.execution.Client().CallContext(ctx, &result, "trace_transaction", txHash, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) TraceBlock(ctx context.Context, blockNumber uint64) (interface{}, error) {
	var result interface{}
	err := c.execution.Client().CallContext(ctx, &result, "trace_block", blockNumber, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}
