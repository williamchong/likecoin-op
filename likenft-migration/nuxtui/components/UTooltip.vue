<template>
  <div
    ref="trigger"
    :class="[_ui.wrapper]"
    v-bind="$attrs"
    @mouseenter="onMouseEnter"
    @mouseleave="onMouseLeave"
  >
    <div :class="['relative']">
      <slot :open="open"> Hover </slot>
      <Transition appear v-bind="_ui.transition">
        <div
          v-if="open && !prevent && isVisible"
          ref="container"
          :class="[_ui.container, _ui.width, 'absolute', 'top-0', 'left-full']"
        >
          <div
            :class="[
              _ui.base,
              _ui.background,
              _ui.color,
              _ui.rounded,
              _ui.shadow,
              _ui.ring,
            ]"
          >
            <slot name="text">
              {{ text }}
            </slot>
          </div>
        </div>
      </Transition>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

interface Data {
  open: boolean;
  openTimeout: ReturnType<typeof setTimeout> | null;
  closeTimeout: ReturnType<typeof setTimeout> | null;
}

export default Vue.extend({
  props: {
    text: {
      type: String,
      default: null,
    },
    prevent: {
      type: Boolean,
      default: false,
    },
    openDelay: {
      type: Number,
      default: 0,
    },
    closeDelay: {
      type: Number,
      default: 0,
    },
    ui: {
      type: Object,
      default: () => ({}),
    },
  },
  data(): Data {
    return {
      open: false,
      openTimeout: null,
      closeTimeout: null,
    };
  },
  computed: {
    isVisible() {
      return !!this.text;
    },
    _ui() {
      return {
        wrapper: 'relative inline-flex',
        container: 'z-20 group',
        width: 'max-w-xs',
        background: 'bg-white dark:bg-gray-900',
        color: 'text-gray-900 dark:text-white',
        shadow: 'shadow',
        rounded: 'rounded',
        ring: 'ring-1 ring-gray-200 dark:ring-gray-800',
        base: '[@media(pointer:coarse)]:hidden h-6 px-2 py-1 text-xs font-normal truncate relative',
        shortcuts: 'hidden md:inline-flex flex-shrink-0 gap-0.5',
        middot: 'mx-1 text-gray-700 dark:text-gray-200',
        // Syntax for `<Transition>` component https://vuejs.org/guide/built-ins/transition.html#css-based-transitions
        ...this.ui,
        transition: {
          enterActiveClass: 'transition ease-out duration-200',
          enterFromClass: 'opacity-0 translate-y-1',
          enterToClass: 'opacity-100 translate-y-0',
          leaveActiveClass: 'transition ease-in duration-150',
          leaveFromClass: 'opacity-100 translate-y-0',
          leaveToClass: 'opacity-0 translate-y-1',
          ...this.ui.transition,
        },
      };
    },
  },
  methods: {
    handleTriggerMouseOver() {
      this.open = true;
    },

    onMouseEnter() {
      // cancel programmed closing
      if (this.closeTimeout) {
        clearTimeout(this.closeTimeout);
        this.closeTimeout = null;
      }
      // dropdown already open
      if (this.open) {
        return;
      }
      this.openTimeout =
        this.openTimeout ||
        setTimeout(() => {
          this.open = true;
          this.openTimeout = null;
        }, this.openDelay);
    },

    onMouseLeave() {
      // cancel programmed opening
      if (this.openTimeout) {
        clearTimeout(this.openTimeout);
        this.openTimeout = null;
      }
      // dropdown already closed
      if (!this.open) {
        return;
      }
      this.closeTimeout =
        this.closeTimeout ||
        setTimeout(() => {
          this.open = false;
          this.closeTimeout = null;
        }, this.closeDelay);
    },
  },
});
</script>
