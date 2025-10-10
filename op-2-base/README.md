# Op-2-base

https://dba.stackexchange.com/questions/90482/export-postgres-table-as-json


## Prepare nft classes json

```
SELECT count(*) from nft_classes 
WHERE nft_classes.metadata::json->>'@type' = 'Book';

\o nft_classes

SELECT array_to_json(array_agg(temp)) AS ok_json FROM (

SELECT
	nft_classes.id,
	nft_classes.address,
	name,
	metadata,
	metadata::json->'potentialAction'->'target'->0->'url' as "salt",
	metadata::json->'name' as "salt2",
    nft_classes.max_supply,
	C.count,
    accounts.evm_address as owner_address
FROM nft_classes LEFT JOIN (
	SELECT nfts.contract_address, COUNT(*) FROM nfts 
	GROUP BY nfts.contract_address) AS C ON C.contract_address = nft_classes.address
LEFT JOIN accounts ON accounts.id = nft_classes.account_nft_classes
WHERE
	nft_classes.metadata::json->>'@type' = 'Book'

) temp;

\o
```

For formating the output into json
```
sed -n 3p nft_classes | jq > nft_classes.json
# Checking count is correct
jq length nft_classes.json
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

For formating the output into json
```
sed -n 3p nfts | jq > nfts.json
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

For formating the output into json
```
sed -n 3p transaction_memos | jq > transaction_memos.json
```

## Prepare minter and updater json

```
SELECT count(*) from (

select
    address as "booknft",
    topic0 as "event",
    topic1 as "role_byte_array_string",
    topic2 as "to",
    topic3 as "by",
    block_number,
    transaction_index,
    log_index
from evm_events
where ("topic0" = 'RoleGranted' or "topic0" = 'RoleRevoked')
order by
    block_number asc,
    transaction_index asc,
    log_index asc
);

\o evm_events_booknft_role_changed

SELECT array_to_json(array_agg(temp)) AS ok_json FROM (

select
    address as "booknft",
    topic0 as "event",
    topic1 as "role_byte_array_string",
    topic2 as "to",
    topic3 as "by",
    block_number,
    transaction_index,
    log_index
from evm_events
where ("topic0" = 'RoleGranted' or "topic0" = 'RoleRevoked')
order by
    block_number asc,
    transaction_index asc,
    log_index asc
) temp;

\o
```

For formating the output into json

```
sed -n 3p evm_events_booknft_role_changed | jq > evm_events_booknft_role_changed.json
```

## 

```bash
go build ./cmd/cli
```

## Precompute addresses

```bash
go run ./cmd/cli workflow compute-address nft_classes.json | jq > addresses.json
```

### Find duplicated new address

```bash
jq -r 'group_by(.new_address) | map({new_address: .[0] | .new_address, addresses: .}) | map(select(.addresses | length > 1))' addresses.json > duplicated_new_addresses.json
```

Retrieving plain addresses list

```bash
jq -r '.[] | .addresses[] | .old_address' duplicated_new_addresses.json
```

### Replace salt of duplicated addresses

```bash
python ./workflow/replace_salt.py \
    nft_classes.json \
    duplicated_new_addresses.json | jq > nft_classes.alt.json
```

```bash
diff nft_classes.json nft_classes.alt.json
mv nft_classes.json nft_classes.bak.json
mv nft_classes.alt.json nft_classes.json
```

### Compute addresses again

```bash
./cli workflow compute-address nft_classes.alt.json | jq > addresses.alt.json
```

```bash
diff addresses.json addresses.alt.json
```

Inspecting against the new addresses should have no duplications.

```bash
jq -r 'group_by(.new_address) | map({new_address: .[0] | .new_address, addresses: .}) | map(select(.addresses | length > 1))' addresses.alt.json
```

In case the addresses are resolved, move the address.alt.json to address

```
mv addresses.json addresses.bak.json
mv addresses.alt.json addresses.json
```

## Prepare migration actions json

```bash
mkdir -p migration-actions
jq -r '.[] | .old_address' addresses.json | xargs -n 1 -I {} bash -c './cli workflow prepare-actions nft_classes.json nfts.json transaction_memos.json evm_events_booknft_role_changed.json {} | jq > migration-actions/{}.json'

go run ./cmd/cli workflow prepare-actions nft_classes.json nfts.json transaction_memos.json evm_events_booknft_role_changed.json 0x00DD2ec446cC9Ea9FA40dd484feBb6B0217cA4b4
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
