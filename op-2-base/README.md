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

## 

```bash
go build ./cmd/cli
```

## Precompute addresses

```bash
go run ./cmd/cli workflow compute-address nft_classes.json | jq > addresses.json
```

## Prepare migration actions json

```bash
mkdir -p migration-actions
jq -r '.[] | .old_address' addresses.json | xargs -n 1 -I {} bash -c './cli workflow prepare-actions nft_classes.json nfts.json transaction_memos.json {} | jq > migration-actions/{}.json'

go run ./cmd/cli workflow prepare-actions nft_classes.json nfts.json transaction_memos.json 0x00DD2ec446cC9Ea9FA40dd484feBb6B0217cA4b4
```

## Prepare airdrop param json

```bash
mkdir -p airdrop-params
ls migration-actions | xargs -n 1 -I {} bash -c './cli workflow prepare-params migration-actions/{} | jq > airdrop-params/{}'
```

## Perform actual airdrop

```bash
mkdir -p airdrop-outputs
./cli workflow airdrop airdrop-params/0xF8307083bC727DfBB9067Cfa46DF5C5Bd68872b4.json | jq > airdrop-outputs/0xF8307083bC727DfBB9067Cfa46DF5C5Bd68872b4.json
```

Note that the command should be run with access to signer backend api

e.g.

```bash
docker compose run --rm op-2-base-cli go run ./cmd/cli workflow airdrop airdrop-params/0x2D28c4154c56488f608394f9B3d3d45932c3F1c9.json | jq > airdrop-outputs/0x2D28c4154c56488f608394f9B3d3d45932c3F1c9.json
```

0x3d8003100b87BaD41c1fb2c9343BA4B5d312E9b7

## Prepare migrate db sqls

```bash
go run ./cmd/cli workflow migratedb airdrop-output.json > migratedb.sql
./cli workflow migratedb airdrop-outputs/0x2D28c4154c56488f608394f9B3d3d45932c3F1c9.json > migratedb.sql
```
