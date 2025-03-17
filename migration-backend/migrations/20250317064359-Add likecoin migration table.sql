
-- +migrate Up
CREATE TABLE likecoin_migration
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT now(),
  user_cosmos_address TEXT NOT NULL,
  burning_cosmos_address TEXT NOT NULL,
  minting_eth_address VARCHAR(42) NOT NULL,
  user_eth_address VARCHAR(42) NOT NULL,
  amount TEXT NOT NULL,
  evm_signature VARCHAR(132) NOT NULL,
  evm_signature_message TEXT NOT NULL,
  status TEXT NOT NULL,
  cosmos_tx_hash TEXT NULL,
  evm_tx_hash VARCHAR(66) NULL,
  failed_reason TEXT NULL,
  UNIQUE NULLS NOT DISTINCT (cosmos_tx_hash)
);

-- +migrate Down
DROP TABLE likecoin_migration;
