package indexer

import (
	"context"
	"math/big"

	"github.com/rabbitprincess/eth-indexer/indexer/client"
	"github.com/rabbitprincess/eth-indexer/indexer/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type RunConfig struct {
	NetworkID     *big.Int
	NetworkName   string
	VerifyBalance bool
	From          uint64
	To            uint64
}

type Indexer struct {
	cfg    *RunConfig
	logger *zerolog.Logger

	client *client.Client
	db     db.DbController

	dto *DTO
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

	dto := &DTO{}
	dto.Init(0)

	return &Indexer{
		logger: logger,
		client: c,
		db:     d,
		dto:    dto,
	}, nil
}

func (i *Indexer) Run(ctx context.Context, cfg *RunConfig) error {
	var err error
	i.cfg = cfg

	// start indexing
	err = i.RunPreAlloc(ctx)
	if err != nil {
		return err
	}

	err = i.RunTraceBlock(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (i *Indexer) Stop() {
	// stop indexing
}
