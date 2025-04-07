
-- +migrate Up
ALTER TABLE likecoin_migration DROP CONSTRAINT "likecoin_migration_cosmos_tx_hash_key";
ALTER TABLE likecoin_migration ADD CONSTRAINT "likecoin_migration_cosmos_tx_hash_key" UNIQUE NULLS DISTINCT (cosmos_tx_hash);

-- +migrate Down
ALTER TABLE likecoin_migration DROP CONSTRAINT "likecoin_migration_cosmos_tx_hash_key";
ALTER TABLE likecoin_migration ADD CONSTRAINT "likecoin_migration_cosmos_tx_hash_key" UNIQUE NULLS NOT DISTINCT (cosmos_tx_hash);
