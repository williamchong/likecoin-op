import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers } from "hardhat";

async function main() {
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");
  const [deployer] = await ethers.getSigners();
  console.log("Deployer:", deployer.address);
  const proxyAddress = process.env.ERC721_PROXY_ADDRESS!
  console.log(
    "Preparing Upgrade of LikeProtocol...",
    proxyAddress,
  );

  const newImplementationAddress = await upgrades.prepareUpgrade(
    proxyAddress,
    LikeProtocol,
    {
      timeout: 0,
      verifySourceCode: true,
      kind: "uups",
      redeployImplementation: "always"
    },
  );
  console.log(
    "LikeProtocol new implementation is deployed to:",
    newImplementationAddress,
  );

  // Too many time the block-explorer not yet catch the contract, not calling here.
  console.log("Run following to verify after block-explorer catch the deployment")
  console.log(`
NEW_LIKEPROTOCOL=${newImplementationAddress} \\\n\
    npm run script:${hardhat.network.name} scripts/verifyLikeProtocol.ts`)
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
