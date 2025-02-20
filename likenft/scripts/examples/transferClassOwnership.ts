import { ethers } from "hardhat";

async function transferClassOwnership() {
  const classId = "0x84ce8AaB5aceCaE283083761498440539a5DD8dE";
  const newOwner = "0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199";
  const signer = await ethers.provider.getSigner();

  const LikeNFTClass = await ethers.getContractAt("LikeNFTClass", classId);
  const likeNFTClass = LikeNFTClass.connect(signer);

  const tx = await likeNFTClass.transferOwnership(newOwner);
  console.log(await tx.wait());
  console.log(
    "Please update the DEPLOY_WALLET_PRIVATE_KEY in .env to operate on the class afterwards",
  );
}

transferClassOwnership().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
