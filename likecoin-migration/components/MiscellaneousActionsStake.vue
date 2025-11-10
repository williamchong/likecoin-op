<template>
  <div :class="['flex', 'flex-row', 'gap-2']">
    <p :class="['flex-1', 'py-[10px]', 'text-sm', 'text-likecoin-darkgrey']">
      {{ $t('miscellaneous-actions.stake.description') }}
    </p>
    <AppButton
      variant="primary"
      :class="['self-start']"
      href="https://3ook.com/account?utm_source=likecoin&utm_medium=miscellaneous_actions_stake_button&utm_campaign=likecoin_migration"
      target="_blank"
    >
      {{ $t('miscellaneous-actions.stake.stake') }}
    </AppButton>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'MiscellaneousActionsStake',
});
</script>
