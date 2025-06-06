import { ethers } from "hardhat";

async function getRoyaltyInfo() {
  const classId = "0x1D146390C1D4E03C74b87D896b254a5468EDF804";
  const signer = await ethers.provider.getSigner();

  const LikeNFTClass = await ethers.getContractAt("BookNFT", classId);
  const likeNFTClass = LikeNFTClass.connect(signer);

  console.log(await likeNFTClass.royaltyInfo(1n, 1000000n));
}

getRoyaltyInfo().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
