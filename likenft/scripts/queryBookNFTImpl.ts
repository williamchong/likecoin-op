import { ethers, upgrades } from "hardhat";

async function getTokenURI() {
  const proxyAddress = process.env.ERC721_PROXY_ADDRESS!;
  const signer = await ethers.provider.getSigner();

  const LikeProtocol = await ethers.getContractAt("LikeProtocol", proxyAddress);
  const likeProtocol = LikeProtocol.connect(signer);

  const implementation = await likeProtocol.implementation();
  console.log("Implementation address:", implementation);
}

getTokenURI().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
