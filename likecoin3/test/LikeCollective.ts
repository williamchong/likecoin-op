import {
  time,
  loadFixture,
} from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem, ignition } from "hardhat";
import { parseEther } from "viem";
import { deployCollective } from "./factory";

describe("LikeCollective", async function () {
  describe("Upgradeable Storage", async function () {
    it("should have the correct STORAGE_SLOT", async function () {
      const { likeCollective, deployer } = await loadFixture(deployCollective);
      const likeCollectiveMock =
        await viem.deployContract("LikeCollectiveMock");
      likeCollective.write.upgradeTo(likeCollectiveMock.address, {
        account: deployer.account,
      });
      const newLikeCollective = await viem.getContractAt(
        "LikeCollectiveMock",
        likeCollectiveMock.address,
      );
      expect(await newLikeCollective.read.dataStorage()).to.equal(
        "0xe9c9d9e1df02920d747aa7516ca1d4362d70267096e6330bcfb24b265ac2ee00",
      );
    });
  });

  describe("Basic contract functionality", async function () {
    it("should be initialized with correct owner", async function () {
      const { likeCollective, deployer } = await loadFixture(deployCollective);
      const owner = await likeCollective.read.owner();
      expect(owner.toLowerCase()).to.equal(
        deployer.account.address.toLowerCase(),
      );
    });

    it("should be paused by default", async function () {
      const { likeCollective } = await loadFixture(deployCollective);
      const paused = await likeCollective.read.paused();
      expect(paused).to.be.false;
    });

    it("should allow owner to pause and unpause", async function () {
      const { likeCollective, deployer } = await loadFixture(deployCollective);

      // Pause the contract
      await likeCollective.write.pause();
      expect(await likeCollective.read.paused()).to.be.true;

      // Unpause the contract
      await likeCollective.write.unpause();
      expect(await likeCollective.read.paused()).to.be.false;
    });

    it("should not allow non-owner to pause", async function () {
      const { likeCollective, rick } = await loadFixture(deployCollective);

      await expect(
        likeCollective.write.pause({
          account: rick.account,
        }),
      ).to.be.rejectedWith("OwnableUnauthorizedAccount");
    });
  });

  describe("Paused state behavior", async function () {
    it("should prevent staking when paused", async function () {
      const { likeCollective, rick } = await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const amount = parseEther("100");

      // Pause the contract
      await likeCollective.write.pause();

      // This should revert due to whenNotPaused modifier
      await expect(
        likeCollective.write.newStakePosition([mockBookNFT, amount], {
          account: rick.account,
        }),
      ).to.be.rejectedWith("EnforcedPause");
    });

    it("should prevent unstaking when paused", async function () {
      const { likeCollective, rick } = await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const amount = parseEther("50");

      // Pause the contract
      await likeCollective.write.pause();

      // This should revert due to whenNotPaused modifier
      await expect(
        likeCollective.write.removeStakePosition([1n], {
          account: rick.account,
        }),
      ).to.be.rejectedWith("EnforcedPause");
    });
  });

  describe("Stake and unstake", async function () {
    it("should allow staking and unstaking", async function () {
      const { likeCollective, rick, likeStakePosition, likecoin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";

      const nextTokenId = await likeStakePosition.read.getNextTokenId();

      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        10000n * 10n ** 6n,
      );
      await likecoin.write.approve(
        [likeCollective.address, 6000n * 10n ** 6n],
        {
          account: rick.account,
        },
      );
      await likeCollective.write.newStakePosition(
        [mockBookNFT, 6000n * 10n ** 6n],
        {
          account: rick.account,
        },
      );
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        4000n * 10n ** 6n,
      );
      const tokenOwner = await likeStakePosition.read.ownerOf([nextTokenId]);
      expect(tokenOwner.toLowerCase()).to.equal(
        rick.account.address.toLowerCase(),
      );
      await likeCollective.write.removeStakePosition([nextTokenId], {
        account: rick.account,
      });
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        10000n * 10n ** 6n,
      );
    });

    it("should correctly allow claiming new rewards and unstaking for one user", async function () {
      const { likeCollective, rick, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const amount = 8000n * 10n ** 6n;
      const reward = 1000n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();
      await likecoin.write.approve([likeCollective.address, amount], {
        account: rick.account,
      });

      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: rick.account,
      });
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        2000n * 10n ** 6n,
      );
      const owner = await likeStakePosition.read.ownerOf([nextTokenId]);
      expect(owner.toLowerCase()).to.equal(rick.account.address.toLowerCase());
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(0n);
      expect(
        await likeCollective.read.getStakeForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(amount);

      await likeCollective.write.claimRewards([nextTokenId], {
        account: rick.account,
      });
      // No operation as no reard deposited
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        2000n * 10n ** 6n,
      );

      // Deposit reward
      await likecoin.write.approve([likeCollective.address, reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(reward);

      // Claim rewards
      await likeCollective.write.claimRewards([nextTokenId], {
        account: rick.account,
      });
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        3000n * 10n ** 6n,
      );

      await likeCollective.write.removeStakePosition([nextTokenId], {
        account: rick.account,
      });
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        11000n * 10n ** 6n,
      );
      await expect(
        likeStakePosition.read.ownerOf([nextTokenId]),
      ).to.be.rejectedWith("ERC721NonexistentToken");
    });

    it("should not allow random user to unstake position", async function () {
      const { likeCollective, rick, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const amount = 2000n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();
      await likecoin.write.approve([likeCollective.address, amount], {
        account: rick.account,
      });

      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: rick.account,
      });
      const owner = await likeStakePosition.read.ownerOf([nextTokenId]);
      expect(owner.toLowerCase()).to.equal(rick.account.address.toLowerCase());

      await expect(
        likeCollective.write.removeStakePosition([nextTokenId], {
          account: kin.account,
        }),
      ).to.be.rejectedWith("ErrInvalidOwner()");
    });

    it("should auto claim on unstake position", async function () {
      const { likeCollective, rick, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const amount = 2000n * 10n ** 6n;
      const reward = 1000n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();
      await likecoin.write.approve([likeCollective.address, 4n * amount], {
        account: rick.account,
      });

      await likeCollective.write.newStakePosition([mockBookNFT, 4n * amount], {
        account: rick.account,
      });

      // Reward
      await likecoin.write.approve([likeCollective.address, reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });

      // Claim rewards
      await likeCollective.write.removeStakePosition([nextTokenId], {
        account: rick.account,
      });
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        11000n * 10n ** 6n,
      );
    });

    it("should allow multiple users to stake and unstake", async function () {
      const { likeCollective, rick, bob, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const amount = 2000n * 10n ** 6n;
      const reward = 1000n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();
      await likecoin.write.approve([likeCollective.address, 4n * amount], {
        account: rick.account,
      });
      await likecoin.write.approve([likeCollective.address, amount], {
        account: bob.account,
      });

      await likeCollective.write.newStakePosition([mockBookNFT, 4n * amount], {
        account: rick.account,
      });
      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: bob.account,
      });

      // Reward
      await likecoin.write.approve([likeCollective.address, reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });

      // Claim rewards
      await likeCollective.write.removeStakePosition([nextTokenId], {
        account: rick.account,
      });
    });
  });

  describe("Increase and decrease stake", async function () {
    it("should increase stake without rewards and update balances/stake", async function () {
      const { likeCollective, rick, likeStakePosition, likecoin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const baseStakeAmount = 2000n * 10n ** 6n;
      const additionalStakeAmount = 500n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();

      // Initial balances
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        10000n * 10n ** 6n,
      );

      // Stake baseStakeAmount
      await likecoin.write.approve(
        [likeCollective.address, baseStakeAmount + additionalStakeAmount],
        {
          account: rick.account,
        },
      );
      await likeCollective.write.newStakePosition(
        [mockBookNFT, baseStakeAmount],
        {
          account: rick.account,
        },
      );
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        10000n * 10n ** 6n - baseStakeAmount,
      );

      // Increase by additionalStakeAmount
      await likeCollective.write.increaseStakeToPosition(
        [nextTokenId, additionalStakeAmount],
        {
          account: rick.account,
        },
      );

      // Assertions
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        10000n * 10n ** 6n - baseStakeAmount - additionalStakeAmount,
      );
      expect(
        await likeCollective.read.getStakeForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(baseStakeAmount + additionalStakeAmount);
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(0n);
      expect(await likeCollective.read.getTotalStake([mockBookNFT])).to.equal(
        baseStakeAmount + additionalStakeAmount,
      );
    });

    it("should increase stake and compound existing rewards into position", async function () {
      const { likeCollective, rick, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const baseStakeAmount = 3000n * 10n ** 6n;
      const additionalStakeAmount = 700n * 10n ** 6n;
      const rewardAmount = 900n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();

      await likecoin.write.approve(
        [likeCollective.address, baseStakeAmount + additionalStakeAmount],
        {
          account: rick.account,
        },
      );
      await likeCollective.write.newStakePosition(
        [mockBookNFT, baseStakeAmount],
        {
          account: rick.account,
        },
      );

      // Deposit reward rewardAmount
      await likecoin.write.approve([likeCollective.address, rewardAmount], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, rewardAmount], {
        account: kin.account,
      });

      // Precondition: pending rewards == rewardAmount
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(rewardAmount);

      // Increase by additionalStakeAmount — should roll rewardAmount into principal, spend only additionalStakeAmount
      await likeCollective.write.increaseStakeToPosition(
        [nextTokenId, additionalStakeAmount],
        {
          account: rick.account,
        },
      );

      // User balance decreases only by additionalStakeAmount
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        10000n * 10n ** 6n - baseStakeAmount - additionalStakeAmount,
      );

      // Position stake includes baseStakeAmount + additionalStakeAmount + rewardAmount;
      // pool total stake must track the same total, since rewards are compounded into principal
      expect(
        await likeCollective.read.getStakeForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(baseStakeAmount + additionalStakeAmount + rewardAmount);
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(0n);
      expect(await likeCollective.read.getTotalStake([mockBookNFT])).to.equal(
        baseStakeAmount + additionalStakeAmount + rewardAmount,
      );
      // Pool pending rewards should drained
      expect(
        await likeCollective.read.getPendingRewardsPool([mockBookNFT]),
      ).to.equal(0n);
    });

    // Regression: increaseStakeToPosition compounds pending rewards into the position
    // principal, so pool.totalStaked must grow by amount + pendingRewards too. Otherwise it
    // drifts below the staked total, a later depositReward divides by an undersized total,
    // and getPendingRewardsForUser explodes. See on-chain book 0xa750...bed78d.
    //
    // Two stakers are required to expose the bug: with a single staker, the old
    // `pool.totalStaked = newAmount` and the fixed `pool.totalStaked += additionalAmount`
    // produce the same value, so a solo-staker test passes on the buggy code too.
    it("should keep totalStaked in sync when compounding rewards, so later rewards stay sane", async function () {
      const { likeCollective, rick, bob, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const rickStake = 1000n * 10n ** 6n;
      const bobStake = 2000n * 10n ** 6n;
      const additionalStakeAmount = 500n * 10n ** 6n;
      // 3000 splits 1:2 between rick and bob (1000 each for rick, 2000 for bob)
      const firstReward = 3000n * 10n ** 6n;
      // Equal to post-compound totalStaked so each staked unit earns exactly 1 LIKE
      const secondReward = 4500n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();

      // Rick and bob both stake into the same pool
      await likecoin.write.approve(
        [likeCollective.address, rickStake + additionalStakeAmount],
        { account: rick.account },
      );
      await likeCollective.write.newStakePosition([mockBookNFT, rickStake], {
        account: rick.account,
      });
      await likecoin.write.approve([likeCollective.address, bobStake], {
        account: bob.account,
      });
      await likeCollective.write.newStakePosition([mockBookNFT, bobStake], {
        account: bob.account,
      });

      // First reward: split 1:2 so rick accumulates 1000, bob accumulates 2000
      await likecoin.write.approve([likeCollective.address, firstReward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, firstReward], {
        account: kin.account,
      });
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(rickStake); // rick's 1/3 share = 1000

      // Compound: folds rick's 1000 pending reward into principal and adds additionalStakeAmount.
      // rick's additionalAmount = additionalStakeAmount + pendingRewards = 500 + 1000 = 1500
      // rick's new position    = rickStake + additionalAmount            = 1000 + 1500 = 2500
      // Expected pool total    = rick(2500) + bob(2000)                              = 4500
      //
      // pool.totalStaked += 1500    = 3000 + 1500 = 4500  ✓
      await likeCollective.write.increaseStakeToPosition(
        [nextTokenId, additionalStakeAmount],
        { account: rick.account },
      );

      const rickExpectedStake = rickStake + additionalStakeAmount + rickStake; // 2500
      const expectedPoolTotal = rickExpectedStake + bobStake; // 4500

      // Rick's individual position must reflect the compounded amount
      expect(
        await likeCollective.read.getStakeForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(rickExpectedStake);

      // Pool total must account for ALL stakers, not just rick's position.
      // Pre-fix this equalled only rick's newAmount (2500), dropping bob entirely.
      expect(await likeCollective.read.getTotalStake([mockBookNFT])).to.equal(
        expectedPoolTotal,
      );

      // The indexer derives staked totals by summing Staked event amounts, so the
      // compounded pending reward must be emitted as a Staked event too. Without it the
      // event sum lags getTotalStake and the indexer drifts out of sync with the chain.
      const stakedEvents = await likeCollective.getEvents.Staked(undefined, {
        fromBlock: 0n,
      });
      const totalStakedFromEvents = stakedEvents.reduce(
        (sum, e) => sum + (e.args.stakedAmount ?? 0n),
        0n,
      );
      expect(totalStakedFromEvents).to.equal(expectedPoolTotal);

      // Second reward (4500) divides by totalStaked (4500), so rick's proportional share
      // is rick(2500) / total(4500) * 4500 = 2500.
      // Pre-fix: totalStaked was 2500, giving rick(2500)/2500 * 4500 = 4500 (inflated).
      await likecoin.write.approve([likeCollective.address, secondReward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, secondReward], {
        account: kin.account,
      });
      expect(
        await likeCollective.read.getPendingRewardsForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(rickExpectedStake); // 2500, not the inflated 4500
    });

    it("should decrease stake without rewards and update balances/stake", async function () {
      const { likeCollective, rick, likeStakePosition, likecoin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const baseStakeAmount = 4000n * 10n ** 6n;
      const unstakeAmount = 1500n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();

      await likecoin.write.approve([likeCollective.address, baseStakeAmount], {
        account: rick.account,
      });
      await likeCollective.write.newStakePosition(
        [mockBookNFT, baseStakeAmount],
        {
          account: rick.account,
        },
      );

      // Decrease by unstakeAmount (no rewards pending)
      await likeCollective.write.decreaseStakePosition(
        [nextTokenId, unstakeAmount],
        {
          account: rick.account,
        },
      );

      // User receives unstakeAmount back
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        10000n * 10n ** 6n - baseStakeAmount + unstakeAmount,
      );
      expect(
        await likeCollective.read.getStakeForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(baseStakeAmount - unstakeAmount);
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(0n);
      expect(await likeCollective.read.getTotalStake([mockBookNFT])).to.equal(
        baseStakeAmount - unstakeAmount,
      );
    });

    it("should decrease stake and auto-claim existing rewards", async function () {
      const { likeCollective, rick, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const baseStakeAmount = 5000n * 10n ** 6n;
      const unstakeAmount = 1000n * 10n ** 6n;
      const rewardAmount = 1200n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();

      await likecoin.write.approve([likeCollective.address, baseStakeAmount], {
        account: rick.account,
      });
      await likeCollective.write.newStakePosition(
        [mockBookNFT, baseStakeAmount],
        {
          account: rick.account,
        },
      );
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        10000n * 10n ** 6n - baseStakeAmount,
      );

      // Deposit reward rewardAmount
      await likecoin.write.approve([likeCollective.address, rewardAmount], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, rewardAmount], {
        account: kin.account,
      });

      // Precondition: pending rewards == rewardAmount
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(rewardAmount);

      // Decrease by unstakeAmount — should transfer unstakeAmount + rewardAmount back to user
      await likeCollective.write.decreaseStakePosition(
        [nextTokenId, unstakeAmount],
        {
          account: rick.account,
        },
      );

      // Position reduced by unstakeAmount; rewards cleared
      expect(
        await likeCollective.read.getStakeForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(baseStakeAmount - unstakeAmount);
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(0n);

      // User receives unstakeAmount + rewardAmount; initial debit was baseStakeAmount
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        10000n * 10n ** 6n - baseStakeAmount + unstakeAmount + rewardAmount,
      );

      // Pool accounting: total stake decreased by unstakeAmount; pending pool reduced by rewardAmount
      expect(await likeCollective.read.getTotalStake([mockBookNFT])).to.equal(
        baseStakeAmount - unstakeAmount,
      );
      expect(
        await likeCollective.read.getPendingRewardsPool([mockBookNFT]),
      ).to.equal(0n);
    });
  });

  describe("Reward", async function () {
    it("should correctly correlate to get total stake for a bookNFT", async function () {
      const { likeCollective, rick, bob, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const amount = 2000n * 10n ** 6n;
      const reward = 500n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();
      await likecoin.write.approve([likeCollective.address, 2n * amount], {
        account: rick.account,
      });
      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: rick.account,
      });

      await likecoin.write.approve([likeCollective.address, 2n * reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });

      expect(await likeCollective.read.getTotalStake([mockBookNFT])).to.equal(
        amount,
      );
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(reward);
      expect(
        await likeCollective.read.getPendingRewardsForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(reward);

      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: rick.account,
      });
      await likecoin.write.approve([likeCollective.address, amount], {
        account: bob.account,
      });
      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: bob.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });

      expect(await likeCollective.read.getTotalStake([mockBookNFT])).to.equal(
        3n * amount,
      );
      expect(
        await likeCollective.read.getPendingRewardsPool([mockBookNFT]),
      ).to.equal(2n * reward);
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(reward + reward / 3n);
      expect(
        await likeCollective.read.getPendingRewardsForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(reward + (2n * reward) / 3n - 1n);

      await likeCollective.write.removeStakePosition([nextTokenId], {
        account: rick.account,
      });
      expect(
        await likeCollective.read.getPendingRewardsPool([mockBookNFT]),
      ).to.equal((2n * reward) / 3n + 1n);
      expect(
        await likeCollective.read.getPendingRewardsForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(reward / 3n);
    });

    it("should allow to claim single reward for a user", async function () {
      const { likeCollective, rick, bob, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const amount = 2000n * 10n ** 6n;
      const reward = 500n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();
      await likecoin.write.approve([likeCollective.address, amount], {
        account: rick.account,
      });
      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: rick.account,
      });

      await likecoin.write.approve([likeCollective.address, reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });

      await likeCollective.write.claimRewards([nextTokenId], {
        account: rick.account,
      });
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        8500n * 10n ** 6n,
      );
    });

    it("should allow to claim multiple positions of a BookNFT reward for a user", async function () {
      const { likeCollective, rick, bob, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const amount = 2000n * 10n ** 6n;
      const reward = 500n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();
      await likecoin.write.approve([likeCollective.address, 2n * amount], {
        account: rick.account,
      });
      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: rick.account,
      });
      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: rick.account,
      });

      await likecoin.write.approve([likeCollective.address, reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });

      await likeCollective.write.claimAllRewards([rick.account.address], {
        account: rick.account,
      });
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        6500n * 10n ** 6n,
      );
    });

    it("should allow to be claimed all rewards for a user", async function () {
      const { likeCollective, rick, bob, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const mockBookNFT2 = "0x2345678901234567890123456789012345678901";
      const amount = 2000n * 10n ** 6n;
      const reward = 1000n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();
      await likecoin.write.approve([likeCollective.address, 8n * amount], {
        account: rick.account,
      });
      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: rick.account,
      });
      await likeCollective.write.newStakePosition([mockBookNFT2, amount], {
        account: rick.account,
      });

      await likecoin.write.approve([likeCollective.address, 2n * reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT2, reward], {
        account: kin.account,
      });

      const pendingRewards = await likeCollective.read.getRewardsOfPosition([
        nextTokenId,
      ]);
      const pendingRewards2 = await likeCollective.read.getRewardsOfPosition([
        nextTokenId + 1n,
      ]);
      expect(pendingRewards).to.equal(reward);
      expect(pendingRewards2).to.equal(reward);

      await likeCollective.write.claimAllRewards([rick.account.address], {
        account: rick.account,
      });

      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        8000n * 10n ** 6n,
      );
    });
  });

  describe("Restake reward", async function () {
    it("should allow to restake reward", async function () {
      const { likeCollective, rick, bob, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901245678901";
      const amount = 2000n * 10n ** 6n;
      const reward = 1000n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();
      await likecoin.write.approve([likeCollective.address, amount], {
        account: rick.account,
      });
      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: rick.account,
      });

      await likecoin.write.approve([likeCollective.address, reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });

      await likeCollective.write.restakeRewardPosition([nextTokenId], {
        account: rick.account,
      });
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        8000n * 10n ** 6n,
      );
      expect(
        await likeCollective.read.getRewardsOfPosition([nextTokenId]),
      ).to.equal(0n);
      expect(
        await likeCollective.read.getStakeForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(amount + reward);
    });

    it("should not allow to restake reward for non-owner", async function () {
      const { likeCollective, rick, bob, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901245678901";
      const amount = 2000n * 10n ** 6n;
      const reward = 1000n * 10n ** 6n;

      const nextTokenId = await likeStakePosition.read.getNextTokenId();
      await likecoin.write.approve([likeCollective.address, amount], {
        account: rick.account,
      });
      await likeCollective.write.newStakePosition([mockBookNFT, amount], {
        account: rick.account,
      });

      await likecoin.write.approve([likeCollective.address, reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });

      await expect(
        likeCollective.write.restakeRewardPosition([nextTokenId], {
          account: bob.account,
        }),
      ).to.be.rejectedWith("ErrInvalidOwner()");
    });
  });
});
