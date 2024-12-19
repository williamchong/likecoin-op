import "@openzeppelin/hardhat-upgrades";
import { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const LikeNFT = await ethers.getContractFactory("LikeNFT");
  console.log("Upgrading LikeNFT...");
  const likeNFT = await upgrades.upgradeProxy(
    process.env.PROXY_ADDRESS!,
    LikeNFT,
    {
      timeout: 0,
    },
  );
  console.log("LikeNFT deployed to:", await likeNFT.getAddress());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
