import type { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox-viem";

import "@nomicfoundation/hardhat-ethers";
import "@openzeppelin/hardhat-upgrades";
import dotenv from "dotenv";
dotenv.config();

const config: HardhatUserConfig = {
  solidity: "0.8.28",
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts: [`0x${process.env.DEPLOY_WALLET_PRIVATE_KEY_LOCALHOST}`],
    },
    "optimism-sepolia": {
      url: "https://sepolia.optimism.io",
      chainId: 11155420,
      accounts: [`0x${process.env.DEPLOY_WALLET_PRIVATE_KEY_OP_SEPOLIA}`],
    },
  },
};

export default config;
