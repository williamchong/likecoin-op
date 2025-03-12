import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const newOwner = "0xC71fe89e4C0e5458a793fc6548EF6B392417A7Fb";
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");
  const [operator] = await ethers.getSigners();
  console.log("Operator:", operator.address);

  // TODO: Prepare an upgrade proposal to safe
  const likeProtocol = LikeProtocol.attach(process.env.ERC721_PROXY_ADDRESS!);
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
