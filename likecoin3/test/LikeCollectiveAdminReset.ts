import { loadFixture } from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { deployCollective } from "./factory";

// Proves the admin reset/sweep functions that `hardhat resetLedger` builds
// `cast send` commands for behave exactly as the runbook assumes:
//   adminResetPool(bookNFT, tokenIds, expectedTotalStaked) -> clean slate
//   adminSweep(to, amount)                                 -> reward LIKE out
const LIKE = (n: bigint) => n * 10n ** 6n; // Likecoin: 6 decimals
const bookA = "0x1234567890123456789012345678901234567890";
const bookB = "0x000000000000000000000000000000000000bEEF";

// Two stakers + a deposited reward in bookA. Mirrors what a drifted pool looks
// like at reset time: positions with non-zero reward indices and a funded pool.
async function fundedPool() {
  const ctx = await loadFixture(deployCollective);
  const { likeCollective, likecoin, likeStakePosition, rick, bob, kin } = ctx;

  const rickStake = LIKE(1000n);
  const bobStake = LIKE(2000n);
  const reward = LIKE(3000n); // splits 1:2 -> rick 1000, bob 2000

  const rickToken = await likeStakePosition.read.getNextTokenId();
  await likecoin.write.approve([likeCollective.address, rickStake], {
    account: rick.account,
  });
  await likeCollective.write.newStakePosition([bookA, rickStake], {
    account: rick.account,
  });

  const bobToken = await likeStakePosition.read.getNextTokenId();
  await likecoin.write.approve([likeCollective.address, bobStake], {
    account: bob.account,
  });
  await likeCollective.write.newStakePosition([bookA, bobStake], {
    account: bob.account,
  });

  await likecoin.write.approve([likeCollective.address, reward], {
    account: kin.account,
  });
  await likeCollective.write.depositReward([bookA, reward], {
    account: kin.account,
  });

  return {
    ...ctx,
    rickToken,
    bobToken,
    rickStake,
    bobStake,
    reward,
    totalStaked: rickStake + bobStake,
  };
}

describe("LikeCollective admin reset (resetLedger runbook)", function () {
  it("adminResetPool zeroes every reward index, pending, and restores totalStaked", async function () {
    const { likeCollective, rickToken, bobToken, reward, totalStaked } =
      await fundedPool();

    // Sanity: rewards accrued before reset.
    expect(
      await likeCollective.read.getRewardsOfPosition([rickToken]),
    ).to.equal(LIKE(1000n));
    expect(await likeCollective.read.getPendingRewardsPool([bookA])).to.equal(
      reward,
    );

    // Operator protocol: pause LikeCollective first (LikeStakePosition stays
    // unpaused). adminResetPool must still work while paused.
    await likeCollective.write.pause();
    expect(await likeCollective.read.paused()).to.be.true;

    await likeCollective.write.adminResetPool([
      bookA,
      [rickToken, bobToken],
      totalStaked,
    ]);

    // PoolReset event with the recomputed total + position count.
    const events = await likeCollective.getEvents.PoolReset(undefined, {
      fromBlock: 0n,
    });
    expect(events.length).to.equal(1);
    expect(events[0].args.trueTotalStaked).to.equal(totalStaked);
    expect(events[0].args.positionsReset).to.equal(2n);

    // Clean slate: no claimable, no pending, totalStaked == Σ stakedAmount.
    expect(
      await likeCollective.read.getRewardsOfPosition([rickToken]),
    ).to.equal(0n);
    expect(await likeCollective.read.getRewardsOfPosition([bobToken])).to.equal(
      0n,
    );
    expect(await likeCollective.read.getPendingRewardsPool([bookA])).to.equal(
      0n,
    );
    expect(await likeCollective.read.getTotalStake([bookA])).to.equal(
      totalStaked,
    );
  });

  it("adminSweep sends the orphaned reward LIKE out of the contract", async function () {
    const {
      likeCollective,
      likecoin,
      kin,
      rickToken,
      bobToken,
      reward,
      totalStaked,
    } = await fundedPool();

    await likeCollective.write.adminResetPool([
      bookA,
      [rickToken, bobToken],
      totalStaked,
    ]);

    // After reset the pool's rewardPending is 0 but the reward LIKE still sits
    // in the contract — sweep it to the downstream book store (here: kin).
    const before = await likecoin.read.balanceOf([kin.account.address]);
    const cBefore = await likecoin.read.balanceOf([likeCollective.address]);

    await likeCollective.write.adminSweep([kin.account.address, reward]);

    expect(await likecoin.read.balanceOf([kin.account.address])).to.equal(
      before + reward,
    );
    // Only the reward left; staked principal (totalStaked) stays in the contract.
    expect(await likecoin.read.balanceOf([likeCollective.address])).to.equal(
      cBefore - reward,
    );
    expect(await likecoin.read.balanceOf([likeCollective.address])).to.equal(
      totalStaked,
    );
  });

  it("reverts when the tokenId list is incomplete (ErrIncompletePositionSet)", async function () {
    const { likeCollective, rickToken, totalStaked } = await fundedPool();

    // Pass only one of two positions but claim the full true total: the
    // on-chain Σ stakedAmount (1000) != expectedTotalStaked (3000) -> revert.
    await expect(
      likeCollective.write.adminResetPool([bookA, [rickToken], totalStaked]),
    ).to.be.rejectedWith("ErrIncompletePositionSet");
  });

  it("reverts when a tokenId belongs to another pool (ErrTokenNotInPool)", async function () {
    const {
      likeCollective,
      likecoin,
      likeStakePosition,
      rick,
      rickToken,
      bobToken,
      totalStaked,
    } = await fundedPool();

    // A position in a different book.
    const otherToken = await likeStakePosition.read.getNextTokenId();
    await likecoin.write.approve([likeCollective.address, LIKE(10n)], {
      account: rick.account,
    });
    await likeCollective.write.newStakePosition([bookB, LIKE(10n)], {
      account: rick.account,
    });

    await expect(
      likeCollective.write.adminResetPool([
        bookA,
        [rickToken, bobToken, otherToken],
        totalStaked + LIKE(10n),
      ]),
    ).to.be.rejectedWith("ErrTokenNotInPool");
  });

  it("only the owner may call adminResetPool / adminSweep", async function () {
    const { likeCollective, rick, rickToken, bobToken, totalStaked } =
      await fundedPool();

    await expect(
      likeCollective.write.adminResetPool(
        [bookA, [rickToken, bobToken], totalStaked],
        { account: rick.account },
      ),
    ).to.be.rejectedWith("OwnableUnauthorizedAccount");

    await expect(
      likeCollective.write.adminSweep([rick.account.address, 1n], {
        account: rick.account,
      }),
    ).to.be.rejectedWith("OwnableUnauthorizedAccount");
  });
});
