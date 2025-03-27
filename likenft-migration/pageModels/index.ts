import {
  CompletedLikeNFTAssetMigration,
  LikeNFTAssetMigration,
} from '~/apis/models/likenftAssetMigration';
import { LikeNFTAssetSnapshot } from '~/apis/models/likenftAssetSnapshot';

export interface StepStateStep1 {
  step: 1;
}

export interface StepStateStep2Init {
  step: 2;
  state: 'Init';
}

export interface StepStateStep2CosmosConnected {
  step: 2;
  state: 'CosmosConnected';
  cosmosAddress: string;
}

export interface StepStateStep2LikerIdResolved {
  step: 2;
  state: 'LikerIdResolved';
  cosmosAddress: string;
  avatar: string | null;
  likerId: string | null;
}

export interface StepStateStep2LikerIdEvmConnected {
  step: 2;
  state: 'LikerIdEvmConnected';
  cosmosAddress: string;
  avatar: string | null;
  likerId: string | null;
  ethAddress: string;
}

export interface StepStateStep2EthConnected {
  step: 2;
  state: 'EthConnected';
  ethAddress: string;
}

export interface StepStateStep3Init {
  step: 3;
  state: 'Init';
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
}

export interface StepStateStep3MigrationPreview {
  step: 3;
  state: 'MigrationPreview';
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
  migrationPreview: LikeNFTAssetSnapshot;
}

export interface StepStateStep4MigrationResult {
  step: 4;
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
  migration: LikeNFTAssetMigration;
}

export interface StepStateEnd {
  step: 99999;
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
  migration: CompletedLikeNFTAssetMigration;
}

export type StepState =
  | StepStateStep1
  | StepStateStep2Init
  | StepStateStep2CosmosConnected
  | StepStateStep2LikerIdResolved
  | StepStateStep2LikerIdEvmConnected
  | StepStateStep2EthConnected
  | StepStateStep3Init
  | StepStateStep3MigrationPreview
  | StepStateStep4MigrationResult
  | StepStateEnd;

export function introductionConfirmed(_: StepStateStep1): StepStateStep2Init {
  return {
    step: 2,
    state: 'Init',
  };
}

export function initCosmosConnected(
  _: Exclude<StepState, StepStateStep1>,
  cosmosAddress: string
): StepStateStep2CosmosConnected {
  return {
    step: 2,
    state: 'CosmosConnected',
    cosmosAddress,
  };
}

export function initEvmConnected(
  _: StepStateStep2Init,
  ethAddress: string
): StepStateStep2EthConnected {
  return {
    step: 2,
    state: 'EthConnected',
    ethAddress,
  };
}

export function likerIdResolved(
  prev: StepStateStep2CosmosConnected,
  avatar: string | null,
  likerId: string | null
): StepStateStep2LikerIdResolved {
  return {
    step: 2,
    state: 'LikerIdResolved',
    cosmosAddress: prev.cosmosAddress,
    avatar,
    likerId,
  };
}

export function likerIdEvmConnected(
  prev: StepStateStep2LikerIdResolved,
  ethAddress: string
): StepStateStep2LikerIdEvmConnected {
  return {
    step: 2,
    state: 'LikerIdEvmConnected',
    cosmosAddress: prev.cosmosAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    ethAddress,
  };
}

export function likerIdMigrated(
  prev: StepStateStep2CosmosConnected | StepStateStep2LikerIdEvmConnected,
  likerId: string | null,
  avatar: string | null,
  ethAddress: string
): StepStateStep3Init {
  return {
    step: 3,
    state: 'Init',
    cosmosAddress: prev.cosmosAddress,
    ethAddress,
    avatar,
    likerId,
  };
}

export function migrationPreviewFetched(
  prev: StepStateStep3Init | StepStateStep3MigrationPreview,
  snapshot: LikeNFTAssetSnapshot
): StepStateStep3MigrationPreview {
  return {
    step: 3,
    state: 'MigrationPreview',
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    migrationPreview: snapshot,
  };
}

export function migrationResultFetched(
  prev:
    | StepStateStep3Init
    | StepStateStep3MigrationPreview
    | StepStateStep4MigrationResult,
  migration: LikeNFTAssetMigration
): StepStateStep4MigrationResult {
  return {
    step: 4,
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    migration,
  };
}

export function migrationCompleted(
  prev: StepStateStep3Init | StepStateStep4MigrationResult,
  completedMigration: CompletedLikeNFTAssetMigration
): StepStateEnd {
  return {
    step: 99999,
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    migration: completedMigration,
  };
}
