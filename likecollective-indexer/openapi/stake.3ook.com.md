# APIs for stake.3ook.com

## Bookshelf

### Queries

#### 1. Get stakings by account

`GET /account/{evm_address}/stakings`

#### 2. Get total staking by class id

`GET /book-nfts?filter_book_nft_in=0x0001&filter_book_nft_in=0x0002`

## Dashboard

### Queries

#### 1. List all staked books (and amount?) from a wallet

`GET /account/{evm_address}/book-nfts`

#### 2. Total staked amount of a wallet across all book

`GET /accounts/{evm_address}`

#### 3. Total unclaimed reward of a wallet across all book

`GET /accounts/{evm_address}`

#### 4. Total claimed reward of a wallet across all book

`GET /accounts/{evm_address}`

#### 5. Historical stake, unstake, claim, reward events by any combination below

`GET /account/{evm_address}/staking-events/{staked|unstaked|reward-added|reward-claimed|reward-deposited|all}`
`GET /book-nft/{evm_address}/staking-events/{staked|unstaked|reward-added|reward-claimed|reward-deposited|all}`

## Explore Books

### Queries

#### 1. List top staked books by staked amount

`GET /book-nfts?time_frame_sort_by=staked_amount`

#### 2. List most staked books by number of staker

`GET /book-nfts?time_frame_sort_by=number_of_stakers`

#### 3. List recently staked book

`GET /book-nfts?time_frame_sort_by=last_staked_at`

#### 4. Get stakings by account

`GET /account/{evm_address}/stakings`

#### 5. Get total staking by class id

`GET /book-nfts?filter_book_nft_in=0x0001&filter_book_nft_in=0x0002`

#### 6. Filters

##### 1. Asc and desc sort

`GET /book-nfts?time_frame_sort_by=&time_frame_sort_order={asc|desc}`

##### 2. By date range

`GET /book-nfts?time_frame_sort_by=&time_frame={7d|30d|1y|all}`
