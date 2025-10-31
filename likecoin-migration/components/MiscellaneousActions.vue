<template>
  <ol
    :class="[
      'pl-4',
      'flex',
      'flex-col',
      'gap-2',
      'list-decimal',
      'marker:text-likecoin-darkgrey',
      'marker:text-sm',
    ]"
  >
    <li>
      <MiscellaneousActionsStake />
    </li>
    <li>
      <MiscellaneousActionsAddChains />
    </li>
    <li><MiscellaneousActionsAddTokens /></li>
  </ol>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'MiscellaneousActions',
});
</script>
