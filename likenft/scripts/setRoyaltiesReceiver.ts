import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  const [operator] = await ethers.getSigners();
  console.log("Operator:", operator.address);

  const proxyAddress = process.env.ERC721_PROXY_ADDRESS!;
  const LikeProtocol = await ethers.getContractAt("LikeProtocol", proxyAddress);
  console.log(
    "Operating on LikeProtocol at:",
    proxyAddress,
  );
  const likeProtocol = LikeProtocol.connect(operator);

  console.log("On chain royalty receiver:", await likeProtocol.getRoyaltyReceiver());
  await likeProtocol.setRoyaltyReceiver(process.env.RECEIVER);
  console.log("New royalty receiver:", await likeProtocol.getRoyaltyReceiver());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
