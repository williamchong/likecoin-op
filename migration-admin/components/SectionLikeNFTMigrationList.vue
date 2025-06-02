<template>
  <div class="relative flex flex-col flex-1">
    <UCard
      :ui="{
        base: '',
        ring: '',
        divide:
          'divide-y divide-gray-200 dark:divide-gray-700 flex flex-1 flex-col',
        header: { padding: 'px-4 py-5' },
        body: {
          padding: '',
          base: 'divide-y divide-gray-200 dark:divide-gray-700 flex flex-1 flex-col',
        },
      }"
    >
      <div
        :class="[
          'flex',
          'items-center',
          'justify-between',
          'gap-3',
          'px-4',
          'py-3',
        ]"
      >
        <!-- Search Bar -->
        <div :class="['flex-1', 'max-w-md']">
          <div :class="['relative']">
            <div :class="['flex', 'items-center']">
              <input
                ref="searchInput"
                v-model="searchKeyword"
                type="text"
                placeholder="Search..."
                :class="[
                  'w-full',
                  'py-2',
                  'px-3',
                  'pr-10',
                  'border',
                  'border-gray-300',
                  'rounded-md',
                  'text-sm',
                  'focus:outline-none',
                  'focus:ring-2',
                  'focus:ring-likecoin-votecolor-yes',
                  'focus:border-transparent',
                  loading ? 'opacity-75' : '',
                ]"
              />
              <div
                :class="[
                  'flex',
                  'items-center',
                  'justify-center',
                  'w-10',
                  'h-full',
                  '-ml-10',
                ]"
              >
                <FontAwesomeIcon icon="search" :class="['text-gray-400']" />
              </div>
            </div>
          </div>
        </div>

        <!-- Status Filter -->
        <div :class="['w-64']">
          <USelectMenu
            v-model="selectedStatus"
            :options="itemsFilterOptions"
            :ui="{
              base: [
                'relative block w-full disabled:cursor-not-allowed disabled:opacity-75 focus:outline-none border-0',
                'text-sm',
                'leading-[20px]',
              ],
            }"
            value-attribute="value"
            :disabled="loading"
          />
        </div>
      </div>
      <UTable
        v-if="tableData.length > 0 || loading"
        :class="['w-full', 'flex', 'flex-1', 'flex-col']"
        :ui="{
          base: ['table-fixed', 'w-full'].join(' '),
          divide: '',
          tr: {
            base: ['hover:bg-gray-50', 'cursor-pointer'].join(' '),
          },
          th: {
            base: [
              'relative',
              'text-left',
              'rtl:text-right',
              'sticky',
              'top-0',
              'bg-white',
              'after:absolute',
              'after:w-full',
              'after:h-px',
              'after:bg-gray-300',
              'after:left-0',
              'after:bottom-0',
            ].join(' '),
            padding: ['py-3.5', 'px-4'].join(' '),
          },
          td: {
            base: [
              'overflow-hidden',
              'whitespace-nowrap',
              'text-ellipsis',
            ].join(' '),
            padding: ['py-3.5', 'px-4'].join(' '),
          },
        }"
        :columns="columns"
        :rows="tableData"
        @row-select="handleRowClick"
      >
        <template #created_at-data="{ row }">
          {{ formatDate(row.created_at) }}
        </template>
        <template #user_cosmos_address-data="{ row }">
          <span :class="['text-xs', 'font-mono']">
            {{ truncateAddress(row.user_cosmos_address) }}
          </span>
        </template>
        <template #user_eth_address-data="{ row }">
          <span :class="['text-xs', 'font-mono']">
            {{ truncateAddress(row.user_eth_address) }}
          </span>
        </template>
        <template #cosmos_tx_hash-data="{ row }">
          <span v-if="row.cosmos_tx_hash" :class="['text-xs', 'font-mono']">
            {{ truncateAddress(row.cosmos_tx_hash) }}
          </span>
          <span v-else>-</span>
        </template>
        <template #evm_tx_hash-data="{ row }">
          <span v-if="row.evm_tx_hash" :class="['text-xs', 'font-mono']">
            {{ truncateAddress(row.evm_tx_hash) }}
          </span>
          <span v-else>-</span>
        </template>
        <template #status-data="{ row }">
          <span
            :class="[
              {
                ['text-[#C19869]']: row.status === migrationStatus.Init,
                ['text-[#4195D2]']: row.status === migrationStatus.InProgress,
                ['text-[#8AB470]']: row.status === migrationStatus.Completed,
                ['text-[#C72F2F]']: row.status === migrationStatus.Failed,
              },
            ]"
          >
            {{ statusTranslation[row.status] }}
          </span>
        </template>
      </UTable>

      <div
        v-else
        class="flex-1 flex items-center justify-center bg-white bg-opacity-70"
      >
        <div class="flex flex-col items-center text-gray-500">
          <FontAwesomeIcon icon="inbox" class="text-4xl mb-2" />
          <p class="text-sm">{{ $t("common.empty") }}</p>
        </div>
      </div>

      <!-- Pagination Controls -->
      <div
        :class="[
          'flex',
          'items-center',
          'justify-between',
          'px-4',
          'py-3',
          'border-t',
          'border-gray-200',
        ]"
      >
        <div :class="['flex', 'items-center', 'gap-2']">
          <span :class="['text-sm', 'text-gray-700']">Page {{ page }}</span>
        </div>
        <div :class="['flex', 'items-center', 'gap-2']">
          <button
            :class="[
              'p-2',
              'rounded-md',
              'border',
              'border-gray-300',
              'text-gray-700',
              'hover:bg-gray-50',
              'focus:outline-none',
              'focus:ring-2',
              'focus:ring-likecoin-votecolor-yes',
              'focus:border-transparent',
              'disabled:opacity-50',
              'disabled:cursor-not-allowed',
            ]"
            :disabled="page <= 1 || loading"
            @click="handlePreviousPage"
          >
            <FontAwesomeIcon icon="arrow-left" />
          </button>
          <button
            :class="[
              'p-2',
              'rounded-md',
              'border',
              'border-gray-300',
              'text-gray-700',
              'hover:bg-gray-50',
              'focus:outline-none',
              'focus:ring-2',
              'focus:ring-likecoin-votecolor-yes',
              'focus:border-transparent',
              'disabled:opacity-50',
              'disabled:cursor-not-allowed',
            ]"
            :disabled="!hasMoreItems || loading"
            @click="handleNextPage"
          >
            <FontAwesomeIcon icon="arrow-right" />
          </button>
        </div>
      </div>
    </UCard>

    <!-- Loading overlay -->
    <div
      v-if="loading"
      class="absolute inset-0 flex items-center justify-center bg-white bg-opacity-70 z-10"
    >
      <div class="flex flex-col items-center">
        <LoadingIcon />
        <span class="mt-2 text-likecoin-darkgrey">{{
          $t("common.loading")
        }}</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { format } from "date-fns";
import Vue, { PropType } from "vue";
import type { TranslateResult } from "vue-i18n";

import {
  LikeNFTMigration,
  LikeNFTMigrationStatus,
} from "~/apis/models/likenftMigration";
import FontAwesomeIcon from "~/components/FontAwesomeIcon.vue";
import LoadingIcon from "~/components/LoadingIcon.vue";

import UCard from "../nuxtui/components/UCard.vue";
import USelectMenu from "../nuxtui/components/USelectMenu.vue";
import UTable from "../nuxtui/components/UTable.vue";

interface Data {
  selectedStatus: "all-items" | LikeNFTMigrationStatus;
  migrationStatus: typeof LikeNFTMigrationStatus;
  searchKeyword: string;
  searchTimeout?: NodeJS.Timeout;
  wasSearchInputFocused: boolean;
}

export default Vue.extend({
  components: {
    UCard,
    USelectMenu,
    UTable,
    LoadingIcon,
    FontAwesomeIcon,
  },
  props: {
    loading: {
      type: Boolean,
      default: false,
    },
    migrations: {
      type: Array as PropType<LikeNFTMigration[]>,
      default: () => [],
    },
    page: {
      type: Number,
      default: 1,
    },
    limit: {
      type: Number,
      default: 20,
    },
  },
  data(): Data {
    return {
      selectedStatus: "all-items",
      migrationStatus: LikeNFTMigrationStatus,
      searchKeyword: "",
      searchTimeout: undefined,
      wasSearchInputFocused: false,
    };
  },
  computed: {
    statusTranslation(): { [key in LikeNFTMigrationStatus]: TranslateResult } {
      return {
        [LikeNFTMigrationStatus.Init]: this.$t(
          "section.likenft-migration.table.data.status.init"
        ),
        [LikeNFTMigrationStatus.InProgress]: this.$t(
          "section.likenft-migration.table.data.status.in_progress"
        ),
        [LikeNFTMigrationStatus.Completed]: this.$t(
          "section.likenft-migration.table.data.status.completed"
        ),
        [LikeNFTMigrationStatus.Failed]: this.$t(
          "section.likenft-migration.table.data.status.failed"
        ),
      };
    },
    itemsFilterOptions(): {}[] {
      return [
        {
          key: "all-items",
          label: this.$t("section.likenft-migration.table.filter.all-items"),
          value: "all-items",
        },
        {
          key: LikeNFTMigrationStatus.Init,
          label: this.$t("section.likenft-migration.table.filter.init"),
          value: LikeNFTMigrationStatus.Init,
        },
        {
          key: LikeNFTMigrationStatus.InProgress,
          label: this.$t("section.likenft-migration.table.filter.in_progress"),
          value: LikeNFTMigrationStatus.InProgress,
        },
        {
          key: LikeNFTMigrationStatus.Completed,
          label: this.$t("section.likenft-migration.table.filter.completed"),
          value: LikeNFTMigrationStatus.Completed,
        },
        {
          key: LikeNFTMigrationStatus.Failed,
          label: this.$t("section.likenft-migration.table.filter.failed"),
          value: LikeNFTMigrationStatus.Failed,
        },
      ];
    },
    columns() {
      return [
        {
          key: "likenft_asset_snapshot_id",
          label: this.$t(
            "section.likenft-migration.table.header.asset-snapshot-id"
          ),
          class: "w-[10%]",
          rowClass: "w-[10%]",
        },
        {
          key: "cosmos_address",
          label: this.$t(
            "section.likenft-migration.table.header.cosmos-address"
          ),
          class: "w-[20%]",
          rowClass: "w-[20%]",
        },
        {
          key: "eth_address",
          label: this.$t("section.likenft-migration.table.header.eth-address"),
          class: "w-[20%]",
          rowClass: "w-[20%]",
        },
        {
          key: "created_at",
          label: this.$t("section.likenft-migration.table.header.created-at"),
          class: "w-[15%]",
          rowClass: "w-[15%]",
        },
        {
          key: "status",
          label: this.$t("section.likenft-migration.table.header.status"),
          class: "w-[10%]",
          rowClass: "w-[10%]",
        },
        {
          key: "failed_reason",
          label: this.$t(
            "section.likenft-migration.table.header.failed_reason"
          ),
          class: "w-[25%]",
          rowClass: "w-[25%]",
        },
      ];
    },
    tableData(): LikeNFTMigration[] {
      return this.migrations;
    },
    hasMoreItems(): boolean {
      // If we have fewer items than the limit, there are no more pages
      return this.migrations.length >= this.limit;
    },
  },
  watch: {
    selectedStatus() {
      this.$emit(
        "status-change",
        this.$data.selectedStatus === "all-items"
          ? null
          : this.$data.selectedStatus
      );
    },
    searchKeyword() {
      if (this.searchTimeout != null) {
        clearTimeout(this.searchTimeout);
      }
      // Check if search input is focused before emitting search event
      this.wasSearchInputFocused =
        document.activeElement === this.$refs.searchInput;
      this.searchTimeout = setTimeout(() => {
        this.$emit("search", this.searchKeyword);
      }, 500);
    },
    loading(newVal, oldVal) {
      // If loading has finished and search input was focused before, refocus it
      if (oldVal === true && newVal === false && this.wasSearchInputFocused) {
        this.$nextTick(() => {
          const searchInput = this.$refs.searchInput as HTMLInputElement;
          if (searchInput) {
            searchInput.focus();
          }
        });
      }
    },
  },
  methods: {
    formatDate(date: Date): string {
      return format(date, "yyyy-MM-dd HH:mm:ss");
    },
    truncateAddress(address: string): string {
      if (!address) return "-";
      return `${address.substring(0, 6)}...${address.substring(
        address.length - 4
      )}`;
    },
    handleRowClick(row: LikeNFTMigration) {
      this.$emit("row-select", row);
    },
    handlePreviousPage() {
      if (this.page > 1) {
        this.$emit("page-change", this.page - 1);
      }
    },
    handleNextPage() {
      this.$emit("page-change", this.page + 1);
    },
  },
});
</script>
