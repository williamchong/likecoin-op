import "@openzeppelin/hardhat-upgrades";
import { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const EkilCoin = await ethers.getContractFactory("EkilCoin");
  console.log("Deploying EkilCoin...");

  const ekilCoin = await upgrades.deployProxy(
    EkilCoin,
    [process.env.INITIAL_OWNER_ADDRESS, process.env.INITIAL_MINTER_ADDRESS],
    {
      initializer: "initialize",
      timeout: 0,
      verifySourceCode: true,
    },
  );

  await ekilCoin.waitForDeployment();
  console.log("EkilCoin deployed to:", await ekilCoin.getAddress());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
