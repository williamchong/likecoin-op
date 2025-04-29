<!-- Ref: https://github.com/nuxt/ui/blob/dev/src/runtime/components/data/Table.vue -->

<template>
  <div :class="[_ui.wrapper]">
    <table :class="[_ui.base, _ui.divide]">
      <thead :class="_ui.thead">
        <tr :class="_ui.tr.base">
          <th
            v-for="(column, index) in columns"
            :key="index"
            scope="col"
            :class="[
              _ui.th.base,
              _ui.th.padding,
              _ui.th.color,
              _ui.th.font,
              _ui.th.size,
              column.key === 'select' && _ui.checkbox.padding,
              column.class,
            ]"
          >
            <span>{{ column[columnAttribute] }}</span>
          </th>
        </tr>
      </thead>
      <tbody :class="_ui.tbody">
        <template v-if="rows?.length">
          <tr
            v-for="(row, index) in rows"
            :key="index"
            :class="[_ui.tr.base, $attrs.onSelect && _ui.tr.active, row?.class]"
          >
            <td
              v-for="(column, subIndex) in columns"
              :key="subIndex"
              :class="[
                _ui.td.base,
                _ui.td.padding,
                _ui.td.color,
                _ui.td.font,
                _ui.td.size,
                column?.rowClass,
                row[column.key]?.class,
                column.key === 'select' && _ui.checkbox.padding,
              ]"
            >
              <slot
                :name="`${column.key}-data`"
                :column="column"
                :row="row"
                :index="index"
              >
                {{ getRowData(row, column.key) }}
              </slot>
            </td>
          </tr>
        </template>
        <tr v-else-if="loading && !!$slots['loading']">
          <td :colspan="columns?.length" :class="[]">
            <slot name="loading" />
          </td>
        </tr>
        <tr v-else>
          <td :colspan="columns?.length" :class="[]">
            <slot name="empty">
              {{ empty || $t('table.noData') }}
            </slot>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue';

export interface TableRow {
  [key: string]: any;
}

export interface TableColumn {
  key: string;
  sortable?: boolean;
  sort?: (a: any, b: any, direction: 'asc' | 'desc') => number;
  direction?: 'asc' | 'desc';
  class?: string;
  rowClass?: string;
  [key: string]: any;
}

export default Vue.extend({
  props: {
    rows: {
      type: Array as PropType<TableRow[]>,
      default: () => [],
    },
    columns: {
      type: Array as PropType<TableColumn[] | null>,
      default: null,
    },
    columnAttribute: {
      type: String,
      default: 'label',
    },
    loading: {
      type: Boolean as PropType<boolean | undefined>,
      required: false,
      default: undefined,
    },
    empty: {
      type: String as PropType<string | undefined>,
      required: false,
      default: undefined,
    },
    ui: {
      type: Object,
      default: () => ({}),
    },
  },
  computed: {
    _ui() {
      // https://github.com/nuxt/ui/blob/dev/src/runtime/ui.config/data/table.ts
      return {
        wrapper: 'relative overflow-x-auto',
        base: 'min-w-full table-fixed',
        divide: 'divide-y divide-gray-300 dark:divide-gray-700',
        thead: 'relative',
        tbody: 'divide-y divide-gray-200 dark:divide-gray-800',
        caption: 'sr-only',
        ...this.ui,
        tr: {
          base: '',
          selected: 'bg-gray-50 dark:bg-gray-800/50',
          expanded: 'bg-gray-50 dark:bg-gray-800/50',
          active: 'hover:bg-gray-50 dark:hover:bg-gray-800/50 cursor-pointer',
          ...this.ui.tr,
        },
        th: {
          base: 'text-left rtl:text-right',
          padding: 'px-4 py-3.5',
          color: 'text-gray-900 dark:text-white',
          font: 'font-semibold',
          size: 'text-sm',
          ...this.ui.th,
        },
        td: {
          base: 'whitespace-nowrap',
          padding: 'px-4 py-4',
          color: 'text-gray-500 dark:text-gray-400',
          font: '',
          size: 'text-sm',
          ...this.ui.td,
        },
      };
    },
  },

  methods: {
    getRowData(
      row: TableRow,
      columnKey: string,
      defaultValue?: string
    ): string | undefined {
      return row[columnKey] ?? defaultValue;
    },
  },
});
</script>
