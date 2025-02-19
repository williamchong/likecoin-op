import "@openzeppelin/hardhat-upgrades";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");
  console.log("Deploying LikeProtocol... Network:", hardhat.network.name);
  console.log("Owner:", process.env.INITIAL_OWNER_ADDRESS);
  console.log("Expecting Proxy Address:", process.env.ERC721_PROXY_ADDRESS);

  const likeProtocol = await upgrades.deployProxy(
    LikeProtocol,
    [process.env.INITIAL_OWNER_ADDRESS!],
    {
      initializer: "initialize",
      timeout: 0,
      verifySourceCode: true,
    },
  );

  await likeProtocol.waitForDeployment();

  const proxyAddress = await likeProtocol.getAddress();
  console.log("LikeProtocol proxy is deployed to:", proxyAddress);

  const implementationAddress =
    await upgrades.erc1967.getImplementationAddress(proxyAddress);
  console.log("LikeProtocol implementation is deployed to:", implementationAddress);

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
