import { SigningStargateClient } from '@cosmjs/stargate';
import { LikeCoinWalletConnectorConnectionResult } from '@likecoin/wallet-connector';

import {
  Completed,
  Failed,
  LikeCoinMigration,
  Pending,
  Polling,
} from '~/apis/models/likeCoinMigration';
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

export interface StepStateStep3AwaitSignature {
  step: 3;
  state: 'AwaitSignature';
  cosmosAddress: string;
  connection: LikeCoinWalletConnectorConnectionResult;
  avatar: string | null;
  likerId: string | null;
  signingStargateClient: SigningStargateClient;
  ethAddress: string;
  ethSigningMessage: string;
  gasEstimation: number;
  currentBalance: ChainCoin;
  estimatedBalance: ChainCoin;
}

export interface StepStateStep4Pending {
  step: 4;
  state: 'Pending';
  cosmosAddress: string;
  connection: LikeCoinWalletConnectorConnectionResult;
  avatar: string | null;
  likerId: string | null;
  signingStargateClient: SigningStargateClient;
  ethAddress: string;
  evmSignature: string;
  gasEstimation: number;
  currentBalance: ChainCoin;
  estimatedBalance: ChainCoin;
  migration: Pending<LikeCoinMigration>;
}

export interface StepStateStep4Polling {
  step: 4;
  state: 'Polling';
  cosmosAddress: string;
  connection: LikeCoinWalletConnectorConnectionResult;
  avatar: string | null;
  likerId: string | null;
  signingStargateClient: SigningStargateClient;
  ethAddress: string;
  evmSignature: string;
  gasEstimation: number;
  currentBalance: ChainCoin;
  estimatedBalance: ChainCoin;
  migration: Polling<LikeCoinMigration>;
}

export interface StepStateStep4PendingCosmosSignCancelled {
  step: 4;
  state: 'PendingCosmosSignCancelled';
  cosmosAddress: string;
  connection: LikeCoinWalletConnectorConnectionResult;
  avatar: string | null;
  likerId: string | null;
  signingStargateClient: SigningStargateClient;
  ethAddress: string;
  evmSignature: string;
  gasEstimation: number;
  currentBalance: ChainCoin;
  estimatedBalance: ChainCoin;
  migration: Pending<LikeCoinMigration>;
  cancelReason: string;
}

export interface StepStateStep4Failed {
  step: 4;
  state: 'Failed';
  cosmosAddress: string;
  connection: LikeCoinWalletConnectorConnectionResult;
  avatar: string | null;
  likerId: string | null;
  signingStargateClient: SigningStargateClient;
  ethAddress: string;
  evmSignature: string;
  gasEstimation: number;
  currentBalance: ChainCoin;
  estimatedBalance: ChainCoin;
  migration: Failed<LikeCoinMigration>;
  failedReason: string;
}

export interface StepStateStepEnd {
  step: 99999;
  cosmosAddress: string;
  connection: LikeCoinWalletConnectorConnectionResult;
  avatar: string | null;
  likerId: string | null;
  signingStargateClient: SigningStargateClient;
  ethAddress: string;
  evmSignature: string;
  gasEstimation: number;
  currentBalance: ChainCoin;
  estimatedBalance: ChainCoin;
  migration: Completed<LikeCoinMigration>;
}

export type StepState =
  | StepStateStep1
  | EitherEthConnected<StepStateStep2Init>
  | EitherEthConnected<StepStateStep2CosmosConnected>
  | EitherEthConnected<StepStateStep2LikerIdResolved>
  | EitherEthConnected<StepStateStep2GasEstimated>
  | StepStateStep3AwaitSignature
  | StepStateStep4Pending
  | StepStateStep4Polling
  | StepStateStep4PendingCosmosSignCancelled
  | StepStateStep4Failed
  | StepStateStepEnd;

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

export function ethSignConfirming(
  prev: EthConnected<StepStateStep2GasEstimated>,
  ethSigningMessage: string
): StepStateStep3AwaitSignature {
  return {
    step: 3,
    state: 'AwaitSignature',
    cosmosAddress: prev.cosmosAddress,
    connection: prev.connection,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signingStargateClient: prev.signingStargateClient,
    ethAddress: prev.ethAddress,
    ethSigningMessage,
    currentBalance: prev.currentBalance,
    gasEstimation: prev.gasEstimation,
    estimatedBalance: prev.estimatedBalance,
  };
}

export function pendingMigrationResolved(
  prev:
    | EitherEthConnected<StepStateStep2GasEstimated>
    | StepStateStep3AwaitSignature
    | StepStateStep4Pending
    | StepStateStep4PendingCosmosSignCancelled
    | StepStateStep4Polling,
  migration: Pending<LikeCoinMigration>
): StepStateStep4Pending {
  return {
    step: 4,
    state: 'Pending',
    cosmosAddress: prev.cosmosAddress,
    connection: prev.connection,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signingStargateClient: prev.signingStargateClient,
    ethAddress: migration.user_eth_address,
    evmSignature: migration.evm_signature,
    gasEstimation: prev.gasEstimation,
    currentBalance: prev.currentBalance,
    estimatedBalance: prev.estimatedBalance,
    migration,
  };
}

export function pollingMigrationResolved(
  prev:
    | EitherEthConnected<StepStateStep2GasEstimated>
    | StepStateStep3AwaitSignature
    | StepStateStep4Pending
    | StepStateStep4Polling,
  migration: Polling<LikeCoinMigration>
): StepStateStep4Polling {
  return {
    step: 4,
    state: 'Polling',
    cosmosAddress: prev.cosmosAddress,
    connection: prev.connection,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signingStargateClient: prev.signingStargateClient,
    ethAddress: migration.user_eth_address,
    evmSignature: migration.evm_signature,
    gasEstimation: prev.gasEstimation,
    currentBalance: prev.currentBalance,
    estimatedBalance: prev.estimatedBalance,
    migration,
  };
}

export function completedMigrationResolved(
  prev:
    | EitherEthConnected<StepStateStep2GasEstimated>
    | StepStateStep3AwaitSignature
    | StepStateStep4Pending
    | StepStateStep4Polling,
  migration: Completed<LikeCoinMigration>
): StepStateStepEnd {
  return {
    step: 99999,
    cosmosAddress: prev.cosmosAddress,
    connection: prev.connection,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signingStargateClient: prev.signingStargateClient,
    ethAddress: migration.user_eth_address,
    evmSignature: migration.evm_signature,
    gasEstimation: prev.gasEstimation,
    currentBalance: prev.currentBalance,
    estimatedBalance: prev.estimatedBalance,
    migration,
  };
}

export function failedMigrationResolved(
  prev:
    | EitherEthConnected<StepStateStep2GasEstimated>
    | StepStateStep3AwaitSignature
    | StepStateStep4Pending
    | StepStateStep4Polling,
  migration: Failed<LikeCoinMigration>
): StepStateStep4Failed {
  return {
    step: 4,
    state: 'Failed',
    cosmosAddress: prev.cosmosAddress,
    connection: prev.connection,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signingStargateClient: prev.signingStargateClient,
    ethAddress: migration.user_eth_address,
    evmSignature: migration.evm_signature,
    gasEstimation: prev.gasEstimation,
    currentBalance: prev.currentBalance,
    estimatedBalance: prev.estimatedBalance,
    migration,
    failedReason: migration.failed_reason,
  };
}

export function migrationCreated(
  prev: StepStateStep3AwaitSignature,
  migration: Pending<LikeCoinMigration>
): StepStateStep4Pending {
  return {
    step: 4,
    state: 'Pending',
    cosmosAddress: prev.cosmosAddress,
    connection: prev.connection,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signingStargateClient: prev.signingStargateClient,
    ethAddress: prev.ethAddress,
    evmSignature: migration.evm_signature,
    gasEstimation: prev.gasEstimation,
    currentBalance: prev.currentBalance,
    estimatedBalance: prev.estimatedBalance,
    migration,
  };
}

export function migrationCancelledByCosmosNotSigned(
  prev: StepStateStep4Pending,
  cancelReason: string
): StepStateStep4PendingCosmosSignCancelled {
  return {
    step: 4,
    state: 'PendingCosmosSignCancelled',
    cosmosAddress: prev.cosmosAddress,
    connection: prev.connection,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signingStargateClient: prev.signingStargateClient,
    ethAddress: prev.ethAddress,
    evmSignature: prev.migration.evm_signature,
    gasEstimation: prev.gasEstimation,
    currentBalance: prev.currentBalance,
    estimatedBalance: prev.estimatedBalance,
    migration: prev.migration,
    cancelReason,
  };
}

export function migrationRetryFailed(
  prev: StepStateStep4Failed
): EthNotConnected<StepStateStep2GasEstimated> {
  return {
    step: 2,
    state: 'GasEstimated',
    avatar: prev.avatar,
    connection: prev.connection,
    cosmosAddress: prev.cosmosAddress,
    currentBalance: prev.currentBalance,
    estimatedBalance: prev.estimatedBalance,
    ethAddress: null,
    gasEstimation: prev.gasEstimation,
    likerId: prev.likerId,
    signingStargateClient: prev.signingStargateClient,
  };
}

export function migrationRetryCosmosSign(
  prev: StepStateStep4PendingCosmosSignCancelled
): StepStateStep4Pending {
  return {
    step: 4,
    state: 'Pending',
    cosmosAddress: prev.cosmosAddress,
    connection: prev.connection,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signingStargateClient: prev.signingStargateClient,
    ethAddress: prev.ethAddress,
    evmSignature: prev.migration.evm_signature,
    gasEstimation: prev.gasEstimation,
    currentBalance: prev.currentBalance,
    estimatedBalance: prev.estimatedBalance,
    migration: prev.migration,
  };
}

export function migrationRefreshed(
  prev: StepStateStep4Polling,
  migration: Polling<LikeCoinMigration>
): StepStateStep4Polling {
  return {
    step: 4,
    state: 'Polling',
    cosmosAddress: prev.cosmosAddress,
    connection: prev.connection,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signingStargateClient: prev.signingStargateClient,
    ethAddress: prev.ethAddress,
    evmSignature: migration.evm_signature,
    gasEstimation: prev.gasEstimation,
    currentBalance: prev.currentBalance,
    estimatedBalance: prev.estimatedBalance,
    migration,
  };
}

export function restart(
  prev: StepStateStepEnd
): EthNotConnected<StepStateStep2LikerIdResolved> {
  return {
    step: 2,
    state: 'LikerIdResolved',
    connection: prev.connection,
    cosmosAddress: prev.cosmosAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    ethAddress: null,
  };
}
