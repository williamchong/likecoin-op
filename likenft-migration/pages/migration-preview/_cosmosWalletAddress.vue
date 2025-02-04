<template>
  <div>
    {{ cosmosWalletAddress }}
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  computed: {
    cosmosWalletAddress() {
      return this.$route.params.cosmosWalletAddress;
    },
  },
});
</script>
