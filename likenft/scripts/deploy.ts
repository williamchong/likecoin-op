import "@openzeppelin/hardhat-upgrades";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const BookNFT = await ethers.getContractFactory("BookNFT");
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");

  console.log("Deploying LikeProtocol... Network:", hardhat.network.name);
  console.log("Owner:", process.env.INITIAL_OWNER_ADDRESS);
  console.log("Expecting Proxy Address:", process.env.ERC721_PROXY_ADDRESS);

  const bookNFT = await BookNFT.deploy({
    creator: process.env.INITIAL_OWNER_ADDRESS!,
    updaters: [],
    minters: [],
    config: {
      name: "BookNFT Implementation",
      symbol: "BOOKNFTV0",
      metadata: "{}",
      max_supply: 10n,
    },
  });

  const bookNFTAddress = await bookNFT.getAddress();
  console.log("BookNFT implementation is deployed to:", bookNFTAddress);

  const likeProtocol = await upgrades.deployProxy(
    LikeProtocol,
    [process.env.INITIAL_OWNER_ADDRESS!, bookNFTAddress],
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
  console.log(
    "LikeProtocol implementation is deployed to:",
    implementationAddress,
  );
  console.log(
    "LikeProtocol using BookNFT implementation at:",
    await likeProtocol.implementation(),
  );

  if (hardhat.network.name === "localhost") {
    console.log("Skipping verification on localhost");
    return;
  }

  await hardhat.run("verify:verify", {
    address: bookNFTAddress,
  });
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
