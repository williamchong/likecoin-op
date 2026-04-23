-- atlas:txmode none

SET statement_timeout = 0;
SET lock_timeout = 0;
CREATE INDEX CONCURRENTLY IF NOT EXISTS "nft_owner_address_lower" ON "nfts" (LOWER("owner_address"));
CREATE INDEX CONCURRENTLY IF NOT EXISTS "nftclass_owner_address_lower" ON "nft_classes" (LOWER("owner_address"));
RESET statement_timeout;
RESET lock_timeout;
