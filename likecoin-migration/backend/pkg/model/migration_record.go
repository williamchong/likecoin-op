package model

type MigrationRecord struct {
	CosmosTxHash  string
	EthTxHash     string
	CosmosAddress string
	EthAddress    string
}
