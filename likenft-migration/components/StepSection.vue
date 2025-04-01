<template>
  <div :class="['flex', 'flex-row', 'gap-6']">
    <div
      :class="[
        'self-stretch',

        'flex',
        'flex-col',
        'items-center',
        'gap-0.5',

        'after:flex-1',
        'after:w-px',
        'after:bg-likecoin-grey',
      ]"
    >
      <div
        :class="[
          {
            'text-likecoin-darkgreen':
              stepState === 'current' || stepState === 'past',
            'text-likecoin-votecolor-abstain': stepState === 'future',
          },
          'text-2xl',
        ]"
      >
        <FontAwesomeIcon :icon="icon" :class="['font-black']" />
      </div>
    </div>
    <div :class="['flex-1', 'min-w-0']">
      <slot />
      <slot v-if="stepState === 'past'" name="past" />
      <slot v-if="stepState === 'current'" name="current" />
      <slot v-if="stepState === 'future'" name="future" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

import { SupportedIcon } from '~/models/faIcon';

type StepState = 'past' | 'current' | 'future';

const stepIconMapping: { [key in number]: SupportedIcon | undefined } = {
  1: 'circle-1',
  2: 'circle-2',
  3: 'circle-3',
  4: 'circle-4',
  5: 'circle-5',
};

export default Vue.extend({
  props: {
    step: {
      type: Number,
      required: true,
    },
    currentStep: {
      type: Number,
      required: true,
    },
  },
  computed: {
    stepState(): StepState {
      if (this.currentStep === this.step) {
        return 'current';
      }
      if (this.currentStep < this.step) {
        return 'future';
      }
      return 'past';
    },
    icon(): SupportedIcon | undefined {
      if (this.stepState === 'past') {
        return 'circle-check';
      }

      return stepIconMapping[this.step];
    },
  },
});
</script>
