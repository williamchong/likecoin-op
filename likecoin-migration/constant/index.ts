export function LIKECOIN_CHAIN_ID(isTestnet: boolean): string {
  return isTestnet ? 'likecoin-public-testnet-5' : 'likecoin-mainnet-2';
}

export function LIKECOIN_CHAIN_NAME(isTestnet: boolean): string {
  return isTestnet ? 'LikeCoin public test chain' : 'LikeCoin';
}

export function LIKECOIN_CHAIN_NFT_RPC(isTestnet: boolean): string {
  return isTestnet
    ? 'https://node.testnet.like.co/rpc/'
    : 'https://mainnet-node.like.co/rpc/';
}

export function LIKECOIN_CHAIN_API(isTestnet: boolean): string {
  return isTestnet
    ? 'https://node.testnet.like.co'
    : 'https://mainnet-node.like.co';
}

export function LIKECOIN_CHAIN_DENOM(isTestnet: boolean): string {
  return isTestnet ? 'EKIL' : 'LIKE';
}

export function LIKECOIN_CHAIN_MIN_DENOM(isTestnet: boolean): string {
  return isTestnet ? 'nanoekil' : 'nanolike';
}

export function LIKECOIN_CHAIN_GECKO_ID(isTestnet: boolean): string {
  return isTestnet ? '' : 'likecoin';
}

export function CIVIC_LIKER_V3_STAKING_ENDPOINT(isTestnet: boolean): string {
  return isTestnet
    ? 'https://likecoin-public-testnet-5.netlify.app/validators'
    : 'https://dao.like.co/validators';
}

export function AUTHCORE_API_HOST(isTestnet: boolean): string {
  return isTestnet
    ? 'https://likecoin-integration-test.authcore.io'
    : 'https://authcore.like.co';
}

export function EVM_CHAIN_ID(isTestnet: boolean): number {
  return isTestnet ? 11155420 : 10;
}

export function EVM_RPC_URL(isTestnet: boolean): string {
  return isTestnet
    ? 'https://sepolia.optimism.io'
    : 'https://mainnet.optimism.io';
}

export function EVM_MAGIC_LINK_API_KEY(isTestnet: boolean): string {
  return isTestnet ? 'pk_live_5E14E3184484268D' : 'pk_live_583D0D54D78940DA';
}
