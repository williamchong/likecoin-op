<template>
  <div
    :class="[
      'flex',
      'flex-col',
      'flex-1',
      'min-h-0',
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
          {{ $t("app.likenft.title") }}
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
      <SectionLikeNFTMigrationList
        :loading="loading"
        :migrations="migrations"
        :page="page"
        :limit="limit"
        @status-change="handleStatusChange"
        @search="handleSearch"
        @page-change="handlePageChange"
        @row-select="handleRowClick"
      />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

import { makeListLikeNFTMigrationsAPI } from "~/apis/ListLikeNFTMigrations";
import {
  LikeNFTMigration,
  LikeNFTMigrationStatus,
} from "~/apis/models/likenftMigration";
import HeroBanner from "~/components/HeroBanner.vue";
import SectionLikeNFTMigrationList from "~/components/SectionLikeNFTMigrationList.vue";

interface Data {
  migrations: LikeNFTMigration[];
  page: number;
  limit: number;
  loading: boolean;
  status: LikeNFTMigrationStatus | null;
  keyword: string | null;
}

export default Vue.extend({
  components: {
    HeroBanner,
    SectionLikeNFTMigrationList,
  },

  data(): Data {
    return {
      migrations: [],
      page: 1,
      limit: 10,
      loading: false,
      status: null,
      keyword: null,
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
    handleRowClick(row: LikeNFTMigration) {
      this.$router.push(`/likenft/${row.id}`);
    },
    handleSearch(keyword: string) {
      this.keyword = keyword === "" ? null : keyword;
      this.page = 1; // Reset to first page on new search
      this.fetchMigrations();
    },
    handleStatusChange(status: LikeNFTMigrationStatus | null) {
      this.status = status;
      this.page = 1; // Reset to first page on status change
      this.fetchMigrations();
    },
    handlePageChange(newPage: number) {
      this.page = newPage;
      this.fetchMigrations();
    },

    async fetchMigrations() {
      this.loading = true;

      try {
        const resp = await makeListLikeNFTMigrationsAPI(this.$apiClient)({
          offset: this.offset,
          limit: this.limit,
          status: this.status,
          q: this.keyword,
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
