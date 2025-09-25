import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import LikeProtocolV0Module from "./LikeProtocolV0";
import BookNFTModule from "./BookNFTV1";

const LikeProtocolV1Module = buildModule("LikeProtocolV1Module", (m) => {
  const { likeProtocolV0 } = m.useModule(LikeProtocolV0Module);
  const { bookNFTImpl } = m.useModule(BookNFTModule);

  const initOwner = m.staticCall(likeProtocolV0, "owner");
  const likeProtocolImpl = m.contract("LikeProtocol", [], {
    id: "LikeProtocolV1Impl",
  });

  const upgradeToData = m.encodeFunctionCall(likeProtocolImpl, "upgradeTo", [
    bookNFTImpl,
  ]);
  m.call(likeProtocolV0, "upgradeToAndCall", [likeProtocolImpl, upgradeToData]);

  const likeProtocol = m.contractAt("LikeProtocol", likeProtocolV0);

  return {
    likeProtocolImpl,
    likeProtocol,
    bookNFTImpl,
  };
});

export default LikeProtocolV1Module;
