import { ethers } from "hardhat";
import { keccak256, toUtf8Bytes } from "ethers";

async function checkRole() {
  const classId = "0x1D146390C1D4E03C74b87D896b254a5468EDF804";
  const signer = await ethers.provider.getSigner();

  const LikeNFTClass = await ethers.getContractAt("BookNFT", classId);
  const likeNFTClass = LikeNFTClass.connect(signer);

  const role = "UPDATER_ROLE";
  const roleHash = keccak256(toUtf8Bytes(role));

  console.log({
    role,
    roleHash,
    address: signer.address,
    hasRole: await likeNFTClass.hasRole(roleHash, signer.address),
  });

  console.log({
    role,
    roleHash,
    address: classId,
    hasRole: await likeNFTClass.hasRole(
      roleHash,
      "0xaaAC16B7c910f9e5236DC12eF8Db72Abc90D62c4",
    ),
  });
}

checkRole().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
