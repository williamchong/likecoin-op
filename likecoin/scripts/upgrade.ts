import "@openzeppelin/hardhat-upgrades";
import { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const EkilCoin = await ethers.getContractFactory("EkilCoin");
  console.log("Upgrading EkilCoin...");
  const ekilCoin = await upgrades.upgradeProxy(
    process.env.PROXY_ADDRESS!,
    EkilCoin,
    {
      timeout: 0,
    },
  );
  console.log("EkilCoin deployed to:", await ekilCoin.getAddress());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
