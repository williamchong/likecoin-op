
-- +migrate Up
ALTER TABLE likenft_migration_action_new_class
    ADD default_royalty_fraction NUMERIC NULL;
UPDATE likenft_migration_action_new_class
    SET default_royalty_fraction = 0;
ALTER TABLE likenft_migration_action_new_class
    ALTER default_royalty_fraction SET NOT NULL;

-- +migrate Down
ALTER TABLE likenft_migration_action_new_class
    DROP default_royalty_fraction;
