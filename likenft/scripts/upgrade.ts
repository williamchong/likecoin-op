import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");
  const [owner] = await ethers.getSigners();
  console.log("Owner:", owner.address);
  console.log("Upgrading LikeProtocol...", process.env.ERC721_PROXY_ADDRESS!);

  const newImplementationAddress = await upgrades.prepareUpgrade(
    process.env.ERC721_PROXY_ADDRESS!,
    LikeProtocol,
    {
      timeout: 0,
      verifySourceCode: true,
      kind: "uups",
    },
  );

  console.log(
    "LikeProtocol new implementation is deployed to:",
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
        "LikeProtocol new implementation is already verified:",
        newImplementationAddress,
      );
    } else {
      if (hardhat.network.name !== "localhost") {
        throw e;
      }
    }
  }

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

  // TODO: Prepare an upgrade proposal to safe
  const likeProtocol = LikeProtocol.attach(process.env.ERC721_PROXY_ADDRESS!);
  console.log("Owner:", await likeProtocol.owner());
  await likeProtocol.upgradeToAndCall(
    newImplementationAddress,
    upgradeToAndCallData,
    {
      gasLimit: 1500000,
    },
  );

  console.log(
    "LikeProtocol upgraded implementation to:",
    newImplementationAddress,
  );
  console.log(
    "LikeProtocol proxy address is:",
    await likeProtocol.getAddress(),
  );
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
