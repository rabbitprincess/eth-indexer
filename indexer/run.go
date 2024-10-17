package indexer

import (
	"context"
	"embed"
	"encoding/json"
	"math"

	"github.com/ethereum/go-ethereum/core/types"
)

func (i *Indexer) RunPreAlloc(ctx context.Context) error {
	var filename string = "allocs/" + i.cfg.NetworkName + ".json"
	ga, err := readPrealloc(filename)
	if err != nil {
		i.logger.Error().Err(err).Msg("failed to read prealloc")
		return nil
	}

	for address, account := range ga {
		// save to db
		_ = address
		_ = account
	}

	return nil
}

//go:embed allocs
var allocs embed.FS

func readPrealloc(filename string) (types.GenesisAlloc, error) {
	f, err := allocs.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	ga := make(types.GenesisAlloc)
	err = decoder.Decode(&ga)
	if err != nil {
		return nil, err
	}
	return ga, nil
}

func (i *Indexer) RunTraceBlock(ctx context.Context) error {
	if i.cfg.To <= 0 {
		i.cfg.To = math.MaxInt
	}

	for blockNumber := i.cfg.From; blockNumber <= i.cfg.To; blockNumber++ {

	}

	return nil
}
