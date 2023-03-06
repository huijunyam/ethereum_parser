package models

type Transaction struct {
	Type                 string   `json:"type"`
	BlockHash            string   `json:"blockHash"`
	BlockNumber          string   `json:"blockNumber"`
	From                 string   `json:"from"`
	Gas                  string   `json:"gas"`
	Hash                 string   `json:"hash"`
	Input                string   `json:"input"`
	Nonce                string   `json:"nonce"`
	To                   string   `json:"to"`
	TransactionIndex     string   `json:"transactionIndex"`
	Value                string   `json:"value"`
	V                    string   `json:"v"`
	R                    string   `json:"r"`
	S                    string   `json:"s"`
	GasPrice             string   `json:"gasPrice"`
	MaxFeePerGas         string   `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
	ChainId              string   `json:"chainId"`
	AccessList           []Access `json:"accessList"`
}

type Access struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}
