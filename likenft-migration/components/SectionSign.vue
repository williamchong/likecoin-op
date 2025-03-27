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
import { sortedJsonStringify } from '@cosmjs/amino/build/signdoc';
import { OfflineAminoSigner } from '@keplr-wallet/types';
import Vue, { PropType } from 'vue';

import { isEthersError } from '@/utils/ethersError';
import { LIKECOIN_WALLET_CONNECTOR_CONFIG } from '~/constant/network';

interface Data {
  isLoading: boolean;
  failedReason: string | null;
}

export default Vue.extend({
  name: 'SectionSign',
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
      this.failedReason = null;

      try {
        const connection =
          await this.$likeCoinWalletConnector.initIfNecessary();
        if (connection == null) {
          this.failedReason = 'cannot get wallet connector connection';
          return;
        }
        const {
          accounts: [account],
        } = connection;

        // FIXME: only works on keplr
        const offlineSigner: OfflineAminoSigner =
          connection.offlineSigner as OfflineAminoSigner;
        if (!offlineSigner.signAmino) {
          this.failedReason = 'cannot sign message';
          return;
        }

        const { chainId, coinMinimalDenom } = LIKECOIN_WALLET_CONNECTOR_CONFIG(
          this.$appConfig.isTestnet
        );

        const signingPayload = {
          chain_id: chainId,
          memo: this.signingMessage,
          msgs: [],
          fee: {
            gas: '0',
            amount: [
              {
                denom: coinMinimalDenom,
                amount: '0',
              },
            ],
          },
          sequence: '0',
          account_number: '0',
        };

        const cosmosSigningMessage = sortedJsonStringify(signingPayload);

        const { signature: cosmosSignature } = await offlineSigner.signAmino(
          account.address,
          signingPayload
        );

        const ethSignature =
          await this.$likeCoinEVMWalletConnector.connector.signMessage(
            this.signingMessage
          );

        this.$emit(
          'signed',
          cosmosSigningMessage,
          cosmosSignature,
          ethSignature
        );
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
