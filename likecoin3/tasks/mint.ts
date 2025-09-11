import { task } from "hardhat/config";
import fs from "fs";

task("mint", "Mint tokens")
  .addParam("likecoin", "The address of the likecoin contract")
  .addParam("amount", "The amount of tokens to mint")
  .addParam("to", "The address to mint tokens to")
  .setAction(async ({ likecoin, amount, to }, { ethers, network }) => {
    const [operator] = await ethers.getSigners();
    console.log("Operator:", operator.address);
    console.log("Likecoin address:", likecoin);

    const LikecoinContract = await ethers.getContractAt("Likecoin", likecoin);
    const likecoinContract = LikecoinContract.connect(operator);

    const mintAmount = ethers.parseUnits(
      amount,
      await likecoinContract.decimals(),
    );
    console.log("Mint amount:", mintAmount);

    const tx = await likecoinContract.mint(to, mintAmount);
    await tx.wait();
  });
