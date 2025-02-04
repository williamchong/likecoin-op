import { ActionContext } from 'vuex';

import { Config, ConfigSchema } from '~/models/config';

export interface State {
  config: Config | null;
}

export const state = (): State => ({
  config: null,
});

export const getters = {
  getConfig(state: State): Config | null {
    return state.config;
  },

  mustGetConfig(state: State): Config {
    if (state.config == null) {
      throw new Error('config is null');
    }
    return state.config;
  },
};

export const mutations = {
  setConfig(state: State, config: Config) {
    state.config = config;
  },
};

export const actions = {
  async fetchConfig(context: ActionContext<State, void>): Promise<Config> {
    const d = await fetch('/config.json');
    const c = await ConfigSchema.parseAsync(await d.json());
    context.commit('setConfig', c);
    return c;
  },
};
