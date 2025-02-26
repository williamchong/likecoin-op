
-- +migrate Up

CREATE TABLE likenft_migration_action_mint_nft
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT now(),
  evm_class_id text NOT NULL,
  cosmos_nft_id text NOT NULL,
  initial_batch_mint_owner text NOT NULL,
  evm_owner text NOT NULL,
  status text NOT NULL,
  evm_tx_hash text NULL,
  failed_reason text NULL,
  UNIQUE (evm_class_id, cosmos_nft_id)
);

CREATE TABLE likenft_migration_action_new_class
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT now(),
  cosmos_class_id text NOT NULL,
  initial_owner text NOT NULL,
  initial_minter text NOT NULL,
  initial_updater text NOT NULL,
  status text NOT NULL,
  evm_class_id text NULL,
  evm_tx_hash text NULL,
  failed_reason text NULL,
  UNIQUE (cosmos_class_id)
);

CREATE TABLE likenft_migration_action_transfer_class
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT now(),
  evm_class_id text NOT NULL,
  cosmos_owner text NOT NULL,
  evm_owner text NOT NULL,
  status text NOT NULL,
  evm_tx_hash text NULL,
  failed_reason text NULL,
  UNIQUE (evm_class_id, evm_owner)
);

-- +migrate Down

DROP TABLE likenft_migration_action_transfer_class;
DROP TABLE likenft_migration_action_new_class;
DROP TABLE likenft_migration_action_mint_nft;
