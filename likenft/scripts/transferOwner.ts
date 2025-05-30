import "@openzeppelin/hardhat-upgrades";
import { ethers } from "hardhat";

async function main() {
  const newOwner = process.env.PROTOCOL_OWNER_ADDRESS!;
  console.log("TransferOwnership to", newOwner);

  const proxyAddress = process.env.ERC721_PROXY_ADDRESS!;
  const [operator] = await ethers.getSigners();
  console.log("Operator:", operator.address);

  const LikeProtocol = await ethers.getContractAt("LikeProtocol", proxyAddress);
  console.log("Operating on LikeProtocol at:", proxyAddress);
  const likeProtocol = LikeProtocol.connect(operator);

  console.log("On chain current owner:", await likeProtocol.owner());
  await likeProtocol.transferOwnership(newOwner);
  console.log("New on chain owner:", await likeProtocol.owner());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
