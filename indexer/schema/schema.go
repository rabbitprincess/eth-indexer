package schema

// DocType is an interface for structs to be used as database documents
type DocType interface {
	GetID() string
	SetID(string)
}

// BaseEsType implements DocType and contains the document's id
type BaseEsType struct {
	Id string `json:"-" db:"id"`
}

// GetID returns the document's id
func (m *BaseEsType) GetID() string {
	return m.Id
}

// SetID sets the document's id
func (m *BaseEsType) SetID(id string) {
	m.Id = id
}

type AccountBalance struct {
	*BaseEsType
	Account        string `json:"account" db:"account"`
	BlockNumber    uint64 `json:"block_number" db:"block_number"`
	BlockTimestamp uint64 `json:"block_timestamp" db:"block_timestamp"`
	Balance        string `json:"balance" db:"balance"`
}

type BalanceCHangeHistory struct {
	*BaseEsType
	Account        string `json:"account" db:"account"`
	BlockNumber    uint64 `json:"block_number" db:"block_number"`
	BlockTimestamp uint64 `json:"block_timestamp" db:"block_timestamp"`
	ChangeType     uint64 `json:"change_type" db:"change_type"`
	BalanceBefore  string `json:"balance_before" db:"balance_before"`
	BalanceAfter   string `json:"balance_after" db:"balance_after"`
	BalanceChange  string `json:"change_balance" db:"balance_change"`
	Txid           string `json:"txid" db:"txid"`
	TxIndex        uint64 `json:"txindex" db:"txindex"`
}

var (
	EsSchema                  map[string]string
	TableAccountBalance       = "account_balance"
	TableBalanceChangeHistory = "balance_change_history"
)

func init() {
	EsSchema = map[string]string{}

	EsSchema[TableAccountBalance] = `{
	"settings": {
		"number_of_shards": 10,
		"number_of_replicas": 1,
		"index.max_result_window": 100000
	},
	"mappings": {
		"properties": {
			"id": {
				"type": "keyword"
			},
			"account": {
				"type": "keyword"
			},
			"block_number": {
				"type": "long"
			},
			"block_timestamp": {
				"type": "date"
			},
			"balance": {
				"type": "keyword"
			}
		}
	}
}`

	EsSchema[TableBalanceChangeHistory] = `{
	"settings": {
		"number_of_shards": 10,
		"number_of_replicas": 1,
		"index.max_result_window": 100000
	},
	"mappings": {
		"properties": {
			"id": {
				"type": "keyword"
			},
			"account": {
				"type": "keyword"
			},
			"block_number": {
				"type": "long"
			},
			"block_timestamp": {
				"type": "date"
			},
			"change_type": {
				"type": "long"
			},
			"balance_before": {
				"type": "keyword"
			},
			"balance_after": {
				"type": "keyword"
			},
			"balance_change": {
				"type": "keyword"
			},
			"txid": {
				"type": "keyword"
			},
			"txindex": {
				"type": "long"
			}
		}
	}
}`

}
