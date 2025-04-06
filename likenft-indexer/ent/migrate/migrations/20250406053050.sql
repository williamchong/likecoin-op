-- Modify "nfts" table
ALTER TABLE "nfts" ADD COLUMN "token_id_num" numeric NULL;
UPDATE "nfts" SET "token_id_num" = "token_id"::numeric;
ALTER TABLE "nfts" DROP COLUMN "token_id";
ALTER TABLE "nfts" RENAME COLUMN "token_id_num" to "token_id";
ALTER TABLE "nfts" ALTER COLUMN "token_id" SET NOT NULL;
ALTER TABLE "nfts" ADD CONSTRAINT "uint64_token_id_check" CHECK (token_id >= (0)::numeric);
CREATE UNIQUE INDEX "nft_contract_address_token_id" ON "nfts" ("contract_address", "token_id");
