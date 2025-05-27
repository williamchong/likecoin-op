import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers } from "hardhat";

async function main() {
  const bookNFT = await ethers.getContractAt("BookNFT", process.env.BOOKNFT_ADDRESS!);
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");
  const likeProtocol = LikeProtocol.attach(process.env.ERC721_PROXY_ADDRESS!);

  console.log(
    "Current bookNFT Implementation is:",
    await likeProtocol.implementation(),
  );
  const bookNFTImplementationAddress = await bookNFT.getAddress();
  console.log(
    "Target bookNFT Implementation is:",
    bookNFTImplementationAddress,
  );

  await likeProtocol.upgradeTo(bookNFTImplementationAddress);

  console.log(
    "Latest BookNFT Implementation address in LikeProtocol is:",
    await likeProtocol.implementation(),
  );
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
