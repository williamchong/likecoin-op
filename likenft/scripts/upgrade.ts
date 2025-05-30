import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  const [owner] = await ethers.getSigners();
  console.log("Operator:", owner.address);
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");
  const protocolAddress = process.env.ERC721_PROXY_ADDRESS!;
  console.log("Upgrading LikeProtocol...", protocolAddress);

  const newImplementationAddress = await upgrades.prepareUpgrade(
    protocolAddress,
    LikeProtocol,
    {
      timeout: 0,
      verifySourceCode: true,
      kind: "uups",
      redeployImplementation: "always",
    },
  );

  console.log(
    "LikeProtocol new implementation is deployed to:",
    newImplementationAddress,
  );

  console.log("Deploying BookNFT...");
  const BookNFT = await ethers.getContractFactory("BookNFT");
  const bookNFT = await BookNFT.deploy();
  await bookNFT.initialize({
    creator: process.env.INITIAL_OWNER_ADDRESS!,
    updaters: [],
    minters: [],
    config: {
      name: "BookNFT Implementation",
      symbol: "BOOKNFTV0",
      metadata: "{}",
      max_supply: 1n,
    },
  });
  const newBookNFTImplementationAddress = await bookNFT.getAddress();
  console.log(
    "New BookNFT implementation is deployed to:",
    newBookNFTImplementationAddress,
  );

  const upgradeToFunctionValues = [newBookNFTImplementationAddress];
  const upgradeToFunctionFragment = LikeProtocol.interface.getFunction(
    "upgradeTo",
    upgradeToFunctionValues,
  );
  if (upgradeToFunctionFragment == null) {
    throw new Error("upgradeTo function not found");
  }
  const upgradeToAndCallData = LikeProtocol.interface.encodeFunctionData(
    upgradeToFunctionFragment,
    upgradeToFunctionValues,
  );

  const likeProtocol = LikeProtocol.attach(protocolAddress);
  console.log("On chain owner:", await likeProtocol.owner());
  await likeProtocol.upgradeToAndCall(
    newImplementationAddress,
    upgradeToAndCallData,
    {
      gasLimit: 1500000,
    },
  );

  const protocolImplementationAddress =
    await upgrades.erc1967.getImplementationAddress(proxyAddress);
  console.log(
    "New onchain LikeProtocol Implementation address:",
    protocolImplementationAddress,
  );

  if (hardhat.network.name === "localhost") {
    console.log("Skipping verification on localhost");
    return;
  }

  console.log("Verifying on block-explorer...");
  try {
    await hardhat.run("verify:verify", {
      address: newImplementationAddress,
    });
  } catch (e) {
    if (e instanceof ContractAlreadyVerifiedError) {
      // There may be the same implementation contract verified due to code revert
      console.log(
        "LikeProtocol new implementation is already verified:",
        newImplementationAddress,
      );
    } else {
      if (hardhat.network.name !== "localhost") {
        throw e;
      }
    }
  }

  try {
    await hardhat.run("verify:verify", {
      address: newBookNFTImplementationAddress,
    });
  } catch (e) {
    if (e instanceof ContractAlreadyVerifiedError) {
      // There may be the same implementation contract verified due to code revert
      console.log(
        "BookNFT new implementation is already verified:",
        newBookNFTImplementationAddress,
      );
    } else {
      if (hardhat.network.name !== "localhost") {
        throw e;
      }
    }
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
