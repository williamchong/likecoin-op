<template>
  <div
    :class="[
      'flex-1',
      'min-h-0',
      'flex',
      'flex-col',
      'bg-likecoin-lightergrey',
    ]"
  >
    <div :class="['-mb-[150px]']">
      <HeroBanner>
        <h1
          :class="[
            '-mt-[75px]',
            'text-3xl',
            'font-inter',
            'font-semibold',
            'text-likecoin-votecolor-yes',
          ]"
        >
          {{ $t("app.likecoin.title") }}
        </h1>
      </HeroBanner>
    </div>
    <div
      :class="[
        'ml-16',
        'mr-16',
        'my-16',
        'relative',
        'px-4',
        'mx-auto',
        'flex',
        'flex-1',
      ]"
    >
      <SectionLikeCoinMigrationList
        :isLoading="loading"
        :migrations="migrations"
        @status-change="handleStatusChange"
      />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { makeListLikeCoinMigrationsAPI } from "~/apis/ListLikeCoinMigrations";
import {
  LikeCoinMigration,
  LikeCoinMigrationStatus,
} from "~/apis/models/likecoinMigration";
import HeroBanner from "~/components/HeroBanner.vue";
import SectionLikeCoinMigrationList from "~/components/SectionLikeCoinMigrationList.vue";

interface Data {
  migrations: LikeCoinMigration[];
  page: number;
  limit: number;
  loading: boolean;
  status: LikeCoinMigrationStatus | null;
}

export default Vue.extend({
  components: {
    HeroBanner,
    SectionLikeCoinMigrationList,
  },

  data(): Data {
    return {
      migrations: [],
      page: 1,
      limit: 10,
      loading: false,
      status: null,
    };
  },

  computed: {
    offset(): number {
      return (this.page - 1) * this.limit;
    },
  },

  mounted() {
    this.fetchMigrations();
  },

  methods: {
    handleStatusChange(status: LikeCoinMigrationStatus | null) {
      this.status = status;
      this.fetchMigrations();
    },

    async fetchMigrations() {
      this.loading = true;

      try {
        const resp = await makeListLikeCoinMigrationsAPI(this.$apiClient)({
          offset: this.offset,
          limit: this.limit,
          status: this.status,
          keyword: null,
        });

        this.migrations = resp.migrations;
      } catch (error) {
        // Handle error appropriately
      } finally {
        this.loading = false;
      }
    },
  },
});
</script>
