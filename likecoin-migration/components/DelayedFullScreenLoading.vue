<template>
  <div
    v-if="didBlock"
    :class="[
      'fixed',
      'top-0',
      'left-0',
      'w-full',
      'h-full',
      'flex',
      'items-center',
      'justify-center',
      'bg-white/70',
      'transition-opacity',
      'ease-in',
      'duration-500',
      'opacity-0',
      {
        'opacity-100': didDisplay,
      },
    ]"
  >
    <LoadingIcon />
  </div>
</template>
<script lang="ts">
import Vue from 'vue';

interface Data {
  effectTimeout: ReturnType<typeof setTimeout> | null;
  didBlock: boolean;
  didDisplay: boolean;
}

export default Vue.extend({
  props: {
    isLoading: {
      type: Boolean,
      required: true,
    },
  },

  data(): Data {
    return {
      effectTimeout: null,
      didBlock: false,
      didDisplay: false,
    };
  },

  computed: {
    displayDelayMs(): number {
      return this.$appConfig.transitionLoadingScreenDelayMs;
    },
  },

  watch: {
    isLoading(isLoading: boolean) {
      if (this.effectTimeout != null) {
        clearTimeout(this.effectTimeout);
        this.effectTimeout = null;
      }

      if (!isLoading) {
        this.didBlock = false;
      } else {
        this.effectTimeout = setTimeout(() => {
          this.didBlock = true;
          this.didDisplay = false;
          setTimeout(() => {
            this.didDisplay = true;
          });
        }, this.displayDelayMs);
      }
    },
  },
});
</script>
