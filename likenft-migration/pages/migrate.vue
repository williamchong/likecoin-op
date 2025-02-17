<template>
  <div
    class="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]"
  >
    <main class="flex flex-col gap-8 row-start-2 items-center sm:items-start">
      <primary-button @click="handleConnectCosmosWalletClick">
        {{ $t('migrate.connect-cosmos-wallet') }}
      </primary-button>
      <p>{{ $t('migrate.cosmos-wallet-address', { cosmosWalletAddress }) }}</p>
      <p>{{ $t('migrate.liker-id', { likerID }) }}</p>
      <primary-button
        v-if="!isEthAddressMigrated"
        @click="handleConnectEVMWalletClick"
      >
        {{ $t('migrate.connect-evm-wallet') }}
      </primary-button>
      <p>{{ $t('migrate.evm-wallet-address', { evmWalletAddress }) }}</p>
      <primary-button
        v-if="
          !isEthAddressMigrated &&
          cosmosWalletAddress != null &&
          evmWalletAddress != null
        "
        @click="handleMigrateLikerIDClick"
      >
        {{ $t('migrate.migrate-likerid') }}
      </primary-button>
      <div v-if="migration == null && migrationPreview != null" class="w-full">
        <h2 class="text-[32px] font-bold">{{ $t('migrate.preview') }}</h2>
        <div
          v-if="
            migrationPreview.status === 'init' ||
            migrationPreview.status === 'in_progress'
          "
        >
          Loading...
        </div>
        <div class="max-h-40 overflow-auto">
          <div v-if="migrationPreview.classes.length > 0">
            <h3 class="text-[20px]">{{ $t('migrate.classes') }}</h3>
            <ol class="list-decimal pl-10">
              <li
                v-for="c in migrationPreview.classes"
                :key="c.cosmos_class_id"
              >
                <a :href="getLikerlandUrlForClass(c)">{{ c.name }}</a>
              </li>
            </ol>
          </div>
          <div v-if="migrationPreview.nfts.length > 0">
            <h3 class="text-[20px]">{{ $t('migrate.nfts') }}</h3>
            <ol class="list-decimal pl-10">
              <li
                v-for="n in migrationPreview.nfts"
                :key="n.cosmos_class_id + '/' + n.cosmos_nft_id"
              >
                <a :href="getLikerlandUrlForNFT(n)"
                  >{{ n.name }}({{ n.cosmos_nft_id }})</a
                >
              </li>
            </ol>
          </div>
        </div>
        <primary-button
          v-if="
            migrationPreview.status === 'completed' ||
            migrationPreview.status === 'failed'
          "
          class="mt-8"
          @click="handleReloadMigrationPreview"
          >{{ $t('migrate.reload-preview') }}</primary-button
        >
        <primary-button
          v-if="migrationPreview.status === 'completed'"
          class="mt-8"
          @click="handleMigrateAssetsClick"
          >{{ $t('migrate.migrate-assets') }}</primary-button
        >
      </div>
      <div v-else-if="migration != null">
        <h2 class="text-[32px] font-bold">{{ $t('migrate.result') }}</h2>
        <div class="max-h-40 overflow-auto">
          <div v-if="migration.classes.length > 0">
            <h3 class="text-[20px]">{{ $t('migrate.classes') }}</h3>
            <ol class="list-decimal pl-10">
              <li v-for="c in migration.classes" :key="c.cosmos_class_id">
                <a :href="getLikerlandUrlForClass(c)">{{ c.name }}</a>
              </li>
            </ol>
          </div>
          <div v-if="migration.nfts.length > 0">
            <h3 class="text-[20px]">{{ $t('migrate.nfts') }}</h3>
            <ol class="list-decimal pl-10">
              <li
                v-for="n in migration.nfts"
                :key="n.cosmos_class_id + '/' + n.cosmos_nft_id"
              >
                <a :href="getLikerlandUrlForNFT(n)"
                  >{{ n.name }}({{ n.cosmos_nft_id }})</a
                >
              </li>
            </ol>
          </div>
        </div>
      </div>
    </main>
    <div
      v-if="isLoading"
      class="fixed top-0 left-0 w-full h-full bg-white/90 flex items-center justify-center"
    >
      Loading
    </div>
  </div>
</template>

<script lang="ts">
import {
  LikeCoinWalletConnector,
  LikeCoinWalletConnectorMethodType,
  LikeCoinWalletConnectorSession,
} from '@likecoin/wallet-connector';
import { isAxiosError } from 'axios';
import { Eip1193Provider } from 'ethers';
import Vue from 'vue';
import Web3 from 'web3';
import { z } from 'zod';

import { makeCreateMigrationAPI } from '~/apis/createMigration';
import { makeCreateMigrationPreviewAPI } from '~/apis/createMigrationPreview';
import { makeGetMigrationAPI } from '~/apis/getMigration';
import { makeGetMigrationPreviewAPI } from '~/apis/getMigrationPreview';
import { getSignMessage } from '~/apis/getSignMessage';
import { makeGetUserProfileAPI } from '~/apis/getUserProfile';
import { makeMigrateLikerIDAPI } from '~/apis/migrateLikerID';
import {
  LikeNFTAssetMigration,
  LikeNFTAssetMigrationClass,
  LikeNFTAssetMigrationNFT,
} from '~/apis/models/likenftAssetMigration';
import {
  LikeNFTAssetSnapshot,
  LikeNFTAssetSnapshotClass,
  LikeNFTAssetSnapshotNFT,
} from '~/apis/models/likenftAssetSnapshot';
import { LIKECOIN_WALLET_CONNECTOR_CONFIG } from '~/constant/network';

async function getEthereumAccount(
  ethereum: Eip1193Provider
): Promise<string | null> {
  const web3 = new Web3(ethereum);
  await ethereum.request({ method: 'eth_requestAccounts' });
  const accounts = await web3.eth.getAccounts();
  if (accounts.length > 0) {
    return accounts[0];
  }
  return null;
}

async function signEthereumMessage(
  message: string,
  ethereum: Eip1193Provider,
  ethereumAddress: string
) {
  const web3 = new Web3(ethereum);
  const sign = await web3.eth.personal.sign(
    message,
    ethereumAddress,
    'Password!'
  );

  return sign;
}

interface Data {
  cosmosWalletAddress: string | null;
  likerID: string | null;
  evmWalletAddress: string | null;
  isEthAddressMigrated: boolean;
  migrationPreview: LikeNFTAssetSnapshot | null;
  isLoading: boolean;
  migrationPreviewFetchTimeout: ReturnType<typeof setTimeout> | null;
  migration: LikeNFTAssetMigration | null;
}

export default Vue.extend({
  data(): Data {
    return {
      cosmosWalletAddress: null,
      likerID: null,
      evmWalletAddress: null,
      isEthAddressMigrated: false,
      migrationPreview: null,
      migrationPreviewFetchTimeout: null,
      isLoading: false,
      migration: null,
    };
  },

  computed: {
    connector() {
      return new LikeCoinWalletConnector(
        LIKECOIN_WALLET_CONNECTOR_CONFIG(this.$appConfig.isTestnet)
      );
    },
    getSignMessage() {
      return getSignMessage(this.$apiClient);
    },
    migrateLikerID() {
      return makeMigrateLikerIDAPI(this.$apiClient);
    },
  },

  watch: {
    migrationPreview(migrationPreview: LikeNFTAssetSnapshot | null) {
      if (this.migrationPreviewFetchTimeout != null) {
        clearTimeout(this.migrationPreviewFetchTimeout);
        this.migrationPreviewFetchTimeout = null;
      }
      if (migrationPreview == null) {
        return;
      }
      if (
        migrationPreview.status === 'init' ||
        migrationPreview.status === 'in_progress'
      ) {
        this.migrationPreviewFetchTimeout = setTimeout(async () => {
          if (this.cosmosWalletAddress == null) {
            return;
          }
          const migrationPreview = await this.fetchMigrationPreview(
            this.cosmosWalletAddress
          );
          this.migrationPreviewFetchTimeout = null;
          this.migrationPreview = migrationPreview;
        }, 1000);
      }
    },
  },

  methods: {
    async handleConnectCosmosWalletClick() {
      const connection =
        await this.connector?.openConnectionMethodSelectionDialog({});
      this.handleConnection(connection);
    },

    async handleConnection(
      connection: LikeCoinWalletConnectorSession | undefined
    ) {
      if (!connection) return;

      this.migrationPreview = null;
      this.migration = null;

      const {
        accounts: [account],
      } = connection;
      this.cosmosWalletAddress = account.address;

      if (this.cosmosWalletAddress != null) {
        this.isLoading = true;
        try {
          const userProfile = await makeGetUserProfileAPI(
            this.cosmosWalletAddress
          )(this.$apiClient)();
          this.likerID = userProfile.user_profile.liker_id;
          this.evmWalletAddress = userProfile.user_profile.eth_wallet_address;
          this.isEthAddressMigrated = this.evmWalletAddress != null;
          if (this.isEthAddressMigrated) {
            const migration = await this.fetchMigration(
              this.cosmosWalletAddress
            );
            if (migration != null) {
              this.migration = migration;
            } else {
              let migrationPreview = await this.fetchMigrationPreview(
                this.cosmosWalletAddress
              );
              if (migrationPreview == null) {
                migrationPreview = await this.createMigrationPreview(
                  this.cosmosWalletAddress
                );
              }
              this.migrationPreview = migrationPreview;
            }
          }
        } finally {
          this.isLoading = false;
        }
      }

      this.connector?.once('account_change', this.handleAccountChange);
    },

    async handleAccountChange(method: LikeCoinWalletConnectorMethodType) {
      const connection = await this.connector?.init(method);
      this.handleConnection(connection);
    },

    async handleConnectEVMWalletClick() {
      this.isLoading = true;
      if (window.ethereum == null) {
        alert('Please install metamask extension');
        return;
      }
      try {
        this.evmWalletAddress = await getEthereumAccount(window.ethereum);
        if (this.cosmosWalletAddress != null) {
          let migrationPreview = await this.fetchMigrationPreview(
            this.cosmosWalletAddress
          );
          if (migrationPreview == null) {
            migrationPreview = await this.createMigrationPreview(
              this.cosmosWalletAddress
            );
          }
          this.migrationPreview = migrationPreview;
        }
      } catch (e) {
        console.error(e);
      } finally {
        this.isLoading = false;
      }
    },

    async handleMigrateLikerIDClick() {
      const S = z.object({
        cosmosWalletAddress: z.string(),
        evmWalletAddress: z.string(),
      });
      const s = S.safeParse({
        cosmosWalletAddress: this.cosmosWalletAddress,
        evmWalletAddress: this.evmWalletAddress,
      });
      if (s.data == null) {
        return;
      }
      if (window.ethereum == null) {
        alert('Please install metamask extension');
        return;
      }
      if (!this.isEthAddressMigrated) {
        const signMessage = await this.getSignMessage({
          cosmos_address: s.data.cosmosWalletAddress,
          eth_address: s.data.evmWalletAddress,
          liker_id: this.likerID,
        });
        const connection = await this.connector.initIfNecessary();
        if (connection == null) {
          alert('cannot get wallet connector connection');
          return;
        }
        const {
          accounts: [account],
          offlineSigner,
        } = connection;

        if (!offlineSigner.signArbitrary) {
          alert('signArbitrary not supported');
          return;
        }
        const result = await offlineSigner.signArbitrary(
          this.connector.options.chainId,
          account.address,
          signMessage.message
        );
        const cosmosSignature = result.signature;
        const signedMessage = await signEthereumMessage(
          signMessage.message,
          window.ethereum,
          s.data.evmWalletAddress
        );

        await this.migrateLikerID({
          cosmos_pub_key: result.pub_key.value,
          cosmos_signature: cosmosSignature,
          eth_address: s.data.evmWalletAddress,
          eth_signature: signedMessage,
          like_id: this.likerID,
          signing_message: signMessage.message,
        });

        this.isEthAddressMigrated = true;
      }
    },

    async fetchMigrationPreview(cosmosWalletAddress: string) {
      try {
        const migrationResponse = await makeGetMigrationPreviewAPI(
          cosmosWalletAddress
        )(this.$apiClient)();
        return migrationResponse.preview;
      } catch (e) {
        if (isAxiosError(e)) {
          if (e.status === 404) {
            return null;
          }
        }
        throw e;
      }
    },

    async createMigrationPreview(cosmosWalletAddress: string) {
      const migrationResponse = await makeCreateMigrationPreviewAPI(
        this.$apiClient
      )({ cosmos_address: cosmosWalletAddress });
      return migrationResponse.preview;
    },

    async fetchMigration(cosmosWalletAddress: string) {
      try {
        const migration = await makeGetMigrationAPI(cosmosWalletAddress)(
          this.$apiClient
        )();
        return migration.migration;
      } catch (e) {
        if (isAxiosError(e)) {
          if (e.status === 404) {
            return null;
          }
        }
        throw e;
      }
    },

    async createMigration(
      snapshotId: number,
      cosmosWalletAddress: string,
      ethAddress: string
    ) {
      const migrationResponse = await makeCreateMigrationAPI(this.$apiClient)({
        asset_snapshot_id: snapshotId,
        cosmos_address: cosmosWalletAddress,
        eth_address: ethAddress,
      });
      return migrationResponse.migration;
    },

    async handleReloadMigrationPreview() {
      if (this.cosmosWalletAddress == null) {
        return;
      }
      const migrationPreview = await this.createMigrationPreview(
        this.cosmosWalletAddress
      );
      this.migrationPreview = migrationPreview;
    },

    async handleMigrateAssetsClick() {
      if (this.migrationPreview == null) {
        return;
      }
      if (this.cosmosWalletAddress == null || this.evmWalletAddress == null) {
        alert('Please connect cosmos wallet and evm wallet');
        return;
      }
      const migration = await this.createMigration(
        this.migrationPreview.id,
        this.cosmosWalletAddress,
        this.evmWalletAddress
      );
      this.migration = migration;
    },

    getLikerlandUrlForClass(
      c: LikeNFTAssetSnapshotClass | LikeNFTAssetMigrationClass
    ): string {
      return `${this.$appConfig.likerlandUrlBase}/nft/class/${c.cosmos_class_id}`;
    },

    getLikerlandUrlForNFT(
      n: LikeNFTAssetSnapshotNFT | LikeNFTAssetMigrationNFT
    ): string {
      return `${this.$appConfig.likerlandUrlBase}/nft/class/${n.cosmos_class_id}/${n.cosmos_nft_id}`;
    },
  },
});
</script>
