
-- +migrate Up
CREATE TABLE likenft_asset_migration
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT now(),
  likenft_asset_snapshot_id INT NOT NULL,
  cosmos_address text NOT NULL,
  eth_address text NOT NULL,
  status text NOT NULL,
  failed_reason text NULL,
  CONSTRAINT fk_likenft_asset_snapshot_id
    FOREIGN KEY(likenft_asset_snapshot_id)
      REFERENCES likenft_asset_snapshot(id),
  UNIQUE (cosmos_address),
  UNIQUE (likenft_asset_snapshot_id)
);

CREATE TABLE likenft_asset_migration_class
(
  id SERIAL PRIMARY KEY,
  likenft_asset_migration_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  cosmos_class_id text NOT NULL,
  name text NOT NULL,
  image text NOT NULL,
  status text NOT NULL,
  enqueue_time TIMESTAMP NULL,
  finish_time TIMESTAMP NULL,
  evm_tx_hash text NULL,
  failed_reason text NULL,
  CONSTRAINT fk_likenft_asset_migration
    FOREIGN KEY(likenft_asset_migration_id)
      REFERENCES likenft_asset_migration(id),
  UNIQUE (likenft_asset_migration_id, cosmos_class_id)
);

CREATE TABLE likenft_asset_migration_nft
(
  id SERIAL PRIMARY KEY,
  likenft_asset_migration_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  cosmos_class_id text NOT NULL,
  cosmos_nft_id text NOT NULL,
  name text NOT NULL,
  image text NOT NULL,
  status text NOT NULL,
  enqueue_time TIMESTAMP NULL,
  finish_time TIMESTAMP NULL,
  evm_tx_hash text NULL,
  failed_reason text NULL,
  CONSTRAINT fk_likenft_asset_migration
    FOREIGN KEY(likenft_asset_migration_id)
      REFERENCES likenft_asset_migration(id),
  UNIQUE (likenft_asset_migration_id, cosmos_class_id, cosmos_nft_id)
);

-- +migrate Down

DROP TABLE likenft_asset_migration_nft;
DROP TABLE likenft_asset_migration_class;
DROP TABLE likenft_asset_migration;
