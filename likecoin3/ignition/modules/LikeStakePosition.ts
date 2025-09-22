import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import LikeStakePositionV0Module from "./LikeStakePositionV0";

const LikeStakePositionModule = buildModule("LikeStakePositionModule", (m) => {
  const { likeStakePositionV0 } = m.useModule(LikeStakePositionV0Module);

  const likeStakePositionImpl = m.contract("LikeStakePosition", [], {
    id: "LikeStakePositionImpl",
  });

  m.call(likeStakePositionV0, "upgradeToAndCall", [
    likeStakePositionImpl,
    "0x",
  ]);

  const likeStakePosition = m.contractAt(
    "LikeStakePosition",
    likeStakePositionV0,
  );

  return {
    likeStakePositionImpl,
    likeStakePosition,
  };
});

export default LikeStakePositionModule;
