import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers } from "hardhat";

async function main() {
  console.log(
    "Trying to verify BookNFT implementation at:",
    process.env.BOOKNFT_ADDRESS!,
  );
  const bookNFT = await ethers.getContractAt(
    "BookNFT",
    process.env.BOOKNFT_ADDRESS!,
  );
  const bookNFTAddress = await bookNFT.getAddress();

  try {
    await hardhat.run("verify:verify", {
      address: bookNFTAddress,
    });
  } catch (e) {
    if (e instanceof ContractAlreadyVerifiedError) {
      // There may be the same implementation contract verified due to code revert
      console.log(
        "BookNFT new implementation is already verified:",
        bookNFTAddress,
      );
    } else {
      if (hardhat.network.name !== "localhost") {
        throw e;
      }
    }
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
