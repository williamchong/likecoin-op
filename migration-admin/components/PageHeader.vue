<template>
  <div
    :class="[
      'max-w-[1440px]',
      'bg-lightcoin-lightergrey',
      'pt-4',
      'pb-3.5',
      'px-15',
      'px-[60px]',
      'mx-auto',
    ]"
  >
    <div :class="['flex', 'items-center', 'justify-between']">
      <img :class="['h-[50px]']" src="~/assets/images/logo.svg" />
      <div :class="['flex', 'gap-4']">
        <AppButton
          to="/likecoin"
          :variant="isLikeCoinRoute ? 'primary' : 'secondary'"
        >
          LikeCoin
        </AppButton>
        <AppButton
          to="/likenft"
          :variant="isLikeNFTRoute ? 'primary' : 'secondary'"
        >
          LikeNFT
        </AppButton>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import AppButton from "~/components/AppButton.vue";

export default Vue.extend({
  components: {
    AppButton,
  },
  computed: {
    isLikeCoinRoute(): boolean {
      return this.$route.path.includes("/likecoin");
    },
    isLikeNFTRoute(): boolean {
      return this.$route.path.includes("/likenft");
    },
  },
});
</script>
