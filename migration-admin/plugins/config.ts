import { Inject, NuxtApp } from "@nuxt/types/app";

import { Config, ConfigSchema } from "~/models/config";

export default async function (_: NuxtApp, inject: Inject) {
  const d = await fetch("/config.json");
  const c = await ConfigSchema.parseAsync(await d.json());
  inject("appConfig", c);
}

declare module "vue/types/vue" {
  interface Vue {
    $appConfig: Config;
  }
}
