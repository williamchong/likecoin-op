<template>
  <span> {{ amountString }} {{ viewCoin.denom }} </span>
</template>

<script lang="ts">
import numeral from 'numeral';
import Vue, { PropType } from 'vue';

import {
  ChainCoin,
  convertChainCoinToViewCoin,
  isChainCoin,
  ViewCoin,
} from '~/models/cosmosNetworkConfig';

export default Vue.extend({
  props: {
    coin: {
      type: Object as PropType<ViewCoin | ChainCoin>,
      required: true,
    },
  },
  computed: {
    viewCoin(): ViewCoin {
      if (isChainCoin(this.coin)) {
        return convertChainCoinToViewCoin(this.coin, this.$cosmosNetworkConfig);
      }
      return this.coin;
    },
    amountString(): string {
      return numeral(this.viewCoin.amount).format('0,0.[0000000000000000]');
    },
  },
});
</script>
