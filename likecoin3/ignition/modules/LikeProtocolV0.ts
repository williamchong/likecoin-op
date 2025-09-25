import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const LikeProtocolV0Module = buildModule("LikeProtocolV0Module", (m) => {
  const initOwner = m.getParameter("initOwner");
  const bookNFTImplementation = m.getParameter("bookNFTImplementation");

  const likeProtocolV0Impl = m.contract("LikeProtocolV0", [], {
    id: "LikeProtocolV0Impl",
  });

  const initDataV0 = m.encodeFunctionCall(likeProtocolV0Impl, "initialize", [
    initOwner,
  ]);

  const likeProtocolProxy = m.contract("ERC1967Proxy", [
    likeProtocolV0Impl,
    initDataV0,
  ]);

  const likeProtocolV0 = m.contractAt("LikeProtocolV0", likeProtocolProxy);

  return {
    likeProtocolV0,
    likeProtocolV0Impl,
    likeProtocolProxy,
  };
});

export default LikeProtocolV0Module;
