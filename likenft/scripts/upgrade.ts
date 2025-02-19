import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");
  console.log("Upgrading LikeProtocol...");

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
      if (network.name !== "localhost") {
        throw e;
      }
    }
  }

  // TODO: Prepare an upgrade proposal to safe
  const likeProtocol = LikeProtocol.attach(process.env.ERC721_PROXY_ADDRESS!);
  await likeProtocol.upgradeToAndCall(newImplementationAddress, "0x");

  console.log("LikeProtocol upgraded implementation to:", newImplementationAddress);
  console.log("LikeProtocol proxy address is:", await likeProtocol.getAddress());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
