import { SigningStargateClient } from '@cosmjs/stargate';
import { LikeCoinWalletConnectorConnectionResult } from '@likecoin/wallet-connector';

import { ChainCoin } from '~/models/cosmosNetworkConfig';

export interface StepStateStep1 {
  step: 1;
}

export type EthConnected<S extends { step: 2 }> = S & {
  ethAddress: string;
};

export type EthNotConnected<S extends { step: 2 }> = S & {
  ethAddress: null;
};

export type EitherEthConnected<S extends { step: 2 }> =
  | EthConnected<S>
  | EthNotConnected<S>;

export function isEthConnected<S extends { step: 2 }>(
  s: EitherEthConnected<S>
): s is EthConnected<S> {
  return s.ethAddress != null;
}

function applyEitherEthConnected<
  S1 extends { step: 2 },
  S2 extends { step: 2 }
>(
  currentState: EitherEthConnected<S1>,
  fn: (s: S1) => S2
): EitherEthConnected<S2> {
  if (isEthConnected(currentState)) {
    return {
      ...fn(currentState),
      ethAddress: currentState.ethAddress,
    };
  }
  return {
    ...fn(currentState),
    ethAddress: null,
  };
}

export interface StepStateStep2Init {
  step: 2;
  state: 'Init';
}

export interface StepStateStep2CosmosConnected {
  step: 2;
  state: 'CosmosConnected';
  cosmosAddress: string;
  connection: LikeCoinWalletConnectorConnectionResult;
}

export interface StepStateStep2LikerIdResolved {
  step: 2;
  state: 'LikerIdResolved';
  cosmosAddress: string;
  connection: LikeCoinWalletConnectorConnectionResult;
  avatar: string | null;
  likerId: string | null;
}

export interface StepStateStep2GasEstimated {
  step: 2;
  state: 'GasEstimated';
  cosmosAddress: string;
  connection: LikeCoinWalletConnectorConnectionResult;
  avatar: string | null;
  likerId: string | null;
  signingStargateClient: SigningStargateClient;
  gasEstimation: number;
  currentBalance: ChainCoin;
  estimatedBalance: ChainCoin;
}

export type StepState =
  | StepStateStep1
  | EitherEthConnected<StepStateStep2Init>
  | EitherEthConnected<StepStateStep2CosmosConnected>
  | EitherEthConnected<StepStateStep2LikerIdResolved>
  | EitherEthConnected<StepStateStep2GasEstimated>;

export function introductionConfirmed(
  _: StepStateStep1
): EthNotConnected<StepStateStep2Init> {
  return {
    step: 2,
    state: 'Init',
    ethAddress: null,
  };
}

export function evmConnected<S extends { step: 2 }>(
  s: S,
  ethAddress: string
): EthConnected<S> {
  return {
    ...s,
    ethAddress,
  };
}

export function initCosmosConnected(
  prev: Exclude<StepState, StepStateStep1>,
  cosmosAddress: string,
  connection: LikeCoinWalletConnectorConnectionResult
): EitherEthConnected<StepStateStep2CosmosConnected> {
  if (prev.step === 2) {
    return applyEitherEthConnected<{ step: 2 }, StepStateStep2CosmosConnected>(
      prev,
      () => ({
        step: 2,
        state: 'CosmosConnected',
        cosmosAddress,
        connection,
      })
    );
  }
  return {
    step: 2,
    state: 'CosmosConnected',
    cosmosAddress,
    connection,
    ethAddress: null,
  };
}

export function likerIdResolved(
  prev: EitherEthConnected<StepStateStep2CosmosConnected>,
  avatar: string | null,
  likerId: string | null
): EitherEthConnected<StepStateStep2LikerIdResolved> {
  return applyEitherEthConnected(prev, (prev) => ({
    step: 2,
    state: 'LikerIdResolved',
    cosmosAddress: prev.cosmosAddress,
    connection: prev.connection,
    avatar,
    likerId,
  }));
}

export function gasEstimated(
  prev: EitherEthConnected<StepStateStep2LikerIdResolved>,
  signingStargateClient: SigningStargateClient,
  currentBalance: ChainCoin,
  gasEstimation: number,
  estimatedBalance: ChainCoin
): EitherEthConnected<StepStateStep2GasEstimated> {
  return applyEitherEthConnected(prev, (prev) => ({
    step: 2,
    state: 'GasEstimated',
    cosmosAddress: prev.cosmosAddress,
    connection: prev.connection,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signingStargateClient,
    currentBalance,
    gasEstimation,
    estimatedBalance,
  }));
}
