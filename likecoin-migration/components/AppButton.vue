<template>
  <div v-if="disabled" :class="outerClass">
    <div :class="innerClass">
      <slot />
    </div>
  </div>
  <a v-else-if="href" :href="href" :class="outerClass" v-on="$listeners">
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
    disabled: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    outerClass() {
      switch (this.variant) {
        case 'secondary':
          return [
            {
              'bg-likecoin-white': !this.disabled,
              'bg-likecoin-grey': this.disabled,
            },
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
            {
              'bg-like-cyan-light': !this.disabled,
              'bg-likecoin-grey': this.disabled,
              'text-like-green': !this.disabled,
              'text-likecoin-black': this.disabled,
            },
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
            'px-[16px]',
            'py-[10px]',
            'w-full',
            'font-semibold',
            ...(!this.disabled
              ? [
                  'hover:bg-white',
                  'hover:bg-opacity-30',
                  'active:bg-opacity-20',
                  'transition duration-200',
                  'active:bg-like-green',
                ]
              : []),
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
            'px-[16px]',
            'py-[10px]',
            'w-full',
            'font-semibold',
            ...(!this.disabled
              ? [
                  'hover:bg-white',
                  'hover:bg-opacity-30',
                  'active:bg-opacity-20',
                  'transition duration-200',
                  'active:bg-like-green',
                ]
              : []),
          ];
      }
    },
  },
});
</script>
