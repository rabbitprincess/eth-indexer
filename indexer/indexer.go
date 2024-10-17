package indexer

import (
	"context"

	"github.com/rabbitprincess/eth-indexer/indexer/client"
	"github.com/rabbitprincess/eth-indexer/indexer/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Indexer struct {
	logger *zerolog.Logger

	client *client.Client
	db     db.DbController
}

func NewIndexer(ctx context.Context, logger *zerolog.Logger, executionURL, beaconURL, esURL string) (*Indexer, error) {
	if logger == nil {
		logger = &log.Logger
	}

	// init client
	c, err := client.NewClient(ctx, logger, executionURL, beaconURL)
	if err != nil {
		return nil, err
	}

	// init db
	d, err := db.NewElasticsearchDbController(ctx, logger, esURL)
	if err != nil {
		return nil, err
	}

	return &Indexer{
		logger: logger,
		client: c,
		db:     d,
	}, nil
}
