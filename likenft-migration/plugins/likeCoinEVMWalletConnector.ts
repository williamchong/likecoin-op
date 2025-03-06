// @ts-expect-error
import { LikeCoinEVMWalletConnector } from '@likecoin/evm-wallet-connector';
import type { LikeCoinEVMWalletConnector as ILikeCoinEVMWalletConnector } from '@likecoin/evm-wallet-connector/dist/index';
import { Inject, NuxtApp } from '@nuxt/types/app';
import { EventEmitter } from 'events';

import { EVM_WALLET_CONNECTOR_CONFIG } from '~/constant/network';

const eventEmitter = new EventEmitter<{
  connect: [{ walletAddress: string; providerId: string }];
}>();

interface _LikeCoinWalletConnector {
  eventEmitter: typeof eventEmitter;
  connector: ILikeCoinEVMWalletConnector;
}

export default function (app: NuxtApp, inject: Inject) {
  const likeCoinWalletConnector = new LikeCoinEVMWalletConnector({
    ...EVM_WALLET_CONNECTOR_CONFIG(app.$appConfig.isTestnet),
    onConnect: (payload: { walletAddress: string; providerId: string }) => {
      eventEmitter.emit('connect', payload);
    },
  });

  inject('likeCoinEVMWalletConnector', {
    eventEmitter,
    connector: likeCoinWalletConnector,
  });
}

declare module 'vue/types/vue' {
  interface Vue {
    $likeCoinEVMWalletConnector: _LikeCoinWalletConnector;
  }
}
