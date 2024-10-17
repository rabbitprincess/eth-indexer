package client

import (
	"context"
	"fmt"

	beacon "github.com/attestantio/go-eth2-client/http"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Client struct {
	logger    *zerolog.Logger
	execution *ethclient.Client
	beacon    *beacon.Service
}

func NewClient(ctx context.Context, logger *zerolog.Logger, executionURL, beaconURL string) (*Client, error) {
	if logger == nil {
		logger = &log.Logger
	}

	execClient, err := ethclient.Dial(executionURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to execution client: %w", err)
	}

	c, err := beacon.New(ctx, beacon.WithAddress(beaconURL))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to beacon client: %w", err)
	}

	return &Client{
		logger:    logger,
		execution: execClient,
		beacon:    c.(*beacon.Service),
	}, nil
}
