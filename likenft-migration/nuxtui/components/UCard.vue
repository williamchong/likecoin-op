<!-- https://github.com/nuxt/ui/blob/dev/src/runtime/components/layout/Card.vue -->

<template>
  <div
    :class="[
      _ui.base,
      _ui.rounded,
      _ui.divide,
      _ui.ring,
      _ui.shadow,
      _ui.background,
    ]"
  >
    <div
      v-if="$slots.header"
      :class="[_ui.header.base, _ui.header.padding, _ui.header.background]"
    >
      <slot name="header" />
    </div>
    <div
      v-if="$slots.default"
      :class="[_ui.body.base, _ui.body.padding, _ui.body.background]"
    >
      <slot />
    </div>
    <div
      v-if="$slots.footer"
      :class="[_ui.footer.base, _ui.footer.padding, _ui.footer.background]"
    >
      <slot name="footer" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  props: {
    as: {
      type: String,
      default: 'div',
    },
    ui: {
      type: Object,
      default: () => ({}),
    },
  },
  computed: {
    _ui() {
      // https://github.com/nuxt/ui/blob/dev/src/runtime/ui.config/layout/card.ts
      return {
        base: '',
        background: 'bg-white dark:bg-gray-900',
        divide: 'divide-y divide-gray-200 dark:divide-gray-800',
        ring: 'ring-1 ring-gray-200 dark:ring-gray-800',
        rounded: 'rounded-lg',
        shadow: 'shadow',
        ...this.ui,
        body: {
          base: '',
          background: '',
          padding: 'px-4 py-5 sm:p-6',
          ...this.ui.body,
        },
        header: {
          base: '',
          background: '',
          padding: 'px-4 py-5 sm:px-6',
          ...this.ui.header,
        },
        footer: {
          base: '',
          background: '',
          padding: 'px-4 py-4 sm:px-6',
          ...this.ui.footer,
        },
      };
    },
  },
});
</script>
