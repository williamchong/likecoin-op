
-- +migrate Up

CREATE TABLE likenft_asset_snapshot
(
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT now(),
  cosmos_address text NOT NULL,
  block_height text NULL,
  block_time TIMESTAMP NULL,
  status text NOT NULL,
  failed_reason text NULL,
  UNIQUE (cosmos_address, block_height)
);

CREATE TABLE likenft_asset_snapshot_class
(
  id SERIAL PRIMARY KEY,
  likenft_asset_snapshot_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  cosmos_class_id text NOT NULL,
  name text NOT NULL,
  image text NOT NULL,
  CONSTRAINT fk_likenft_asset_snapshot
    FOREIGN KEY(likenft_asset_snapshot_id)
      REFERENCES likenft_asset_snapshot(id),
  UNIQUE (likenft_asset_snapshot_id, cosmos_class_id)
);

CREATE TABLE likenft_asset_snapshot_nft
(
  id SERIAL PRIMARY KEY,
  likenft_asset_snapshot_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  cosmos_class_id text NOT NULL,
  cosmos_nft_id text NOT NULL,
  name text NOT NULL,
  image text NOT NULL,
  CONSTRAINT fk_likenft_asset_snapshot
    FOREIGN KEY(likenft_asset_snapshot_id)
      REFERENCES likenft_asset_snapshot(id),
  UNIQUE (likenft_asset_snapshot_id, cosmos_class_id, cosmos_nft_id)
);

-- +migrate Down

DROP TABLE likenft_asset_snapshot_nft;
DROP TABLE likenft_asset_snapshot_class;
DROP TABLE likenft_asset_snapshot;
