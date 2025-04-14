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
        :loading="loading"
        :columns="columns"
        :rows="tableData"
      >
        <template #loading>
          <div
            :class="[
              'absolute',
              'top-0',
              'left-0',
              'w-full',
              'h-full',
              'flex',
              'flex-row',
              'justify-center',
              'items-center',
            ]"
          >
            <LoadingIcon />
          </div>
        </template>
        <template #empty>
          <div
            :class="[
              'absolute',
              'top-0',
              'left-0',
              'w-full',
              'h-full',
              'flex',
              'flex-row',
              'justify-center',
              'items-center',
              'bg-white/70',
              'text-likecoin-darkgrey',
            ]"
          >
            {{ $t('section.asset-preview.no-data') }}
          </div>
        </template>
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
      </UTable>
    </UCard>
    <div :class="['mt-4', 'flex', 'flex-row', 'justify-end']">
      <AppButton
        v-if="allTableRows.length > 0"
        @click="handleConfirmMigrationClick"
      >
        {{ $t('section.asset-preview.confirm-migration') }}
      </AppButton>
    </div>
  </div>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue';
import type { TranslateResult } from 'vue-i18n';

import { LikeNFTAssetMigrationClass } from '~/apis/models/likenftAssetMigration';
import { LikeNFTAssetSnapshot } from '~/apis/models/likenftAssetSnapshot';
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
  status: LikeNFTAssetMigrationClass['status'] | '-';
};

interface Data {
  selectedStatus: 'all-item' | 'publishing' | 'books';
}

function makeTableDataRows(
  assetUrlBase: string,
  snapshot: LikeNFTAssetSnapshot
): TableData[] {
  const res: TableData[] = [];
  for (const c of snapshot.classes) {
    res.push({
      type: 'class',
      image: makeImageUrl(c.image),
      name: c.name,
      status: '-',
      txHash: '-',
      assetUrl: cosmosClassUrl(assetUrlBase, c.cosmos_class_id),
    });
  }
  for (const n of snapshot.nfts) {
    res.push({
      type: 'book',
      image: makeImageUrl(n.image),
      name: n.name,
      status: '-',
      txHash: '-',
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
    loading: {
      type: Object as PropType<boolean | undefined>,
      required: false,
      default: undefined,
    },
    snapshot: {
      type: Object as PropType<LikeNFTAssetSnapshot | null>,
      required: true,
    },
  },
  data(): Data {
    return {
      selectedStatus: 'all-item',
    };
  },
  computed: {
    typeTranslation(): { [key in TableData['type']]: TranslateResult } {
      return {
        book: this.$t('section.asset-preview.table.data.type.book'),
        class: this.$t('section.asset-preview.table.data.type.class'),
      };
    },
    itemsFilterOptions(): {}[] {
      return [
        {
          key: 'all-item',
          label: this.$t('section.asset-preview.table.filter.all-items', {
            count: this.allTableRows.length,
          }),
          value: 'all-item',
        },
        {
          key: 'publishing',
          label: this.$t('section.asset-preview.table.filter.publishing', {
            count: this.publishingTableRows.length,
          }),
          value: 'publishing',
        },
        {
          key: 'books',
          label: this.$t('section.asset-preview.table.filter.books', {
            count: this.bookTableRows.length,
          }),
          value: 'books',
        },
      ];
    },
    columns() {
      return [
        {
          key: 'image',
          label: this.$t('section.asset-preview.table.header.cover'),
          class: 'w-[7.994757536%] px-0 text-center',
          rowClass: 'w-[7.994757536%]',
        },
        {
          key: 'name',
          label: this.$t('section.asset-preview.table.header.title'),
          class: 'w-[35.7798165138%]',
          rowClass: 'w-[35.7798165138%] pl-0',
        },
        {
          key: 'txHash',
          label: this.$t('section.asset-preview.table.header.tx-hash'),
          class: 'w-[26.2123197903%]',
          rowClass: 'w-[26.2123197903%]',
        },
        {
          key: 'type',
          label: this.$t('section.asset-preview.table.header.type'),
          class: 'w-[17.3001310616%] text-center',
          rowClass: 'w-[17.3001310616%] text-center',
        },
        {
          key: 'status',
          label: this.$t('section.asset-preview.table.header.status'),
          class: 'w-[12.7129750983%]',
          rowClass: 'w-[12.7129750983%]',
        },
      ];
    },
    allTableRows(): TableData[] {
      if (this.snapshot != null) {
        return makeTableDataRows(
          this.$appConfig.likerlandUrlBase,
          this.snapshot
        );
      }
      return [];
    },
    publishingTableRows(): TableData[] {
      return this.allTableRows.filter((r) => r.type === 'class');
    },
    bookTableRows(): TableData[] {
      return this.allTableRows.filter((r) => r.type === 'book');
    },
    tableData(): TableData[] {
      switch (this.selectedStatus) {
        case 'all-item':
          return this.allTableRows;
        case 'publishing':
          return this.publishingTableRows;
        case 'books':
          return this.bookTableRows;
        default:
          return [];
      }
    },
  },
  methods: {
    handleConfirmMigrationClick() {
      this.$emit('confirmMigration');
    },
  },
});
</script>
