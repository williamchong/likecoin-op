<template>
  <div :class="['flex-1', 'min-h-0', 'bg-likecoin-lightergrey', 'pb-4']">
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
          <template #default="{ isCurrent, isPast }">
            <h2
              :class="[
                'text-base',
                {
                  ['font-semibold']: isCurrent() || isPast(),
                },
                'leading-[30px]',
                'text-likecoin-darkgrey',
              ]"
            >
              {{ $t('section.introduction.title') }}
            </h2>
          </template>
          <template #current>
            <div :class="['mt-2', 'mb-5']">
              <SectionIntroduction
                @confirmClicked="handleIntroductionSectionConfirmClick"
              />
            </div>
          </template>
        </StepSection>
        <StepSection :step="2" :current-step="currentStep.step">
          <template #default="{ isCurrent, isPast }">
            <h2
              :class="[
                'text-base',
                {
                  ['font-semibold']: isCurrent() || isPast(),
                },
                'leading-[30px]',
                'text-likecoin-darkgrey',
              ]"
            >
              {{ $t('section.connect-wallet.title') }}
            </h2>
          </template>
          <template #current>
            <SectionWalletConnect
              :class="['mt-2.5', 'mb-4']"
              :liker-id="likerId"
              :avatar="avatar"
              :cosmos-address="cosmosAddress"
              :current-balance="currentBalance"
              :eth-address="ethAddress"
              :estimated-balance="estimatedBalance"
              @likeCoinWalletConnected="handleLikeCoinWalletConnected"
              @likeCoinEVMWalletConnected="handleLikeCoinEVMWalletConnected"
            />
          </template>
          <template #past>
            <SectionWalletConnect
              :class="['mt-2.5', 'mb-4']"
              :liker-id="likerId"
              :avatar="avatar"
              :cosmos-address="cosmosAddress"
              :current-balance="currentBalance"
              :eth-address="ethAddress"
              :estimated-balance="estimatedBalance"
            />
          </template>
        </StepSection>
        <StepSection :step="3" :current-step="currentStep.step">
          <template #future>
            <h2
              :class="['text-base', 'leading-[30px]', 'text-likecoin-darkgrey']"
            >
              {{ $t('section.confirm-by-signing.title') }}
            </h2>
          </template>
          <template #current>
            <SectionEvmSign
              :signing-message="
                currentStep.step === 3 ? currentStep.ethSigningMessage : ''
              "
              @signed="handleEvmSigned"
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
            </SectionEvmSign>
          </template>
          <template #past>
            <h2
              :class="[
                'text-base',
                'leading-[30px]',
                'font-semibold',
                'text-likecoin-darkgrey',
              ]"
            >
              {{ $t('section.confirm-by-signing.title') }}
            </h2>
          </template>
        </StepSection>
        <StepSection
          v-slot="{ isCurrent, isPast }"
          :step="4"
          :current-step="currentStep.step"
        >
          <h2
            :class="[
              'text-base',
              {
                ['font-semibold']: isCurrent() || isPast(),
              },
              'leading-[30px]',
              'text-likecoin-darkgrey',
            ]"
          >
            {{ $t('section.start-migration.title') }}
          </h2>
        </StepSection>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { SigningStargateClient } from '@cosmjs/stargate';
import { LikeCoinWalletConnectorConnectionResult } from '@likecoin/wallet-connector';
import { Decimal } from 'decimal.js';
import Vue from 'vue';

import {
  CreateEthSigningMessage,
  makeCreateEthSigningMessageAPI,
} from '~/apis/createEthSigningMessage';
import { GetUserProfile, makeGetUserProfileAPI } from '~/apis/getUserProfile';
import { ChainCoin } from '~/models/cosmosNetworkConfig';
import {
  EitherEthConnected,
  EthConnected,
  ethSignConfirming,
  evmConnected,
  gasEstimated,
  initCosmosConnected,
  introductionConfirmed,
  isEthConnected,
  likerIdResolved,
  StepState,
  StepStateStep2CosmosConnected,
  StepStateStep2GasEstimated,
  StepStateStep2LikerIdResolved,
  StepStateStep3AwaitSignature,
} from '~/pageModels';

interface Data {
  isTransitioning: boolean;

  currentStep: StepState;
}

export default Vue.extend({
  data(): Data {
    return {
      isTransitioning: false,

      currentStep: { step: 1 },
    };
  },

  computed: {
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
    currentBalance(): ChainCoin | null {
      if ('currentBalance' in this.currentStep) {
        return this.currentStep.currentBalance;
      }
      return null;
    },
    ethAddress(): string | null {
      if ('ethAddress' in this.currentStep) {
        return this.currentStep.ethAddress;
      }
      return null;
    },
    estimatedBalance(): ChainCoin | null {
      if ('estimatedBalance' in this.currentStep) {
        return this.currentStep.estimatedBalance;
      }
      return null;
    },

    getUserProfile(): GetUserProfile {
      if (this.cosmosAddress == null) {
        throw new Error('cosmos address not connected');
      }
      return makeGetUserProfileAPI(this.cosmosAddress)(this.$apiClient);
    },
    createEthSigningMessage(): CreateEthSigningMessage {
      return makeCreateEthSigningMessageAPI(this.$apiClient);
    },
  },

  methods: {
    handleIntroductionSectionConfirmClick() {
      if (this.currentStep.step !== 1) {
        return;
      }
      this.currentStep = introductionConfirmed(this.currentStep);
    },

    async handleLikeCoinWalletConnected(
      cosmosAddress: string,
      connection: LikeCoinWalletConnectorConnectionResult
    ) {
      if (this.currentStep.step === 1) {
        return;
      }

      this.currentStep = initCosmosConnected(
        this.currentStep,
        cosmosAddress,
        connection
      );

      this.currentStep = await this._asyncStateTransition(
        this.currentStep,
        this._resolveLikerId
      );

      if (this.currentStep.state === 'LikerIdResolved') {
        this.currentStep = await this._asyncStateTransition(
          this.currentStep,
          this._estimateBalance
        );
      }

      if (
        this.currentStep.step === 2 &&
        this.currentStep.state === 'GasEstimated' &&
        isEthConnected(this.currentStep)
      ) {
        this.currentStep = await this._asyncStateTransition(
          this.currentStep,
          this._resolveEvmSigningMessage
        );
      }
    },

    async handleLikeCoinEVMWalletConnected(ethAddress: string) {
      if (this.currentStep.step === 2) {
        this.currentStep = evmConnected(this.currentStep, ethAddress);
      }

      if (
        this.currentStep.step === 2 &&
        this.currentStep.state === 'GasEstimated' &&
        isEthConnected(this.currentStep)
      ) {
        this.currentStep = await this._asyncStateTransition(
          this.currentStep,
          this._resolveEvmSigningMessage
        );
      }
    },

    handleEvmSigned(signature: string) {
      if (this.currentStep.step !== 3) {
        return;
      }
      alert(`TODO create migration ${signature}`);
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

    async _resolveLikerId(
      s: EitherEthConnected<StepStateStep2CosmosConnected>
    ): Promise<EitherEthConnected<StepStateStep2LikerIdResolved>> {
      const userProfile = await this.getUserProfile();
      return likerIdResolved(
        s,
        userProfile.user_profile.avatar,
        userProfile.user_profile.liker_id
      );
    },

    async _estimateBalance(
      s: EitherEthConnected<StepStateStep2LikerIdResolved>
    ): Promise<EitherEthConnected<StepStateStep2GasEstimated>> {
      const { offlineSigner } = s.connection;
      const client = await SigningStargateClient.connectWithSigner(
        this.$likeCoinWalletConnector.options.rpcURL,
        // @ts-expect-error
        offlineSigner
      );

      const balance = (await client.getBalance(
        s.cosmosAddress,
        this.$cosmosNetworkConfig.coinLookup[0].chainDenom
      )) as unknown as ChainCoin;

      const gasEstimation = await this._estimateGas(
        client,
        s.cosmosAddress,
        balance,
        ''
      );

      // this is the tier selection from keplr sign dialog
      const tierMultiplier = {
        low: 1000,
        average: 10000,
        high: 1000000,
      };

      const gasFee = Decimal(
        gasEstimation *
          // Assume worst case user select high without insufficient fund
          // in the signing ui
          tierMultiplier.high
      );

      const estimatedBalance: ChainCoin = {
        denom: balance.denom,
        amount: Decimal(balance.amount).minus(gasFee).toString(),
      };

      return gasEstimated(s, client, balance, gasEstimation, estimatedBalance);
    },

    async _resolveEvmSigningMessage(
      prev: EthConnected<StepStateStep2GasEstimated>
    ): Promise<StepStateStep3AwaitSignature> {
      const m = await this.createEthSigningMessage({
        amount: prev.estimatedBalance,
      });
      return ethSignConfirming(prev, m.signing_message);
    },

    async _estimateGas(
      signingStargateClient: SigningStargateClient,
      cosmosAddress: string,
      amount: ChainCoin,
      memo: string
    ) {
      const msg = {
        typeUrl: '/cosmos.bank.v1beta1.MsgSend',
        value: {
          fromAddress: cosmosAddress,
          toAddress: this.$appConfig.cosmosDepositAddress,
          amount: [amount],
        },
      };
      const fluctuationMultiplier = 1.3;
      const gasEstimation = await signingStargateClient.simulate(
        cosmosAddress,
        [msg],
        memo
      );
      return Math.floor(gasEstimation * fluctuationMultiplier);
    },
  },
});
</script>
