import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const LikeStakePositionV0Module = buildModule(
  "LikeStakePositionV0Module",
  (m) => {
    const initOwner = m.getParameter("initOwner");

    const likeStakePositionV0Impl = m.contract("LikeStakePositionV0", [], {
      id: "LikeStakePositionV0Impl",
    });

    const initDataV0 = m.encodeFunctionCall(
      likeStakePositionV0Impl,
      "initialize",
      [initOwner],
    );
    const likeStakePositionProxy = m.contract("ERC1967Proxy", [
      likeStakePositionV0Impl,
      initDataV0,
    ]);

    const likeStakePositionV0 = m.contractAt(
      "LikeStakePositionV0",
      likeStakePositionProxy,
    );

    return {
      likeStakePositionV0,
      likeStakePositionV0Impl,
      likeStakePositionProxy,
    };
  },
);

export default LikeStakePositionV0Module;
