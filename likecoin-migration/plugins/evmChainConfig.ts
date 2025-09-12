import { Inject, NuxtApp } from '@nuxt/types/app';

import { EVMChainConfig, EVMChainConfigSchema } from '~/models/evmChainConfig';

export default async function (app: NuxtApp, inject: Inject) {
  const d = await fetch(app.$appConfig.evmLikeCoinChainConfigPath);
  const c = await EVMChainConfigSchema.parseAsync(await d.json());
  inject('evmChainConfig', c);
}

declare module 'vue/types/vue' {
  interface Vue {
    $evmChainConfig: EVMChainConfig;
  }
}
