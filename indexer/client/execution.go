package client

import (
	"context"
)

func (c *Client) GetLatestBlockNumber(ctx context.Context) (uint64, error) {
	return c.execution.BlockNumber(ctx)
}
