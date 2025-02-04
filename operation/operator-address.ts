import { ethers } from "hardhat";

async function getOperatorAddress() {
  const signer = await ethers.provider.getSigner();

  console.log(signer.address);
  const signerETH = await ethers.provider.getBalance(signer.address);
  console.log(signerETH);
}

getOperatorAddress().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
