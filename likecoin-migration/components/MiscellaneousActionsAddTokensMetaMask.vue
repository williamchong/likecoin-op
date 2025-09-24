<template>
  <div
    v-if="addTokenRequestState.type === 'added'"
    :class="['flex', 'flex-row', 'gap-2']"
  >
    <div :class="['flex-1', 'py-[10px]']">
      <p :class="['text-sm', 'text-likecoin-grey']">
        {{
          $t('miscellaneous-actions.add-tokens.metamask.description', {
            tokenSymbol: addTokenRequestState.tokenInfoState.tokenSymbol,
          })
        }}
      </p>
    </div>
    <FontAwesomeIcon
      icon="circle-check"
      :class="[
        'py-[10px]',
        'text-sm',
        'font-black',
        'text-likecoin-votecolor-yes',
      ]"
    />
  </div>
  <div
    v-else-if="tokenInfoState.type === 'resolved'"
    :class="['flex', 'flex-row', 'gap-2']"
  >
    <div :class="['flex-1', 'flex', 'flex-col', 'gap-1', 'py-[10px]']">
      <p :class="['text-sm', 'text-likecoin-darkgrey']">
        {{
          $t('miscellaneous-actions.add-tokens.metamask.description', {
            tokenSymbol: tokenInfoState.tokenSymbol,
          })
        }}
      </p>
      <span
        v-if="chainAddedState.type !== 'chain-added'"
        :class="['text-xs', 'text-likecoin-votecolor-no']"
      >
        {{
          $t('miscellaneous-actions.add-tokens.metamask.chain-not-added', {
            chainName,
            chainId: chainId.toString(),
          })
        }}
      </span>
    </div>
    <AppButton
      variant="primary"
      :class="['self-start']"
      :href="tokenBlockExplorerUrl"
      target="_blank"
    >
      {{
        $t('miscellaneous-actions.add-tokens.metamask.view-on-block-explorer')
      }}
    </AppButton>
    <AppButton
      v-if="chainAddedState.type === 'chain-added'"
      variant="primary"
      :class="['self-start']"
      @click="handleAddTokenClick(chainAddedState, tokenInfoState)"
    >
      <LoadingIcon
        v-if="addTokenRequestState.type === 'loading'"
        :class="['w-5', 'h-5']"
      />
      <span v-else>
        {{ $t('miscellaneous-actions.add-tokens.metamask.add') }}
      </span>
    </AppButton>
    <AppButton
      v-else
      variant="primary"
      :class="['self-start']"
      :disabled="true"
    >
      {{ $t('miscellaneous-actions.add-tokens.metamask.add') }}
    </AppButton>
  </div>
</template>

<script lang="ts">
import { Decimal } from 'decimal.js';
import { ethers, Network } from 'ethers';
import Vue from 'vue';

const likeCoinTokenABI = [
  'function symbol() view returns (string)',
  'function decimals() view returns (uint8)',
] as const;

interface LikeCoinTokenInterface {
  symbol(): Promise<string>;
  decimals(): Promise<bigint>;
}

interface TokenInfoStateInit {
  type: 'init';
}
interface TokenInfoStateLoading {
  type: 'loading';
}
interface TokenInfoStateResolved {
  type: 'resolved';
  tokenAddress: string;
  tokenSymbol: string;
  tokenDecimals: number;
}
interface TokenInfoStateError {
  type: 'error';
  error: Error;
}

type TokenInfoState =
  | TokenInfoStateInit
  | TokenInfoStateLoading
  | TokenInfoStateResolved
  | TokenInfoStateError;

interface ChainAddedStateInit {
  type: 'init';
}

interface ChainAddedStateNoMetaMask {
  type: 'no-meta-mask';
}

interface ChainAddedStateLoading {
  type: 'loading';
}

interface ChainAddedStateChainAdded {
  type: 'chain-added';
  metaMaskProvider: ethers.BrowserProvider;
}

interface ChainNotAddedStateChainNotAdded {
  type: 'chain-not-added';
}

type ChainAddedState =
  | ChainAddedStateInit
  | ChainAddedStateNoMetaMask
  | ChainAddedStateLoading
  | ChainAddedStateChainAdded
  | ChainNotAddedStateChainNotAdded;

interface AddTokenRequestStateInit {
  type: 'init';
}

interface AddTokenRequestStateLoading {
  type: 'loading';
}

interface AddTokenRequestStateAdded {
  type: 'added';
  tokenInfoState: TokenInfoStateResolved;
}

interface AddTokenRequestStateError {
  type: 'error';
  error: Error;
}

type AddTokenRequestState =
  | AddTokenRequestStateInit
  | AddTokenRequestStateLoading
  | AddTokenRequestStateAdded
  | AddTokenRequestStateError;

interface Data {
  chainName: string;
  chainId: Decimal;
  tokenInfoState: TokenInfoState;
  chainAddedState: ChainAddedState;
  addTokenRequestState: AddTokenRequestState;
}

export default Vue.extend({
  name: 'MiscellaneousActionsAddTokensMetaMask',

  data(): Data {
    return {
      chainName: this.$evmChainConfig.chainName,
      chainId: new Decimal(this.$evmChainConfig.chainId),
      tokenInfoState: {
        type: 'init',
      },
      chainAddedState: {
        type: 'init',
      },
      addTokenRequestState: {
        type: 'init',
      },
    };
  },

  computed: {
    tokenBlockExplorerUrl() {
      const [blockExplorerUrl] = this.$evmChainConfig.blockExplorerUrls;
      return new URL(
        `/token/${this.$appConfig.evmTokenAddress}`,
        blockExplorerUrl
      ).toString();
    },
  },

  mounted() {
    if (window.ethereum == null) {
      return;
    }
    window.ethereum.on('chainChanged', () => {
      this.checkChainAdded();
    });
    this.checkChainAdded();
    this.loadToken();
  },

  methods: {
    async checkChainAdded() {
      if (window.ethereum == null) {
        this.chainAddedState = {
          type: 'no-meta-mask',
        };
        return;
      }
      this.chainAddedState = {
        type: 'loading',
      };
      const metaMaskProvider = new ethers.BrowserProvider(
        window.ethereum,
        new Network(
          this.$evmChainConfig.chainName,
          this.$evmChainConfig.chainId
        )
      );
      try {
        await metaMaskProvider.send('wallet_switchEthereumChain', [
          { chainId: this.$evmChainConfig.chainId },
        ]);
        this.chainAddedState = {
          type: 'chain-added',
          metaMaskProvider,
        };
      } catch (error) {
        this.chainAddedState = {
          type: 'chain-not-added',
        };
      }
    },

    async loadToken() {
      const tokenAddress = this.$appConfig.evmTokenAddress;
      const jsonRpcProvider = new ethers.JsonRpcProvider(
        this.$evmChainConfig.rpcUrls[0]
      );
      const tokenContract = ethers.BaseContract.from<LikeCoinTokenInterface>(
        tokenAddress,
        likeCoinTokenABI,
        jsonRpcProvider
      );
      this.tokenInfoState = {
        type: 'loading',
      };
      try {
        const symbol = await tokenContract.symbol();
        const decimals = await tokenContract.decimals();
        this.tokenInfoState = {
          type: 'resolved',
          tokenAddress,
          tokenSymbol: symbol,
          tokenDecimals: Number(decimals),
        };
      } catch (error) {
        this.tokenInfoState = {
          type: 'error',
          error: error instanceof Error ? error : new Error(String(error)),
        };
      }
    },

    async handleAddTokenClick(
      { metaMaskProvider }: ChainAddedStateChainAdded,
      tokenInfoState: TokenInfoStateResolved
    ) {
      const { tokenAddress, tokenDecimals, tokenSymbol } = tokenInfoState;
      this.addTokenRequestState = {
        type: 'loading',
      };
      try {
        const wasAdded = await metaMaskProvider.send('wallet_watchAsset', {
          type: 'ERC20',
          options: {
            address: tokenAddress,
            symbol: tokenSymbol,
            decimals: tokenDecimals,
          },
        });
        if (wasAdded) {
          this.addTokenRequestState = {
            type: 'added',
            tokenInfoState,
          };
        } else {
          this.addTokenRequestState = {
            type: 'init',
          };
        }
      } catch (error) {
        this.addTokenRequestState = {
          type: 'error',
          error: error instanceof Error ? error : new Error(String(error)),
        };
      }
    },
  },
});
</script>
