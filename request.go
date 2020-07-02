package zabbix

type Request struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Auth    string      `json:"auth"`
	ID      uint64      `json:"id,omitempty"`
}
