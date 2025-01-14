import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const EkilCoin = await ethers.getContractFactory("EkilCoin");
  console.log("Upgrading EkilCoin...");

  const newImplementationAddress = await upgrades.prepareUpgrade(
    process.env.ERC20_PROXY_ADDRESS!,
    EkilCoin,
    {
      timeout: 0,
      verifySourceCode: true,
      kind: "uups",
    },
  );

  console.log(
    "EkilCoin new implementation is deployed to:",
    newImplementationAddress,
  );

  try {
    await hardhat.run("verify:verify", {
      address: newImplementationAddress,
    });
  } catch (e) {
    if (e instanceof ContractAlreadyVerifiedError) {
      // There may be the same implementation contract verified due to code revert
      console.log(
        "LikeNFT new implementation is already verified:",
        newImplementationAddress,
      );
    } else {
      throw e;
    }
  }

  // TODO: Prepare an upgrade proposal to safe
  const ekilCoin = EkilCoin.attach(process.env.PROXY_ADDRESS!);
  await ekilCoin.upgradeToAndCall(newImplementationAddress, "0x");

  console.log("EkilCoin upgraded implementation to:", newImplementationAddress);
  console.log("EkilCoin proxy address is:", await ekilCoin.getAddress());
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
