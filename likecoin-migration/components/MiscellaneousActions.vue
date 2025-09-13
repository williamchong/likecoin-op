<template>
  <div :class="['flex', 'flex-col', 'gap-2']">
    <MiscellaneousActionsAddChains />
    <MiscellaneousActionsAddTokens />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'MiscellaneousActions',
});
</script>
