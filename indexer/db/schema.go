package db

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
	BaseEsType
	Account        string `json:"account" db:"account"`
	BlockNumber    uint64 `json:"block_number" db:"block_number"`
	BlockTimestamp string `json:"block_timestamp" db:"block_timestamp"`
	Balance        uint64 `json:"balance" db:"balance"`
}

type AccountBalanceHistory struct {
	BaseEsType
	Account        string `json:"account" db:"account"`
	BlockNumber    uint64 `json:"block_number" db:"block_number"`
	BlockTimestamp string `json:"block_timestamp" db:"block_timestamp"`
	ChangeType     uint64 `json:"change_type" db:"change_type"`
	ChangeBalance  uint64 `json:"change_balance" db:"change_balance"`
	Txid           uint64 `json:"txid" db:"txid"`
}

var EsSchema map[string]string

func init() {
	EsSchema = map[string]string{}

	EsSchema["account_balance"] = `{
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
				"type": "long"
			}
		}
	}
}`

	EsSchema["account_balance_history"] = `{
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
				"type": "long"
			},
			"balance_after": {
				"type": "long"
			},
			"balance_change": {
				"type": "long"
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
