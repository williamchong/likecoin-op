import { viem, ignition } from "hardhat";
import LikeProtocolV1Module from "../ignition/modules/LikeProtocolV1";

export async function deployProtocol() {
  const [deployer, classOwner, likerLand, randomSigner, randomSigner2] =
    await viem.getWalletClients();
  const publicClient = await viem.getPublicClient();

  const { likeProtocolImpl, likeProtocol, bookNFTImpl } = await ignition.deploy(
    LikeProtocolV1Module,
    {
      parameters: {
        LikeProtocolV0Module: {
          initOwner: deployer.account.address,
        },
      },
      defaultSender: deployer.account.address,
      strategy: "create2",
    },
  );

  return {
    likeProtocolImpl,
    likeProtocol,
    bookNFTImpl,
    deployer,
    classOwner,
    likerLand,
    randomSigner,
    randomSigner2,
    publicClient,
  };
}
