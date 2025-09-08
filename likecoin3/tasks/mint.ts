import { task } from "hardhat/config";
import fs from "fs";

task("mint", "Mint tokens")
  .addParam("amount", "The amount of tokens to mint")
  .addParam("to", "The address to mint tokens to")
  .setAction(async ({ amount, to }, { ethers, network }) => {
    const [operator] = await ethers.getSigners();
    console.log("Operator:", operator.address);

    let likecoinAddress = "";
    if (
      fs.existsSync(
        `ignition/deployments/chain-${network.config.chainId}/deployed_addresses.json`,
      )
    ) {
      const deployedAddresses = JSON.parse(
        fs.readFileSync(
          `ignition/deployments/chain-${network.config.chainId}/deployed_addresses.json`,
          "utf8",
        ),
      );
      likecoinAddress = deployedAddresses["LikecoinModule#Likecoin"];
    } else {
      throw new Error(
        `Deployed addresses not found for chain ${network.config.chainId}`,
      );
    }

    console.log("Likecoin address:", likecoinAddress);

    const Likecoin = await ethers.getContractAt("Likecoin", likecoinAddress);
    const likecoin = Likecoin.connect(operator);

    const mintAmount = ethers.parseUnits(amount, await likecoin.decimals());
    console.log("Mint amount:", mintAmount);

    const tx = await likecoin.mint(to, mintAmount);
    await tx.wait();
  });
