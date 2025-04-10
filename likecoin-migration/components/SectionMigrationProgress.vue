<template>
  <div>
    <div :class="['flex', 'flex-row', 'items-center']">
      <div
        :class="[
          'flex-1',
          'flex',
          'flex-col',
          'min-h-[30px]',
          'justify-center',
        ]"
      >
        <slot name="title" />
        <p
          v-if="errorMessage"
          :class="['text-likecoin-votecolor-no', 'text-xs', 'mt-1']"
        >
          <FontAwesomeIcon :class="['font-bold']" icon="triangle-exclamation" />
          {{ errorMessage }}
        </p>
        <p
          v-else-if="completedAndHasEstimatedBalance"
          :class="['text-like-green', 'text-xs', 'mt-1']"
        >
          {{ $t('section.migration-progress.balance-remaining-message') }}
        </p>
      </div>
      <div v-if="errorMessage">
        <AppButton @click="handleRetryClick">{{
          $t('section.migration-progress.retry')
        }}</AppButton>
      </div>
      <div v-else-if="completedAndHasEstimatedBalance">
        <AppButton @click="handleRestartClick">{{
          $t('section.migration-progress.restart')
        }}</AppButton>
      </div>
    </div>
    <div :class="['bg-likecoin-lightergrey', 'p-8', 'rounded', 'mt-2.5']">
      <div :class="['flex', 'flex-col', 'gap-2.5']">
        <h3 :class="['text-base', 'font-semibold', 'text-likecoin-darkgreen']">
          {{ $t('section.migration-progress.status') }}
        </h3>
        <div
          :class="[
            'h-[3px]',
            'bg-likecoin-grey',
            'flex',
            'flex-row',
            'items-stretch',
            'rounded-full',
          ]"
        >
          <div
            :class="[
              'bg-gradient-to-r',
              'from-likecoin-buttontext',
              'to-likecoin-votecolor-yes',
              'rounded-full',
            ]"
            :style="{ width: `${viewState.progress}%` }"
          ></div>
        </div>
        <ul :class="['grid', 'grid-cols-4', 'gap-4']">
          <li>
            <div :class="['flex', 'flex-row', 'gap-2.5']">
              <FontAwesomeIcon
                :class="[
                  'font-bold',
                  {
                    'text-likecoin-votecolor-yes':
                      viewState.stepState[0] === 'success',
                    'text-likecoin-votecolor-no':
                      viewState.stepState[0] === 'failed',
                    'text-likecoin-grey': viewState.stepState[0] === 'pending',
                  },
                ]"
                :icon="stepIcon[0]"
              />
              <span
                :class="[
                  'font-bold',
                  {
                    'text-likecoin-darkgrey': viewState.progress >= 25,
                    'text-likecoin-grey': viewState.progress < 25,
                  },
                ]"
              >
                {{ $t('section.migration-progress.pending.title') }}
              </span>
            </div>
            <p :class="['mt-1', 'text-sm', 'text-likecoin-darkgrey']">
              {{ $t('section.migration-progress.pending.description') }}
            </p>
          </li>
          <li>
            <div :class="['flex', 'flex-row', 'gap-2.5']">
              <FontAwesomeIcon
                :class="[
                  'font-bold',
                  {
                    'text-likecoin-votecolor-yes':
                      viewState.stepState[1] === 'success',
                    'text-likecoin-votecolor-no':
                      viewState.stepState[1] === 'failed',
                    'text-likecoin-grey': viewState.stepState[1] === 'pending',
                  },
                ]"
                :icon="stepIcon[1]"
              />
              <span
                :class="[
                  'font-bold',
                  {
                    'text-likecoin-darkgrey': viewState.progress >= 50,
                    'text-likecoin-grey': viewState.progress < 50,
                  },
                ]"
              >
                {{ $t('section.migration-progress.receiving.title') }}
              </span>
            </div>
            <p :class="['mt-1', 'text-sm', 'text-likecoin-darkgrey']">
              {{ $t('section.migration-progress.receiving.description') }}
            </p>
            <div
              :class="['flex', 'flex-row', 'items-center', 'mt-2', 'gap-2.5']"
            >
              <a
                :href="cosmosTxUrl"
                target="_blank"
                :class="[
                  'text-base',
                  'text-likecoin-votecolor-yes',
                  'overflow-hidden',
                  'text-ellipsis',
                ]"
              >
                {{ cosmosTxHash }}
              </a>
              <button
                v-if="cosmosTxHash != null"
                type="button"
                @click="handleTxHashCopyClick(cosmosTxHash)"
              >
                <FontAwesomeIcon :class="['text-base']" icon="copy" />
              </button>
            </div>
          </li>
          <li>
            <div :class="['flex', 'flex-row', 'gap-2.5']">
              <FontAwesomeIcon
                :class="[
                  'font-bold',
                  {
                    'text-likecoin-votecolor-yes':
                      viewState.stepState[2] === 'success',
                    'text-likecoin-votecolor-no':
                      viewState.stepState[2] === 'failed',
                    'text-likecoin-grey': viewState.stepState[2] === 'pending',
                  },
                ]"
                :icon="stepIcon[2]"
              />
              <span
                :class="[
                  'font-bold',
                  {
                    'text-likecoin-darkgrey': viewState.progress >= 75,
                    'text-likecoin-grey': viewState.progress < 75,
                  },
                ]"
              >
                {{ $t('section.migration-progress.minting.title') }}
              </span>
            </div>
            <p :class="['mt-1', 'text-sm', 'text-likecoin-darkgrey']">
              {{ $t('section.migration-progress.minting.description') }}
            </p>
            <div
              :class="['flex', 'flex-row', 'items-center', 'mt-2', 'gap-2.5']"
            >
              <a
                :href="evmTxUrl"
                target="_blank"
                :class="[
                  'text-base',
                  'text-likecoin-votecolor-yes',
                  'overflow-hidden',
                  'text-ellipsis',
                ]"
              >
                {{ evmTxHash }}
              </a>
              <button
                v-if="evmTxHash != null"
                type="button"
                @click="handleTxHashCopyClick(evmTxHash)"
              >
                <FontAwesomeIcon :class="['text-base']" icon="copy" />
              </button>
            </div>
          </li>
          <li>
            <div :class="['flex', 'flex-row', 'gap-2.5']">
              <FontAwesomeIcon
                :class="[
                  'font-bold',
                  {
                    'text-likecoin-votecolor-yes':
                      viewState.stepState[3] === 'success',
                    'text-likecoin-votecolor-no':
                      viewState.stepState[3] === 'failed',
                    'text-likecoin-grey': viewState.stepState[3] === 'pending',
                  },
                ]"
                :icon="stepIcon[3]"
              />
              <span
                :class="[
                  'font-bold',
                  {
                    'text-likecoin-darkgrey': viewState.progress >= 100,
                    'text-likecoin-grey': viewState.progress < 100,
                  },
                ]"
              >
                {{ $t('section.migration-progress.verified.title') }}
              </span>
            </div>
            <div :class="['flex', 'flex-row']">
              <p :class="['mt-1', 'text-sm', 'text-likecoin-darkgrey']">
                {{ $t('section.migration-progress.verified.description') }}
              </p>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue';

import { LikeCoinMigrationStatus } from '~/apis/models/likeCoinMigration';
import { ChainCoin } from '~/models/cosmosNetworkConfig';
import { SupportedIcon } from '~/models/faIcon';

type StepState = 'pending' | 'success' | 'failed';

type ViewState = {
  stepState: [StepState, StepState, StepState, StepState];
  progress: number;
};

function computeViewState(
  migrationStatus: LikeCoinMigrationStatus,
  cosmosTxHash: string | null,
  evmTxHash: string | null
): ViewState {
  switch (migrationStatus) {
    case 'failed':
      if (cosmosTxHash != null && evmTxHash != null) {
        return {
          stepState: ['success', 'success', 'success', 'failed'],
          progress: 75,
        };
      }
      if (cosmosTxHash != null) {
        return {
          stepState: ['success', 'success', 'failed', 'pending'],
          progress: 50,
        };
      }
      return {
        stepState: ['success', 'failed', 'pending', 'pending'],
        progress: 25,
      };
    case 'completed':
      return {
        stepState: ['success', 'success', 'success', 'success'],
        progress: 100,
      };
    case 'pending_cosmos_tx_hash':
      return {
        stepState: ['success', 'pending', 'pending', 'pending'],
        progress: 25,
      };
    case 'verifying_cosmos_tx':
      return {
        stepState: ['success', 'pending', 'pending', 'pending'],
        progress: 25,
      };
    case 'evm_minting':
      return {
        stepState: ['success', 'success', 'pending', 'pending'],
        progress: 50,
      };
    case 'evm_verifying':
      return {
        stepState: ['success', 'success', 'success', 'pending'],
        progress: 75,
      };
  }
}

export default Vue.extend({
  props: {
    enableRetry: {
      type: Boolean,
      default: true,
    },
    estimatedBalance: {
      type: Object as PropType<ChainCoin | null>,
      default: null,
    },
    errorMessage: {
      type: String as PropType<string | null>,
      default: null,
    },
    migrationStatus: {
      type: String as PropType<LikeCoinMigrationStatus>,
      required: true,
    },
    cosmosTxHash: {
      type: String as PropType<string | null>,
      default: null,
    },
    evmTxHash: {
      type: String as PropType<string | null>,
      default: null,
    },
  },
  computed: {
    viewState(): ViewState {
      return computeViewState(
        this.migrationStatus,
        this.cosmosTxHash,
        this.evmTxHash
      );
    },

    stepIcon(): { [key in 0 | 1 | 2 | 3]: SupportedIcon } {
      return {
        0:
          this.viewState.stepState[0] === 'failed'
            ? 'circle-xmark'
            : 'circle-check',
        1:
          this.viewState.stepState[1] === 'failed'
            ? 'circle-xmark'
            : 'circle-check',
        2:
          this.viewState.stepState[2] === 'failed'
            ? 'circle-xmark'
            : 'circle-check',
        3:
          this.viewState.stepState[3] === 'failed'
            ? 'circle-xmark'
            : 'circle-check',
      };
    },

    completedAndHasEstimatedBalance() {
      return (
        this.enableRetry &&
        this.migrationStatus === 'completed' &&
        this.estimatedBalance != null &&
        this.estimatedBalance.amount !== '0'
      );
    },

    cosmosTxUrl(): string {
      return new URL(
        `cosmos/tx/v1beta1/txs/${this.cosmosTxHash}`,
        this.$appConfig.cosmosExplorerBaseURL
      ).toString();
    },

    evmTxUrl(): string {
      return new URL(
        `/tx/${this.evmTxHash}`,
        this.$appConfig.evmExplorerBaseURL
      ).toString();
    },
  },

  methods: {
    async handleTxHashCopyClick(txHash: string) {
      await window.navigator.clipboard.writeText(txHash);
      alert('Transaction hash is copied to clipboard!');
    },
    handleRetryClick() {
      this.$emit('retry');
    },
    handleRestartClick() {
      this.$emit('restart');
    },
  },
});
</script>
