import "@openzeppelin/hardhat-upgrades";
import { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const LikeNFT = await ethers.getContractFactory("LikeNFT");
  console.log("Deploying LikeNFT...");

  const likeNFT = await upgrades.deployProxy(
    LikeNFT,
    [process.env.INITIAL_OWNER_ADDRESS!],
    {
      initializer: "initialize",
      timeout: 0,
      verifySourceCode: true,
    },
  );

  await likeNFT.waitForDeployment();
  console.log("LikeNFT deployed to:", await likeNFT.getAddress());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
