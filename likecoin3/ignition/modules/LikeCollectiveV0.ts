import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const LikeCollectiveV0Module = buildModule("LikeCollectiveV0Module", (m) => {
  const initOwner = m.getParameter("initOwner");

  const likeCollectiveV0Impl = m.contract("LikeCollectiveV0", [], {
    id: "LikeCollectiveV0Impl",
  });
  const initDataV0 = m.encodeFunctionCall(likeCollectiveV0Impl, "initialize", [
    initOwner,
  ]);
  const likeCollectiveProxy = m.contract("ERC1967Proxy", [
    likeCollectiveV0Impl,
    initDataV0,
  ]);
  const likeCollectiveV0 = m.contractAt(
    "LikeCollectiveV0",
    likeCollectiveProxy,
  );

  return {
    likeCollectiveV0,
    likeCollectiveV0Impl,
    likeCollectiveProxy,
  };
});

export default LikeCollectiveV0Module;
