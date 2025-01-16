import { ethers } from "hardhat";

async function getTokenURI() {
  const signer = await ethers.provider.getSigner();

  const Class = await ethers.getContractFactory("Class", {
    signer,
  });

  // Extract and update the class id from newClass's NewClass event
  const class_ = Class.attach("0x14CE6632272552E676b53FE6202edA8F1Be4992c");

  console.log(await class_.tokenURI(0n));
}

getTokenURI().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
