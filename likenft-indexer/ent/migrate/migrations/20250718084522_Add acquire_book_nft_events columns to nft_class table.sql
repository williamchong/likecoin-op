-- Modify "nft_classes" table
ALTER TABLE "nft_classes"
    ADD COLUMN "acquire_book_nft_events_weight" double precision NOT NULL DEFAULT 1,
    ADD COLUMN "acquire_book_nft_events_last_processed_time" timestamptz NULL,
    ADD COLUMN "acquire_book_nft_events_score" numeric NULL,
    ADD COLUMN "acquire_book_nft_events_status" character varying NULL,
    ADD COLUMN "acquire_book_nft_events_failed_reason" character varying NULL,
    ADD COLUMN "acquire_book_nft_events_failed_count" bigint NOT NULL DEFAULT 0;
-- Create index "nftclass_acquire_book_nft_events_score" to table: "nft_classes"
CREATE INDEX "nftclass_acquire_book_nft_events_score" ON "nft_classes" ("acquire_book_nft_events_score");
-- Create index "nftclass_acquire_book_nft_events_status" to table: "nft_classes"
CREATE INDEX "nftclass_acquire_book_nft_events_status" ON "nft_classes" ("acquire_book_nft_events_status");
