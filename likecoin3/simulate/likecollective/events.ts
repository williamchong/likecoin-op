import { PublicClient } from "@nomicfoundation/hardhat-viem/src/types";
import { LikeCollective, LikeStakePosition } from "./setup";

export class EventRetriever {
  publicClient: PublicClient;
  likeCollective: LikeCollective;
  likeStakePosition: LikeStakePosition;

  constructor(
    publicClient: PublicClient,
    likeCollective: LikeCollective,
    likeStakePosition: LikeStakePosition,
  ) {
    this.publicClient = publicClient;
    this.likeCollective = likeCollective;
    this.likeStakePosition = likeStakePosition;
  }

  async getEvents(toBlock: bigint, fromBlock?: bigint) {
    const logs = await Promise.all([
      this.likeCollective.getEvents.Staked(undefined, { toBlock, fromBlock }),
      this.likeCollective.getEvents.Unstaked(undefined, { toBlock, fromBlock }),
      this.likeCollective.getEvents.RewardClaimed(undefined, {
        toBlock,
        fromBlock,
      }),
      this.likeCollective.getEvents.RewardDeposited(undefined, {
        toBlock,
        fromBlock,
      }),
      this.likeCollective.getEvents.AllRewardClaimed(undefined, {
        toBlock,
        fromBlock,
      }),
      this.likeStakePosition.getEvents.Transfer(undefined, {
        toBlock,
        fromBlock,
      }),
    ])
      .then((logsOfEvent) =>
        logsOfEvent.flatMap((logs) => logs.map((log) => log)),
      )
      .then((logs) =>
        Promise.all(
          logs.map(async (log) => {
            const block = await this.publicClient.getBlock({
              blockHash: log.blockHash,
            });
            return {
              log,
              block,
            };
          }),
        ),
      );
    return logs.sort((a, b) => {
      if (a.log.blockNumber !== b.log.blockNumber) {
        return a.log.blockNumber > b.log.blockNumber ? 1 : -1;
      }
      if (a.log.transactionIndex !== b.log.transactionIndex) {
        return a.log.transactionIndex > b.log.transactionIndex ? 1 : -1;
      }
      return a.log.logIndex > b.log.logIndex ? 1 : -1;
    });
  }
}
