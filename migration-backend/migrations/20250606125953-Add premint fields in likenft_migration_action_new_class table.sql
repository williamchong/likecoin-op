
-- +migrate Up
ALTER TABLE likenft_migration_action_new_class
    ADD should_premint_all_nfts BOOLEAN NULL;
ALTER TABLE likenft_migration_action_new_class
    ADD initial_batch_mint_owner VARCHAR(42) NULL;
ALTER TABLE likenft_migration_action_new_class
    ADD number_of_cosmos_nfts_found NUMERIC NULL;

UPDATE likenft_migration_action_new_class
    SET should_premint_all_nfts = false;
UPDATE likenft_migration_action_new_class
    SET initial_batch_mint_owner = '0x';

ALTER TABLE likenft_migration_action_new_class
    ALTER should_premint_all_nfts SET NOT NULL;
ALTER TABLE likenft_migration_action_new_class
    ALTER initial_batch_mint_owner SET NOT NULL;

-- +migrate Down
ALTER TABLE likenft_migration_action_new_class
    DROP should_premint_all_nfts;
ALTER TABLE likenft_migration_action_new_class
    DROP initial_batch_mint_owner;
ALTER TABLE likenft_migration_action_new_class
    DROP number_of_cosmos_nfts_found;
