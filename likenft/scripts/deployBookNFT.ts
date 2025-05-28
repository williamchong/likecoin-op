import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers } from "hardhat";

async function main() {
  const BookNFT = await ethers.getContractFactory("BookNFT");
  const [deployer] = await ethers.getSigners();
  console.log("Deployer:", deployer.address);
  const initOwner = process.env.INITIAL_OWNER_ADDRESS!;
  console.log("Deploying with initial owner with:", initOwner);

  const bookNFT = await BookNFT.deploy();
  await bookNFT.initialize({
    creator: initOwner,
    updaters: [],
    minters: [],
    config: {
      name: "BookNFT Implementation",
      symbol: "BOOKNFTV0",
      metadata: "{}",
      max_supply: 1n,
    },
  });

  const newImplementationAddress = await bookNFT.getAddress();
  console.log(
    "New BookNFT implementation is deployed to:",
    newImplementationAddress,
  );

// Too many time the block-explorer not yet catch the contract, not calling here.
  console.log("Run following to verify after block-explorer catch the deployment")
  console.log(`
BOOKNFT_ADDRESS=${newImplementationAddress} \\\n\
    npm run script:${hardhat.network.name} scripts/verifyBookNFT.ts`)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
