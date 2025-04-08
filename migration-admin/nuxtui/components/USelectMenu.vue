<!-- https://github.com/nuxt/ui/blob/dev/src/runtime/components/forms/SelectMenu.vue -->

<template>
  <div :class="[_ui.wrapper]">
    <select
      :value="value"
      v-bind="$attrs"
      :class="[
        _ui.base,
        _ui.form,
        'rounded-md',
        _ui.variant.outline,
        'ring-gray-300',
        _ui.padding.xs,
      ]"
      @input="$emit('input', $event.currentTarget.value)"
    >
      <option
        v-for="(option, index) in options"
        :key="index"
        :value="
          valueAttribute && option instanceof Object
            ? accessor(option, valueAttribute)
            : option
        "
      >
        {{
          option instanceof Object ? accessor(option, optionAttribute) : option
        }}
      </option>
    </select>
  </div>
</template>

<script lang="ts">
import Vue, { PropType } from "vue";

export default Vue.extend({
  props: {
    value: {
      type: [String, Number, Object, Array, Boolean] as PropType<
        string | number | readonly string[] | undefined
      >,
      default: "",
    },
    options: {
      type: Array as PropType<
        { [key: string]: any; disabled?: boolean }[] | string[]
      >,
      default: () => [],
    },
    optionAttribute: {
      type: String,
      default: "label",
    },
    valueAttribute: {
      type: String,
      default: null,
    },
    ui: {
      type: Object,
      default: () => ({}),
    },
  },
  computed: {
    _ui() {
      const input = {
        wrapper: "relative",
        base: "relative block w-full disabled:cursor-not-allowed disabled:opacity-75 focus:outline-none border-0",
        form: "form-input",
        rounded: "rounded-md",
        placeholder: "placeholder-gray-400 dark:placeholder-gray-500",
        ...this.ui,
        size: {
          "2xs": "text-xs",
          xs: "text-xs",
          sm: "text-sm",
          md: "text-sm",
          lg: "text-sm",
          xl: "text-base",
        },
        gap: {
          "2xs": "gap-x-1",
          xs: "gap-x-1.5",
          sm: "gap-x-1.5",
          md: "gap-x-2",
          lg: "gap-x-2.5",
          xl: "gap-x-2.5",
        },
        padding: {
          "2xs": "px-2 py-1",
          xs: "px-2.5 py-1.5",
          sm: "px-2.5 py-1.5",
          md: "px-3 py-2",
          lg: "px-3.5 py-2.5",
          xl: "px-3.5 py-2.5",
        },
        leading: {
          padding: {
            "2xs": "ps-7",
            xs: "ps-8",
            sm: "ps-9",
            md: "ps-10",
            lg: "ps-11",
            xl: "ps-12",
          },
        },
        trailing: {
          padding: {
            "2xs": "pe-7",
            xs: "pe-8",
            sm: "pe-9",
            md: "pe-10",
            lg: "pe-11",
            xl: "pe-12",
          },
        },
        color: {
          white: {
            outline:
              "shadow-sm bg-white dark:bg-gray-900 text-gray-900 dark:text-white ring-1 ring-inset ring-gray-300 dark:ring-gray-700 focus:ring-2 focus:ring-primary-500 dark:focus:ring-primary-400",
          },
          gray: {
            outline:
              "shadow-sm bg-gray-50 dark:bg-gray-800 text-gray-900 dark:text-white ring-1 ring-inset ring-gray-300 dark:ring-gray-700 focus:ring-2 focus:ring-primary-500 dark:focus:ring-primary-400",
          },
        },
        variant: {
          outline:
            "shadow-sm bg-transparent text-gray-900 dark:text-white ring-1 ring-inset ring-{color}-500 dark:ring-{color}-400 focus:ring-2 focus:ring-{color}-500 dark:focus:ring-{color}-400",
          none: "bg-transparent focus:ring-0 focus:shadow-none",
        },
        default: {
          size: "sm",
          color: "white",
          variant: "outline",
          loadingIcon: "i-heroicons-arrow-path-20-solid",
        },
      };
      return {
        ...input,
        form: "form-select",
        placeholder: "text-gray-400 dark:text-gray-500",
        default: {
          size: "sm",
          color: "white",
          variant: "outline",
          loadingIcon: "i-heroicons-arrow-path-20-solid",
          trailingIcon: "i-heroicons-chevron-down-20-solid",
        },
      };
    },
    uiMenu() {
      const inputMenu = {
        container: "z-20 group",
        trigger: "flex items-center w-full",
        width: "w-full",
        height: "max-h-60",
        base: "relative focus:outline-none overflow-y-auto scroll-py-1",
        background: "bg-white dark:bg-gray-800",
        shadow: "shadow-lg",
        rounded: "rounded-md",
        padding: "p-1",
        ring: "ring-1 ring-gray-200 dark:ring-gray-700",
        empty: "text-sm text-gray-400 dark:text-gray-500 px-2 py-1.5",
        option: {
          base: "cursor-default select-none relative flex items-center justify-between gap-1",
          rounded: "rounded-md",
          padding: "px-1.5 py-1.5",
          size: "text-sm",
          color: "text-gray-900 dark:text-white",
          container: "flex items-center gap-1.5 min-w-0",
          active: "bg-gray-100 dark:bg-gray-900",
          inactive: "",
          selected: "pe-7",
          disabled: "cursor-not-allowed opacity-50",
          empty: "text-sm text-gray-400 dark:text-gray-500 px-2 py-1.5",
          icon: {
            base: "flex-shrink-0 h-5 w-5",
            active: "text-gray-900 dark:text-white",
            inactive: "text-gray-400 dark:text-gray-500",
          },
          selectedIcon: {
            wrapper: "absolute inset-y-0 end-0 flex items-center",
            padding: "pe-2",
            base: "h-5 w-5 text-gray-900 dark:text-white flex-shrink-0",
          },
          avatar: {
            base: "flex-shrink-0",
          },
          chip: {
            base: "flex-shrink-0 w-2 h-2 mx-1 rounded-full",
          },
        },
        // Syntax for `<Transition>` component https://vuejs.org/guide/built-ins/transition.html#css-based-transitions
        transition: {
          leaveActiveClass: "transition ease-in duration-100",
          leaveFromClass: "opacity-100",
          leaveToClass: "opacity-0",
        },
        popper: {
          placement: "bottom-end",
        },
        default: {
          selectedIcon: "i-heroicons-check-20-solid",
          trailingIcon: "i-heroicons-chevron-down-20-solid",
          empty: {
            label: "No options.",
          },
          optionEmpty: {
            label: 'No results for "{query}".',
          },
        },
        arrow: {
          ring: "before:ring-1 before:ring-gray-200 dark:before:ring-gray-700",
          background: "before:bg-white dark:before:bg-gray-700",
        },
      };
      return {
        ...inputMenu,
        select: "inline-flex items-center text-left cursor-default",
        input:
          "block w-[calc(100%+0.5rem)] focus:ring-transparent text-sm px-3 py-1.5 text-gray-700 dark:text-gray-200 bg-white dark:bg-gray-800 border-0 border-b border-gray-200 dark:border-gray-700 sticky -top-1 -mt-1 mb-1 -mx-1 z-10 placeholder-gray-400 dark:placeholder-gray-500 focus:outline-none",
        required: "absolute inset-0 w-px opacity-0 cursor-default",
        label: "block truncate",
        ...this.ui.inputMenu,
        option: {
          ...inputMenu.option,
          create: "block truncate",
        },
        // Syntax for `<Transition>` component https://vuejs.org/guide/built-ins/transition.html#css-based-transitions
        transition: {
          leaveActiveClass: "transition ease-in duration-100",
          leaveFromClass: "opacity-100",
          leaveToClass: "opacity-0",
        },
        popper: {
          placement: "bottom-end",
        },
        default: {
          selectedIcon: "i-heroicons-check-20-solid",
          clearSearchOnClose: false,
          showCreateOptionWhen: "empty",
          searchablePlaceholder: {
            label: "Search...",
          },
          empty: {
            label: "No options.",
          },
          optionEmpty: {
            label: 'No results for "{query}".',
          },
        },
        arrow: {
          ring: "before:ring-1 before:ring-gray-200 dark:before:ring-gray-700",
          background: "before:bg-white dark:before:bg-gray-700",
        },
      };
    },
  },
  methods: {
    onUpdate(value: any) {
      this.$emit("update:modelValue", value);
    },
    accessor<T extends Record<string, any>>(obj: T, key: string) {
      return obj[key];
    },
  },
});
</script>
