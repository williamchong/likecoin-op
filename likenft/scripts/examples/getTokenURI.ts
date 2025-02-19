import { ethers } from "hardhat";

async function getTokenURI() {
  // Extract and update the class id from newClass's NewClass event
  const classId = "0x84ce8AaB5aceCaE283083761498440539a5DD8dE";
  const signer = await ethers.provider.getSigner();

  const LikeNFTClass = await ethers.getContractAt(
    "LikeNFTClass",
    classId,
  );
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
