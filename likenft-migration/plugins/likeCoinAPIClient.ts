import { Inject, NuxtApp } from '@nuxt/types/app';
import type { NuxtAxiosInstance } from '@nuxtjs/axios';

export default function (app: NuxtApp, inject: Inject) {
  const apiClient = app.$axios.create({});
  apiClient.setBaseURL(app.$appConfig.likeCoinAPIBaseURL);
  inject('likeCoinApiClient', apiClient);
}

declare module 'vue/types/vue' {
  interface Vue {
    $likeCoinApiClient: NuxtAxiosInstance;
  }
}
