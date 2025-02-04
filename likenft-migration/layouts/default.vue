<template>
  <div>
    <Nuxt v-if="config != null" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

import { Config } from '~/models/config';

export default Vue.extend({
  computed: {
    config(): Config | null {
      return this.$store.state.config;
    },
  },
  mounted() {
    this.$store.dispatch('config/fetchConfig');
  },
});
</script>
