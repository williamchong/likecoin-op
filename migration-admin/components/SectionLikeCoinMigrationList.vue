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
      <UTable
        :class="['w-full', 'flex', 'flex-1', 'flex-col']"
        :ui="{
          base: ['table-fixed', 'w-full'].join(' '),
          divide: '',
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
                ['text-[#C19869]']:
                  row.status === migrationStatus.PendingCosmosTxHash,
                ['text-[#4195D2]']:
                  row.status === migrationStatus.VerifyingCosmosTx ||
                  row.status === migrationStatus.EvmMinting ||
                  row.status === migrationStatus.EvmVerifying,
                ['text-[#8AB470]']: row.status === migrationStatus.Completed,
                ['text-[#C72F2F]']: row.status === migrationStatus.Failed,
              },
            ]"
          >
            {{ statusTranslation[row.status] }}
          </span>
        </template>
      </UTable>
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
import Vue, { PropType } from "vue";
import type { TranslateResult } from "vue-i18n";
import { format } from "date-fns";

import {
  LikeCoinMigration,
  LikeCoinMigrationStatus,
} from "~/apis/models/likecoinMigration";

import LoadingIcon from "~/components/LoadingIcon.vue";
import UCard from "../nuxtui/components/UCard.vue";
import USelectMenu from "../nuxtui/components/USelectMenu.vue";
import UTable from "../nuxtui/components/UTable.vue";
import FontAwesomeIcon from "~/components/FontAwesomeIcon.vue";

interface Data {
  selectedStatus: "all-items" | LikeCoinMigrationStatus;
  migrationStatus: typeof LikeCoinMigrationStatus;
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
      type: Array as PropType<LikeCoinMigration[]>,
      default: () => [],
    },
  },
  data(): Data {
    return {
      selectedStatus: "all-items",
      migrationStatus: LikeCoinMigrationStatus,
    };
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
  },
  computed: {
    statusTranslation(): { [key in LikeCoinMigrationStatus]: TranslateResult } {
      return {
        [LikeCoinMigrationStatus.PendingCosmosTxHash]: this.$t(
          "section.likecoin-migration.table.data.status.pending_cosmos_tx_hash"
        ),
        [LikeCoinMigrationStatus.VerifyingCosmosTx]: this.$t(
          "section.likecoin-migration.table.data.status.verifying_cosmos_tx"
        ),
        [LikeCoinMigrationStatus.EvmMinting]: this.$t(
          "section.likecoin-migration.table.data.status.evm_minting"
        ),
        [LikeCoinMigrationStatus.EvmVerifying]: this.$t(
          "section.likecoin-migration.table.data.status.evm_verifying"
        ),
        [LikeCoinMigrationStatus.Completed]: this.$t(
          "section.likecoin-migration.table.data.status.completed"
        ),
        [LikeCoinMigrationStatus.Failed]: this.$t(
          "section.likecoin-migration.table.data.status.failed"
        ),
      };
    },
    itemsFilterOptions(): {}[] {
      return [
        {
          key: "all-items",
          label: this.$t("section.likecoin-migration.table.filter.all-items"),
          value: "all-items",
        },
        {
          key: LikeCoinMigrationStatus.PendingCosmosTxHash,
          label: this.$t(
            "section.likecoin-migration.table.filter.pending_cosmos_tx_hash"
          ),
          value: LikeCoinMigrationStatus.PendingCosmosTxHash,
        },
        {
          key: LikeCoinMigrationStatus.VerifyingCosmosTx,
          label: this.$t(
            "section.likecoin-migration.table.filter.verifying_cosmos_tx"
          ),
          value: LikeCoinMigrationStatus.VerifyingCosmosTx,
        },
        {
          key: LikeCoinMigrationStatus.EvmMinting,
          label: this.$t("section.likecoin-migration.table.filter.evm_minting"),
          value: LikeCoinMigrationStatus.EvmMinting,
        },
        {
          key: LikeCoinMigrationStatus.EvmVerifying,
          label: this.$t(
            "section.likecoin-migration.table.filter.evm_verifying"
          ),
          value: LikeCoinMigrationStatus.EvmVerifying,
        },
        {
          key: LikeCoinMigrationStatus.Completed,
          label: this.$t("section.likecoin-migration.table.filter.completed"),
          value: LikeCoinMigrationStatus.Completed,
        },
        {
          key: LikeCoinMigrationStatus.Failed,
          label: this.$t("section.likecoin-migration.table.filter.failed"),
          value: LikeCoinMigrationStatus.Failed,
        },
      ];
    },
    columns() {
      return [
        {
          key: "cosmos_tx_hash",
          label: this.$t(
            "section.likecoin-migration.table.header.cosmos-tx-hash"
          ),
          class: "w-[10%]",
          rowClass: "w-[10%]",
        },
        {
          key: "evm_tx_hash",
          label: this.$t("section.likecoin-migration.table.header.evm-tx-hash"),
          class: "w-[10%]",
          rowClass: "w-[10%]",
        },
        {
          key: "user_cosmos_address",
          label: this.$t(
            "section.likecoin-migration.table.header.cosmos-address"
          ),
          class: "w-[10%]",
          rowClass: "w-[10%]",
        },
        {
          key: "user_eth_address",
          label: this.$t("section.likecoin-migration.table.header.eth-address"),
          class: "w-[10%]",
          rowClass: "w-[10%]",
        },
        {
          key: "amount",
          label: this.$t("section.likecoin-migration.table.header.amount"),
          class: "w-[5%]",
          rowClass: "w-[5%]",
        },
        {
          key: "created_at",
          label: this.$t("section.likecoin-migration.table.header.created-at"),
          class: "w-[10%]",
          rowClass: "w-[10%]",
        },
        {
          key: "status",
          label: this.$t("section.likecoin-migration.table.header.status"),
          class: "w-[10%]",
          rowClass: "w-[10%]",
        },
        {
          key: "failed_reason",
          label: this.$t(
            "section.likecoin-migration.table.header.failed_reason"
          ),
          class: "w-[30%]",
          rowClass: "w-[30%]",
        },
      ];
    },
    tableData(): LikeCoinMigration[] {
      return this.migrations;
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
  },
});
</script>
