import "@openzeppelin/hardhat-upgrades";
import hardhat, { ethers, upgrades } from "hardhat";

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
  const proxyAddress = await ekilCoin.getAddress();
  console.log("EkilCoin proxy is deployed to:", proxyAddress);

  const implementationAddress =
    await upgrades.erc1967.getImplementationAddress(proxyAddress);
  console.log("EkilCoin implementation is deployed to:", implementationAddress);

  await hardhat.run("verify:verify", {
    address: implementationAddress,
  });
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
