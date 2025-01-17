import "@nomicfoundation/hardhat-toolbox-viem";
import type { HardhatUserConfig } from "hardhat/config";

import "@nomicfoundation/hardhat-ethers";
import "@nomicfoundation/hardhat-verify";
import "@openzeppelin/hardhat-upgrades";
import dotenv from "dotenv";
dotenv.config();

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.28",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
        details: {
          yulDetails: {
            optimizerSteps: "u",
          },
        },
      },
      viaIR: true,
    },
  },
  etherscan: {
    apiKey: {
      "optimism-sepolia":
        "Is not required by blockscout. Can be any non-empty string",
      sepolia: "Is not required by blockscout. Can be any non-empty string",
    },
    customChains: [
      {
        network: "sepolia",
        chainId: 11155111,
        urls: {
          apiURL: "https://eth-sepolia.blockscout.com/api",
          browserURL: "https://eth-sepolia.blockscout.com/",
        },
      },
      {
        network: "optimism-sepolia",
        chainId: 11155420,
        urls: {
          apiURL: "https://optimism-sepolia.blockscout.com/api",
          browserURL: "https://optimism-sepolia.blockscout.com/",
        },
      },
    ],
  },
  networks: {
    ...(process.env.DEPLOY_WALLET_PRIVATE_KEY_LOCALHOST != null
      ? {
          localhost: {
            url: "http://127.0.0.1:8545",
            accounts: [`0x${process.env.DEPLOY_WALLET_PRIVATE_KEY_LOCALHOST}`],
          },
        }
      : {}),
    ...(process.env.DEPLOY_WALLET_PRIVATE_KEY_OP_SEPOLIA != null
      ? {
          sepolia: {
            url: "https://sepolia.drpc.org",
            chainId: 11155111,
            accounts: [`0x${process.env.DEPLOY_WALLET_PRIVATE_KEY_OP_SEPOLIA}`],
          },
        }
      : {}),
    ...(process.env.DEPLOY_WALLET_PRIVATE_KEY_OP_SEPOLIA != null
      ? {
          "optimism-sepolia": {
            url: "https://sepolia.optimism.io",
            chainId: 11155420,
            accounts: [`0x${process.env.DEPLOY_WALLET_PRIVATE_KEY_OP_SEPOLIA}`],
          },
        }
      : {}),
  },
};

export default config;
