import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { ethers, upgrades } from "hardhat";

export async function createProtocol(ownerSigner: SignerWithAddress) {
  const BookNFT = await ethers.getContractFactory("BookNFT");
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");

  const bookNFT = await upgrades.deployProxy(
    BookNFT,
    [
      {
        creator: ownerSigner.address,
        updaters: [],
        minters: [],
        config: {
          name: "BookNFT Implementation",
          symbol: "BOOKNFTV0",
          metadata: "{}",
          max_supply: 10,
        },
      },
    ],
    {
      initializer: "initialize",
    },
  );
  const bookNFTDeployment = await bookNFT.waitForDeployment();
  let bookNFTAddress = await bookNFTDeployment.getAddress();
  const bookNFTContract = await ethers.getContractAt("BookNFT", bookNFTAddress);

  bookNFTAddress =
    await upgrades.erc1967.getImplementationAddress(bookNFTAddress);

  const likeProtocol = await upgrades.deployProxy(
    LikeProtocol,
    [ownerSigner.address, bookNFTAddress],
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

  return {
    likeProtocol,
    likeProtocolDeployment,
    likeProtocolAddress,
    likeProtocolContract,
    bookNFT,
    bookNFTDeployment,
    bookNFTAddress,
    bookNFTContract,
  };
}
