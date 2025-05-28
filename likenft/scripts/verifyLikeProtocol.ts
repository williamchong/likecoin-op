import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers } from "hardhat";

async function main() {
  if (hardhat.network.name == "localhost") {
    throw "No verification at local network";
  }
  console.log(
    "Trying to verify LikeProtocol implementation at:",
    process.env.NEW_LIKEPROTOCOL!,
  );
  const likeProtocol = await ethers.getContractAt(
    "LikeProtocol",
    process.env.NEW_LIKEPROTOCOL!,
  );
  // Verify it already on chain
  const protcolAddress = await likeProtocol.getAddress();

  try {
    await hardhat.run("verify:verify", {
      address: protcolAddress,
    });
  } catch (e) {
    if (e instanceof ContractAlreadyVerifiedError) {
      // There may be the same implementation contract verified due to code revert
      console.log(
        "LikeProtocol new implementation is already verified:",
        protcolAddress,
      );
    }
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
