import { ethers } from "hardhat";

async function getTokenURI() {
  const signer = await ethers.provider.getSigner();

  const Class = await ethers.getContractFactory("Class", {
    signer,
  });

  const class_ = Class.attach("0x7b43F8aB14A983FF3D8831081F62370a14967f7B");

  console.log(await class_.tokenURI(0n));
}

getTokenURI().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
