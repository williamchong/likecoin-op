import type { Address } from "viem";

// The vault and reward proxies are deployed with CREATE2 (fixed salt), so their
// addresses are identical on every network — these defaults work on both base
// mainnet and baseSepolia.
export const VELIKE_DEFAULT: Address =
  "0xE55C2b91E688BE70e5BbcEdE3792d723b4766e2B";
export const REWARD_DEFAULT: Address =
  "0x5806B7fb388C2D3D894fE40Ba4598938b01496BA"; // veLikeRewardNoLock proxy

// Deploy block of the veLike vault proxy (veLikeV0Module#ERC1967Proxy) per chain.
// veLIKE cannot exist before this block, so it is the floor for holder discovery.
// The address is the same across chains, but the deploy block is NOT — sourced
// from each network's ignition journal (TRANSACTION_CONFIRM receipt blockNumber).
const VAULT_DEPLOY_BLOCK_BY_CHAIN: Record<number, bigint> = {
  8453: 37568967n, // Base mainnet
  84532: 33038431n, // Base Sepolia
};

export function vaultDeployBlock(chainId: number): bigint {
  const block = VAULT_DEPLOY_BLOCK_BY_CHAIN[chainId];
  if (block === undefined) {
    throw new Error(
      `No vault deploy block known for chainId ${chainId}. Add it to ` +
        `VAULT_DEPLOY_BLOCK_BY_CHAIN in tasks/velikeConstants.ts (the block of ` +
        `veLikeV0Module#ERC1967Proxy on that network).`,
    );
  }
  return block;
}
