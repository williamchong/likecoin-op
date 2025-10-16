import { PublicClient, WriteContractReturnType } from "viem";
import {
  CallParams,
  ClaimAllRewardsCallParams,
  ClaimRewardsCallParams,
  DecreaseStakePositionCallParams,
  DepositRewardCallParams,
  IncreaseStakeToPositionCallParams,
  NewStakePositionCallParams,
  RemoveStakePositionCallParams,
  RestakeRewardPositionCallParams,
  TransferStakePositionCallParams,
} from "./models/schema";
import { LikeCoin, LikeCollective, LikeStakePosition } from "./setup";

class NewStakePositionCallExecutor {
  constructor(
    private readonly likeCollective: LikeCollective,
    private readonly likeCoin: LikeCoin,
  ) {}

  async execute({
    bookNFT,
    account,
    amount,
  }: NewStakePositionCallParams): Promise<WriteContractReturnType> {
    await this.likeCoin.write.approve([this.likeCollective.address, amount], {
      account: account,
    });
    return this.likeCollective.write.newStakePosition([bookNFT, amount], {
      account: account,
    });
  }
}

class IncreaseStakeToPositionCallExecutor {
  constructor(
    private readonly likeCollective: LikeCollective,
    private readonly likeCoin: LikeCoin,
  ) {}

  async execute({
    tokenID,
    account,
    amount,
  }: IncreaseStakeToPositionCallParams): Promise<WriteContractReturnType> {
    await this.likeCoin.write.approve([this.likeCollective.address, amount], {
      account: account,
    });
    return this.likeCollective.write.increaseStakeToPosition(
      [tokenID, amount],
      {
        account: account,
      },
    );
  }
}

class DecreaseStakePositionCallExecutor {
  constructor(private readonly likeCollective: LikeCollective) {}

  async execute({
    tokenID,
    account,
    amount,
  }: DecreaseStakePositionCallParams): Promise<WriteContractReturnType> {
    return this.likeCollective.write.decreaseStakePosition([tokenID, amount], {
      account: account,
    });
  }
}

class RemoveStakePositionCallExecutor {
  constructor(private readonly likeCollective: LikeCollective) {}

  async execute({
    tokenID,
    account,
  }: RemoveStakePositionCallParams): Promise<WriteContractReturnType> {
    return this.likeCollective.write.removeStakePosition([tokenID], {
      account: account,
    });
  }
}

class ClaimRewardsCallExecutor {
  constructor(private readonly likeCollective: LikeCollective) {}

  async execute({
    tokenID,
    account,
  }: ClaimRewardsCallParams): Promise<WriteContractReturnType> {
    return this.likeCollective.write.claimRewards([tokenID], {
      account: account,
    });
  }
}

class ClaimAllRewardsCallExecutor {
  constructor(private readonly likeCollective: LikeCollective) {}

  async execute({
    account,
  }: ClaimAllRewardsCallParams): Promise<WriteContractReturnType> {
    return this.likeCollective.write.claimAllRewards([account], {
      account: account,
    });
  }
}

class RestakeRewardPositionCallExecutor {
  constructor(
    private readonly publicClient: PublicClient,
    private readonly likeCollective: LikeCollective,
    private readonly likeCoin: LikeCoin,
  ) {}

  async execute({
    tokenID,
    account,
  }: RestakeRewardPositionCallParams): Promise<WriteContractReturnType> {
    const currentBlock = await this.publicClient.getBlockNumber();
    const reward = await this.likeCollective.read.getRewardsOfPosition(
      [tokenID],
      { blockNumber: currentBlock },
    );
    await this.likeCoin.write.approve([this.likeCollective.address, reward], {
      account: account,
    });
    return this.likeCollective.write.restakeRewardPosition([tokenID], {
      account: account,
    });
  }
}

class DepositRewardCallExecutor {
  constructor(
    private readonly likeCollective: LikeCollective,
    private readonly likeCoin: LikeCoin,
  ) {}

  async execute({
    bookNFT,
    amount,
    account,
  }: DepositRewardCallParams): Promise<WriteContractReturnType> {
    await this.likeCoin.write.approve([this.likeCollective.address, amount], {
      account: account,
    });
    return this.likeCollective.write.depositReward([bookNFT, amount], {
      account: account,
    });
  }
}

class TransferStakePositionCallExecutor {
  constructor(private readonly likeStakePosition: LikeStakePosition) {}

  async execute({
    from,
    to,
    tokenID,
  }: TransferStakePositionCallParams): Promise<WriteContractReturnType> {
    return this.likeStakePosition.write.transferFrom([from, to, tokenID], {
      account: from,
    });
  }
}

export class CallExecutor {
  newStakePositionCallExecutor: NewStakePositionCallExecutor;
  increaseStakeToPositionCallExecutor: IncreaseStakeToPositionCallExecutor;
  decreaseStakePositionCallExecutor: DecreaseStakePositionCallExecutor;
  removeStakePositionCallExecutor: RemoveStakePositionCallExecutor;
  claimRewardsCallExecutor: ClaimRewardsCallExecutor;
  claimAllRewardsCallExecutor: ClaimAllRewardsCallExecutor;
  restakeRewardPositionCallExecutor: RestakeRewardPositionCallExecutor;
  depositRewardCallExecutor: DepositRewardCallExecutor;
  transferStakePositionCallExecutor: TransferStakePositionCallExecutor;

  constructor(
    publicClient: PublicClient,
    likeCoin: LikeCoin,
    likeCollective: LikeCollective,
    likeStakePosition: LikeStakePosition,
  ) {
    this.newStakePositionCallExecutor = new NewStakePositionCallExecutor(
      likeCollective,
      likeCoin,
    );
    this.increaseStakeToPositionCallExecutor =
      new IncreaseStakeToPositionCallExecutor(likeCollective, likeCoin);
    this.decreaseStakePositionCallExecutor =
      new DecreaseStakePositionCallExecutor(likeCollective);
    this.removeStakePositionCallExecutor = new RemoveStakePositionCallExecutor(
      likeCollective,
    );
    this.claimRewardsCallExecutor = new ClaimRewardsCallExecutor(
      likeCollective,
    );
    this.claimAllRewardsCallExecutor = new ClaimAllRewardsCallExecutor(
      likeCollective,
    );
    this.restakeRewardPositionCallExecutor =
      new RestakeRewardPositionCallExecutor(
        publicClient,
        likeCollective,
        likeCoin,
      );
    this.depositRewardCallExecutor = new DepositRewardCallExecutor(
      likeCollective,
      likeCoin,
    );
    this.transferStakePositionCallExecutor =
      new TransferStakePositionCallExecutor(likeStakePosition);
  }

  async executeSingle(callParam: CallParams): Promise<WriteContractReturnType> {
    switch (callParam.type) {
      case "newStakePosition":
        return this.newStakePositionCallExecutor.execute(callParam);
      case "increaseStakeToPosition":
        return this.increaseStakeToPositionCallExecutor.execute(callParam);
      case "decreaseStakePosition":
        return this.decreaseStakePositionCallExecutor.execute(callParam);
      case "removeStakePosition":
        return this.removeStakePositionCallExecutor.execute(callParam);
      case "claimRewards":
        return this.claimRewardsCallExecutor.execute(callParam);
      case "claimAllRewards":
        return this.claimAllRewardsCallExecutor.execute(callParam);
      case "restakeRewardPosition":
        return this.restakeRewardPositionCallExecutor.execute(callParam);
      case "depositReward":
        return this.depositRewardCallExecutor.execute(callParam);
      case "transferStakePosition":
        return this.transferStakePositionCallExecutor.execute(callParam);
    }
    throw new Error(`Unknown call params: ${callParam}`);
  }

  async execute(
    callParam: CallParams,
    ...callParams: CallParams[]
  ): Promise<WriteContractReturnType> {
    let txHash = await this.executeSingle(callParam);
    for (const callParam of callParams) {
      txHash = await this.executeSingle(callParam);
    }
    return txHash;
  }
}
