<template>
  <div :class="['flex-1', 'min-h-0', 'bg-likecoin-lightergrey']">
    <div :class="['-mb-[65px]']">
      <HeroBanner>
        <h1
          :class="[
            'text-3xl',
            'font-inter',
            'font-semibold',
            'text-likecoin-votecolor-yes',
          ]"
        >
          {{ $t("app.title") }}
        </h1>
      </HeroBanner>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

interface Data {}

export default Vue.extend({
  components: {},

  filters: {},
  data(): Data {
    return {};
  },
  computed: {},
  watch: {},
  methods: {},
});
</script>
