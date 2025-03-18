<template>
  <div :class="['p-4', 'rounded-xl', 'border', 'border-likecoin-grey']">
    <p
      :class="['text-sm', 'text-likecoin-darkgrey', 'whitespace-pre-line']"
      v-text="$t('section.introduction.text')"
    />
    <div :class="['flex', 'flex-row', 'justify-end', 'mt-6']">
      <AppButton @click="handleConfirmClick">{{
        $t('section.introduction.confirm')
      }}</AppButton>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  methods: {
    handleConfirmClick() {
      this.$emit('confirmClicked');
    },
  },
});
</script>
