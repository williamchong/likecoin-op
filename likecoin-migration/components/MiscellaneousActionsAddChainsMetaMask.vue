<template>
  <div
    v-if="addChainState.type === 'chain-added'"
    :class="['flex', 'flex-row', 'gap-2']"
  >
    <p :class="['flex-1', 'py-[10px]', 'text-sm', 'text-likecoin-grey']">
      {{
        $t('miscellaneous-actions.add-chain.metamask.description', {
          chainName,
          chainId: chainId.toString(),
        })
      }}
    </p>
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
  <div v-else :class="['flex', 'flex-row', 'gap-2']">
    <p :class="['flex-1', 'py-[10px]', 'text-sm', 'text-likecoin-darkgrey']">
      {{
        $t('miscellaneous-actions.add-chain.metamask.description', {
          chainName,
          chainId: chainId.toString(),
        })
      }}
    </p>
    <AppButton
      v-if="addChainState.type === 'ready' || addChainState.type === 'error'"
      variant="primary"
      :class="['self-start']"
      @click="handleAddChainClick(addChainState)"
    >
      {{ $t('miscellaneous-actions.add-chain.metamask.add-chain') }}
    </AppButton>
    <AppButton v-else variant="primary" :class="['self-start']" disabled>
      <LoadingIcon :class="['w-5', 'h-5']" />
    </AppButton>
  </div>
</template>

<script lang="ts">
import { Decimal } from 'decimal.js';
import { ethers, Network } from 'ethers';
import Vue from 'vue';

interface AddChainStateInit {
  type: 'init';
}

interface AddChainStateNoMetaMask {
  type: 'no-meta-mask';
}

interface AddChainStateChecking {
  type: 'checking';
  metaMaskProvider: ethers.BrowserProvider;
}

interface AddChainStateReady {
  type: 'ready';
  metaMaskProvider: ethers.BrowserProvider;
}

interface AddChainStateLoading {
  type: 'loading';
  metaMaskProvider: ethers.BrowserProvider;
}

interface AddChainStateChainAdded {
  type: 'chain-added';
  metaMaskProvider: ethers.BrowserProvider;
}

interface AddChainStateError {
  type: 'error';
  error: Error;
  metaMaskProvider: ethers.BrowserProvider;
}

type AddChainState =
  | AddChainStateInit
  | AddChainStateNoMetaMask
  | AddChainStateReady
  | AddChainStateChecking
  | AddChainStateLoading
  | AddChainStateChainAdded
  | AddChainStateError;

interface Data {
  addChainState: AddChainState;
  chainName: string;
  chainId: Decimal;
}

export default Vue.extend({
  name: 'MiscellaneousActionsAddChainsMetaMask',

  data(): Data {
    return {
      addChainState: {
        type: 'init',
      },
      chainName: this.$evmChainConfig.chainName,
      chainId: new Decimal(this.$evmChainConfig.chainId),
    };
  },

  mounted() {
    if (window.ethereum == null) {
      return;
    }
    window.ethereum.on('chainChanged', () => {
      this.checkReady();
    });
    this.checkReady();
  },

  methods: {
    async checkReady() {
      if (window.ethereum == null) {
        this.addChainState = {
          type: 'no-meta-mask',
        };
        return;
      }
      const metaMaskProvider = new ethers.BrowserProvider(
        window.ethereum,
        new Network(
          this.$evmChainConfig.chainName,
          this.$evmChainConfig.chainId
        )
      );

      this.addChainState = {
        type: 'checking',
        metaMaskProvider,
      };
      try {
        await metaMaskProvider.send('wallet_switchEthereumChain', [
          { chainId: this.$evmChainConfig.chainId },
        ]);
        this.addChainState = {
          type: 'chain-added',
          metaMaskProvider,
        };
      } catch (error) {
        this.addChainState = {
          type: 'ready',
          metaMaskProvider,
        };
      }
    },

    async handleAddChainClick(state: AddChainStateReady | AddChainStateError) {
      this.addChainState = {
        type: 'loading',
        metaMaskProvider: state.metaMaskProvider,
      };
      try {
        await state.metaMaskProvider.send('wallet_addEthereumChain', [
          this.$evmChainConfig,
        ]);
        this.addChainState = {
          type: 'chain-added',
          metaMaskProvider: state.metaMaskProvider,
        };
      } catch (error) {
        this.addChainState = {
          type: 'error',
          error: error instanceof Error ? error : new Error(String(error)),
          metaMaskProvider: state.metaMaskProvider,
        };
      }
    },
  },
});
</script>
