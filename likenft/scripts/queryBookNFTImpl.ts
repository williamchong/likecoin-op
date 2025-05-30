import { ethers, upgrades } from "hardhat";

async function getTokenURI() {
  const [operator] = await ethers.getSigners();
  console.log("Operator:", operator.address);

  const proxyAddress = process.env.ERC721_PROXY_ADDRESS!;
  const LikeProtocol = await ethers.getContractAt("LikeProtocol", proxyAddress);
  console.log("Operating on LikeProtocol at:", proxyAddress);
  const likeProtocol = LikeProtocol.connect(operator);

  const owner = await likeProtocol.owner();
  console.log("Protocol owner:", owner);

  const protocolImplementationAddress =
    await upgrades.erc1967.getImplementationAddress(proxyAddress);
  console.log(
    "LikeProtocol Implementation address:",
    protocolImplementationAddress,
  );

  const implementation = await likeProtocol.implementation();
  console.log("BookNFT Implementation address:", implementation);
}

getTokenURI().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
