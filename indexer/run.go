package indexer

import (
	"context"
	"embed"
	"encoding/json"
	"math"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rabbitprincess/eth-indexer/indexer/schema"
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
		addr := address.String()
		bal := account.Balance.String()

		i.dto.AddAccountBalance(0, 0, addr, bal)
		i.dto.AddBalanceChange(0, 0, address.String(), schema.PreAlloc, "0", bal, bal, "", 0)
	}

	if i.cfg.VerifyBalance {
		err = i.dto.VerifyBalance(ctx, 0, i.client)
		if err != nil {
			return err
		}
	}

	err = i.dto.Commit(i.db)
	if err != nil {
		return err
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
	if i.cfg.To == 0 {
		i.cfg.To = math.MaxInt
	}

	blockNumber := i.cfg.From
	for {
		i.logger.Info().Uint64("blockNumber", blockNumber).Msg("Trace block")
		if blockNumber > i.cfg.To {
			i.logger.Info().Uint64("to", blockNumber).Msg("Trace block finished")
			break
		}

		// wait new block
		latestBlock, err := i.client.GetLatestBlockNumber(ctx)
		if err != nil {
			return err
		} else if blockNumber > latestBlock {
			i.logger.Info().Uint64("latestBlock", latestBlock).Uint64("blockNumber", blockNumber).Msg("waiting for new block...")
			time.Sleep(time.Second * 10)
			continue
		}

		// init dto
		i.dto.Init(blockNumber)

		// trace balance
		trace, err := i.client.TraceBlock(ctx, blockNumber)
		if err != nil {
			return err
		}
		_ = trace

		// verify balance
		if i.cfg.VerifyBalance {
			err := i.dto.VerifyBalance(ctx, blockNumber, i.client)
			if err != nil {
				return err
			}
		}

		err = i.dto.Commit(i.db)
		if err != nil {
			return err
		}

		blockNumber++
	}

	return nil
}
