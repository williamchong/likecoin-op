import type { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox-viem";
import "@openzeppelin/hardhat-upgrades";

import dotenv from "dotenv";
dotenv.config({ path: process.env.DOTENV_CONFIG_PATH || ".env.local" });
const alchemyApiKey = process.env.ALCHEMY_API_KEY;
let signerKey = process.env.DEPLOY_WALLET_PRIVATE_KEY ||
  "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80";
signerKey = `0x${signerKey}`;
const create2Salt = process.env.CREATE2_SALT || "0x0000000000000000000000000000000000000000000000000000000000000000"

const config: HardhatUserConfig = {
  solidity: "0.8.28",
  networks: {
    localhost: {
      url: "http://localhost:8545",
      chainId: 900,
      accounts: [signerKey],
    },
    superism1: {
      chainId: 901,
      url: "http://localhost:9545",
      accounts: [signerKey],
    },
    superism2: {
      chainId: 902,
      url: "http://localhost:9546",
      accounts: [signerKey],
    },
    mainnet: {
      chainId: 1,
      url: `https://eth-mainnet.g.alchemy.com/v2/${alchemyApiKey}`,
      accounts: [signerKey],
    },
    sepolia: {
      chainId: 11155111,
      url: `https://eth-sepolia.g.alchemy.com/v2/${alchemyApiKey}`,
      accounts: [signerKey],
    },
    optimism: {
      chainId: 10,
      url: `https://opt-mainnet.g.alchemy.com/v2/${alchemyApiKey}`,
      accounts: [signerKey],
    },
    optimismSepolia: {
      chainId: 11155420,
      url: `https://opt-sepolia.g.alchemy.com/v2/${alchemyApiKey}`,
      accounts: [signerKey],
    },
    base: {
      chainId: 8453,
      url: `https://base-mainnet.g.alchemy.com/v2/${alchemyApiKey}`,
      accounts: [signerKey],
    },
    baseSepolia: {
      chainId: 84532,
      url: `https://base-sepolia.g.alchemy.com/v2/${alchemyApiKey}`,
      accounts: [signerKey],
    },
    unichain: {
      chainId: 130,
      url: `https://unichain-mainnet.g.alchemy.com/v2/${alchemyApiKey}`,
      accounts: [signerKey],
    },
    unichainSepolia: {
      chainId: 1301,
      url: `https://unichain-sepolia.g.alchemy.com/v2/${alchemyApiKey}`,
      accounts: [signerKey],
    },
  },
  ignition: {
    strategyConfig: {
      create2: {
        salt: create2Salt,
      },
    },
  },
};

export default config;
