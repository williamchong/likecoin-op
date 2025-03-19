import { Inject, NuxtApp } from '@nuxt/types/app';

import {
  CosmosNetworkConfig,
  CosmosNetworkConfigSchema,
} from '~/models/cosmosNetworkConfig';

export default async function (app: NuxtApp, inject: Inject) {
  const d = await fetch(app.$appConfig.cosmosLikeCoinNetworkConfigPath);
  const c = await CosmosNetworkConfigSchema.parseAsync(await d.json());
  inject('cosmosNetworkConfig', c);
}

declare module 'vue/types/vue' {
  interface Vue {
    $cosmosNetworkConfig: CosmosNetworkConfig;
  }
}
