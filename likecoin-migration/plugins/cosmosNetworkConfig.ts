import { Inject, NuxtApp } from '@nuxt/types/app';

import {
  CosmosNetworkConfig,
  CosmosNetworkConfigSchema,
} from '~/models/cosmosNetworkConfig';

export default async function (_: NuxtApp, inject: Inject) {
  const d = await fetch('/cosmos-network-config.json');
  const c = await CosmosNetworkConfigSchema.parseAsync(await d.json());
  inject('cosmosNetworkConfig', c);
}

declare module 'vue/types/vue' {
  interface Vue {
    $cosmosNetworkConfig: CosmosNetworkConfig;
  }
}
