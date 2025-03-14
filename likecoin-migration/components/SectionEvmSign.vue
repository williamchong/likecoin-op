<template>
  <div :class="['flex', 'flex-row', 'items-center']">
    <div :class="['flex-1']">
      <slot name="title" />
      <p
        v-if="failedReason != null"
        :class="['text-likecoin-votecolor-no', 'text-xs', 'mt-1']"
      >
        {{ failedReason }}
      </p>
      <p v-else :class="['text-likecoin-votecolor-abstain', 'text-xs', 'mt-1']">
        {{ $t('section.confirm-by-signing.message') }}
      </p>
    </div>
    <div v-if="isLoading">
      <LoadingIcon />
    </div>
    <div v-else-if="failedReason != null">
      <AppButton @click="handleRetryClick">{{
        $t('section.confirm-by-signing.retry')
      }}</AppButton>
    </div>
  </div>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue';

import { isEthersError } from '~/utils/ethersError';

interface Data {
  isLoading: boolean;
  failedReason: string | null;
}

export default Vue.extend({
  props: {
    signingMessage: {
      type: String as PropType<string>,
      required: true,
    },
  },
  data(): Data {
    return {
      isLoading: false,
      failedReason: null,
    };
  },
  mounted() {
    this.initSign();
  },
  methods: {
    async initSign() {
      this.isLoading = true;
      try {
        const signature =
          await this.$likeCoinEVMWalletConnector.connector.signMessage(
            this.signingMessage
          );
        this.$emit('signed', signature);
      } catch (e: unknown) {
        console.warn({ e });
        let message: string;
        if (isEthersError(e)) {
          message = e.shortMessage;
        } else if (e instanceof Error) {
          message = e.message;
        } else {
          message = `${e}`;
        }
        this.failedReason = message;
      } finally {
        this.isLoading = false;
      }
    },
    handleRetryClick() {
      this.initSign();
    },
  },
});
</script>
