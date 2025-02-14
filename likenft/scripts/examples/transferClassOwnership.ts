import { ethers } from "hardhat";

async function mintNFTInClass() {
  const signer = await ethers.provider.getSigner();

  const Class_ = await ethers.getContractFactory("Class", {
    signer,
  });

  // Extract and update the class id from newClass's NewClass event
  const class_ = Class_.attach("0xCafac3dD18aC6c6e92c921884f9E4176737C052c");

  const tx = await class_.transferOwnership(
    "0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199",
  );
  console.log(await tx.wait());
  console.log(
    "Please update the DEPLOY_WALLET_PRIVATE_KEY in .env to operate on the class afterwards",
  );
}

mintNFTInClass().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
