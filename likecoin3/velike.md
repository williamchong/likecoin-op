# veLike deployment notes

Since the deployment of veLike is a bit complicated.

## Addresses (Base Mainnet)

| Contract                                             | Address                                      |
| ---------------------------------------------------- | -------------------------------------------- |
| LIKE token                                           | `0x1EE5DD1794C28F559f94d2cc642BaE62dC3be5cf` |
| veLike vault (proxy)                                 | `0xE55C2b91E688BE70e5BbcEdE3792d723b4766e2B` |
| veLikeReward (1st period, **legacy** after rotation) | `0x465629cedF312B77C48602D5AfF1Ecb4FEb1Bf62` |

Set env vars before running cast commands:

```bash
export RPC=https://mainnet.base.org   # or your Alchemy/Infura URL
export VELIKE=0xE55C2b91E688BE70e5BbcEdE3792d723b4766e2B
export LIKE=0x1EE5DD1794C28F559f94d2cc642BaE62dC3be5cf
```

---

## Initial Deployment Reference (commit 80a60ec)

Period 1: **Nov 3 2025 → Feb 1 2026** (timestamps `1762164000` → `1769940000`)

```
DOTENV_CONFIG_PATH=.env \
    npx hardhat ignition deploy \
    ignition/modules/veLike.ts \
    --verify --strategy create2 \
    --parameters ignition/parameters.json --network baseSepolia
DOTENV_CONFIG_PATH=.env \
    npx hardhat ignition deploy \
    ignition/modules/veLikeReward.ts \
    --verify --strategy create2 \
    --parameters ignition/parameters.json --network baseSepolia
```

Post-deploy setup calls that were made (for reference):

```bash
# Wire reward to vault and LIKE token
cast send 0x465629cedF312B77C48602D5AfF1Ecb4FEb1Bf62 \
  "setVault(address)" 0xE55C2b91E688BE70e5BbcEdE3792d723b4766e2B \
  --account likecoin-deployer.eth --rpc-url $RPC

cast send 0x465629cedF312B77C48602D5AfF1Ecb4FEb1Bf62 \
  "setLikecoin(address)" 0x1EE5DD1794C28F559f94d2cc642BaE62dC3be5cf \
  --account likecoin-deployer.eth --rpc-url $RPC

# Fund the reward period (drawer: 0x1f135ca20cE4d5Abb53dB27c7F981b43a5734419)
cast send 0x465629cedF312B77C48602D5AfF1Ecb4FEb1Bf62 \
  "addReward(address,uint256,uint256,uint256)" \
  0x1f135ca20cE4d5Abb53dB27c7F981b43a5734419 \
  10000000000000 1762164000 1769940000 \
  --account likecoin-deployer.eth --rpc-url $RPC

# Set vault's active reward and lock (users can't withdraw before lock time)
cast send $VELIKE "setRewardContract(address)" 0x465629cedF312B77C48602D5AfF1Ecb4FEb1Bf62 \
  --account likecoin-deployer.eth --rpc-url $RPC

cast send $VELIKE "setLockTime(uint256)" 1769940000 \
  --account likecoin-deployer.eth --rpc-url $RPC
```

---

## Upgrading veLike & veLikeReward (→ V2)

The initial deployment (commit `80a60ec`) used `veLikeV0` and the original `veLikeReward`.
Before rotating to `veLikeRewardNoLock`, both contracts must be upgraded in-place:

| Contract         | Change                                                                                                               |
| ---------------- | -------------------------------------------------------------------------------------------------------------------- |
| **veLike**       | Adds partial withdraw support, `setLockTime`, `setLegacyRewardContract`, `claimLegacyReward`, non-transferable token |
| **veLikeReward** | Updates `withdraw(address)` → `withdraw(address, uint256)` to match the `IRewardContract` interface in veLike V1     |

Both upgrades are handled by a single ignition module so the proxy implementations stay
in sync.

```bash
DOTENV_CONFIG_PATH=.env \
    npx hardhat ignition deploy \
    ignition/modules/veLikeUpgradeV2.ts \
    --verify --strategy create2 \
    --parameters ignition/parameters.json --network baseSepolia
```

### Prerequisites

- The current reward period has ended (or will end before any user withdraws after the
  upgrade).
- You have the deployer account that owns both proxies.

### Step 1: Deploy the upgrade

```bash
DOTENV_CONFIG_PATH=.env \
  npx hardhat ignition deploy \
  ignition/modules/veLikeUpgradeV1.ts \
  --verify --strategy create2 \
  --parameters ignition/parameters.json \
  --network base
```

This deploys new implementation contracts for both `veLike` and `veLikeReward`, then calls
`upgradeToAndCall` on each proxy.

### Step 2: Verify

Check the new implementation addresses in
`ignition/deployments/chain-8453/deployed_addresses.json`:

- `"veLikeUpgradeV1Module#veLikeV1Impl"` — new veLike implementation
- `"veLikeUpgradeV1Module#veLikeRewardV1Impl"` — new veLikeReward implementation

Confirm on-chain:

```bash
# veLike proxy should point to the new impl
cast storage $VELIKE 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc --rpc-url $RPC

# veLikeReward proxy should point to the new impl
export REWARD=0x465629cedF312B77C48602D5AfF1Ecb4FEb1Bf62
cast storage $REWARD 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc --rpc-url $RPC
```

### Next steps

After the upgrade, proceed to **"Rotating to a New veLikeRewardNoLock"** below to deploy
the no-lock reward contract for the next period.

---

## Rotating to a New veLikeRewardNoLock (2nd period onwards)

Use this when the current reward period has ended and you want to start a new one with a
fresh contract. The `veLikeRewardNoLock` contract supports auto-enrollment of existing
stakers without requiring them to re-deposit.

### Before you start

- Confirm the current period has ended (`endTime` has passed).
- Decide the new period parameters: `startTime`, `endTime`, `rewardAmount`, `drawer`.
- The `drawer` account must hold enough LIKE and pre-approve the new reward contract.

### Step 1: (Optional but recommended) Sync stakers on the old contract

Call `syncStakers` on the **current** (old) reward contract to freeze staker balances before
anyone withdraws between periods. Must be called while the old period is still active.

```bash
cast send $OLD_REWARD \
  "syncStakers(address[])" \
  "[0xAddr1,0xAddr2,...]" \
  --account likecoin-deployer.eth --rpc-url $RPC
```

Skip this if all stakers have already interacted with the old contract (they are already
synced) or if the period has already ended.

### Step 2: Add parameters for the new ignition module

Add to `ignition/parameters.json`:

```json
"veLikeRewardNoLockModule": {
  "initOwner": "0x2dd2253cd5bef4ea6d74efdfad9718a73a7d7ec7"
}
```

For a 3rd/4th reward, duplicate `ignition/modules/veLikeRewardNoLock.ts` as
`veLikeRewardNoLockV2.ts` (change the module ID string to `"veLikeRewardNoLockV2Module"`)
and add a matching `"veLikeRewardNoLockV2Module"` entry in `parameters.json`.

### Step 3: Deploy the new reward contract

The ignition module calls `setVault` and `setLikecoin` during deployment.

```bash
DOTENV_CONFIG_PATH=.env \
  npx hardhat ignition deploy \
  ignition/modules/veLikeRewardNoLock.ts \
  --verify --strategy create2 \
  --parameters ignition/parameters.json \
  --network base
```

Note the new contract address from
`ignition/deployments/chain-8453/deployed_addresses.json` — it will appear as
`"veLikeRewardNoLockModule#veLikeRewardNoLock"`.

### Step 4: Fund and start the new reward period

The drawer must first approve the reward contract, then the owner calls `addReward`.

`startTime` must be strictly greater than the old contract's `lastRewardTime` (which equals
the old period's `startTime`). Use `date -d "..." +%s` to convert a human date to Unix
timestamp.

```bash
export NEW_REWARD=<address from step 3>

# Drawer approves (run as the drawer account, not deployer)
# it should be on a multisig account.
cast send $LIKE \
  "approve(address,uint256)" $NEW_REWARD $REWARD_AMOUNT \
  --account <drawer-account> --rpc-url $RPC

# Start the period
cast send $NEW_REWARD \
  "addReward(address,uint256,uint256,uint256)" \
  $DRAWER $REWARD_AMOUNT $START_TIME $END_TIME \
  --account likecoin-deployer.eth --rpc-url $RPC
```

### Step 5: Wire the vault to the new reward contract

```bash
export OLD_REWARD=0x465629cedF312B77C48602D5AfF1Ecb4FEb1Bf62  # period 1 for first rotation
export MULTICALL3=0xcA11bde05977b3631167028862bE2a173976CA11
```

Remove the lock (only needed when rotating from the original locked `veLikeReward`):

```bash
cast send $VELIKE "setLockTime(uint256)" 0 \
  --account likecoin-deployer.eth --rpc-url $RPC
```

Atomically snapshot `totalSupply` and switch the active reward in one transaction via
Multicall3, so no withdrawal can land in between:

```bash
cast send $MULTICALL3 \
  "aggregate3((address,bool,bytes)[])" \
  "[($NEW_REWARD,false,$(cast calldata "initTotalStaked()")),($VELIKE,false,$(cast calldata "setRewardContract(address)" $NEW_REWARD))]" \
  --account likecoin-deployer.eth --rpc-url $RPC
```

Register the old contract as legacy so users can still claim accrued rewards:

```bash
cast send $VELIKE "setLegacyRewardContract(address,bool)" $OLD_REWARD true \
  --account likecoin-deployer.eth --rpc-url $RPC
```

### Step 6: (Optional) Eager staker sync on new contract

Pre-rotation stakers are lazily auto-enrolled on their first deposit/withdraw/claim.
To eagerly sync them (useful to front-run any stale-balance edge cases), call during
the active period (`startTime` ≤ now ≤ `endTime`):

```bash
cast send $NEW_REWARD \
  "syncStakers(address[])" \
  "[0xAddr1,0xAddr2,...]" \
  --account likecoin-deployer.eth --rpc-url $RPC
```

---

## Adding Another Period to the Same Contract (no rotation)

If you want to reuse an existing `veLikeRewardNoLock` contract for the next period (no new
deployment needed), just repeat the funding step after the current period ends:

```bash
# Drawer approves
cast send $LIKE "approve(address,uint256)" $REWARD_CONTRACT $REWARD_AMOUNT \
  --account <drawer-account> --rpc-url $RPC

# Add the next period (startTime must be > lastRewardTime of the contract)
cast send $REWARD_CONTRACT \
  "addReward(address,uint256,uint256,uint256)" \
  $DRAWER $REWARD_AMOUNT $START_TIME $END_TIME \
  --account likecoin-deployer.eth --rpc-url $RPC
```

---

## Claiming Legacy Rewards

After rotation, users can claim their accrued rewards from the old contract via the vault.
The legacy reward contract must be registered with `setLegacyRewardContract` (step 4 above).

Using the hardhat task (permissionless — anyone can trigger a claim on behalf of any account):

```bash
DOTENV_CONFIG_PATH=.env npx hardhat claimLegacyReward \
  --velike $VELIKE \
  --legacyreward $OLD_REWARD \
  --account $USER_ADDRESS \
  --network base
```

Or via cast:

```bash
cast send $VELIKE \
  "claimLegacyReward(address,address)" \
  $OLD_REWARD $USER_ADDRESS \
  --account likecoin-deployer.eth --rpc-url $RPC
```
