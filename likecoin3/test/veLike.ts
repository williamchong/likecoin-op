import {
  time,
  loadFixture,
} from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem, ignition } from "hardhat";

import "./setup";
import { deployVeLike, initialMint, initialCondition } from "./factory";

describe("veLike ", async function () {
  describe("as ERC1967 Proxy", async function () {
    it("should have the correct owner", async function () {
      const { veLike, deployer } = await loadFixture(deployVeLike);
      expect(await veLike.read.owner()).to.equalAddress(
        deployer.account.address,
      );
    });

    it("should have the correct asset address", async function () {
      const { veLike, likecoin } = await loadFixture(deployVeLike);
      expect(await veLike.read.asset()).to.equalAddress(likecoin.address);
    });

    it("should have the correct name and symbol", async function () {
      const { veLike } = await loadFixture(deployVeLike);
      expect(await veLike.read.name()).to.equal("vote-escrowed LikeCoin");
      expect(await veLike.read.symbol()).to.equal("veLIKE");
    });

    it("should have the correct decimals", async function () {
      const { veLike } = await loadFixture(deployVeLike);
      expect(await veLike.read.decimals()).to.equal(6);
    });

    it("should have the correct STORAGE_SLOT", async function () {
      const { veLike, deployer } = await loadFixture(deployVeLike);
      const veLikeMock = await viem.deployContract("veLikeMock");
      veLike.write.upgradeTo(veLikeMock.address, {
        account: deployer.account,
      });
      const newVeLike = await viem.getContractAt(
        "veLikeMock",
        veLikeMock.address,
      );
      expect(await newVeLike.read.version()).to.equal(2n);
      expect(await newVeLike.read.dataStorage()).to.equal(
        "0xb9e14b2a89d227541697d62a06ecbf5ccc9ad849800745b40b2826662a177600",
      );
    });
  });

  describe("as pausable contract", async function () {
    it("should be paused by default", async function () {
      const { veLike } = await loadFixture(initialMint);
      const paused = await veLike.read.paused();
      expect(paused).to.be.false;
    });

    it("should not allow non-owner to pause", async function () {
      const { veLike, rick } = await loadFixture(initialMint);
      await expect(
        veLike.write.pause({ account: rick.account.address }),
      ).to.be.rejectedWith("OwnableUnauthorizedAccount");
    });

    it("should not able to deposit when pausehave the correct owner", async function () {
      const { veLike, likecoin, deployer, rick } =
        await loadFixture(initialCondition);
      await veLike.write.pause();

      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });
      await expect(
        veLike.write.deposit([100n * 10n ** 6n, rick.account.address], {
          account: rick.account.address,
        }),
      ).to.be.rejectedWith("EnforcedPause");
    });
  });

  describe("as ERC20", async function () {
    it("should have the zero initial supply supply", async function () {
      const { veLike } = await loadFixture(initialMint);
      expect(await veLike.read.totalSupply()).to.equal(0n);
    });

    it("should not able to transfer between account", async function () {
      const { veLike, likecoin, deployer, rick, kin } =
        await loadFixture(initialMint);
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });
      await veLike.write.deposit([100n * 10n ** 6n, rick.account.address], {
        account: rick.account.address,
      });

      expect(await veLike.read.balanceOf([rick.account.address])).to.equal(
        100n * 10n ** 6n,
      );
      await expect(
        veLike.write.transfer([kin.account.address, 100n * 10n ** 6n], {
          account: rick.account.address,
        }),
      ).to.be.rejectedWith("ErrNonTransferable");
    });
  });

  describe("as ERC4626 vault", async function () {
    it("should have increase supply after deposit", async function () {
      const { veLike, likecoin, deployer, rick } =
        await loadFixture(initialMint);
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });
      await veLike.write.deposit([100n * 10n ** 6n, rick.account.address], {
        account: rick.account.address,
      });
      expect(await veLike.read.balanceOf([rick.account.address])).to.equal(
        100n * 10n ** 6n,
      );
      expect(await veLike.read.totalAssets()).to.equal(100n * 10n ** 6n);
      expect(await veLike.read.totalSupply()).to.equal(100n * 10n ** 6n);
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        9900n * 10n ** 6n,
      );
    });

    it("should have increase supply after multiple deposit", async function () {
      const { veLike, likecoin, deployer, rick, kin } =
        await loadFixture(initialMint);
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });
      await veLike.write.deposit([100n * 10n ** 6n, rick.account.address], {
        account: rick.account.address,
      });
      expect(await veLike.read.balanceOf([rick.account.address])).to.equal(
        100n * 10n ** 6n,
      );
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        9900n * 10n ** 6n,
      );

      await likecoin.write.approve([veLike.address, 200n * 10n ** 6n], {
        account: kin.account.address,
      });
      await veLike.write.deposit([200n * 10n ** 6n, kin.account.address], {
        account: kin.account.address,
      });
      expect(await veLike.read.balanceOf([kin.account.address])).to.equal(
        200n * 10n ** 6n,
      );
      expect(await likecoin.read.balanceOf([kin.account.address])).to.equal(
        9800n * 10n ** 6n,
      );

      expect(await veLike.read.totalAssets()).to.equal(300n * 10n ** 6n);
      expect(await veLike.read.totalSupply()).to.equal(300n * 10n ** 6n);
    });

    it("should keep total assets as zero after adding reward", async function () {
      const { veLike, likecoin, deployer, publicClient } =
        await loadFixture(initialMint);
      await likecoin.write.approve([veLike.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      const block = await publicClient.getBlock();
      const startTime = block.timestamp;
      const endTime = block.timestamp + 1000n;
      expect(await veLike.read.totalAssets()).to.equal(0n);
    });

    it("should have increase supply after deposit and add reward", async function () {
      const { veLike, likecoin, deployer, rick, publicClient } =
        await loadFixture(initialMint);
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });
      await veLike.write.deposit([100n * 10n ** 6n, rick.account.address], {
        account: rick.account.address,
      });
      expect(await veLike.read.balanceOf([rick.account.address])).to.equal(
        100n * 10n ** 6n,
      );
      expect(await veLike.read.totalAssets()).to.equal(100n * 10n ** 6n);
      expect(await veLike.read.totalSupply()).to.equal(100n * 10n ** 6n);
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        9900n * 10n ** 6n,
      );

      const block = await publicClient.getBlock();
      const startTime = block.timestamp;
      const endTime = block.timestamp + 1000n;
      await likecoin.write.approve([veLike.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      expect(await veLike.read.totalAssets()).to.equal(100n * 10n ** 6n);
    });
  });

  describe("as a no-lock vault", async function () {
    it("should have lock time as zero (no lock)", async function () {
      const { veLike } = await loadFixture(initialCondition);
      expect(await veLike.read.getLockTime()).to.equal(0n);
    });

    it("should allow withdraw during active reward period", async function () {
      const { veLike, veLikeReward, likecoin, rick, testClient, startTime } =
        await loadFixture(initialCondition);
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });
      await veLike.write.deposit([100n * 10n ** 6n, rick.account.address], {
        account: rick.account.address,
      });
      // Withdraw during active period should succeed
      await veLike.write.withdraw(
        [100n * 10n ** 6n, rick.account.address, rick.account.address],
        {
          account: rick.account.address,
        },
      );
      expect(await veLike.read.balanceOf([rick.account.address])).to.equal(0n);
      // Rick gets back 10000 LIKE + any auto-claimed reward from the time between deposit and withdraw
      const rickBalance = await likecoin.read.balanceOf([rick.account.address]);
      expect(rickBalance >= 10000n * 10n ** 6n).to.be.true;
    });

    it("should support partial withdraw and keep remaining stake", async function () {
      const { veLike, veLikeReward, likecoin, rick, testClient, startTime } =
        await loadFixture(initialCondition);
      await likecoin.write.approve([veLike.address, 200n * 10n ** 6n], {
        account: rick.account.address,
      });
      await veLike.write.deposit([200n * 10n ** 6n, rick.account.address], {
        account: rick.account.address,
      });
      expect(await veLike.read.balanceOf([rick.account.address])).to.equal(
        200n * 10n ** 6n,
      );

      // Partial withdraw: take out 50 LIKE
      await veLike.write.withdraw(
        [50n * 10n ** 6n, rick.account.address, rick.account.address],
        {
          account: rick.account.address,
        },
      );
      expect(await veLike.read.balanceOf([rick.account.address])).to.equal(
        150n * 10n ** 6n,
      );
      // Rick started with 10000, deposited 200, withdrew 50 => 9850 LIKE + any claimed reward
      // (reward claimed auto on withdraw is 0 since just deposited same block)
      const rickBalance = await likecoin.read.balanceOf([rick.account.address]);
      expect(rickBalance >= 9850n * 10n ** 6n).to.be.true;
    });

    it("should correctly track reward after partial withdraw", async function () {
      const {
        veLike,
        veLikeReward,
        likecoin,
        bob,
        rick,
        testClient,
        startTime,
        endTime,
      } = await loadFixture(initialCondition);
      // Bob already has 100 LIKE deposited from fixture.
      // Advance 500 seconds (half the period).
      await testClient.setNextBlockTimestamp({
        timestamp: startTime + 500n,
      });
      await testClient.mine({ blocks: 1 });

      // Bob's pending reward at half-period: ~5000 LIKE (sole staker, 10000 total over 1000s)
      const pendingBefore = await veLike.read.getPendingReward([
        bob.account.address,
      ]);
      expect(pendingBefore).to.equal(5000n * 10n ** 6n);

      // Bob partial withdraws 50 LIKE (keeps 50 staked) at the exact half-point
      await testClient.setNextBlockTimestamp({
        timestamp: startTime + 500n + 1n,
      });
      await veLike.write.withdraw(
        [50n * 10n ** 6n, bob.account.address, bob.account.address],
        {
          account: bob.account.address,
        },
      );

      // After withdraw, bob should have 50 veLIKE remaining
      expect(await veLike.read.balanceOf([bob.account.address])).to.equal(
        50n * 10n ** 6n,
      );

      // Advance to end of period
      await testClient.setNextBlockTimestamp({
        timestamp: endTime,
      });
      await testClient.mine({ blocks: 1 });

      // Bob should have earned additional reward on remaining 50 LIKE for the rest of the period
      // Bob is still the sole staker, so gets all remaining reward
      const pendingAfter = await veLike.read.getPendingReward([
        bob.account.address,
      ]);
      // Remaining period: endTime - (startTime+501) = 499 seconds
      // Reward rate: 10000 LIKE / 1000s = 10 LIKE/s
      // Expected: 499 * 10 = 4990 LIKE
      expect(pendingAfter).to.equal(4990n * 10n ** 6n);
    });

    it("should be able to withdraw after the reward period ends", async function () {
      const { veLike, veLikeReward, likecoin, bob, testClient, endTime } =
        await loadFixture(initialCondition);
      await testClient.setNextBlockTimestamp({
        timestamp: endTime + 100n,
      });
      // Must mine for follow read command to work.
      await testClient.mine({
        blocks: 1,
      });
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        9900n * 10n ** 6n,
      );
      expect(await veLike.read.balanceOf([bob.account.address])).to.equal(
        100n * 10n ** 6n,
      );
      const pendingReward = await veLike.read.getPendingReward([
        bob.account.address,
      ]);
      expect(pendingReward).to.equal(10000n * 10n ** 6n);
      await veLike.write.claimReward([bob.account.address], {
        account: bob.account.address,
      });
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        10000n * 10n ** 6n + 10000n * 10n ** 6n - 100n * 10n ** 6n,
      );
      await veLike.write.withdraw(
        [100n * 10n ** 6n, bob.account.address, bob.account.address],
        {
          account: bob.account.address,
        },
      );
      expect(await veLike.read.balanceOf([bob.account.address])).to.equal(0n);
      // Bob receive all the reward from the reward contract.
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        10000n * 10n ** 6n + 10000n * 10n ** 6n,
      );
    });
  });

  describe("as a lockable vault for whole period", async function () {
    async function lockedCondition() {
      const result = await loadFixture(initialCondition);
      await result.veLike.write.setLockTime([result.endTime], {
        account: result.deployer.account.address,
      });
      return result;
    }

    it("should set the lock time same as the condition", async function () {
      const { veLike, likecoin, rick, endTime } = await lockedCondition();
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });
      await veLike.write.deposit([100n * 10n ** 6n, rick.account.address], {
        account: rick.account.address,
      });
      expect(await veLike.read.getLockTime()).to.equal(endTime);
    });

    it("should be revert on withdraw when the condition is active", async function () {
      const { veLike, likecoin, rick } = await lockedCondition();
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });
      await veLike.write.deposit([100n * 10n ** 6n, rick.account.address], {
        account: rick.account.address,
      });
      await expect(
        veLike.write.withdraw(
          [100n * 10n ** 6n, rick.account.address, rick.account.address],
          {
            account: rick.account.address,
          },
        ),
      ).to.be.rejectedWith("ErrWithdrawLocked");
    });

    it("should be able to withdraw when the block timestamp is after the lock time", async function () {
      const { veLike, likecoin, bob, testClient, endTime } =
        await lockedCondition();
      await testClient.setNextBlockTimestamp({
        timestamp: endTime + 100n,
      });
      await testClient.mine({
        blocks: 1,
      });
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        9900n * 10n ** 6n,
      );
      expect(await veLike.read.balanceOf([bob.account.address])).to.equal(
        100n * 10n ** 6n,
      );
      const pendingReward = await veLike.read.getPendingReward([
        bob.account.address,
      ]);
      expect(pendingReward).to.equal(10000n * 10n ** 6n);
      await veLike.write.claimReward([bob.account.address], {
        account: bob.account.address,
      });
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        10000n * 10n ** 6n + 10000n * 10n ** 6n - 100n * 10n ** 6n,
      );
      await veLike.write.withdraw(
        [100n * 10n ** 6n, bob.account.address, bob.account.address],
        {
          account: bob.account.address,
        },
      );
      expect(await veLike.read.balanceOf([bob.account.address])).to.equal(0n);
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        10000n * 10n ** 6n + 10000n * 10n ** 6n,
      );
    });
  });

  describe("reward distribution on empty account and condition", async function () {
    it("should have return zero on account never deposited", async function () {
      const { veLike, rick } = await loadFixture(initialMint);
      expect(
        await veLike.read.getPendingReward([rick.account.address]),
      ).to.equal(0n);
    });

    it("should be revert on claimReward on empty account", async function () {
      const { veLike, rick } = await loadFixture(initialMint);
      await expect(
        veLike.write.claimReward([rick.account.address], {
          account: rick.account.address,
        }),
      ).to.be.rejectedWith("ErrNoRewardToClaim");
    });

    it("should be revert on restakeReward on empty account", async function () {
      const { veLike, rick } = await loadFixture(initialMint);
      await expect(
        veLike.write.restakeReward([rick.account.address], {
          account: rick.account.address,
        }),
      ).to.be.rejectedWith("ErrNoRewardToClaim");
    });
  });
});
