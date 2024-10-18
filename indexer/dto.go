package indexer

import (
	"context"

	"github.com/rabbitprincess/eth-indexer/indexer/client"
	"github.com/rabbitprincess/eth-indexer/indexer/db"
	"github.com/rabbitprincess/eth-indexer/indexer/schema"
	"github.com/rs/zerolog/log"
)

type DTO struct {
	blockNumber uint64

	accountBalance map[string]*schema.AccountBalance
	balanceChange  []*schema.BalanceCHangeHistory
}

func (d *DTO) Init(blockNumber uint64) {
	d.blockNumber = blockNumber
	d.accountBalance = make(map[string]*schema.AccountBalance)
	if d.balanceChange == nil {
		d.balanceChange = make([]*schema.BalanceCHangeHistory, 0, 1024)
	} else {
		d.balanceChange = d.balanceChange[:0]
	}
}

func (d *DTO) Commit(db db.DbController) error {
	bulk := db.InsertBulk(schema.TableAccountBalance)
	for _, balance := range d.accountBalance {
		bulk.Add(balance)
	}
	err := bulk.Commit()
	if err != nil {
		return err
	}

	bulk = db.InsertBulk(schema.TableBalanceChangeHistory)
	for _, change := range d.balanceChange {
		bulk.Add(change)
	}
	err = bulk.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *DTO) VerifyBalance(ctx context.Context, blockNumber uint64, client *client.Client) error {
	for _, a := range d.accountBalance {
		verifyBalance, err := client.GetAccountBalance(ctx, a.Account, blockNumber)
		if err != nil {
			return err
		}
		if verifyBalance.String() != a.Balance {
			log.Error().Uint64("blockNumber", blockNumber).Str("address", a.Account).Str("balance", a.Balance).Str("verifyBalance", verifyBalance.String()).Msg("balance mismatch")
			return err
		}
	}
	return nil
}

func (d *DTO) AddAccountBalance(blockNumber uint64, blockTimeStamp uint64, account string, balance string) {
	d.accountBalance[account] = &schema.AccountBalance{
		Account:        account,
		BlockNumber:    blockNumber,
		BlockTimestamp: blockTimeStamp,
		Balance:        balance,
	}
}

func (d *DTO) GetAccountBalance(account string, dbController db.DbController, client *client.Client) (*schema.AccountBalance, error) {
	// get from cache
	if accBalance, exist := d.accountBalance[account]; exist {
		return accBalance, nil
	}
	// get from db
	doc, err := dbController.SelectOne(db.QueryParams{
		IndexName: schema.TableAccountBalance,
		StringMatch: &db.StringMatchQuery{
			Field: "account",
			Value: account,
		},
	}, func() schema.DocType {
		balance := new(schema.AccountBalance)
		balance.BaseEsType = new(schema.BaseEsType)
		return balance
	})
	if err != nil {
		return nil, err
	}
	if doc != nil {
		d.accountBalance[account] = doc.(*schema.AccountBalance)
		return doc.(*schema.AccountBalance), nil
	}

	// get from server
	balance, err := client.GetAccountBalance(context.Background(), account, d.blockNumber)
	if err != nil {
		return nil, err
	}
	return &schema.AccountBalance{
		Account:        account,
		BlockNumber:    d.blockNumber,
		BlockTimestamp: 0,
		Balance:        balance.String(),
	}, nil
}

func (d *DTO) AddBalanceChange(blockNumber uint64, blockTimeStamp uint64, account string, changeType schema.BalanceChange, balanceBefore, balanceAfter, balanceChange string, txid string, txIndex uint64) {
	d.balanceChange = append(d.balanceChange, &schema.BalanceCHangeHistory{
		Account:        account,
		BlockNumber:    blockNumber,
		BlockTimestamp: blockTimeStamp,
		ChangeType:     uint64(changeType),
		BalanceBefore:  balanceBefore,
		BalanceAfter:   balanceAfter,
		BalanceChange:  balanceChange,
		Txid:           txid,
		TxIndex:        txIndex,
	})
}
