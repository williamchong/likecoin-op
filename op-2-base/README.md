# Op-2-base

https://dba.stackexchange.com/questions/90482/export-postgres-table-as-json

## Prepare nft classes json

```
SELECT count(*) from nft_classes;

\o nft_classes

SELECT array_to_json(array_agg(temp)) AS ok_json FROM (

SELECT
    id,
    nft_classes.address,
    name, metadata, 
    metadata::json->'potentialAction'->'target'->0->'url' as "salt",
    metadata::json->'name' as "salt2",
    owner_address,
    max_supply,
    count FROM nft_classes, (
	SELECT nft_classes.address, COUNT(*) FROM nft_classes, nfts
	WHERE nft_classes.metadata::json->>'@type' = 'Book' AND nfts.contract_address = nft_classes.address
	GROUP BY nft_classes.address) AS C
WHERE C.address = nft_classes.address

) temp;

\o
```

## Prepare nfts json

```
SELECT count(*) from nfts;

\o nfts

SELECT array_to_json(array_agg(temp)) AS ok_json FROM (

select
    contract_address,
    token_id,
    token_uri,
    owner_address
from nfts

) temp;

\o
```

## Prepare transaction memos json

```
SELECT count(*) from transaction_memos;

\o transaction_memos

SELECT array_to_json(array_agg(temp)) AS ok_json FROM (

SELECT
    book_nft_id,
    token_id,
    memo,
    block_number
FROM transaction_memos

) temp;

\o
```

## Prepare minter and updater json

TODO

## Precompute addresses

```bash
go run ./cmd/cli workflow compute-address nft_classes.json | jq > addresses.json
```

## Prepare migration actions json

```bash
go run ./cmd/cli workflow prepare-actions nft_classes.json nfts.json transaction_memos.json | jq > migration-actions.json
```

## Prepare airdrop param json

```bash
go run ./cmd/cli workflow prepare-params migration-actions.json | jq > airdrop-params.json
```

## Prepare airdrop output json

```bash
go run ./cmd/cli workflow airdrop airdrop-params.json | jq > airdrop-output.json
```

Note that the command should be run with access to signer backend api

e.g.

```bash
docker compose run --rm op-2-base-cli go run ./cmd/cli workflow airdrop airdrop-params.json | jq > airdrop-output.json
```

## Prepare migrate db sqls

```bash
go run ./cmd/cli workflow migratedb airdrop-output.json > migratedb.sql
```
