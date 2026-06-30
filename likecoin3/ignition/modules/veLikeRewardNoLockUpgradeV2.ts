import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

/*
# Command to deploy the upgrade
DOTENV_CONFIG_PATH=.env \
  npx hardhat ignition deploy \
  ignition/modules/veLikeRewardNoLockUpgradeV2.ts \
  --verify --strategy create2 \
  --parameters ignition/parameters.json \
  --network base

# Rerun only this for testing
npx hardhat ignition wipe chain-901 \
  veLikeRewardNoLockUpgradeV2Module#veLikeRewardNoLockV2Impl
npx hardhat ignition wipe chain-901 \
  veLikeRewardNoLockUpgradeV2Module#veLikeRewardNoLockProxy.upgradeToAndCall
*/
const veLikeRewardNoLockUpgradeV2Module = buildModule(
  "veLikeRewardNoLockUpgradeV2Module",
  (m) => {
    const veLikeRewardNoLockAddress = m.getParameter(
      "veLikeRewardNoLockAddress",
    );

    const veLikeRewardNoLock = m.contractAt(
      "contracts/veLikeRewardNoLock.sol:veLikeRewardNoLock",
      veLikeRewardNoLockAddress,
    );

    // Deploy new veLikeRewardNoLock implementation and upgrade the proxy.
    // Adds: finalizeSync() and auto-sync-gated syncStakers() so pre-rotation
    // stakers can be materialized across multiple calls and closed out safely.
    const veLikeRewardNoLockV2Impl = m.contract(
      "contracts/veLikeRewardNoLockV2.sol:veLikeRewardNoLock",
      [],
      {
        id: "veLikeRewardNoLockV2Impl",
      },
    );
    m.call(veLikeRewardNoLock, "upgradeToAndCall", [
      veLikeRewardNoLockV2Impl,
      "0x",
    ]);
    // For testing to force new ABI
    const veLikeRewardNoLockV2 = m.contractAt(
      "contracts/veLikeRewardNoLockV2.sol:veLikeRewardNoLock",
      veLikeRewardNoLockAddress,
      {
        id: "veLikeRewardNoLockV2",
      },
    );

    return {
      veLikeRewardNoLock: veLikeRewardNoLockV2,
      veLikeRewardNoLockV2Impl,
    };
  },
);

export default veLikeRewardNoLockUpgradeV2Module;
