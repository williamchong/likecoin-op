

import { ethers } from "hardhat";

async function fundOperator() {
  console.log("Funding operator");
  const funderKey = process.env.LOCALHOST_FUNDING_WALLET_PRIVATE_KEY;
  if (!funderKey) {
    throw new Error("LOCALHOST_FUNDING_WALLET_PRIVATE_KEY is not set");
  }
  const wallet = new ethers.Wallet(funderKey, ethers.provider);

  const signer = await ethers.provider.getSigner();

  const operatorAddress = process.env.OPERATOR_ADDRESS || signer.address;
  if (!operatorAddress) {
    throw new Error("OPERATOR_ADDRESS is not set");
  }
  console.log("Funding operator address", operatorAddress);
  const payload = {
    to: operatorAddress,
    value: ethers.parseEther("12.0"), // Sending 1 ETH
  };
  const tx = await wallet.sendTransaction(payload);
  console.log("Transaction hash", tx.hash);
  const receipt = await tx.wait();
  console.log("Transaction receipt", receipt);

  const signerETH = await ethers.provider.getBalance(signer.address);
  console.log("Operator address balance", signerETH);
}

export default fundOperator;