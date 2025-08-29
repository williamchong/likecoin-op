import type { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox-viem";

const config: HardhatUserConfig = {
  solidity: "0.8.28",
  networks: {
    localhost: {
      url: "http://localhost:8545",
      chainId: 900,
    },
    superism1: {
      chainId: 901,
      url: "http://localhost:9545",
    },
    superism2: {
      chainId: 902,
      url: "http://localhost:9546",
    },
  },
};

export default config;
