<template>
  <span> {{ amountString }} {{ viewCoin.denom }} </span>
</template>

<script lang="ts">
import { Decimal } from 'decimal.js';
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
    ltLimit: {
      type: String,
      default: '0.01',
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
      const amountDecimal = Decimal(this.viewCoin.amount);
      if (amountDecimal.equals(Decimal('0'))) {
        return `0`;
      }
      if (amountDecimal.lessThan(Decimal(this.ltLimit))) {
        // Too small will be formatted as NaN
        return `< ${this.ltLimit}`;
      }
      return numeral(this.viewCoin.amount).format('0,0.[0000000000000000]');
    },
  },
});
</script>
