<template>
  <div :class="['flex-1', 'min-h-0', 'bg-likecoin-lightergrey']">
    <div :class="['-mb-[65px]']">
      <HeroBanner>
        <h1
          :class="[
            'text-3xl',
            'font-inter',
            'font-semibold',
            'text-likecoin-votecolor-yes',
          ]"
        >
          {{ $t('app.title') }}
        </h1>
      </HeroBanner>
    </div>
    <div :class="['relative', 'max-w-[880px]', 'px-4', 'mx-auto']">
      <div :class="['bg-white', 'p-[30px]', 'rounded-md', 'shadow-md']">
        <StepSection :step="1" :current-step="currentStep.step">
          <h2 :class="['text-base', 'font-semibold', 'text-likecoin-darkgrey']">
            {{ $t('section.introduction.title') }}
          </h2>
          <template #current>
            <div :class="['mt-2', 'mb-5']">
              <SectionIntroduction
                @confirmClicked="handleIntroductionSectionConfirmClick"
              />
            </div>
          </template>
        </StepSection>
        <StepSection :step="2" :current-step="currentStep.step">
          <h2
            :class="[
              'text-base',
              'font-semibold',
              'leading-[30px]',
              'text-likecoin-darkgrey',
            ]"
          >
            {{ $t('section.wallet-connect.title') }}
          </h2>
          <template #current>
            <div :class="['mt-2.5', 'mb-5']">
              <SectionWalletConnect
                :liker-id="likerId"
                :avatar="avatar"
                :cosmos-address="cosmosAddress"
                :eth-address="ethAddress"
                @likeCoinWalletConnected="handleLikeCoinWalletConnected"
                @likeCoinEVMWalletConnected="handleLikeCoinEVMWalletConnected"
              />
            </div>
          </template>
          <template #past>
            <div :class="['mt-2.5', 'mb-5']">
              <SectionWalletConnect
                :liker-id="likerId"
                :avatar="avatar"
                :cosmos-address="cosmosAddress"
                :eth-address="ethAddress"
                @likeCoinWalletConnected="handleLikeCoinWalletConnected"
                @likeCoinEVMWalletConnected="handleLikeCoinEVMWalletConnected"
              />
            </div>
          </template>
        </StepSection>
        <StepSection :step="3" :current-step="currentStep.step">
          <template #future>
            <h2
              :class="[
                'text-base',
                'font-semibold',
                'leading-[30px]',
                'text-likecoin-darkgrey',
              ]"
            >
              {{ $t('section.confirm-by-signing.title') }}
            </h2>
          </template>
          <template v-if="currentStep.step === 3" #current>
            <SectionSign
              :signing-message="currentStep.signMessage"
              @signed="handleSigned"
            >
              <template #title>
                <h2
                  :class="[
                    'text-base',
                    'font-semibold',
                    'leading-[30px]',
                    'text-likecoin-darkgrey',
                  ]"
                >
                  {{ $t('section.confirm-by-signing.title') }}
                </h2>
              </template>
            </SectionSign>
          </template>
          <template #past>
            <h2
              :class="[
                'text-base',
                'font-semibold',
                'leading-[30px]',
                'text-likecoin-darkgrey',
              ]"
            >
              {{ $t('section.confirm-by-signing.title') }}
            </h2>
          </template>
        </StepSection>
        <StepSection :step="4" :current-step="currentStep.step">
          <template #future>
            <h2
              :class="[
                'text-base',
                'font-semibold',
                'leading-[30px]',
                'text-likecoin-darkgrey',
              ]"
            >
              {{ $t('migrate.preview') }}
            </h2>
          </template>
          <template v-if="currentStep.step === 4" #current>
            <div :class="['flex', 'flex-row', 'gap-1']">
              <h2
                :class="[
                  'text-base',
                  'font-semibold',
                  'leading-[30px]',
                  'text-likecoin-darkgrey',
                ]"
              >
                {{ $t('migrate.preview') }}
              </h2>
              <UTooltip
                v-if="
                  currentStep.migrationPreview != null &&
                  currentStep.migrationPreview.block_time != null &&
                  currentStep.migrationPreview.block_height != null
                "
                :text="
                  $t('section.asset-preview.tooltip', {
                    date: _formatDate(currentStep.migrationPreview.block_time),
                    height: _formatNumber(
                      currentStep.migrationPreview.block_height
                    ),
                  })
                "
                :ui="{
                  base: '[@media(pointer:coarse)]:hidden px-2 py-1 text-xs font-normal w-80 relative',
                }"
                ><FontAwesomeIcon
                  icon="circle-exclamation"
                  :class="[
                    'text-sm',
                    'leading-[30px]',
                    'text-likecoin-votecolor-yes',
                  ]"
              /></UTooltip>
            </div>
            <SectionAssetPreview
              v-if="
                currentStep.state === 'MigrationPreview' &&
                currentStep.migrationPreview != null
              "
              :class="['max-w-full', 'mt-2']"
              :snapshot="currentStep.migrationPreview"
              @confirmMigration="handleConfirmMigrate"
            />
            <div
              v-else
              :class="[
                'flex',
                'flex-row',
                'items-center',
                'justify-center',
                'my-4',
              ]"
            >
              <LoadingIcon />
            </div>
            <template v-if="currentStep.state === 'MigrationRetryPreview'">
              <div :class="['flex', 'flex-row', 'gap-1']">
                <h2
                  :class="[
                    'text-base',
                    'font-semibold',
                    'leading-[30px]',
                    'text-likecoin-darkgrey',
                  ]"
                >
                  {{ $t('migrate.preview') }}
                </h2>
              </div>
              <SectionMigrationResult
                :class="['max-w-full', 'mt-2']"
                :migration="currentStep.failedMigration"
                :initial-status="'failed'"
              />
              <div :class="['mt-4', 'flex', 'flex-row', 'justify-end']">
                <AppButton @click="handleConfirmMigrate">
                  {{ $t('section.asset-preview.confirm-retry') }}
                </AppButton>
              </div>
            </template>
          </template>
          <template #past>
            <h2
              :class="[
                'text-base',
                'font-semibold',
                'leading-[30px]',
                'text-likecoin-darkgrey',
              ]"
            >
              {{ $t('migrate.preview') }}
            </h2>
          </template>
        </StepSection>
        <StepSection :step="5" :current-step="currentStep.step">
          <template #future>
            <h2
              :class="[
                'text-base',
                'font-semibold',
                'leading-[30px]',
                'text-likecoin-darkgrey',
              ]"
            >
              {{ $t('section.start-migration.title') }}
            </h2>
          </template>
          <template #current>
            <div
              :class="[
                'flex',
                'flex-row',
                'items-center',
                'justify-between',
                'min-h-[30px]',
              ]"
            >
              <h2
                :class="[
                  'text-base',
                  'font-semibold',
                  'text-likecoin-darkgrey',
                ]"
              >
                {{ $t('section.migration-result.title') }}
                <span
                  v-if="
                    migration != null &&
                    (migration.status === 'in_progress' ||
                      migration.status === 'init')
                  "
                >
                  {{ $t('section.migration-result.in-progress') }}
                </span>
              </h2>
              <LoadingIcon />
            </div>
            <SectionMigrationResult
              v-if="migration != null"
              :class="['max-w-full', 'mt-2']"
              :migration="migration"
            />
          </template>
          <template #past>
            <div
              :class="['flex', 'flex-row', 'items-center', 'justify-between']"
            >
              <div
                :class="[
                  'min-h-[30px]',
                  'flex',
                  'flex-col',
                  'justify-center',
                  'gap-1',
                ]"
              >
                <h2
                  :class="[
                    'text-base',
                    'font-semibold',
                    'leading-[20px]',
                    'text-likecoin-darkgrey',
                  ]"
                >
                  {{ $t('section.migration-result.title') }}
                </h2>
                <p
                  v-if="
                    migration != null &&
                    migration.status === 'failed' &&
                    failedMigrationCount != null &&
                    failedMigrationCount > 0
                  "
                  :class="['text-xs', 'text-likecoin-votecolor-no']"
                >
                  <FontAwesomeIcon icon="triangle-exclamation" />
                  {{
                    $t('section.migration-result.failed-message', {
                      count: failedMigrationCount,
                    })
                  }}
                </p>
              </div>
              <AppButton
                v-if="failedMigrationCount != null && failedMigrationCount > 0"
                :class="['w-[120px]']"
                @click="handleRetryClick"
              >
                {{ $t('section.migration-result.retry') }}
              </AppButton>
            </div>
            <SectionMigrationResult
              v-if="migration != null"
              :class="['max-w-full', 'mt-2']"
              :migration="migration"
            />
          </template>
        </StepSection>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { StdSignature } from '@keplr-wallet/types';
import { LikeCoinWalletConnectorMethodType } from '@likecoin/wallet-connector';
import { isAxiosError } from 'axios';
import { format as formatDate } from 'date-fns/format';
import numeral from 'numeral';
import Vue from 'vue';

import { makeCreateMigrationAPI } from '~/apis/createMigration';
import { makeCreateMigrationPreviewAPI } from '~/apis/createMigrationPreview';
import { makeGetMigrationAPI } from '~/apis/getMigration';
import { makeGetMigrationPreviewAPI } from '~/apis/getMigrationPreview';
import { getSignMessage } from '~/apis/getSignMessage';
import { makeGetUserProfileAPI } from '~/apis/getUserProfile';
import { makeMigrateLikerIDAPI } from '~/apis/migrateLikerID';
import {
  isMigrationCompleted,
  isMigrationFailed,
  LikeNFTAssetMigration,
} from '~/apis/models/likenftAssetMigration';
import { LikeNFTAssetSnapshot } from '~/apis/models/likenftAssetSnapshot';
import {
  makeRetryMigrationAPI,
  RetryMigrationRequest,
} from '~/apis/retryMigration';
import {
  initCosmosConnected,
  initEvmConnected,
  introductionConfirmed,
  likerIdEvmConnected,
  likerIdMigrated,
  likerIdResolved,
  migrationCompleted,
  migrationFailed,
  migrationPreviewFetched,
  migrationResultFetched,
  migrationRetried,
  signMessageRequested,
  StepState,
  StepStateCompleted,
  StepStateFailed,
  StepStateStep2CosmosConnected,
  StepStateStep2LikerIdEvmConnected,
  StepStateStep2LikerIdResolved,
  StepStateStep3Signing,
  StepStateStep4Init,
  StepStateStep4MigrationPreview,
  StepStateStep4MigrationRetryPreview,
  StepStateStep5MigrationResult,
} from '~/pageModels';

import UTooltip from '../nuxtui/components/UTooltip.vue';

interface Data {
  isTransitioning: boolean;

  currentStep: StepState;

  migrationPreviewFetchTimeout: ReturnType<typeof setTimeout> | null;
  migrationFetchTimeout: ReturnType<typeof setTimeout> | null;
}

export default Vue.extend({
  components: {
    UTooltip,
  },

  filters: {},
  data(): Data {
    return {
      isTransitioning: false,

      currentStep: { step: 1 },

      migrationPreviewFetchTimeout: null,
      migrationFetchTimeout: null,
    };
  },
  computed: {
    getSignMessage() {
      return getSignMessage(this.$apiClient);
    },

    migrateLikerID() {
      return makeMigrateLikerIDAPI(this.$apiClient);
    },

    likerId(): string | null {
      if ('likerId' in this.currentStep) {
        return this.currentStep.likerId;
      }
      return null;
    },

    avatar(): string | null {
      if ('avatar' in this.currentStep) {
        return this.currentStep.avatar;
      }
      return null;
    },

    cosmosAddress(): string | null {
      if ('cosmosAddress' in this.currentStep) {
        return this.currentStep.cosmosAddress;
      }
      return null;
    },

    ethAddress(): string | null {
      if ('ethAddress' in this.currentStep) {
        return this.currentStep.ethAddress;
      }
      return null;
    },

    migrationPreview(): LikeNFTAssetSnapshot | null {
      if ('migrationPreview' in this.currentStep) {
        return this.currentStep.migrationPreview;
      }
      return null;
    },

    migration(): LikeNFTAssetMigration | null {
      if ('migration' in this.currentStep) {
        return this.currentStep.migration;
      }
      return null;
    },

    failedMigrationCount(): number | null {
      if ('migration' in this.currentStep) {
        return (
          this.currentStep.migration.classes.filter(
            (c) => c.status === 'failed'
          ).length +
          this.currentStep.migration.nfts.filter((c) => c.status === 'failed')
            .length
        );
      }
      return null;
    },

    retryMigration() {
      return (cosmosAddress: string, req: RetryMigrationRequest) =>
        makeRetryMigrationAPI(cosmosAddress)(this.$apiClient)(req);
    },
  },

  watch: {
    migrationPreview(migrationPreview: LikeNFTAssetSnapshot | null) {
      if (this.migrationPreviewFetchTimeout != null) {
        clearTimeout(this.migrationPreviewFetchTimeout);
        this.migrationPreviewFetchTimeout = null;
      }
      if (migrationPreview == null) {
        return;
      }
      if (
        this.currentStep.step !== 4 ||
        this.currentStep.state !== 'MigrationPreview'
      ) {
        return;
      }
      if (
        migrationPreview.status === 'init' ||
        migrationPreview.status === 'in_progress'
      ) {
        const currentStep = this.currentStep;
        this.migrationPreviewFetchTimeout = setTimeout(async () => {
          this.currentStep = await this._asyncStateTransition(
            currentStep,
            (s) => this._getOrCreateMigrationPreview(s)
          );
        }, 1000);
      }
    },

    migration(_: LikeNFTAssetMigration | null) {
      if (this.migrationFetchTimeout != null) {
        clearTimeout(this.migrationFetchTimeout);
        this.migrationFetchTimeout = null;
      }
      if (this.currentStep.step !== 5) {
        return;
      }
      const { migration } = this.currentStep;
      if (migration.status === 'init' || migration.status === 'in_progress') {
        const currentStep = this.currentStep;
        this.migrationFetchTimeout = setTimeout(async () => {
          this.currentStep = await this._asyncStateTransition(
            currentStep,
            (s) => this._refreshMigration(s)
          );
        }, 1000);
      }
    },
  },

  async mounted() {
    await this.handleMaybeLikeCoinWalletConnectedFromRedirect();
  },

  methods: {
    handleIntroductionSectionConfirmClick() {
      if (this.currentStep.step !== 1) {
        return;
      }
      this.currentStep = introductionConfirmed(this.currentStep);
    },

    async handleMaybeLikeCoinWalletConnectedFromRedirect() {
      const { code, method, ...query } = this.$route.query;
      if (method && code) {
        this.$router.replace({ query });
        const connection = await this.$likeCoinWalletConnector.handleRedirect(
          method as LikeCoinWalletConnectorMethodType,
          { code }
        );
        if (connection != null) {
          if ('method' in connection) {
            const {
              accounts: [account],
            } = connection;
            await this.handleLikeCoinWalletConnected(account.address);
          }
        }
      }
    },

    async handleLikeCoinWalletConnected(cosmosAddress: string) {
      this.currentStep = initCosmosConnected(this.currentStep, cosmosAddress);
      this.currentStep = await this._asyncStateTransition(
        this.currentStep,
        (s) => this._checkLikerID(s, cosmosAddress)
      );

      if (this.currentStep.step === 4 && this.currentStep.state === 'Init') {
        this.currentStep = await this._asyncStateTransition(
          this.currentStep,
          (s) => this._checkMigration(s)
        );
      }
    },

    async handleLikeCoinEVMWalletConnected(ethAddress: string) {
      if (this.currentStep.step !== 2) {
        return;
      }

      switch (this.currentStep.state) {
        case 'Init': {
          this.currentStep = initEvmConnected(this.currentStep, ethAddress);
          break;
        }
        case 'LikerIdResolved': {
          this.currentStep = likerIdEvmConnected(this.currentStep, ethAddress);
          this.currentStep = await this._asyncStateTransition(
            this.currentStep,
            (s) => this._requestSignMessage(s)
          );
        }
      }
    },

    async handleSigned(
      cosmosSigningMessage: string,
      cosmosSignature: StdSignature,
      ethSignature: string
    ) {
      if (this.currentStep.step === 3 && this.currentStep.state === 'Signing') {
        this.currentStep = await this._asyncStateTransition(
          this.currentStep,
          (s) =>
            this._doMigrateLikerID(
              s,
              cosmosSigningMessage,
              cosmosSignature,
              ethSignature
            )
        );
      }

      if (this.currentStep.step === 4 && this.currentStep.state === 'Init') {
        this.currentStep = await this._asyncStateTransition(
          this.currentStep,
          (s) => this._checkMigration(s)
        );
      }
    },

    async handleConfirmMigrate() {
      if (
        this.currentStep.step === 4 &&
        this.currentStep.state === 'MigrationPreview'
      ) {
        this.currentStep = await this._asyncStateTransition(
          this.currentStep,
          (s) => this._createMigration(s)
        );
      }

      if (
        this.currentStep.step === 4 &&
        this.currentStep.state === 'MigrationRetryPreview'
      ) {
        this.currentStep = await this._asyncStateTransition(
          this.currentStep,
          (s) => this._retryMigration(s)
        );
      }
    },

    handleRetryClick() {
      if (
        this.currentStep.step === 99999 &&
        this.currentStep.state === 'Failed'
      ) {
        this.currentStep = migrationRetried(
          this.currentStep,
          this.currentStep.migration
        );
      }
    },

    async _asyncStateTransition<S1 extends StepState, S2 extends StepState>(
      state1: S1,
      asyncT: (state1: S1) => Promise<S2>
    ): Promise<S1 | S2> {
      if (this.isTransitioning) {
        return state1;
      }
      this.isTransitioning = true;
      try {
        return await asyncT(state1);
      } finally {
        this.isTransitioning = false;
      }
    },

    async _checkLikerID(
      currentStep: StepStateStep2CosmosConnected,
      cosmosAddress: string
    ): Promise<StepStateStep2LikerIdResolved | StepStateStep4Init> {
      const userProfile = await makeGetUserProfileAPI(cosmosAddress)(
        this.$apiClient
      )();
      const remoteEthAddress = userProfile.user_profile.eth_wallet_address;
      if (remoteEthAddress != null) {
        return likerIdMigrated(
          currentStep,
          userProfile.user_profile.liker_id,
          userProfile.user_profile.avatar,
          remoteEthAddress
        );
      } else {
        return likerIdResolved(
          currentStep,
          userProfile.user_profile.avatar,
          userProfile.user_profile.liker_id
        );
      }
    },

    async _requestSignMessage(
      currentStep: StepStateStep2LikerIdEvmConnected
    ): Promise<StepStateStep3Signing> {
      const signMessage = await this.getSignMessage({
        cosmos_address: currentStep.cosmosAddress,
        eth_address: currentStep.ethAddress,
        liker_id: currentStep.likerId,
      });
      return signMessageRequested(currentStep, signMessage.message);
    },

    async _doMigrateLikerID(
      currentStep: StepStateStep3Signing,
      cosmosSigningMessage: string,
      cosmosSignature: StdSignature,
      ethSignature: string
    ): Promise<StepStateStep3Signing | StepStateStep4Init> {
      await this.migrateLikerID({
        cosmos_address: currentStep.cosmosAddress,
        cosmos_pub_key: cosmosSignature.pub_key.value,
        cosmos_signature: cosmosSignature.signature,
        eth_address: currentStep.ethAddress,
        eth_signature: ethSignature,
        like_id: currentStep.likerId,
        cosmos_signing_message: cosmosSigningMessage,
        eth_signing_message: currentStep.signMessage,
      });
      const userProfile = await makeGetUserProfileAPI(
        currentStep.cosmosAddress
      )(this.$apiClient)();
      const remoteEthAddress = userProfile.user_profile.eth_wallet_address;
      if (remoteEthAddress != null) {
        return likerIdMigrated(
          currentStep,
          userProfile.user_profile.liker_id,
          userProfile.user_profile.avatar,
          remoteEthAddress
        );
      }
      return currentStep;
    },

    async _getOrCreateMigrationPreview(
      s: StepStateStep4Init | StepStateStep4MigrationPreview
    ): Promise<StepStateStep4MigrationPreview> {
      const migrationPreview = await this._fetchMigrationPreview(
        s.cosmosAddress
      );

      if (migrationPreview == null) {
        const newMigrationPreview = await this._createMigrationPreview(
          s.cosmosAddress
        );
        return migrationPreviewFetched(s, newMigrationPreview);
      } else {
        return migrationPreviewFetched(s, migrationPreview);
      }
    },

    async _fetchMigrationPreview(cosmosWalletAddress: string) {
      try {
        const migrationResponse = await makeGetMigrationPreviewAPI(
          cosmosWalletAddress
        )(this.$apiClient)();
        return migrationResponse.preview;
      } catch (e) {
        if (isAxiosError(e)) {
          if (e.status === 404) {
            return null;
          }
        }
        throw e;
      }
    },

    async _createMigrationPreview(cosmosWalletAddress: string) {
      const migrationResponse = await makeCreateMigrationPreviewAPI(
        this.$apiClient
      )({ cosmos_address: cosmosWalletAddress });
      return migrationResponse.preview;
    },

    async _createMigration(
      s: StepStateStep4MigrationPreview
    ): Promise<StepStateStep5MigrationResult> {
      const migrationResponse = await makeCreateMigrationAPI(this.$apiClient)({
        asset_snapshot_id: s.migrationPreview.id,
        cosmos_address: s.cosmosAddress,
        eth_address: s.ethAddress,
      });
      return migrationResultFetched(s, migrationResponse.migration);
    },

    async _retryMigration(
      s: StepStateStep4MigrationRetryPreview
    ): Promise<StepStateStep5MigrationResult> {
      const migrationResponse = await this.retryMigration(s.cosmosAddress, {
        book_nft_collection: s.failedMigration.classes
          .filter((c) => c.status === 'failed')
          .map((c) => c.cosmos_class_id),
        book_nft: s.failedMigration.nfts
          .filter((n) => n.status === 'failed')
          .map((n) => ({
            class_id: n.cosmos_class_id,
            nft_id: n.cosmos_nft_id,
          })),
      });
      return migrationResultFetched(s, migrationResponse.migration);
    },

    async _refreshMigration(
      s: StepStateStep5MigrationResult
    ): Promise<
      StepStateStep5MigrationResult | StepStateCompleted | StepStateFailed
    > {
      const resp = await makeGetMigrationAPI(s.cosmosAddress)(
        this.$apiClient
      )();
      // expect throw on error
      if (isMigrationCompleted(resp.migration)) {
        return migrationCompleted(s, resp.migration);
      }
      if (isMigrationFailed(resp.migration)) {
        return migrationFailed(s, resp.migration);
      }
      return migrationResultFetched(s, resp.migration);
    },

    async _checkMigration(
      s: StepStateStep4Init
    ): Promise<
      | StepStateStep4MigrationPreview
      | StepStateStep5MigrationResult
      | StepStateCompleted
      | StepStateFailed
    > {
      try {
        const resp = await makeGetMigrationAPI(s.cosmosAddress)(
          this.$apiClient
        )();
        if (isMigrationCompleted(resp.migration)) {
          return migrationCompleted(s, resp.migration);
        }
        if (isMigrationFailed(resp.migration)) {
          return migrationFailed(s, resp.migration);
        }
        return migrationResultFetched(s, resp.migration);
      } catch (e) {
        if (isAxiosError(e)) {
          if (e.status === 404) {
            return this._getOrCreateMigrationPreview(s);
          }
        }
        throw e;
      }
    },

    _formatDate(value: Date) {
      return formatDate(value, 'dd MMM, yyyy HH:mm:ss');
    },

    _formatNumber(value: number | string) {
      return numeral(value).format('0,0');
    },
  },
});
</script>
