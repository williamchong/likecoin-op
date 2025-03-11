import { Inject, NuxtApp } from '@nuxt/types/app';
import type { CrispClass } from 'crisp-sdk-web';
import { ChatboxColors, Crisp } from 'crisp-sdk-web';

export default function (app: NuxtApp, inject: Inject) {
  Crisp.configure(app.$appConfig.crispWebsiteId);
  Crisp.setColorTheme(ChatboxColors.Teal);
  inject('crisp', Crisp);
}

declare module 'vue/types/vue' {
  interface Vue {
    $crisp: CrispClass;
  }
}
