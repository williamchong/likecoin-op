import "@openzeppelin/hardhat-upgrades";
import { ContractAlreadyVerifiedError } from "@nomicfoundation/hardhat-verify/internal/errors";
import hardhat, { ethers, upgrades } from "hardhat";

async function main() {
  // We get the contract to deploy
  const LikeProtocol = await ethers.getContractFactory("LikeProtocol");
  const [owner] = await ethers.getSigners();
  console.log("Operator:", owner.address);
  const proxyAddress = process.env.ERC721_PROXY_ADDRESS!
  console.log("Upgrading LikeProtocol...", proxyAddress);

  // TODO: Prepare an upgrade proposal to safe
  const likeProtocol = LikeProtocol.attach(proxyAddress);
  console.log("Onchain Owner:", await likeProtocol.owner());

  const newImplementationAddress = process.env.NEW_LIKEPROTOCOL;
  const upgradeToAndCallData = process.env.CALLDATA
  await likeProtocol.upgradeToAndCall(
    newImplementationAddress,
    upgradeToAndCallData,
    {
      gasLimit: 1500000,
    },
  );

  const protocolImplementationAddress =
    await upgrades.erc1967.getImplementationAddress(proxyAddress);
  console.log(
    "New onchain LikeProtocol Implementation address:",
    protocolImplementationAddress,
  );
  console.log(
    "New onchain BOOKNFT implementation address is:",
    await likeProtocol.implementation(),
  );
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
