package client

import (
	"context"
)

func (c *Client) GetLatestBlockNumber(ctx context.Context) (uint64, error) {
	return c.execution.BlockNumber(ctx)
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
