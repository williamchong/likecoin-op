
-- +migrate Up
CREATE TABLE likenft_migration_action_batch_mint_nfts
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT now(),
  evm_class_id text NOT NULL,
  current_supply numeric NOT NULL,
  expected_supply numeric NOT NULL,
  batch_mint_amount numeric NOT NULL,
  initial_batch_mint_owner text NOT NULL,
  status text NOT NULL,
  from_id numeric NULL,
  to_id numeric NULL,
  evm_tx_hash text NULL,
  failed_reason text NULL
);


-- +migrate Down
DROP TABLE likenft_migration_action_batch_mint_nfts;
