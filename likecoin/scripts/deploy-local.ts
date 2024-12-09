import { ethers } from "hardhat";

async function main() {
  // We get the contract to deploy
  const EkilCoin = await ethers.getContractFactory("EkilCoin");
  console.log("Deploying EkilCoin...");
  const ekilCoin = await EkilCoin.deploy(1000000, process.env.GIT_HASH);
  await ekilCoin.waitForDeployment();
  console.log("EkilCoin deployed to:", await ekilCoin.getAddress());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
