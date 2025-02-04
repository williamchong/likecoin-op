<template>
  <a
    v-if="href"
    :href="href"
    class="bg-like-cyan-light text-like-green h-44px rounded-[10px] flex box-border overflow-hidden items-center cursor-pointer transition duration-200"
    v-on="$listeners"
  >
    <div
      class="flex items-center text-[16px] leading-[1.35] justify-between h-full text-center whitespace-nowrap hover:bg-white hover:bg-opacity-30 active:bg-opacity-20 transition duration-200 active:bg-like-green px-[16px] py-[10px] w-full"
    >
      <slot />
    </div>
  </a>
  <nuxt-link
    v-else-if="to"
    :to="to"
    class="bg-like-cyan-light text-like-green h-44px rounded-[10px] flex box-border overflow-hidden items-center cursor-pointer transition duration-200"
    v-on="$listeners"
  >
    <div
      class="flex items-center text-[16px] leading-[1.35] justify-between h-full text-center whitespace-nowrap hover:bg-white hover:bg-opacity-30 active:bg-opacity-20 transition duration-200 active:bg-like-green px-[16px] py-[10px] w-full"
    >
      <slot />
    </div>
  </nuxt-link>
  <button
    v-else
    class="bg-like-cyan-light text-like-green h-44px rounded-[10px] flex box-border overflow-hidden items-center cursor-pointer transition duration-200"
    v-on="$listeners"
  >
    <div
      class="flex items-center text-[16px] leading-[1.35] justify-between h-full text-center whitespace-nowrap hover:bg-white hover:bg-opacity-30 active:bg-opacity-20 transition duration-200 active:bg-like-green px-[16px] py-[10px] w-full"
    >
      <slot />
    </div>
  </button>
</template>

<script lang="ts">
import Vue from 'vue';

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
  },
});
</script>
