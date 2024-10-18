package schema

type BalanceChange uint64

const (
	// ethereum balance change type
	PreAlloc BalanceChange = iota
	Transfer
	ContractCall
	FeeDeduction
	MiningReward
	StakingDeposit
	StakingReward
	StakingSlashing
	StakingWithdrawal
)
