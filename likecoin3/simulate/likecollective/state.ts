import { Address, State } from "./models/schema";
import { LikeCoin, LikeCollective, LikeStakePosition } from "./setup";

type BookNFT = Address;
type Account = Address;

export class StateRetriever {
  constructor(
    private readonly likeCollective: LikeCollective,
    private readonly likeStakePosition: LikeStakePosition,
    private readonly likeCoin: LikeCoin,
  ) {}

  async retrieve(): Promise<State> {
    const nextTokenId = await this.likeStakePosition.read.getNextTokenId();

    const owners = new Set<Account>();
    const bookNFTs = new Set<BookNFT>();
    const stakingKeys = new Map<BookNFT, Set<Account>>();

    for (let i = 0n; i < nextTokenId; i++) {
      const positionInfo = await this.likeStakePosition.read.positionInfo([i]);
      let owner: Account;
      try {
        owner = await this.likeStakePosition.read.ownerOf([i]);
      } catch (error) {
        console.error(error);
        continue;
      }
      owners.add(owner);
      bookNFTs.add(positionInfo.bookNFT);
      if (!stakingKeys.has(positionInfo.bookNFT)) {
        stakingKeys.set(positionInfo.bookNFT, new Set<`0x${string}`>());
      }
      stakingKeys.get(positionInfo.bookNFT)!.add(owner);
    }

    const state: State = {
      bookStakedAmounts: {},
      userStakedAmounts: {},
      bookPendingRewards: {},
      userPendingRewards: {},
      userBalance: {},
    };

    for (const owner of owners) {
      state.userBalance[owner] = await this.likeCoin.read.balanceOf([owner]);
      state.userStakedAmounts[owner] = {};
      for (const bookNFT of bookNFTs) {
        state.userStakedAmounts[owner][bookNFT] =
          await this.likeCollective.read.getStakeForUser([owner, bookNFT]);
      }
    }

    for (const bookNFT of bookNFTs) {
      state.bookStakedAmounts[bookNFT] =
        await this.likeCollective.read.getTotalStake([bookNFT]);
      state.bookPendingRewards[bookNFT] =
        await this.likeCollective.read.getPendingRewardsPool([bookNFT]);
    }

    for (const [bookNFT, stakingOwnerKeys] of stakingKeys.entries()) {
      for (const owner of stakingOwnerKeys) {
        if (!state.userStakedAmounts[owner]) {
          state.userStakedAmounts[owner] = {};
        }
        state.userStakedAmounts[owner][bookNFT] =
          await this.likeCollective.read.getStakeForUser([owner, bookNFT]);
        if (!state.userPendingRewards[owner]) {
          state.userPendingRewards[owner] = {};
        }
        state.userPendingRewards[owner][bookNFT] =
          await this.likeCollective.read.getPendingRewardsForUser([
            owner,
            bookNFT,
          ]);
      }
    }

    return state;
  }
}
