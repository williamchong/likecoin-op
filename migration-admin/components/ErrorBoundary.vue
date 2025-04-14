<template>
  <div>
    <slot />
  </div>
</template>

<script lang="ts">
import { isAxiosError } from "axios";
import Vue from "vue";

export default Vue.extend({
  errorCaptured(err): boolean {
    const propagate = handleWalk(err, this.handleAxiosError);
    if (propagate) {
      // eslint-disable-next-line no-console
      console.warn(
        "Error is not handled. Will be propagated to outer components.",
        err
      );
    }
    return propagate;
  },

  methods: {
    handleAxiosError(err: Error): boolean {
      if (isAxiosError(err)) {
        if (err.status === 500 || err.message === "Network Error") {
          alert(err.message);
          return false;
        }
      }
      return true;
    },
  },
});

function handleWalk(
  err: Error,
  ...errorHandlers: ((err: Error) => boolean)[]
): boolean {
  return errorHandlers.reduce<boolean>((prev, curr) => {
    if (!prev) {
      // false, stop error propagation to next handler.
      return prev;
    }
    return curr(err);
  }, true);
}
</script>
