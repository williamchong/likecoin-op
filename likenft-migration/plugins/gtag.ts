import { Context, Inject } from '@nuxt/types/app';
import Vue from 'vue';
import VueGtagPlugin from 'vue-gtag';

export default function (ctx: Context, _: Inject) {
  const { app } = ctx;
  const id = app.$appConfig.googleAnalyticsTagId;
  Vue.use(
    VueGtagPlugin,
    {
      enabled: !!id,
      config: { id },
    },
    app.router
  );
}
