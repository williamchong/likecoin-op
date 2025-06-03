
-- +migrate Up
ALTER TABLE likenft_asset_snapshot_class
    ADD estimated_migration_duration_needed NUMERIC NULL;
ALTER TABLE likenft_asset_snapshot_nft
    ADD estimated_migration_duration_needed NUMERIC NULL;
ALTER TABLE likenft_asset_snapshot
    ADD estimated_migration_duration_needed NUMERIC NULL;

ALTER TABLE likenft_asset_migration_class
    ADD estimated_duration_needed NUMERIC NULL;
ALTER TABLE likenft_asset_migration_nft
    ADD estimated_duration_needed NUMERIC NULL;
ALTER TABLE likenft_asset_migration
    ADD estimated_finished_time TIMESTAMP NULL;

UPDATE likenft_asset_snapshot_class
    SET estimated_migration_duration_needed = 0;
UPDATE likenft_asset_snapshot_nft
    SET estimated_migration_duration_needed = 0;

UPDATE likenft_asset_migration_class
    SET estimated_duration_needed = 0;
UPDATE likenft_asset_migration_nft
    SET estimated_duration_needed = 0;
UPDATE likenft_asset_migration
    SET estimated_finished_time = now();

ALTER TABLE likenft_asset_snapshot_class
    ALTER estimated_migration_duration_needed SET NOT NULL;
ALTER TABLE likenft_asset_snapshot_nft
    ALTER estimated_migration_duration_needed SET NOT NULL;

ALTER TABLE likenft_asset_migration_class
    ALTER estimated_duration_needed SET NOT NULL;
ALTER TABLE likenft_asset_migration_nft
    ALTER estimated_duration_needed SET NOT NULL;
ALTER TABLE likenft_asset_migration
    ALTER estimated_finished_time SET NOT NULL;

-- +migrate Down
ALTER TABLE likenft_asset_migration
    DROP estimated_finished_time;
ALTER TABLE likenft_asset_migration_nft
    DROP estimated_duration_needed;
ALTER TABLE likenft_asset_migration_class
    DROP estimated_duration_needed;

ALTER TABLE likenft_asset_snapshot
    DROP estimated_migration_duration_needed;
ALTER TABLE likenft_asset_snapshot_nft
    DROP estimated_migration_duration_needed;
ALTER TABLE likenft_asset_snapshot_class
    DROP estimated_migration_duration_needed;
