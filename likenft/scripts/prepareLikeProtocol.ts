import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers } from "hardhat";

async function main() {
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");
  const [owner] = await ethers.getSigners();
  console.log("Deployer:", owner.address);
  console.log(
    "Preparing Upgrade of LikeProtocol...",
    process.env.ERC721_PROXY_ADDRESS!,
  );

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
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
