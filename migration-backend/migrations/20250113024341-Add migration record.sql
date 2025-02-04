
-- +migrate Up

CREATE TABLE migration_record
(
  id SERIAL PRIMARY KEY, 
  cosmos_tx_hash text NOT NULL,
  eth_tx_hash text NULL,
  cosmos_address text NULL,
  eth_address text NULL,
  UNIQUE (cosmos_tx_hash)
);

-- +migrate Down

DROP TABLE migration_record;
