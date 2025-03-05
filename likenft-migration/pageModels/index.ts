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

export interface StepStateStep2LikerIdMigrated {
  step: 2;
  state: 'LikerIdMigrated';
  cosmosAddress: string;
  avatar: string | null;
  likerId: string | null;
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

export type StepState =
  | StepStateStep1
  | StepStateStep2Init
  | StepStateStep2CosmosConnected
  | StepStateStep2LikerIdResolved
  | StepStateStep2LikerIdEvmConnected
  | StepStateStep2EthConnected
  | StepStateStep2LikerIdMigrated
  | StepStateStep3Init;

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
): StepStateStep2LikerIdMigrated {
  return {
    step: 2,
    state: 'LikerIdMigrated',
    cosmosAddress: prev.cosmosAddress,
    ethAddress,
    avatar,
    likerId,
  };
}

export function initMigrationPreview(
  prev: StepStateStep2LikerIdMigrated
): StepStateStep3Init {
  return {
    step: 3,
    state: 'Init',
    cosmosAddress: prev.cosmosAddress,
    ethAddress: prev.ethAddress,
    avatar: prev.avatar,
    likerId: prev.likerId,
  };
}
