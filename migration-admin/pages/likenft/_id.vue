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
    <div :class="['-mb-[120px]']">
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
          {{ $t("app.likecoin.transaction_details") }}
        </h1>
      </HeroBanner>
    </div>
    <div
      :class="[
        'mx-16',
        'mb-16',
        'px-4',
        'flex',
        'flex-1',
        'flex-col',
        'gap-6',
        'z-10',
      ]"
    >
      <!-- Loading state -->
      <div
        v-if="loading"
        :class="[
          'flex-1',
          'text-center',
          'flex',
          'justify-center',
          'items-center',
          'flex-col',
        ]"
      >
        <LoadingIcon />
        <p>{{ $t("common.loading") }}</p>
      </div>

      <!-- Error state -->
      <div
        v-else-if="error"
        :class="[
          'flex-1',
          'text-center',
          'flex',
          'justify-center',
          'items-center',
          'flex-col',
        ]"
      >
        <p :class="['text-red-500']">{{ error }}</p>
      </div>

      <!-- Content -->
      <template v-else-if="migration">
        <!-- Transaction Header -->
        <UCard :class="['w-full']">
          <div :class="['flex', 'flex-col', 'gap-4']">
            <!-- Transaction Status -->
            <div :class="['flex', 'items-center', 'justify-between']">
              <div :class="['flex', 'items-center', 'gap-4']">
                <div :class="['flex', 'items-center', 'gap-2']">
                  <div
                    :class="[
                      'w-8',
                      'h-8',
                      'rounded-full',
                      'flex',
                      'items-center',
                      'justify-center',
                      {
                        'bg-red-100': migration.status === 'failed',
                        'bg-green-100': migration.status === 'completed',
                        'bg-yellow-100':
                          migration.status !== 'failed' &&
                          migration.status !== 'completed',
                      },
                    ]"
                  >
                    <span
                      :class="[
                        {
                          'text-red-500': migration.status === 'failed',
                          'text-green-500': migration.status === 'completed',
                          'text-yellow-500':
                            migration.status !== 'failed' &&
                            migration.status !== 'completed',
                        },
                      ]"
                    >
                      <template v-if="migration.status === 'failed'"
                        >✕</template
                      >
                      <template v-else-if="migration.status === 'completed'"
                        >✓</template
                      >
                      <template v-else>⟳</template>
                    </span>
                  </div>
                  <div>
                    <h2 :class="['text-xl', 'font-semibold']">
                      {{ $t("section.likenft-migration.title") }}
                    </h2>
                    <p :class="['text-sm', 'text-gray-500']">
                      {{ formatDate(migration.created_at) }}
                    </p>
                  </div>
                </div>
              </div>
              <div
                :class="[
                  'px-4',
                  'py-2',
                  'rounded-full',
                  'text-sm',
                  'font-medium',
                  {
                    'bg-red-100 text-red-800': migration.status === 'failed',
                    'bg-green-100 text-green-800':
                      migration.status === 'completed',
                    'bg-yellow-100 text-yellow-800':
                      migration.status === 'init',
                    'bg-blue-100 text-blue-800':
                      migration.status === 'in_progress',
                  },
                ]"
              >
                {{ $t(getStatusKey(migration.status)) }}
              </div>
            </div>
          </div>
        </UCard>

        <!-- Transaction Information -->
        <UCard :class="['w-full']">
          <h3 :class="['text-lg', 'font-semibold', 'mb-4']">
            {{ $t("migration.information") }}
          </h3>
          <div :class="['grid', 'grid-cols-1', 'md:grid-cols-2', 'gap-4']">
            <!-- Transaction ID -->
            <div :class="['flex', 'flex-col']">
              <span :class="['text-sm', 'text-gray-500']">{{
                $t("section.likenft-migration.table.header.id")
              }}</span>
              <div :class="['flex', 'items-center', 'gap-2']">
                <span :class="['font-medium']">{{ migration.id }}</span>
              </div>
            </div>

            <!-- User Cosmos Address -->
            <div :class="['flex', 'flex-col']">
              <span :class="['text-sm', 'text-gray-500']">{{
                $t("section.likenft-migration.table.header.asset-snapshot-id")
              }}</span>
              <div :class="['flex', 'items-center', 'gap-2']">
                <span :class="['font-medium', 'truncate']">{{
                  migration.likenft_asset_snapshot_id
                }}</span>
              </div>
            </div>

            <!-- User Cosmos Address -->
            <div :class="['flex', 'flex-col']">
              <span :class="['text-sm', 'text-gray-500']">{{
                $t("section.likenft-migration.table.header.cosmos-address")
              }}</span>
              <div :class="['flex', 'items-center', 'gap-2']">
                <span :class="['font-medium', 'truncate']">{{
                  migration.cosmos_address
                }}</span>
                <button
                  :class="[
                    'p-1.5',
                    'rounded-md',
                    'bg-gray-100',
                    'hover:bg-gray-200',
                    'text-gray-600',
                    'hover:text-gray-800',
                    'transition-colors',
                  ]"
                  @click="copyToClipboard(migration.cosmos_address)"
                >
                  <FontAwesomeIcon icon="copy" />
                </button>
              </div>
            </div>

            <!-- User ETH Address -->
            <div :class="['flex', 'flex-col']">
              <span :class="['text-sm', 'text-gray-500']">{{
                $t("section.likenft-migration.table.header.eth-address")
              }}</span>
              <div :class="['flex', 'items-center', 'gap-2']">
                <span :class="['font-medium', 'truncate']">{{
                  migration.eth_address
                }}</span>
                <button
                  :class="[
                    'p-1.5',
                    'rounded-md',
                    'bg-gray-100',
                    'hover:bg-gray-200',
                    'text-gray-600',
                    'hover:text-gray-800',
                    'transition-colors',
                  ]"
                  @click="copyToClipboard(migration.eth_address)"
                >
                  <FontAwesomeIcon icon="copy" />
                </button>
              </div>
            </div>

            <!-- Network -->
            <div :class="['flex', 'flex-col']">
              <span :class="['text-sm', 'text-gray-500']">{{
                $t("migration.network")
              }}</span>
              <span :class="['font-medium']">{{ chain }}</span>
            </div>
          </div>
        </UCard>

        <UCard>
          <!-- Action Buttons -->
          <div :class="['flex', 'flex-col', 'gap-4']">
            <div>
              <AppButton
                variant="warning"
                :loading="deleting"
                @click="promptRestartFailedMigrations"
              >
                {{ $t("migration.restart_failed") }}
              </AppButton>
            </div>
            <div>
              <AppButton
                variant="warning"
                :loading="deleting"
                @click="promptRestartAllMigrations"
              >
                {{ $t("migration.restart_all") }}
              </AppButton>
            </div>
            <div>
              <AppButton
                variant="warning"
                :loading="deleting"
                @click="promptRemoveMigration"
              >
                {{ $t("migration.delete_migration") }}
              </AppButton>
            </div>
          </div>
        </UCard>

        <!-- Error Details (if failed) -->
        <UCard
          v-if="migration.status === 'failed' && migration.failed_reason"
          :class="['w-full', 'bg-red-50']"
        >
          <h3 :class="['text-lg', 'font-semibold', 'mb-4', 'text-red-700']">
            {{ $t("migration.error_details") }}
          </h3>
          <span :class="['font-medium']">{{ migration.failed_reason }}</span>
        </UCard>

        <!-- Error Details (if failed) -->
        <UCard v-if="migration.classes.length > 0" :class="['w-full']">
          <h3 :class="['text-lg', 'font-semibold', 'mb-4']">
            {{ $t("section.likenft-migration.detail.nft-class.title") }}
          </h3>
          <!-- Class Cards -->
          <div class="space-y-4">
            <div
              v-for="nftClass in migration.classes"
              :key="nftClass.id"
              :class="[
                'border rounded-lg overflow-hidden',
                nftClass.status === 'failed'
                  ? 'border-red-200'
                  : 'border-gray-200',
              ]"
            >
              <div class="flex flex-col md:flex-row">
                <!-- Class Image -->
                <div
                  class="w-full md:w-1/4 h-48 md:h-auto bg-gray-100 flex items-center justify-center"
                >
                  <img
                    v-if="nftClass.image"
                    :src="nftClass.image"
                    alt="NFT Class"
                    class="object-cover h-full w-full"
                  />
                  <div v-else class="text-gray-400">
                    <i class="i-heroicons-photo text-4xl"></i>
                  </div>
                </div>

                <!-- Class Details -->
                <div class="p-4 flex-1">
                  <div class="flex justify-between items-start">
                    <div>
                      <h4 class="text-lg font-medium">
                        {{ nftClass.name }}
                      </h4>
                      <p class="text-sm text-gray-500 break-all mb-2">
                        {{
                          $t(
                            "section.likenft-migration.detail.nft-class.class-id",
                            {
                              class_id: nftClass.cosmos_class_id,
                            }
                          )
                        }}
                      </p>
                    </div>
                    <div class="flex items-center">
                      <span
                        :class="[
                          'px-3 py-1 text-xs rounded-full capitalize',
                          nftClass.status === 'completed'
                            ? 'bg-green-100 text-green-700'
                            : nftClass.status === 'failed'
                            ? 'bg-red-100 text-red-700'
                            : nftClass.status === 'in_progress'
                            ? 'bg-blue-100 text-blue-700'
                            : 'bg-yellow-100 text-yellow-700',
                        ]"
                      >
                        {{ $t(getStatusKey(nftClass.status)) }}
                      </span>
                    </div>
                  </div>

                  <!-- Error message if failed - Prominently displayed -->
                  <div
                    v-if="
                      nftClass.status === 'failed' && nftClass.failed_reason
                    "
                    class="mt-3 p-3 bg-red-50 rounded text-sm text-red-700 border border-red-100"
                  >
                    <p class="font-medium">
                      {{
                        $t(
                          "section.likenft-migration.detail.nft-class.failed-reason"
                        )
                      }}
                    </p>
                    <p>{{ nftClass.failed_reason }}</p>
                  </div>

                  <div class="grid grid-cols-1 md:grid-cols-2 gap-3 mt-4">
                    <div>
                      <p class="text-xs text-gray-500">
                        {{
                          $t(
                            "section.likenft-migration.detail.nft-class.enqueued-time"
                          )
                        }}
                      </p>
                      <p class="text-sm">
                        {{
                          nftClass.enqueue_time
                            ? formatDate(nftClass.enqueue_time)
                            : "N/A"
                        }}
                      </p>
                    </div>
                    <div>
                      <p class="text-xs text-gray-500">
                        {{
                          $t(
                            "section.likenft-migration.detail.nft-class.finished-time"
                          )
                        }}
                      </p>
                      <p class="text-sm">
                        {{
                          nftClass.finish_time
                            ? formatDate(nftClass.finish_time)
                            : "N/A"
                        }}
                      </p>
                    </div>
                    <div class="md:col-span-2">
                      <p class="text-xs text-gray-500">
                        {{
                          $t(
                            "section.likenft-migration.detail.nft-class.evn-tx-hash"
                          )
                        }}
                      </p>
                      <p class="text-sm break-all">
                        {{ nftClass.evm_tx_hash || "N/A" }}
                        <button
                          v-if="nftClass.evm_tx_hash != null"
                          :class="[
                            'p-1.5',
                            'rounded-md',
                            'bg-gray-100',
                            'hover:bg-gray-200',
                            'text-gray-600',
                            'hover:text-gray-800',
                            'transition-colors',
                          ]"
                          @click="copyToClipboard(nftClass.evm_tx_hash)"
                        >
                          <FontAwesomeIcon icon="copy" />
                        </button>
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </UCard>

        <UCard v-if="migration.nfts.length > 0" :class="['w-full']">
          <h3 :class="['text-lg', 'font-semibold', 'mb-4']">
            {{ $t("section.likenft-migration.detail.nft.title") }}
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            <div
              v-for="nft in migration.nfts"
              :key="nft.id"
              :class="[
                'rounded-lg overflow-hidden border',
                nft.status === 'failed' ? 'border-red-200' : 'border-gray-200',
              ]"
            >
              <!-- NFT Header with Image and Status -->
              <div class="h-48 bg-gray-100 relative">
                <img
                  v-if="nft.image"
                  :src="nft.image"
                  alt="NFT"
                  class="object-cover h-full w-full"
                />
                <div
                  v-else
                  class="h-full flex items-center justify-center text-gray-400"
                >
                  <i class="i-heroicons-photo text-4xl"></i>
                </div>
                <div
                  :class="[
                    'absolute top-2 right-2 px-2 py-1 text-xs rounded-full capitalize',
                    nft.status === 'completed'
                      ? 'bg-green-100 text-green-700'
                      : nft.status === 'failed'
                      ? 'bg-red-100 text-red-700'
                      : nft.status === 'in_progress'
                      ? 'bg-blue-100 text-blue-700'
                      : 'bg-yellow-100 text-yellow-700',
                  ]"
                >
                  {{ $t(getStatusKey(nft.status)) }}
                </div>
              </div>

              <!-- NFT Details -->
              <div class="p-3">
                <div class="flex justify-between items-center">
                  <h4 class="font-medium truncate">{{ nft.name }}</h4>
                </div>
                <p class="text-xs text-gray-500 truncate">
                  {{
                    $t(`section.likenft-migration.detail.nft.nft-id`, {
                      nft_id: nft.cosmos_nft_id,
                    })
                  }}
                </p>
                <p class="text-xs text-gray-500 truncate">
                  {{
                    $t(`section.likenft-migration.detail.nft.class-id`, {
                      class_id: nft.cosmos_class_id,
                    })
                  }}
                </p>
                <p class="text-xs text-gray-500 truncate">
                  {{
                    $t(`section.likenft-migration.detail.nft.evn-tx-hash`, {
                      hash: nft.evm_tx_hash ?? "N/A",
                    })
                  }}
                </p>

                <!-- Error message prominently displayed -->
                <div
                  v-if="nft.status === 'failed' && nft.failed_reason"
                  class="mt-2 p-2 bg-red-50 rounded text-xs text-red-700 border border-red-100"
                >
                  <p class="font-medium">
                    {{
                      $t(`section.likenft-migration.detail.nft.failed-reason`)
                    }}
                  </p>
                  <p>{{ nft.failed_reason }}</p>
                </div>

                <div
                  class="flex justify-between items-center mt-2 text-xs text-gray-500"
                >
                  <div>
                    <p class="text-xs text-gray-500">
                      {{
                        $t("section.likenft-migration.detail.nft.created-at")
                      }}
                    </p>
                    <p class="text-sm">
                      {{
                        nft.enqueue_time ? formatDate(nft.created_at) : "N/A"
                      }}
                    </p>
                  </div>
                  <div>
                    <p class="text-xs text-gray-500">
                      {{
                        $t("section.likenft-migration.detail.nft.enqueued-time")
                      }}
                    </p>
                    <p class="text-sm">
                      {{
                        nft.enqueue_time ? formatDate(nft.enqueue_time) : "N/A"
                      }}
                    </p>
                  </div>
                  <div>
                    <p class="text-xs text-gray-500">
                      {{
                        $t("section.likenft-migration.detail.nft.finished-time")
                      }}
                    </p>
                    <p class="text-sm">
                      {{
                        nft.finish_time ? formatDate(nft.finish_time) : "N/A"
                      }}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </UCard>

        <!-- Action Buttons -->
        <div :class="['flex', 'justify-end', 'gap-4', 'mt-4']">
          <AppButton
            variant="warning"
            :loading="deleting"
            @click="promptRemoveMigration"
          >
            {{ $t("migration.delete_migration") }}
          </AppButton>
        </div>
      </template>
    </div>
  </div>
</template>

<script lang="ts">
import { format } from "date-fns";
import Vue from "vue";

import { makeCreateMigrationAPI } from "~/apis/CreateMigration";
import { makeGetLikeNFTMigrationsAPI } from "~/apis/GetLikeNFTMigration";
import {
  LikeNFTMigrationDetail,
  LikeNFTMigrationStatus,
} from "~/apis/models/likenftMigration";
import { makeRemoveLikeNFTMigrationsAPI } from "~/apis/RemoveLikeNFTMigration";
import { makeRetryFailedAssetsMigrationAPI } from "~/apis/RetryFailedAssetsMigration";
import AppButton from "~/components/AppButton.vue";
import HeroBanner from "~/components/HeroBanner.vue";
import LoadingIcon from "~/components/LoadingIcon.vue";
import { LIKECOIN_CHAIN_NAME } from "~/constant";
import UCard from "~/nuxtui/components/UCard.vue";

interface Data {
  migration: LikeNFTMigrationDetail | null;
  loading: boolean;
  deleting: boolean;
  error: string | null;
}

export default Vue.extend({
  components: {
    HeroBanner,
    AppButton,
    UCard,
    LoadingIcon,
  },

  data(): Data {
    return {
      migration: null,
      loading: true,
      deleting: false,
      error: null,
    };
  },

  computed: {
    migrationId(): number {
      return parseInt(this.$route.params.id, 10);
    },
    chain(): string {
      return LIKECOIN_CHAIN_NAME(this.$appConfig.isTestnet);
    },
    failedMigration(): LikeNFTMigrationDetail | null {
      if (this.migration != null && this.migration.status === "failed") {
        return this.migration;
      }
      return null;
    },
  },

  mounted() {
    this.fetchMigration();
  },

  methods: {
    async fetchMigration() {
      this.loading = true;
      this.error = null;

      try {
        const resp = await makeGetLikeNFTMigrationsAPI(this.migrationId)(
          this.$apiClient
        )();
        this.migration = resp.migration;
      } catch (_) {
        this.error = this.$t("migration.error_fetching") as string;
      } finally {
        this.loading = false;
      }
    },

    formatDate(date: Date): string {
      return format(date, "yyyy-MM-dd HH:mm:ss");
    },

    getStatusKey(status: LikeNFTMigrationStatus): string {
      switch (status) {
        case LikeNFTMigrationStatus.Init:
          return "section.likenft-migration.table.data.status.init";
        case LikeNFTMigrationStatus.InProgress:
          return "section.likenft-migration.table.data.status.in_progress";
        case LikeNFTMigrationStatus.Completed:
          return "section.likenft-migration.table.data.status.completed";
        case LikeNFTMigrationStatus.Failed:
          return "section.likenft-migration.table.data.status.failed";
        default:
          throw new Error("Invalid status");
      }
    },

    async retryFailedAssetsMigration() {
      if (this.failedMigration == null) {
        return;
      }
      this.deleting = true;
      try {
        const resp = await makeRetryFailedAssetsMigrationAPI(
          this.failedMigration.cosmos_address
        )(this.$apiClient)({
          book_nft_collection: this.failedMigration.classes
            .filter((c) => c.status === "failed")
            .map((c) => c.cosmos_class_id),
          book_nft: this.failedMigration.nfts
            .filter((n) => n.status === "failed")
            .map((n) => ({
              class_id: n.cosmos_class_id,
              nft_id: n.cosmos_nft_id,
            })),
        });
        this.migration = resp.migration;
      } finally {
        this.deleting = false;
      }
    },

    async retryAllAssetsMigration() {
      if (this.failedMigration == null) {
        return;
      }
      this.deleting = true;
      try {
        await makeRemoveLikeNFTMigrationsAPI(this.migrationId)(
          this.$apiClient
        )();
        const resp = await makeCreateMigrationAPI(this.$apiClient)({
          asset_snapshot_id: this.failedMigration.likenft_asset_snapshot_id,
          cosmos_address: this.failedMigration.cosmos_address,
          eth_address: this.failedMigration.eth_address,
        });
        this.$router.replace(`/likenft/${resp.migration.id}`);
      } finally {
        this.deleting = false;
      }
    },

    async removeMigration() {
      this.deleting = true;
      try {
        await makeRemoveLikeNFTMigrationsAPI(this.migrationId)(
          this.$apiClient
        )();
        this.$router.push("/likenft");
      } finally {
        this.deleting = false;
      }
    },

    async promptRestartFailedMigrations() {
      if (
        confirm(
          this.$t("migration.confirm_restart_failed_migrations") as string
        )
      ) {
        await this.retryFailedAssetsMigration();
      }
    },

    async promptRestartAllMigrations() {
      if (
        confirm(this.$t("migration.confirm_restart_all_migrations") as string)
      ) {
        await this.retryAllAssetsMigration();
      }
    },

    async promptRemoveMigration() {
      if (confirm(this.$t("migration.confirm_delete_migration") as string)) {
        await this.removeMigration();
      }
    },

    async copyToClipboard(text: string) {
      await navigator.clipboard.writeText(text);
      alert(this.$t("common.copied_to_clipboard"));
    },
  },
});
</script>
