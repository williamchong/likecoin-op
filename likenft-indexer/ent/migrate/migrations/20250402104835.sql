-- Modify "evm_event_processed_block_heights" table
ALTER TABLE "evm_event_processed_block_heights" ADD CONSTRAINT "uint64_block_height_check" CHECK (block_height >= (0)::numeric), ALTER COLUMN "block_height" TYPE numeric;
-- Modify "evm_events" table
ALTER TABLE "evm_events" ALTER COLUMN "block_number" TYPE numeric;
ALTER TABLE "evm_events" ADD CONSTRAINT "uint64_block_number_check" CHECK (block_number >= (0)::numeric);
-- Modify "nft_classes" table
ALTER TABLE "nft_classes" ALTER COLUMN "total_supply" TYPE numeric;
ALTER TABLE "nft_classes" ADD CONSTRAINT "uint64_total_supply_check" CHECK (total_supply >= (0)::numeric);

ALTER TABLE "nft_classes" ADD COLUMN "deployed_block_number_num" numeric NULL;
UPDATE "nft_classes" SET "deployed_block_number_num" = "deployed_block_number"::numeric;
ALTER TABLE "nft_classes" DROP COLUMN "deployed_block_number";
ALTER TABLE "nft_classes" RENAME COLUMN "deployed_block_number_num" to "deployed_block_number";
ALTER TABLE "nft_classes" ALTER COLUMN "deployed_block_number" SET NOT NULL;
ALTER TABLE "nft_classes" ADD CONSTRAINT "uint64_deployed_block_number_check" CHECK (deployed_block_number >= (0)::numeric);

ALTER TABLE "nft_classes" ADD COLUMN "max_supply" numeric NOT NULL;
ALTER TABLE "nft_classes" ADD CONSTRAINT "uint64_max_supply_check" CHECK (max_supply >= (0)::numeric);
-- Modify "transaction_memos" table
ALTER TABLE "transaction_memos" ALTER COLUMN "token_id" TYPE numeric, ALTER COLUMN "block_number" TYPE numeric;
ALTER TABLE "transaction_memos" ADD CONSTRAINT "uint64_block_number_check" CHECK (block_number >= (0)::numeric), ADD CONSTRAINT "uint64_token_id_check" CHECK (token_id >= (0)::numeric);

-- Modify "nfts" table
ALTER TABLE "nfts" ALTER COLUMN "name" DROP NOT NULL, ALTER COLUMN "description" DROP NOT NULL, ALTER COLUMN "image" DROP NOT NULL, ALTER COLUMN "token_uri" DROP NOT NULL;
