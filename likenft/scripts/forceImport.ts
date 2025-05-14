import { ethers, upgrades } from "hardhat";

async function main() {
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");

  await upgrades.forceImport(process.env.ERC721_PROXY_ADDRESS!, LikeProtocol);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
