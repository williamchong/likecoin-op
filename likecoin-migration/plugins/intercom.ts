import { Intercom } from '@intercom/messenger-js-sdk';
import { Context, Inject } from '@nuxt/types/app';

export default function (ctx: Context, inject: Inject) {
  const { app } = ctx;

  Intercom({
    app_id: app.$appConfig.intercomAppId,
  });

  const intercomWrapper = (method: string, ...args: any[]) => {
    if (typeof window !== 'undefined' && (window as any).Intercom) {
      return (window as any).Intercom(method, ...args);
    }
  };

  inject('intercom', intercomWrapper);
}

declare module 'vue/types/vue' {
  interface Vue {
    $intercom: (method: string, ...args: any[]) => void;
  }
}
