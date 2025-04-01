package model

type MigrateUserEVMWalletMemoAction string

const (
	MigrateUserEVMWalletMemoActionMigrate MigrateUserEVMWalletMemoAction = "migrate"
)

type MigrateUserEVMWalletMemo struct {
	Action       MigrateUserEVMWalletMemoAction `json:"action"`
	CosmosWallet string                         `json:"cosmosWallet"`
	LikerWallet  string                         `json:"likeWallet"`
	Ts           uint64                         `json:"ts"`
	EvmWallet    string                         `json:"evm_wallet"`
}
