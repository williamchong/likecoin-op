import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const LikeNFT = await ethers.getContractFactory("LikeNFT");
  console.log("Upgrading LikeNFT...");

  const newImplementationAddress = await upgrades.prepareUpgrade(
    process.env.ERC721_PROXY_ADDRESS!,
    LikeNFT,
    {
      timeout: 0,
      verifySourceCode: true,
      kind: "uups",
    },
  );

  console.log(
    "LikeNFT new implementation is deployed to:",
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
        "LikeNFT new implementation is already verified:",
        newImplementationAddress,
      );
    } else {
      throw e;
    }
  }

  // TODO: Prepare an upgrade proposal to safe
  const likeNFT = LikeNFT.attach(process.env.ERC721_PROXY_ADDRESS!);
  await likeNFT.upgradeToAndCall(newImplementationAddress, "0x");

  console.log("LikeNFT upgraded implementation to:", newImplementationAddress);
  console.log("LikeNFT proxy address is:", await likeNFT.getAddress());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
