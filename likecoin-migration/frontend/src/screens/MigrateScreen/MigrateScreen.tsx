import { Coin, SigningStargateClient } from "@cosmjs/stargate";
import { Keplr } from "@keplr-wallet/types";
import { Eip1193Provider } from "ethers";
import { useCallback, useState } from "react";
import { useNavigate } from "react-router";
import Web3 from "web3";

import { useConfig } from "../../hooks/useConfig";

interface Memo {
  signature: string;
  ethAddress: string;
  amount: string;
  denom: string;
}

function Memo(params: Memo): Memo {
  return params;
}

async function getEthereumAccount(
  ethereum: Eip1193Provider,
): Promise<string | null> {
  const web3 = new Web3(ethereum);
  const accounts = await web3.eth.getAccounts();
  if (accounts.length > 0) {
    return accounts[0];
  }
  return null;
}

function getMessage(amount: string, denom: string) {
  return `You are going to deposit ${amount} ${denom} to migration program.

This sign make sure the address is correct.`;
}

async function signEthereumMessage(
  message: string,
  ethereum: Eip1193Provider,
  ethereumAddress: string,
) {
  const web3 = new Web3(ethereum);
  const sign = await web3.eth.personal.sign(
    message,
    ethereumAddress,
    "Password!",
  );

  return sign;
}

async function getKeplrAccount(
  keplr: Keplr,
  chainId: string,
): Promise<string | null> {
  const offlineSigner = keplr.getOfflineSigner(chainId);
  const [account] = await offlineSigner.getAccounts();
  return account?.address ?? null;
}

export default function MigrateScreen() {
  const config = useConfig();
  const navigate = useNavigate();
  const [ethereumAddress, setEthereumAddress] = useState<string | null>(null);
  const [cosmosAddress, setCosmosAddress] = useState<string | null>(null);

  const retrieveEthereumAccount = useCallback(() => {
    if (window.ethereum == null) {
      return;
    }
    getEthereumAccount(window.ethereum).then(setEthereumAddress);
  }, []);

  const retrieveKeplrAccount = useCallback(() => {
    if (window.keplr == null) {
      return;
    }
    getKeplrAccount(window.keplr, config.cosmosChainId).then(setCosmosAddress);
  }, [config.cosmosChainId]);

  const handleConnectMetamaskClick = useCallback(() => {
    if (window.ethereum == null) {
      alert("Please install metamask extension");
      return;
    }
    window.ethereum
      ?.request({ method: "eth_requestAccounts" })
      .then(retrieveEthereumAccount);
  }, [retrieveEthereumAccount]);

  const handleConnectKeplrClick = useCallback(() => {
    if (window.keplr == null) {
      alert("Please install keplr extension");
      return;
    }
    const { keplr } = window;
    keplr.enable(config.cosmosChainId).then(retrieveKeplrAccount);
  }, [config.cosmosChainId, retrieveKeplrAccount]);

  const handleMigrateClick = useCallback(() => {
    if (window.ethereum == null) {
      alert("Please install metamask extension");
      return;
    }
    const { ethereum } = window;
    if (ethereumAddress == null) {
      alert("Please connect metamask");
      return;
    }
    if (window.keplr == null) {
      alert("Please install keplr");
      return;
    }
    const { keplr } = window;
    if (cosmosAddress == null) {
      alert("Please connect keplr");
      return;
    }

    const offlineSigner = keplr.getOfflineSigner(config.cosmosChainId);
    SigningStargateClient.connectWithSigner(
      "https://node.testnet.like.co/rpc/",
      offlineSigner,
      {},
    )
      .then(async (stargate) => {
        const balance = await stargate.getBalance(
          cosmosAddress,
          config.cosmosDenom,
        );
        const deductedBalance: Coin = {
          amount: `${BigInt(balance.amount) - BigInt(config.cosmosFeeAmount) - BigInt(config.cosmosFeeGas)}`,
          denom: balance.denom,
        };
        const message = getMessage(
          deductedBalance.amount,
          deductedBalance.denom,
        );
        const signature = await signEthereumMessage(
          message,
          ethereum,
          ethereumAddress,
        );
        const cosmosTx = await stargate.sendTokens(
          cosmosAddress,
          config.cosmosDepositAddress,
          [deductedBalance],
          {
            amount: [
              {
                amount: `${config.cosmosFeeAmount}`,
                denom: config.cosmosDenom,
              },
            ],
            gas: `${config.cosmosFeeGas}`,
          },
          JSON.stringify(
            Memo({
              signature: signature,
              ethAddress: ethereumAddress,
              amount: deductedBalance.amount,
              denom: deductedBalance.denom,
            }),
          ),
        );

        navigate(`/migrate/${cosmosTx.transactionHash}`);
      })
      .catch(console.error);
  }, [
    config.cosmosChainId,
    config.cosmosDenom,
    config.cosmosDepositAddress,
    config.cosmosFeeAmount,
    config.cosmosFeeGas,
    cosmosAddress,
    ethereumAddress,
    navigate,
  ]);

  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
        <button
          type="button"
          onClick={handleConnectKeplrClick}
          className="inline-flex items-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
        >
          Connect keplr
        </button>
        <p>Cosmos address: {cosmosAddress}</p>
        <button
          type="button"
          onClick={handleConnectMetamaskClick}
          className="inline-flex items-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
        >
          Connect metamask
        </button>
        <p>Eth address: {ethereumAddress}</p>
        <button
          type="button"
          onClick={handleMigrateClick}
          className="inline-flex items-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
        >
          Migrate
        </button>
      </main>
    </div>
  );
}
