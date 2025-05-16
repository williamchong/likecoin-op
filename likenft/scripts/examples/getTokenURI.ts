import { ethers } from "hardhat";

async function getTokenURI() {
  // Extract and update the class id from newBookNFT's NewBookNFT event
  const classId = "0x1D146390C1D4E03C74b87D896b254a5468EDF804";
  const signer = await ethers.provider.getSigner();

  const LikeNFTClass = await ethers.getContractAt("BookNFT", classId);
  const likeNFTClass = LikeNFTClass.connect(signer);

  console.log(await likeNFTClass.getAddress());
  console.log(await likeNFTClass.name());
  console.log(await likeNFTClass.symbol());
  console.log(await likeNFTClass.contractURI());
  console.log(await likeNFTClass.tokenURI(0n));
}

getTokenURI().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
