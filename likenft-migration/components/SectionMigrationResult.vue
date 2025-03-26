<template>
  <div>
    <UCard
      :ui="{
        base: '',
        ring: '',
        divide: 'divide-y divide-gray-200 dark:divide-gray-700',
        header: { padding: 'px-4 py-5' },
        body: {
          padding: '',
          base: 'divide-y divide-gray-200 dark:divide-gray-700',
        },
      }"
    >
      <div
        :class="[
          'flex',
          'items-center',
          'justify-betwee',
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
        />
      </div>
      <UTable
        :class="['w-full', 'h-[313px]']"
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
        <template #image-data="{ row }">
          <div :class="['w-full', 'flex', 'flex-row', 'justify-center']">
            <img :src="row.image" :class="['w-6', 'h-6', 'object-cover']" />
          </div>
        </template>
        <template #name-data="{ row }">
          <a
            :class="['w-full', 'flex', 'flex-row', 'items-center']"
            :href="row.assetUrl"
            target="_blank"
            ><FontAwesomeIcon
              :class="['text-likecoin-votecolor-yes']"
              icon="link-simple"
            /><span
              :class="['flex-1', 'min-w-0', 'overflow-hidden', 'text-ellipsis']"
              >{{ row.name }}</span
            ></a
          >
        </template>
        <template #type-data="{ row }">
          {{
            // @ts-expect-error
            typeTranslation[row.type]
          }}
        </template>
        <template #status-data="{ row }">
          <span
            :class="[
              {
                ['text-[#C19869]']: row.status === 'init',
                ['text-[#4195D2]']: row.status === 'in_progress',
                ['text-[#8AB470]']: row.status === 'completed',
                ['text-[#C72F2F]']: row.status === 'failed',
              },
            ]"
          >
            {{
              // @ts-expect-error
              statusTranslation[row.status]
            }}
          </span>
        </template>
      </UTable>
    </UCard>
  </div>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue';
import { PropValidator } from 'vue/types/options';
import type { TranslateResult } from 'vue-i18n';

import {
  LikeNFTAssetMigration,
  LikeNFTAssetMigrationClass,
} from '~/apis/models/likenftAssetMigration';
import { makeImageUrl } from '~/utils/imageUrl';
import { cosmosClassUrl, cosmosNFTUrl } from '~/utils/nft';

import UCard from '../nuxtui/components/UCard.vue';
import USelectMenu from '../nuxtui/components/USelectMenu.vue';
import UTable from '../nuxtui/components/UTable.vue';

type TableData = {
  type: 'class' | 'book';
  image: string;
  name: string;
  assetUrl: string;
  txHash: string | '-';
  status: LikeNFTAssetMigrationClass['status'];
};

interface Data {
  selectedStatus: 'all-items' | LikeNFTAssetMigrationClass['status'];
}

function makeTableDataRows(
  assetUrlBase: string,
  migration: LikeNFTAssetMigration
): TableData[] {
  const res: TableData[] = [];
  for (const c of migration.classes) {
    res.push({
      type: 'class',
      image: makeImageUrl(c.image),
      name: c.name,
      status: c.status,
      txHash: c.evm_tx_hash || '-',
      assetUrl: cosmosClassUrl(assetUrlBase, c.cosmos_class_id),
    });
  }
  for (const n of migration.nfts) {
    res.push({
      type: 'book',
      image: makeImageUrl(n.image),
      name: n.name,
      status: n.status,
      txHash: n.evm_tx_hash || '-',
      assetUrl: cosmosNFTUrl(assetUrlBase, n.cosmos_class_id, n.cosmos_nft_id),
    });
  }
  return res;
}

export default Vue.extend({
  components: {
    UCard,
    USelectMenu,
    UTable,
  },
  props: {
    initialStatus: {
      type: String as PropType<Data['selectedStatus']>,
      default: 'all-items',
    } as PropValidator<Data['selectedStatus']>,
    migration: {
      type: Object as PropType<LikeNFTAssetMigration>,
      required: true,
    },
  },
  data(): Data {
    return {
      selectedStatus: this.initialStatus,
    };
  },
  computed: {
    typeTranslation(): { [key in TableData['type']]: TranslateResult } {
      return {
        book: this.$t('section.migration-result.table.data.type.book'),
        class: this.$t('section.migration-result.table.data.type.class'),
      };
    },
    statusTranslation(): { [key in TableData['status']]: TranslateResult } {
      return {
        init: this.$t('section.migration-result.table.data.status.init'),
        in_progress: this.$t(
          'section.migration-result.table.data.status.in-progress'
        ),
        completed: this.$t(
          'section.migration-result.table.data.status.completed'
        ),
        failed: this.$t('section.migration-result.table.data.status.failed'),
      };
    },
    itemsFilterOptions(): {}[] {
      return [
        {
          key: 'all-items',
          label: this.$t('section.migration-result.table.filter.all-items', {
            count: this.allTableRows.length,
          }),
          value: 'all-items',
        },
        {
          key: 'init',
          label: this.$t('section.migration-result.table.filter.init', {
            count: this.initTableRows.length,
          }),
          value: 'init',
        },
        {
          key: 'in_progress',
          label: this.$t('section.migration-result.table.filter.in-progress', {
            count: this.inProgressTableRows.length,
          }),
          value: 'in_progress',
        },
        {
          key: 'completed',
          label: this.$t('section.migration-result.table.filter.completed', {
            count: this.completedTableRows.length,
          }),
          value: 'completed',
        },
        {
          key: 'failed',
          label: this.$t('section.migration-result.table.filter.failed', {
            count: this.failedTableRows.length,
          }),
          value: 'failed',
        },
      ];
    },
    columns() {
      return [
        {
          key: 'image',
          label: this.$t('section.migration-result.table.header.cover'),
          class: 'w-[7.994757536%] px-0 text-center',
          rowClass: 'w-[7.994757536%]',
        },
        {
          key: 'name',
          label: this.$t('section.migration-result.table.header.title'),
          class: 'w-[35.7798165138%]',
          rowClass: 'w-[35.7798165138%] pl-0',
        },
        {
          key: 'txHash',
          label: this.$t('section.migration-result.table.header.tx-hash'),
          class: 'w-[26.2123197903%]',
          rowClass: 'w-[26.2123197903%]',
        },
        {
          key: 'type',
          label: this.$t('section.migration-result.table.header.type'),
          class: 'w-[17.3001310616%] text-center',
          rowClass: 'w-[17.3001310616%] text-center',
        },
        {
          key: 'status',
          label: this.$t('section.migration-result.table.header.status'),
          class: 'w-[12.7129750983%]',
          rowClass: 'w-[12.7129750983%]',
        },
      ];
    },
    allTableRows(): TableData[] {
      return makeTableDataRows(
        this.$appConfig.likerlandUrlBase,
        this.migration
      );
    },
    initTableRows(): TableData[] {
      return this.allTableRows.filter((r) => r.status === 'init');
    },
    inProgressTableRows(): TableData[] {
      return this.allTableRows.filter((r) => r.status === 'in_progress');
    },
    completedTableRows(): TableData[] {
      return this.allTableRows.filter((r) => r.status === 'completed');
    },
    failedTableRows(): TableData[] {
      return this.allTableRows.filter((r) => r.status === 'failed');
    },
    tableData(): TableData[] {
      switch (this.selectedStatus) {
        case 'all-items':
          return this.allTableRows;
        case 'init':
          return this.initTableRows;
        case 'in_progress':
          return this.inProgressTableRows;
        case 'completed':
          return this.completedTableRows;
        case 'failed':
          return this.failedTableRows;
        default:
          return [];
      }
    },
  },
});
</script>
