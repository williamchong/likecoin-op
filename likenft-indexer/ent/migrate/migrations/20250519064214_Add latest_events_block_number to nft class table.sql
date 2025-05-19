-- Modify "nft_classes" table
ALTER TABLE "nft_classes" ADD COLUMN "latest_event_block_number" numeric NULL;
ALTER TABLE "nft_classes" ADD CONSTRAINT "uint64_latest_event_block_number_check" CHECK (latest_event_block_number >= (0)::numeric);
UPDATE "nft_classes" SET "latest_event_block_number" = "deployed_block_number";
ALTER TABLE "nft_classes" ALTER COLUMN "latest_event_block_number" SET NOT NULL;
