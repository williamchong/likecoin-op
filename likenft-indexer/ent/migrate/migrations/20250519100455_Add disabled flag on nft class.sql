-- Modify "nft_classes" table
ALTER TABLE "nft_classes" ADD COLUMN "disabled_for_indexing" boolean NOT NULL DEFAULT false, ADD COLUMN "disabled_for_indexing_reason" character varying NULL;
