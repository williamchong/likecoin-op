-- Drop index "nftclass_acquire_book_nft_events_score" from table: "nft_classes"
DROP INDEX "nftclass_acquire_book_nft_events_score";
-- Rename a column from "acquire_book_nft_events_score" to "acquire_book_nft_events_eta"
ALTER TABLE "nft_classes" RENAME COLUMN "acquire_book_nft_events_score" TO "acquire_book_nft_events_eta";
-- Create index "nftclass_acquire_book_nft_events_eta" to table: "nft_classes"
CREATE INDEX "nftclass_acquire_book_nft_events_eta" ON "nft_classes" ("acquire_book_nft_events_eta");
