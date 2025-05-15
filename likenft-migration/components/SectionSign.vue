<template>
  <div :class="['flex', 'flex-row', 'items-center']">
    <div :class="['flex-1']">
      <slot name="title" />
      <p
        v-if="failedReason != null"
        :class="['text-likecoin-votecolor-no', 'text-xs', 'mt-1']"
      >
        {{ $t('section.confirm-by-signing.failed-reason', failedReason) }}
      </p>
    </div>
    <div v-if="isLoading">
      <LoadingIcon />
    </div>
    <div v-else-if="failedReason != null">
      <AppButton @click="failedReason.retryAction">{{
        failedReason.buttonText
      }}</AppButton>
    </div>
  </div>
</template>

<script lang="ts">
import { sortedJsonStringify } from '@cosmjs/amino/build/signdoc';
import { OfflineAminoSigner } from '@keplr-wallet/types';
import Vue, { PropType } from 'vue';
import VueI18n from 'vue-i18n';

import { isEthersError } from '@/utils/ethersError';
import { LIKECOIN_WALLET_CONNECTOR_CONFIG } from '~/constant/network';
import { StepStateStep3SigningFailedReason } from '~/pageModels';

interface Data {
  isLoading: boolean;
  signFailedReason: string | null;
}

export default Vue.extend({
  name: 'SectionSign',
  props: {
    signingMessage: {
      type: String as PropType<string>,
      required: true,
    },
    externalFailedReason: {
      type: Object as PropType<StepStateStep3SigningFailedReason | null>,
      required: false,
      default: null,
    },
  },

  data(): Data {
    return {
      isLoading: false,
      signFailedReason: null,
    };
  },

  computed: {
    failedReason(): {
      type: VueI18n.TranslateResult;
      message: VueI18n.TranslateResult;
      buttonText: VueI18n.TranslateResult;
      retryAction: () => void;
    } | null {
      if (this.signFailedReason) {
        return {
          type: this.$t('section.confirm-by-signing.error-type.signing'),
          message: this.signFailedReason,
          buttonText: this.$t('section.confirm-by-signing.retry'),
          retryAction: this.initSign,
        };
      }
      if (this.externalFailedReason) {
        switch (this.externalFailedReason.type) {
          case 'likerIDMigration':
            if (this.externalFailedReason.error.isEnum) {
              switch (this.externalFailedReason.error.value) {
                case 'EVM_WALLET_USED_BY_OTHER_USER':
                  return {
                    type: this.$t(
                      'section.confirm-by-signing.error-type.liker-id-migration'
                    ),
                    message: this.$t(
                      'errors.liker-id-migration.book-user-error.evm-walletused-by-other-user'
                    ),
                    buttonText: this.$t(
                      'section.confirm-by-signing.restart-message.reconnect-evm-wallet'
                    ),
                    retryAction: () => {
                      return this.emit('reconnect-evm-wallet');
                    },
                  };
              }
            }
            return {
              type: this.$t(
                'section.confirm-by-signing.error-type.liker-id-migration'
              ),
              message: this.externalFailedReason.error.value,
              buttonText: this.$t(
                'section.confirm-by-signing.restart-message.restart'
              ),
              retryAction: () => {
                return this.emit('restart');
              },
            };
        }
        return {
          type: this.$t('section.confirm-by-signing.error-type.unknown'),
          message: this.$t('errors.unknown'),
          buttonText: this.$t(
            'section.confirm-by-signing.restart-message.restart'
          ),
          retryAction: () => {
            return this.emit('restart');
          },
        };
      }
      return null;
    },
  },

  mounted() {
    this.initSign();
  },

  methods: {
    async initSign() {
      this.isLoading = true;
      this.signFailedReason = null;

      try {
        const connection =
          await this.$likeCoinWalletConnector.initIfNecessary();
        if (connection == null) {
          this.signFailedReason = 'cannot get wallet connector connection';
          return;
        }
        const {
          accounts: [account],
        } = connection;

        // FIXME: only works on keplr
        const offlineSigner: OfflineAminoSigner =
          connection.offlineSigner as OfflineAminoSigner;
        if (!offlineSigner.signAmino) {
          this.signFailedReason = 'cannot sign message';
          return;
        }

        const { chainId, coinMinimalDenom } = LIKECOIN_WALLET_CONNECTOR_CONFIG(
          this.$appConfig.isTestnet,
          this.$appConfig.authcoreRedirectUrl
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
        this.signFailedReason = message;
      } finally {
        this.isLoading = false;
      }
    },

    emit(event: 'reconnect-evm-wallet' | 'restart', ...args: any[]) {
      this.$emit(event, ...args);
    },
  },
});
</script>
