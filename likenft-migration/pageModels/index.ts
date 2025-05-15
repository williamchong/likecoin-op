import {
  CompletedLikeNFTAssetMigration,
  FailedLikeNFTAssetMigration,
  LikeNFTAssetMigration,
} from '~/apis/models/likenftAssetMigration';
import {
  EmptyLikeNFTAssetSnapshot,
  LikeNFTAssetSnapshot,
  NonEmptyLikeNFTAssetSnapshot,
} from '~/apis/models/likenftAssetSnapshot';
import { LikerIDMigrationError } from '~/apis/models/likerIDMigration';

export interface StepStateStep1 {
  step: 1;
}

export interface StepStateStep2Init {
  step: 2;
  state: 'Init';
}

export interface StepStateStep2AuthcoreRedirected {
  step: 2;
  state: 'AuthcoreRedirected';
  method: string | (string | null)[];
  code: string | (string | null)[];
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

export interface StepStateStep3Signing {
  step: 3;
  state: 'Signing';
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
  signMessage: string;
}

export type StepStateStep3SigningFailedReason = {
  type: 'likerIDMigration';
  error: LikerIDMigrationError;
};

export interface StepStateStep3SigningFailed {
  step: 3;
  state: 'SigningFailed';
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
  signMessage: string;
  failedReason: StepStateStep3SigningFailedReason;
}

export interface StepStateStep4Init {
  step: 4;
  state: 'Init';
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
}

export interface StepStateStep4EmptyMigrationPreview {
  step: 4;
  state: 'EmptyMigrationPreview';
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
  migrationPreview: EmptyLikeNFTAssetSnapshot;
}

export interface StepStateStep4NonEmptyMigrationPreview {
  step: 4;
  state: 'NonEmptyMigrationPreview';
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
  migrationPreview: NonEmptyLikeNFTAssetSnapshot;
}

export interface StepStateStep4MigrationRetryPreview {
  step: 4;
  state: 'MigrationRetryPreview';
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
  migrationPreview: LikeNFTAssetSnapshot;
  failedMigration: FailedLikeNFTAssetMigration;
}

export interface StepStateStep5MigrationResult {
  step: 5;
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
  migrationPreview: LikeNFTAssetSnapshot;
  migration: LikeNFTAssetMigration;
}

export interface StepStateCompleted {
  step: 99999;
  state: 'Completed';
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
  migrationPreview: LikeNFTAssetSnapshot;
  migration: CompletedLikeNFTAssetMigration;
}

export interface StepStateFailed {
  step: 99999;
  state: 'Failed';
  cosmosAddress: string;
  ethAddress: string;
  avatar: string | null;
  likerId: string | null;
  migrationPreview: LikeNFTAssetSnapshot;
  migration: FailedLikeNFTAssetMigration;
}

export type StepState =
  | StepStateStep1
  | StepStateStep2Init
  | StepStateStep2AuthcoreRedirected
  | StepStateStep2CosmosConnected
  | StepStateStep2LikerIdResolved
  | StepStateStep2LikerIdEvmConnected
  | StepStateStep2EthConnected
  | StepStateStep3Signing
  | StepStateStep3SigningFailed
  | StepStateStep4Init
  | StepStateStep4EmptyMigrationPreview
  | StepStateStep4NonEmptyMigrationPreview
  | StepStateStep4MigrationRetryPreview
  | StepStateStep5MigrationResult
  | StepStateCompleted
  | StepStateFailed;

export function introductionConfirmed(_: StepStateStep1): StepStateStep2Init {
  return {
    step: 2,
    state: 'Init',
  };
}

export function restarted(
  _: Exclude<StepState, { step: 1 }>
): StepStateStep2Init {
  return {
    step: 2,
    state: 'Init',
  };
}

export function authcoreRedirected(
  _: StepState,
  method: string | (string | null)[],
  code: string | (string | null)[]
): StepStateStep2AuthcoreRedirected {
  return {
    step: 2,
    state: 'AuthcoreRedirected',
    method,
    code,
  };
}

export function authcoreRedirectionFailed(
  _: StepStateStep2AuthcoreRedirected
): StepStateStep2Init {
  return {
    step: 2,
    state: 'Init',
  };
}

export function initCosmosConnected(
  _: StepState,
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
  prev: StepStateStep2CosmosConnected | StepStateStep3Signing,
  likerId: string | null,
  avatar: string | null,
  ethAddress: string
): StepStateStep4Init {
  return {
    step: 4,
    state: 'Init',
    cosmosAddress: prev.cosmosAddress,
    ethAddress,
    avatar,
    likerId,
  };
}

export function emptySnapshotRetried(
  prev: StepStateStep4EmptyMigrationPreview
): StepStateStep4Init {
  return {
    step: 4,
    state: 'Init',
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
  };
}

export function signMessageRequested(
  prev: StepStateStep2LikerIdEvmConnected,
  signMessage: string
): StepStateStep3Signing {
  return {
    step: 3,
    state: 'Signing',
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signMessage,
  };
}

export function likerIdMigrationFailed(
  prev: StepStateStep3Signing,
  reason: StepStateStep3SigningFailedReason
): StepStateStep3SigningFailed {
  return {
    step: 3,
    state: 'SigningFailed',
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    signMessage: prev.signMessage,
    failedReason: reason,
  };
}

export function emptyMigrationPreviewFetched(
  prev: StepStateStep4Init | StepStateStep4EmptyMigrationPreview,
  snapshot: EmptyLikeNFTAssetSnapshot
): StepStateStep4EmptyMigrationPreview {
  return {
    step: 4,
    state: 'EmptyMigrationPreview',
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    migrationPreview: snapshot,
  };
}

export function nonEmptyMigrationPreviewFetched(
  prev: StepStateStep4Init | StepStateStep4EmptyMigrationPreview,
  snapshot: NonEmptyLikeNFTAssetSnapshot
): StepStateStep4NonEmptyMigrationPreview {
  return {
    step: 4,
    state: 'NonEmptyMigrationPreview',
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    migrationPreview: snapshot,
  };
}

export function migrationResultFetched(
  prev:
    | StepStateStep4Init
    | StepStateStep4NonEmptyMigrationPreview
    | StepStateStep4MigrationRetryPreview
    | StepStateStep5MigrationResult,
  snapshot: LikeNFTAssetSnapshot,
  migration: LikeNFTAssetMigration
): StepStateStep5MigrationResult {
  return {
    step: 5,
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    migrationPreview: snapshot,
    migration,
  };
}

export function migrationCompleted(
  prev: StepStateStep4Init | StepStateStep5MigrationResult,
  snapshot: LikeNFTAssetSnapshot,
  completedMigration: CompletedLikeNFTAssetMigration
): StepStateCompleted {
  return {
    step: 99999,
    state: 'Completed',
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    migrationPreview: snapshot,
    migration: completedMigration,
  };
}

export function migrationFailed(
  prev: StepStateStep4Init | StepStateStep5MigrationResult,
  snapshot: LikeNFTAssetSnapshot,
  failedMigration: FailedLikeNFTAssetMigration
): StepStateFailed {
  return {
    step: 99999,
    state: 'Failed',
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    migrationPreview: snapshot,
    migration: failedMigration,
  };
}

export function migrationRetried(
  prev: StepStateFailed,
  failedMigration: FailedLikeNFTAssetMigration
): StepStateStep4MigrationRetryPreview {
  return {
    step: 4,
    state: 'MigrationRetryPreview',
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
    migrationPreview: prev.migrationPreview,
    failedMigration,
  };
}
