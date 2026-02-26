import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import veLikeModule from "./veLike";

/*
# Command to deploy the contract for testing
npx hardhat ignition deploy ignition/modules/veLikeRewardNoLock.ts \
  --strategy create2 \
  --parameters ignition/parameters.local.json \
  --network superism1

# Rerun only this for testing
npx hardhat ignition wipe chain-901 \
  veLikeRewardNoLockModule#veLikeRewardNoLock
npx hardhat ignition wipe chain-901 \
  veLikeRewardNoLockModule#ERC1967Proxy
npx hardhat ignition wipe chain-901 \
  "veLikeRewardNoLockModule#encodeFunctionCall(veLikeRewardNoLockModule#veLikeRewardNoLockImpl.initialize)"
npx hardhat ignition wipe chain-901 \
  veLikeRewardNoLockModule#veLikeRewardNoLockImpl
*/
const veLikeRewardNoLockModule = buildModule(
  "veLikeRewardNoLockModule",
  (m) => {
    const initOwner = m.getParameter("initOwner");
    const { veLike, likecoin } = m.useModule(veLikeModule);

    const veLikeRewardNoLockImpl = m.contract("veLikeRewardNoLock", [], {
      id: "veLikeRewardNoLockImpl",
    });

    const initData = m.encodeFunctionCall(
      veLikeRewardNoLockImpl,
      "initialize",
      [initOwner],
    );

    const veLikeRewardNoLockProxy = m.contract("ERC1967Proxy", [
      veLikeRewardNoLockImpl,
      initData,
    ]);

    const veLikeRewardNoLock = m.contractAt(
      "veLikeRewardNoLock",
      veLikeRewardNoLockProxy,
    );

    // Configure the reward contract to point at the vault and likecoin token.
    m.call(veLikeRewardNoLock, "setVault", [veLike]);
    m.call(veLikeRewardNoLock, "setLikecoin", [likecoin]);

    // Initialize totalStaked from vault's totalSupply so the reward accumulator
    // uses the correct denominator for pre-rotation stakers, and enable
    // auto-sync for lazy staker enrollment.
    m.call(veLikeRewardNoLock, "initTotalStaked", []);

    return {
      veLikeRewardNoLock,
      veLikeRewardNoLockProxy,
      veLikeRewardNoLockImpl,
    };
  },
);

export default veLikeRewardNoLockModule;
