package indexer

type balanceChange uint64

const (
	// ethereum balance change type
	PreAlloc balanceChange = iota
	Transfer
	ContractCall
	FeeDeduction
	MiningReward
	StakingDeposit
	StakingReward
	StrkingSlashing
	StakingWithdrawal
)
