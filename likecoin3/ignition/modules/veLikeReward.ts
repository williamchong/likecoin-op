import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import veLikeModule from "./veLike";

/*
# Command to deploy the contract for testing
npx hardhat ignition deploy ignition/modules/veLikeReward.ts \
  --strategy create2 \
  --parameters ignition/parameters.local.json \
  --network superism1

# Rerun only this for testing
npx hardhat ignition wipe chain-901 \
  veLikeRewardModule#veLikeReward
npx hardhat ignition wipe chain-901 \
  veLikeRewardModule#ERC1967Proxy
npx hardhat ignition wipe chain-901 \
  veLikeRewardModule#encodeFunctionCall(veLikeRewardModule#veLikeRewardImpl.initialize)
npx hardhat ignition wipe chain-901 \
  veLikeRewardModule#veLikeRewardImpl
*/
const veLikeRewardModule = buildModule("veLikeRewardModule", (m) => {
  const initOwner = m.getParameter("initOwner");

  const veLikeRewardImpl = m.contract("veLikeReward", [], {
    id: "veLikeRewardImpl",
  });

  const initData = m.encodeFunctionCall(veLikeRewardImpl, "initialize", [
    initOwner,
  ]);

  const veLikeRewardProxy = m.contract("ERC1967Proxy", [
    veLikeRewardImpl,
    initData,
  ]);

  const veLikeReward = m.contractAt("veLikeReward", veLikeRewardProxy);

  return {
    veLikeReward,
    veLikeRewardProxy,
    veLikeRewardImpl,
  };
});

export default veLikeRewardModule;
