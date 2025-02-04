import { ethers } from "hardhat";
import fundOperator from "./src/fund-operator";

async function main() {
  await fundOperator();
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
