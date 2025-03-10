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
              {{ $t('migrate.preview') }}
            </h2>
          </template>
          <template v-if="migrationPreview != null" #current>
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
                  migrationPreview.block_time != null &&
                  migrationPreview.block_height != null
                "
                :text="
                  $t('section.asset-preview.tooltip', {
                    date: _formatDate(migrationPreview.block_time),
                    height: _formatNumber(migrationPreview.block_height),
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
              v-if="migrationPreview != null"
              :class="['max-w-full', 'mt-2']"
              :snapshot="migrationPreview"
            />
          </template>
        </StepSection>
        <StepSection :step="4" :current-step="currentStep.step"></StepSection>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { isAxiosError } from 'axios';
import { format as formatDate } from 'date-fns/format';
import numeral from 'numeral';
import Vue from 'vue';

import { makeCreateMigrationPreviewAPI } from '~/apis/createMigrationPreview';
import { makeGetMigrationPreviewAPI } from '~/apis/getMigrationPreview';
import { getSignMessage } from '~/apis/getSignMessage';
import { makeGetUserProfileAPI } from '~/apis/getUserProfile';
import { makeMigrateLikerIDAPI } from '~/apis/migrateLikerID';
import { LikeNFTAssetSnapshot } from '~/apis/models/likenftAssetSnapshot';
import {
  initCosmosConnected,
  initEvmConnected,
  initMigrationPreview,
  introductionConfirmed,
  likerIdEvmConnected,
  likerIdMigrated,
  likerIdResolved,
  migrationPreviewFetched,
  StepState,
  StepStateStep2CosmosConnected,
  StepStateStep2LikerIdEvmConnected,
  StepStateStep2LikerIdMigrated,
  StepStateStep2LikerIdResolved,
  StepStateStep3Init,
  StepStateStep3MigrationPreview,
} from '~/pageModels';

import UTooltip from '../nuxtui/components/UTooltip.vue';

interface Data {
  isTransitioning: boolean;

  currentStep: StepState;

  migrationPreviewFetchTimeout: ReturnType<typeof setTimeout> | null;
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
        this.currentStep.step !== 3 ||
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
  },
  methods: {
    handleIntroductionSectionConfirmClick() {
      if (this.currentStep.step !== 1) {
        return;
      }
      this.currentStep = introductionConfirmed(this.currentStep);
    },

    async handleLikeCoinWalletConnected(cosmosAddress: string) {
      if (this.currentStep.step === 1) {
        return;
      }

      this.currentStep = initCosmosConnected(this.currentStep, cosmosAddress);
      this.currentStep = await this._asyncStateTransition(
        this.currentStep,
        (s) => this._checkLikerID(s, cosmosAddress)
      );

      if (this.currentStep.state === 'LikerIdMigrated') {
        this.currentStep = initMigrationPreview(this.currentStep);
      }

      if (this.currentStep.step === 3) {
        this.currentStep = await this._asyncStateTransition(
          this.currentStep,
          (s) => this._getOrCreateMigrationPreview(s)
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
            (s) => this._doMigrateLikerID(s, s.cosmosAddress, s.ethAddress)
          );
        }
      }

      if (this.currentStep.state === 'LikerIdMigrated') {
        this.currentStep = initMigrationPreview(this.currentStep);
      }

      if (this.currentStep.step === 3) {
        this.currentStep = await this._asyncStateTransition(
          this.currentStep,
          (s) => this._getOrCreateMigrationPreview(s)
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
    ): Promise<StepStateStep2LikerIdResolved | StepStateStep2LikerIdMigrated> {
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

    async _doMigrateLikerID(
      currentStep: StepStateStep2LikerIdEvmConnected,
      cosmosAddress: string,
      ethAddress: string
    ): Promise<
      StepStateStep2LikerIdEvmConnected | StepStateStep2LikerIdMigrated
    > {
      const signMessage = await this.getSignMessage({
        cosmos_address: cosmosAddress,
        eth_address: ethAddress,
        liker_id: currentStep.likerId,
      });
      const connection = await this.$likeCoinWalletConnector.initIfNecessary();
      if (connection == null) {
        alert('cannot get wallet connector connection');
        return currentStep;
      }
      const {
        accounts: [account],
        offlineSigner,
      } = connection;

      if (!offlineSigner.signArbitrary) {
        alert('signArbitrary not supported');
        return currentStep;
      }

      const result = await offlineSigner.signArbitrary(
        this.$likeCoinWalletConnector.options.chainId,
        account.address,
        signMessage.message
      );
      const cosmosSignature = result.signature;

      const signedMessage =
        await this.$likeCoinEVMWalletConnector.connector.signMessage(
          signMessage.message
        );

      await this.migrateLikerID({
        cosmos_pub_key: result.pub_key.value,
        cosmos_signature: cosmosSignature,
        eth_address: ethAddress,
        eth_signature: signedMessage,
        like_id: currentStep.likerId,
        signing_message: signMessage.message,
      });

      // Check again on likerland to see if eth address is migrated
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
      }

      return currentStep;
    },

    async _getOrCreateMigrationPreview(
      s: StepStateStep3Init | StepStateStep3MigrationPreview
    ): Promise<StepStateStep3MigrationPreview> {
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

    _formatDate(value: Date) {
      return formatDate(value, 'dd MMM, yyyy HH:mm:ss');
    },

    _formatNumber(value: number | string) {
      return numeral(value).format('0,0');
    },
  },
});
</script>
