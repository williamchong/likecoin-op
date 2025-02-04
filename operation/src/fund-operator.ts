

import { ethers } from "hardhat";

async function fundOperator() {
  console.log("Funding operator");
  const funderKey = process.env.LOCALHOST_FUNDING_WALLET_PRIVATE_KEY;
  if (!funderKey) {
    throw new Error("LOCALHOST_FUNDING_WALLET_PRIVATE_KEY is not set");
  }
  const wallet = new ethers.Wallet(funderKey, ethers.provider);

  const signer = await ethers.provider.getSigner();

  console.log("Funding operator address", signer.address);
  const payload = {
    to: signer.address,
    value: ethers.parseEther("1.0"), // Sending 1 ETH
  };
  const tx = await wallet.sendTransaction(payload);
  console.log("Transaction hash", tx.hash);
  const receipt = await tx.wait();
  console.log("Transaction receipt", receipt);

  const signerETH = await ethers.provider.getBalance(signer.address);
  console.log("Operator address balance", signerETH);
}

export default fundOperator;