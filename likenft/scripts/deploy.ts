import "@openzeppelin/hardhat-upgrades";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const LikeNFT = await ethers.getContractFactory("LikeNFT");
  console.log("Deploying LikeNFT... Network:", hardhat.network.name);
  console.log("Owner:", process.env.INITIAL_OWNER_ADDRESS);
  console.log("Expecting Proxy Address:", process.env.ERC721_PROXY_ADDRESS);

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

  const proxyAddress = await likeNFT.getAddress();
  console.log("LikeNFT proxy is deployed to:", proxyAddress);

  const implementationAddress =
    await upgrades.erc1967.getImplementationAddress(proxyAddress);
  console.log("LikeNFT implementation is deployed to:", implementationAddress);

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
