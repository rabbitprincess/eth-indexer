package client

type TraceBlock struct {
	Action              Action `json:"action"`
	BlockHash           string `json:"blockHash"`
	BlockNumber         int    `json:"blockNumber"`
	Result              Result `json:"result"`
	Subtraces           int    `json:"subtraces"`
	TraceAddress        []any  `json:"traceAddress"`
	TransactionHash     string `json:"transactionHash"`
	TransactionPosition int    `json:"transactionPosition"`
	Type                string `json:"type"`
}

type Action struct {
	CallType string `json:"callType"`
	From     string `json:"from"`
	Gas      string `json:"gas"`
	Input    string `json:"input"`
	To       string `json:"to"`
	Value    string `json:"value"`
}

type Result struct {
	GasUsed string `json:"gasUsed"`
	Output  string `json:"output"`
}
