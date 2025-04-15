import type { LikeCoinEVMWalletConnectorOptions } from '@likecoin/evm-wallet-connector/dist/index';
import {
  LikeCoinWalletConnectorConfig,
  LikeCoinWalletConnectorMethodType,
} from '@likecoin/wallet-connector';

import {
  AUTHCORE_API_HOST,
  CIVIC_LIKER_V3_STAKING_ENDPOINT,
  EVM_CHAIN_ID,
  EVM_MAGIC_LINK_API_KEY,
  EVM_RPC_URL,
  LIKECOIN_CHAIN_API,
  LIKECOIN_CHAIN_DENOM,
  LIKECOIN_CHAIN_GECKO_ID,
  LIKECOIN_CHAIN_ID,
  LIKECOIN_CHAIN_MIN_DENOM,
  LIKECOIN_CHAIN_NAME,
  LIKECOIN_CHAIN_NFT_RPC,
} from '.';

export function LIKECOIN_WALLET_CONNECTOR_CONFIG(
  isTestnet: boolean,
  authcoreRedirectUrl: string
): LikeCoinWalletConnectorConfig {
  return {
    chainId: LIKECOIN_CHAIN_ID(isTestnet),
    chainName: LIKECOIN_CHAIN_NAME(isTestnet),
    rpcURL: LIKECOIN_CHAIN_NFT_RPC(isTestnet),
    restURL: LIKECOIN_CHAIN_API(isTestnet),
    coinType: 118,
    coinDenom: LIKECOIN_CHAIN_DENOM(isTestnet),
    coinMinimalDenom: LIKECOIN_CHAIN_MIN_DENOM(isTestnet),
    coinDecimals: 9,
    coinGeckoId: LIKECOIN_CHAIN_GECKO_ID(isTestnet),
    walletURLForStaking: CIVIC_LIKER_V3_STAKING_ENDPOINT(isTestnet),
    bech32PrefixAccAddr: 'like',
    bech32PrefixAccPub: 'likepub',
    bech32PrefixValAddr: 'likevaloper',
    bech32PrefixValPub: 'likevaloperpub',
    bech32PrefixConsAddr: 'likevalcons',
    bech32PrefixConsPub: 'likevalconspub',
    availableMethods: [
      LikeCoinWalletConnectorMethodType.LikerId,
      LikeCoinWalletConnectorMethodType.Keplr,
      [
        LikeCoinWalletConnectorMethodType.KeplrMobile,
        { tier: 1, isRecommended: true },
      ],
      LikeCoinWalletConnectorMethodType.Cosmostation,
      LikeCoinWalletConnectorMethodType.CosmostationMobile,
      LikeCoinWalletConnectorMethodType.LikerLandApp,
      LikeCoinWalletConnectorMethodType.Leap,
      LikeCoinWalletConnectorMethodType.MetaMaskLeap,
      LikeCoinWalletConnectorMethodType.WalletConnectV2,
    ],
    keplrInstallCTAPreset: 'origin',
    likerLandAppWCBridge: 'https://wc-bridge-1.like.co',
    walletConnectProjectId: 'e110ac49451fee41d5bcda1b0dfdb94e',
    walletConnectMetadata: {
      description: 'Turn stories into collectibles',
      url: 'https://liker.land',
      icons: ['https://liker.land/logo.png'],
      name: 'Liker Land',
    },
    cosmostationDirectSignEnabled: true,
    authcoreClientId: 'likecoin-app-hidesocial', // 'likecoin-app' if not hide
    authcoreApiHost: AUTHCORE_API_HOST(isTestnet),
    authcoreRedirectUrl: `${authcoreRedirectUrl}/auth/redirect?method=${LikeCoinWalletConnectorMethodType.LikerId}`,
  };
}

export function EVM_WALLET_CONNECTOR_CONFIG(
  isTestnet: boolean
): LikeCoinEVMWalletConnectorOptions {
  return {
    magicLinkAPIKey: EVM_MAGIC_LINK_API_KEY(isTestnet),
    rpcURL: EVM_RPC_URL(isTestnet),
    chainId: EVM_CHAIN_ID(isTestnet),
  };
}
