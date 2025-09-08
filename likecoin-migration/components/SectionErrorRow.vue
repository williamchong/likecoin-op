<template>
  <div :class="['flex', 'items-center']">
    <div :class="['flex-1']">
      <p
        v-if="errorMessage != null"
        :class="['text-likecoin-votecolor-no', 'text-xs']"
      >
        <FontAwesomeIcon :class="['font-bold']" icon="triangle-exclamation" />
        {{ errorMessage }}
      </p>
    </div>
    <div v-if="!!$listeners['retryClick']">
      <AppButton
        variant="primary"
        :class="['w-full']"
        @click="handleRetryClick"
      >
        {{ retryButtonText ?? $t('section.error-row.retry') }}
      </AppButton>
    </div>
  </div>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue';

export default Vue.extend({
  props: {
    errorMessage: {
      type: String as PropType<string>,
      required: true,
    },
    retryButtonText: {
      type: String as PropType<string>,
      required: true,
    },
  },

  methods: {
    handleRetryClick() {
      this.$emit('retryClick');
    },
  },
});
</script>
