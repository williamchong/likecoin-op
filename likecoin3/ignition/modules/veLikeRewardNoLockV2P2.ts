import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import veLikeV0Module from "./veLikeV0";

/*
# Command to deploy the contract
DOTENV_CONFIG_PATH=.env \
  npx hardhat ignition deploy \
  ignition/modules/veLikeRewardNoLockV2P2.ts \
  --verify --strategy create2 \
  --parameters ignition/parameters.json \
  --network base

# Rerun only this for testing
npx hardhat ignition wipe chain-901 \
  veLikeRewardNoLockV2P2Module#veLikeRewardNoLockV2P2
npx hardhat ignition wipe chain-901 \
  veLikeRewardNoLockV2P2Module#ERC1967Proxy
npx hardhat ignition wipe chain-901 \
  "veLikeRewardNoLockV2P2Module#encodeFunctionCall(veLikeRewardNoLockV2P2Module#veLikeRewardNoLockV2P2Impl.initialize)"
npx hardhat ignition wipe chain-901 \
  veLikeRewardNoLockV2P2Module#veLikeRewardNoLockV2P2Impl
*/
const veLikeRewardNoLockV2P2Module = buildModule(
  "veLikeRewardNoLockV2P2Module",
  (m) => {
    const initOwner = m.getParameter("initOwner");
    const { veLikeV0, likecoin } = m.useModule(veLikeV0Module);

    const veLikeRewardNoLockV2P2Impl = m.contract(
      "contracts/veLikeRewardNoLockV2P2.sol:veLikeRewardNoLock",
      [],
      {
        id: "veLikeRewardNoLockV2P2Impl",
      },
    );

    const initData = m.encodeFunctionCall(
      veLikeRewardNoLockV2P2Impl,
      "initialize",
      [initOwner],
    );

    const veLikeRewardNoLockV2P2Proxy = m.contract("ERC1967Proxy", [
      veLikeRewardNoLockV2P2Impl,
      initData,
    ]);

    const veLikeRewardNoLockV2P2 = m.contractAt(
      "contracts/veLikeRewardNoLockV2P2.sol:veLikeRewardNoLock",
      veLikeRewardNoLockV2P2Proxy,
    );

    // Configure the reward contract to point at the vault and likecoin token.
    m.call(veLikeRewardNoLockV2P2, "setVault", [veLikeV0]);
    m.call(veLikeRewardNoLockV2P2, "setLikecoin", [likecoin]);

    return {
      veLikeRewardNoLockV2P2,
      veLikeRewardNoLockV2P2Proxy,
      veLikeRewardNoLockV2P2Impl,
    };
  },
);

export default veLikeRewardNoLockV2P2Module;
