import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const LikeProtocolModule = buildModule("LikeProtocolModule", (m) => {
  const likeProtocolImpl = m.contract("LikeProtocol", [], {
    id: "LikeProtocolImpl",
  });

  return {
    likeProtocolImpl,
  };
});

export default LikeProtocolModule;
