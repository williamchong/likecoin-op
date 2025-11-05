import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import LikeStakePositionV0Module from "./LikeStakePositionV0";

/*

# Command to deploy the contract for testing
npx hardhat ignition deploy ignition/modules/LikeStakePositionV1.ts \
  --strategy create2 \
  --parameters ignition/parameters.local.json \
  --network superism1

npx hardhat ignition wipe chain-901 \
  LikeStakePositionV1Module#LikeStakePositionV0Module~LikeStakePositionV0.upgradeToAndCall
 npx hardhat ignition wipe chain-901 \
  LikeStakePositionV1Module#LikeStakePosition
npx hardhat ignition wipe chain-901 \
  LikeStakePositionV1Module#LikeStakePositionV1Impl

*/

const LikeStakePositionV1Module = buildModule(
  "LikeStakePositionV1Module",
  (m) => {
    const { likeStakePositionV0 } = m.useModule(LikeStakePositionV0Module);

    const likeStakePositionV1Impl = m.contract("LikeStakePosition", [], {
      id: "LikeStakePositionV1Impl",
    });

    m.call(likeStakePositionV0, "upgradeToAndCall", [
      likeStakePositionV1Impl,
      "0x",
    ]);

    const likeStakePosition = m.contractAt(
      "LikeStakePosition",
      likeStakePositionV0,
    );

    return {
      likeStakePositionV1Impl,
      likeStakePosition,
    };
  },
);

export default LikeStakePositionV1Module;
