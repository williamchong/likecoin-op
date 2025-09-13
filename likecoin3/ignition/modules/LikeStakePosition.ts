import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const LikeStakePositionModule = buildModule("LikeStakePositionModule", (m) => {
  const initOwner = m.getParameter("initOwner");

  const likeStakePositionImpl = m.contract("LikeStakePosition", [], { id: "LikeStakePositionImpl" });

  const initData = m.encodeFunctionCall(likeStakePositionImpl, "initialize", [
    initOwner,
  ]);
  const likeStakePositionProxy = m.contract("ERC1967Proxy", [likeStakePositionImpl, initData]);

  const likeStakePosition = m.contractAt("LikeStakePosition", likeStakePositionProxy);

  return { likeStakePosition, likeStakePositionImpl, likeStakePositionProxy };
});

export default LikeStakePositionModule;


