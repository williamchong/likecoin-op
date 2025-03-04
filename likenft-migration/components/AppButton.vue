<template>
  <a v-if="href" :href="href" :class="outerClass" v-on="$listeners">
    <div :class="innerClass">
      <slot />
    </div>
  </a>
  <nuxt-link v-else-if="to" :to="to" :class="outerClass" v-on="$listeners">
    <div :class="innerClass">
      <slot />
    </div>
  </nuxt-link>
  <button v-else :class="outerClass" v-on="$listeners">
    <div :class="innerClass">
      <slot />
    </div>
  </button>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue';

type ButtonType = 'primary' | 'secondary';

export default Vue.extend({
  props: {
    href: {
      type: String,
      default: '',
    },
    to: {
      type: String,
      default: '',
    },
    variant: {
      type: String as PropType<ButtonType>,
      default: 'primary',
    },
  },
  computed: {
    outerClass() {
      switch (this.variant) {
        case 'secondary':
          return [
            'bg-likecoin-white',
            'text-likecoin-black',
            'h-44px',
            'rounded-[10px]',
            'flex',
            'box-border',
            'overflow-hidden',
            'items-center',
            'cursor-pointer',
            'transition',
            'duration-200',
          ];
        default:
          return [
            'bg-like-cyan-light',
            'text-like-green',
            'h-44px',
            'rounded-[10px]',
            'flex',
            'box-border',
            'overflow-hidden',
            'items-center',
            'cursor-pointer',
            'transition',
            'duration-200',
          ];
      }
    },
    innerClass() {
      switch (this.variant) {
        case 'secondary':
          return [
            'flex',
            'items-center',
            'justify-center',
            'text-[16px]',
            'leading-[1.35]',
            'h-full',
            'text-center',
            'whitespace-nowrap',
            'hover:bg-white',
            'hover:bg-opacity-30',
            'active:bg-opacity-20',
            'transition duration-200',
            'active:bg-like-green',
            'px-[16px]',
            'py-[10px]',
            'w-full',
            'font-semibold',
          ];
        default:
          return [
            'flex',
            'items-center',
            'justify-center',
            'text-[16px]',
            'leading-[1.35]',
            'h-full',
            'text-center',
            'whitespace-nowrap',
            'hover:bg-white',
            'hover:bg-opacity-30',
            'active:bg-opacity-20',
            'transition duration-200',
            'active:bg-like-green',
            'px-[16px]',
            'py-[10px]',
            'w-full',
            'font-semibold',
          ];
      }
    },
  },
});
</script>
