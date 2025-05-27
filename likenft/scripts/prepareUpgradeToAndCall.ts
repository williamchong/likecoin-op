import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");
  const newBookNFTImplementationAddress = process.env.BOOKNFT_ADDRESS!;
  console.log(
    "Target new BookNFT implementation address:",
    newBookNFTImplementationAddress,
  );

  const upgradeToFunctionValues = [newBookNFTImplementationAddress];
  const upgradeToFunctionFragment = LikeProtocol.interface.getFunction(
    "upgradeTo",
    upgradeToFunctionValues,
  );
  if (upgradeToFunctionFragment == null) {
    throw new Error("upgradeTo function not found");
  }
  const upgradeToAndCallData = LikeProtocol.interface.encodeFunctionData(
    upgradeToFunctionFragment,
    upgradeToFunctionValues,
  );
  console.log("Upgrade to and call data:", upgradeToAndCallData);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
