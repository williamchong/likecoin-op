import "@openzeppelin/hardhat-upgrades";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const EkilCoin = await ethers.getContractFactory("EkilCoin");
  console.log("Deploying EkilCoin... Network:", hardhat.network.name);
  console.log("Owner:", process.env.INITIAL_OWNER_ADDRESS);
  console.log("Minter:", process.env.INITIAL_MINTER_ADDRESS);
  console.log("Expecting Proxy Address:", process.env.ERC20_PROXY_ADDRESS);

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

  if (hardhat.network.name === "localhost") {
    console.log("Skipping verification on localhost");
    return;
  }
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
