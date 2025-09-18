-- Modify "accounts" table
ALTER TABLE "accounts" ALTER COLUMN "staked_amount" SET DEFAULT 0, ALTER COLUMN "pending_reward_amount" SET DEFAULT 0, ALTER COLUMN "claimed_reward_amount" SET DEFAULT 0;
-- Modify "nft_classes" table
ALTER TABLE "nft_classes" ALTER COLUMN "staked_amount" SET DEFAULT 0;
-- Modify "staking_events" table
ALTER TABLE "staking_events" ALTER COLUMN "staked_amount_added" SET DEFAULT 0, ALTER COLUMN "staked_amount_removed" SET DEFAULT 0, DROP COLUMN "reward_amount_added", DROP COLUMN "reward_amount_removed", DROP COLUMN "account_id", DROP COLUMN "nft_class_id", ADD COLUMN "transaction_hash" character varying NOT NULL, ADD COLUMN "transaction_index" bigint NOT NULL, ADD COLUMN "block_number" numeric NOT NULL, ADD COLUMN "log_index" bigint NOT NULL, ADD COLUMN "nft_class_address" character varying NOT NULL, ADD COLUMN "account_evm_address" character varying NOT NULL, ADD COLUMN "pending_reward_amount_added" numeric NOT NULL DEFAULT 0, ADD COLUMN "pending_reward_amount_removed" numeric NOT NULL DEFAULT 0, ADD COLUMN "claimed_reward_amount_added" numeric NOT NULL DEFAULT 0, ADD COLUMN "claimed_reward_amount_removed" numeric NOT NULL DEFAULT 0;
-- Modify "stakings" table
ALTER TABLE "stakings" ALTER COLUMN "pool_share" SET DEFAULT '0', ALTER COLUMN "staked_amount" SET DEFAULT 0, ALTER COLUMN "pending_reward_amount" SET DEFAULT 0, ALTER COLUMN "claimed_reward_amount" SET DEFAULT 0;
