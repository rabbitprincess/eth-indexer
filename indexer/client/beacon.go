package client

import (
	"context"
)

func (c *Client) GetBeaconChainHead() ([]byte, error) {
	head, err := c.beacon.BeaconBlockHeader(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return head.Data.Root[:], nil
}
