import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { ethers, upgrades, viem, ignition } from "hardhat";
import LikeProtocolV1Module from "../ignition/modules/LikeProtocolV1";

export async function createProtocol(ownerSigner: SignerWithAddress) {
  const BookNFT = await ethers.getContractFactory("BookNFT");
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");

  const bookNFTDeployment = await BookNFT.deploy();
  let bookNFTAddress = await bookNFTDeployment.getAddress();

  const likeProtocol = await upgrades.deployProxy(
    LikeProtocol,
    [ownerSigner.address],
    {
      initializer: "initialize",
    },
  );
  const likeProtocolDeployment = await likeProtocol.waitForDeployment();
  const likeProtocolAddress = await likeProtocolDeployment.getAddress();
  const likeProtocolContract = await ethers.getContractAt(
    "LikeProtocol",
    likeProtocolAddress,
  );
  await likeProtocolContract.upgradeTo(bookNFTAddress);

  return {
    likeProtocol,
    likeProtocolDeployment,
    likeProtocolAddress,
    likeProtocolContract,
    bookNFTDeployment,
    bookNFTAddress,
  };
}

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
