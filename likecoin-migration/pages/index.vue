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
            {{ $t('section.connect-wallet.title') }}
          </h2>
        </StepSection>
        <StepSection
          v-slot="{ isCurrent, isPast }"
          :step="3"
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
            {{ $t('section.confirm-by-signing.title') }}
          </h2>
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
import Vue from 'vue';

import { StepState } from '~/pageModels';

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

  methods: {
    handleIntroductionSectionConfirmClick() {
      alert('TODO introduction confirmed');
    },
  },
});
</script>
