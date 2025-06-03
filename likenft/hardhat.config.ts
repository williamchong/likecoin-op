import "@nomicfoundation/hardhat-toolbox-viem";
import type { HardhatUserConfig } from "hardhat/config";

import "@nomicfoundation/hardhat-ethers";
import "@nomicfoundation/hardhat-verify";
import "@openzeppelin/hardhat-upgrades";
import "@nomicfoundation/hardhat-ledger";

import dotenv from "dotenv";
dotenv.config({ path: process.env.DOTENV_CONFIG_PATH });
let signerKey =
  process.env.DEPLOY_WALLET_PRIVATE_KEY ||
  "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80";
signerKey = `0x${signerKey}`;

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
      optimism: `${process.env.OPTIMISM_BLOCKSCOUT_KEY}`, // From rickmak.eth account
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
      ledgerAccounts: ["0xB0318A8f049b625dA5DdD184FfFF668Aa6E96261"],
    },
    optimism: {
      url: `${process.env.OPTIMISM_BLOCKSCOUT_URL}`,
      chainId: 10,
      ledgerAccounts: ["0xB0318A8f049b625dA5DdD184FfFF668Aa6E96261"],
    },
    "optimism-sepolia": {
      url: "https://sepolia.optimism.io",
      chainId: 11155420,
      accounts: [signerKey],
    },
    sepolia: {
      url: "https://sepolia.drpc.org",
      chainId: 11155111,
      accounts: [signerKey],
    },
  },
};

export default config;
