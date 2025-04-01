import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers } from "hardhat";

async function main() {
  // We get the contract to deploy
  const BookNFT = await ethers.getContractFactory("BookNFT");
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");
  const likeProtocol = LikeProtocol.attach(process.env.ERC721_PROXY_ADDRESS!);

  console.log(
    "Current bookNFT Implementation is:",
    await likeProtocol.implementation(),
  );

  const bookNFT = await BookNFT.deploy({
    creator: process.env.INITIAL_OWNER_ADDRESS!,
    updaters: [],
    minters: [],
    config: {
      name: "BookNFT Implementation",
      symbol: "BOOKNFTV0",
      metadata: "{}",
      max_supply: 10n,
    },
  });

  const newImplementationAddress = await bookNFT.getAddress();
  console.log(
    "New BookNFT implementation is deployed to:",
    newImplementationAddress,
  );

  try {
    await hardhat.run("verify:verify", {
      address: newImplementationAddress,
    });
  } catch (e) {
    if (e instanceof ContractAlreadyVerifiedError) {
      // There may be the same implementation contract verified due to code revert
      console.log(
        "BookNFT new implementation is already verified:",
        newImplementationAddress,
      );
    } else {
      if (hardhat.network.name !== "localhost") {
        throw e;
      }
    }
  }

  // TODO: Prepare an upgrade proposal to safe
  await likeProtocol.upgradeTo(newImplementationAddress);

  console.log(
    "LikeProtocol upgraded BookNFT implementation to:",
    newImplementationAddress,
  );
  console.log(
    "BookNFT Implementation address in LikeProtocol is:",
    await likeProtocol.implementation(),
  );
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
