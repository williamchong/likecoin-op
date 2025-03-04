<template>
  <span :class="['font-[fa]']">{{ unicode[icon] }}</span>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue';

import m, { SupportedIcon } from '~/models/faIcon';

export default Vue.extend({
  props: {
    icon: {
      type: String as PropType<SupportedIcon>,
      required: true,
    },
  },
  computed: {
    unicode() {
      return m;
    },
  },
});
</script>
