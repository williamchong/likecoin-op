import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

/*
# Command to deploy the upgrade
DOTENV_CONFIG_PATH=.env \
  npx hardhat ignition deploy \
  ignition/modules/veLikeUpgradeV2.ts \
  --verify --strategy create2 \
  --parameters ignition/parameters.json \
  --network base

# Rerun only this for testing
npx hardhat ignition wipe chain-901 \
  veLikeUpgradeV2Module#veLikeV2Impl
npx hardhat ignition wipe chain-901 \
  veLikeUpgradeV2Module#veLikeProxy.upgradeToAndCall
npx hardhat ignition wipe chain-901 \
  veLikeUpgradeV2Module#veLikeRewardV2Impl
npx hardhat ignition wipe chain-901 \
  veLikeUpgradeV2Module#veLikeRewardProxy.upgradeToAndCall
*/
const veLikeUpgradeV2Module = buildModule("veLikeUpgradeV2Module", (m) => {
  const veLikeAddress = m.getParameter("veLikeAddress");
  const veLikeRewardAddress = m.getParameter("veLikeRewardAddress");

  const veLike = m.contractAt("contracts/veLike.sol:veLike", veLikeAddress);
  const veLikeReward = m.contractAt(
    "contracts/veLikeReward.sol:veLikeReward",
    veLikeRewardAddress,
  );

  // Deploy new veLike implementation and upgrade the proxy.
  // Adds: partial withdraw, legacy reward claiming, lock time management.
  const veLikeV2Impl = m.contract("contracts/veLikeV2.sol:veLike", [], {
    id: "veLikeV2Impl",
  });
  m.call(veLike, "upgradeToAndCall", [veLikeV2Impl, "0x"]);
  // For testing to force new ABI
  const veLikeV2 = m.contractAt(
    "contracts/veLikeV2.sol:veLike",
    veLikeAddress,
    {
      id: "veLikeV2",
    },
  );

  // Deploy new veLikeReward implementation and upgrade the proxy.
  // Updates withdraw(address) → withdraw(address, uint256) for partial
  // withdraw support, matching the IRewardContract interface in veLike V2.
  const veLikeRewardV2Impl = m.contract(
    "contracts/veLikeRewardV2.sol:veLikeReward",
    [],
    {
      id: "veLikeRewardV2Impl",
    },
  );
  m.call(veLikeReward, "upgradeToAndCall", [veLikeRewardV2Impl, "0x"]);
  const veLikeRewardV2 = m.contractAt(
    "contracts/veLikeRewardV2.sol:veLikeReward",
    veLikeRewardAddress,
    {
      id: "veLikeRewardV2",
    },
  );

  return {
    veLike: veLikeV2,
    veLikeReward: veLikeRewardV2,
    veLikeV2Impl,
    veLikeRewardV2Impl,
  };
});

export default veLikeUpgradeV2Module;
