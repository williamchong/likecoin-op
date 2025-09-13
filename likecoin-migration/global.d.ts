import { Window as KeplrWindow } from '@keplr-wallet/types';
import { Eip1193Provider } from 'ethers';

declare global {
  interface Window extends KeplrWindow {
    ethereum?: Eip1193Provider & {
      on(event: 'chainChanged', listener: (chainId: string) => void): void;
    };
  }
}
