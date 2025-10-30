import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import veLikeV0Module from "./veLikeV0";

/*
# Command to deploy the contract for testing
npx hardhat ignition deploy ignition/modules/veLike.ts \
  --strategy create2 \
  --parameters ignition/parameters.local.json \
  --network superism1

# Rerun only this for testing
npx hardhat ignition wipe chain-901 \
  veLikeModule#veLikeV0Module~veLikeV0.upgradeToAndCall
npx hardhat ignition wipe chain-901 \
  veLikeModule#veLikeImpl 
*/
const veLikeModule = buildModule("veLikeModule", (m) => {
  const { veLikeV0, veLikeProxy, likecoin } = m.useModule(veLikeV0Module);

  const veLikeImpl = m.contract("veLike", [], {
    id: "veLikeImpl",
  });
  m.call(veLikeV0, "upgradeToAndCall", [veLikeImpl, "0x"]);
  const veLike = m.contractAt("veLike", veLikeProxy);

  return {
    veLike,
    veLikeProxy,
    veLikeImpl,
    likecoin,
  };
});

export default veLikeModule;
