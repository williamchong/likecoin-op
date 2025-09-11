# DEV command for reference

Common command that is useful for local dev with default hardhat account `0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266` as in .env.local

## Operations

### Mint coin for itself

```
npx hardhat mint \
    --network superism1 \
    --likecoin 0x5540d1fdC1b1fec3Ed8AEF26f73951c6ce5F97A1 \
    --amount 10000 \
    --to 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

### Stake

```
npx hardhat stake \
    --network superism1 \
    --collective 0x37f0dE4E511efC5D993c0b6d8cB4B73560b2f127 \
    --amount 1000 --booknft 0x9EC5814c3307862e9647aA0CA6758fF294f6ba26
```

### Unstake

```
npx hardhat unstake \
    --network local \
    --collective 0x37f0dE4E511efC5D993c0b6d8cB4B73560b2f127 \
    --amount 1000 \
    --booknft 0x9EC5814c3307862e9647aA0CA6758fF294f6ba26
```

### Emit Only Reward Added

```
npx hardhat emitOnlyRewardAdded \
    --network local \
    --collective 0x37f0dE4E511efC5D993c0b6d8cB4B73560b2f127 \
    --amount 1000 \
    --booknft 0x9EC5814c3307862e9647aA0CA6758fF294f6ba26
```

### Claim

```
npx hardhat claimRewards \
    --network local \
    --collective 0x37f0dE4E511efC5D993c0b6d8cB4B73560b2f127 \
    --booknft 0x9EC5814c3307862e9647aA0CA6758fF294f6ba26
```

### Claim All Rewards

```
npx hardhat claimAllRewards \
    --network local \
    --collective 0x37f0dE4E511efC5D993c0b6d8cB4B73560b2f127
```

### Restake Reward

```
npx hardhat restakeReward \
    --network local \
    --collective 0x37f0dE4E511efC5D993c0b6d8cB4B73560b2f127 \
    --booknft 0x9EC5814c3307862e9647aA0CA6758fF294f6ba26
```
