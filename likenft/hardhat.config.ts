import "@nomicfoundation/hardhat-toolbox-viem";
import type { HardhatUserConfig } from "hardhat/config";

import "@nomicfoundation/hardhat-ethers";
import "@nomicfoundation/hardhat-verify";
import "@openzeppelin/hardhat-upgrades";
import "@nomicfoundation/hardhat-ledger";

import dotenv from "dotenv";
dotenv.config({ path: process.env.DOTENV_CONFIG_PATH });

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
      optimism: "d1e693b9-8f5f-42b3-bbe4-3c191cb26c06", // From rickmak.eth account
    },
    customChains: [
      {
        network: "optimism",
        chainId: 10,
        urls: {
          apiURL: "https://optimism.blockscout.com/api",
          browserURL: "https://optimism.blockscout.com/",
        },
      },
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
    localhost: {
      url: "http://127.0.0.1:8545",
      accounts: [`0x${process.env.DEPLOY_WALLET_PRIVATE_KEY}`],
    },
    optimism: {
      url: "https://optimism.drpc.org",
      chainId: 10,
      ledgerAccounts: ["0xB0318A8f049b625dA5DdD184FfFF668Aa6E96261"],
    },
    "optimism-sepolia": {
      url: "https://sepolia.optimism.io",
      chainId: 11155420,
      accounts: [`0x${process.env.DEPLOY_WALLET_PRIVATE_KEY}`],
    },
    sepolia: {
      url: "https://sepolia.drpc.org",
      chainId: 11155111,
      accounts: [`0x${process.env.DEPLOY_WALLET_PRIVATE_KEY}`],
    },
  },
};

export default config;
