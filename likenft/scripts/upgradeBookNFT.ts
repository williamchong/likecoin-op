import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers } from "hardhat";

async function main() {
  const [operator] = await ethers.getSigners();
  console.log("Operator:", operator.address);

  const bookNFT = await ethers.getContractAt(
    "BookNFT",
    process.env.BOOKNFT_ADDRESS!,
  );
  // Query once to ensure contract exist on target chain
  const bookNFTImplementationAddress = await bookNFT.getAddress();
  console.log(
    "Target bookNFT Implementation is:",
    bookNFTImplementationAddress,
  );

  const proxyAddress = process.env.ERC721_PROXY_ADDRESS!;
  const likeProtocol = await ethers.getContractAt("LikeProtocol", proxyAddress);
  console.log(
    "Operating on LikeProtocol at:",
    proxyAddress,
  );

  console.log(
    "Current bookNFT Implementation is:",
    await likeProtocol.implementation(),
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
