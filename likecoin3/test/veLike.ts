import {
  time,
  loadFixture,
} from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem, ignition } from "hardhat";
import { encodeFunctionData } from "viem";

import "./setup";
import {
  deployVeLike,
  deployVeLikeReward,
  initialMint,
  initialCondition,
  initialMintNoLock,
  initialConditionNoLock,
} from "./factory";

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
      const { veLike } = await loadFixture(initialConditionNoLock);
      expect(await veLike.read.getLockTime()).to.equal(0n);
    });

    it("should allow withdraw during active reward period", async function () {
      const { veLike, veLikeReward, likecoin, rick, testClient, startTime } =
        await loadFixture(initialConditionNoLock);
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
        await loadFixture(initialConditionNoLock);
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
      } = await loadFixture(initialConditionNoLock);
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
        await loadFixture(initialConditionNoLock);
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
      const result = await loadFixture(initialConditionNoLock);
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

  describe("legacy reward claims", async function () {
    it("should allow owner to set legacy reward contract", async function () {
      const { veLike, veLikeReward, deployer } = await loadFixture(
        initialConditionNoLock,
      );
      // setLegacyRewardContract should succeed for owner
      await veLike.write.setLegacyRewardContract([veLikeReward.address, true], {
        account: deployer.account.address,
      });
    });

    it("should not allow non-owner to set legacy reward contract", async function () {
      const { veLike, veLikeReward, rick } = await loadFixture(
        initialConditionNoLock,
      );
      await expect(
        veLike.write.setLegacyRewardContract([veLikeReward.address, true], {
          account: rick.account.address,
        }),
      ).to.be.rejectedWith("OwnableUnauthorizedAccount");
    });

    it("should revert claimLegacyReward on non-allowlisted contract", async function () {
      const { veLike, veLikeReward, bob } = await loadFixture(
        initialConditionNoLock,
      );
      await expect(
        veLike.write.claimLegacyReward(
          [veLikeReward.address, bob.account.address],
          {
            account: bob.account.address,
          },
        ),
      ).to.be.rejectedWith("ErrNotLegacyRewardContract");
    });

    it("should allow user to claim legacy reward from allowlisted contract", async function () {
      const {
        veLike,
        veLikeReward,
        likecoin,
        deployer,
        bob,
        testClient,
        endTime,
      } = await loadFixture(initialConditionNoLock);

      // Advance past the reward period end
      await testClient.setNextBlockTimestamp({ timestamp: endTime + 100n });
      await testClient.mine({ blocks: 1 });

      // Bob has 100 LIKE staked and earned all 10000 LIKE reward
      const pendingReward = await veLike.read.getPendingReward([
        bob.account.address,
      ]);
      expect(pendingReward).to.equal(10000n * 10n ** 6n);

      // Now simulate rotation: allowlist the old reward contract
      await veLike.write.setLegacyRewardContract([veLikeReward.address, true], {
        account: deployer.account.address,
      });

      // Set reward contract to address(0) to simulate rotation
      await veLike.write.setRewardContract(
        ["0x0000000000000000000000000000000000000000"],
        { account: deployer.account.address },
      );

      // Bob claims legacy reward
      const balanceBefore = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      await veLike.write.claimLegacyReward(
        [veLikeReward.address, bob.account.address],
        {
          account: bob.account.address,
        },
      );
      const balanceAfter = await likecoin.read.balanceOf([bob.account.address]);

      // Bob should have received the full reward
      expect(balanceAfter - balanceBefore).to.equal(10000n * 10n ** 6n);
    });

    it("should revert claimLegacyReward after removing from allowlist", async function () {
      const { veLike, veLikeReward, deployer, bob, testClient, endTime } =
        await loadFixture(initialConditionNoLock);

      // Advance past the reward period end
      await testClient.setNextBlockTimestamp({ timestamp: endTime + 100n });
      await testClient.mine({ blocks: 1 });

      // Allowlist, then remove
      await veLike.write.setLegacyRewardContract([veLikeReward.address, true], {
        account: deployer.account.address,
      });
      await veLike.write.setLegacyRewardContract(
        [veLikeReward.address, false],
        { account: deployer.account.address },
      );

      // Should revert since no longer allowlisted
      await expect(
        veLike.write.claimLegacyReward(
          [veLikeReward.address, bob.account.address],
          {
            account: bob.account.address,
          },
        ),
      ).to.be.rejectedWith("ErrNotLegacyRewardContract");
    });
  });

  describe("legacy reward with lock", async function () {
    /**
     * Simulates the real production path:
     * - Period 1 uses veLikeReward (original) with lock enabled
     * - Period 1 ends, lock expires
     * - Rotate to veLikeRewardNoLock for period 2
     * - Bob claims legacy reward from the locked veLikeReward
     * - Auto-enrollment works for period 2
     */
    it("should allow claiming legacy reward from a locked veLikeReward after rotation", async function () {
      const {
        veLike,
        veLikeReward: reward1,
        likecoin,
        deployer,
        bob,
        rick,
        publicClient,
        testClient,
      } = await loadFixture(initialMint);

      // --- Period 1 setup with veLikeReward (original) + lock ---
      // Bob deposits 100 LIKE before reward period
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: bob.account.address,
      });
      await veLike.write.deposit([100n * 10n ** 6n, bob.account.address], {
        account: bob.account.address,
      });

      // Fund period 1: 10000 LIKE over 1000 seconds
      await likecoin.write.approve([reward1.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      const block1 = await publicClient.getBlock();
      const start1 = block1.timestamp + 100n;
      const end1 = start1 + 1000n;
      await reward1.write.addReward(
        [deployer.account.address, 10000n * 10n ** 6n, start1, end1],
        { account: deployer.account.address },
      );

      // Set lock time to end of period 1 (users can't withdraw until period ends)
      await veLike.write.setLockTime([end1], {
        account: deployer.account.address,
      });

      // Advance to start of period 1
      await testClient.setNextBlockTimestamp({ timestamp: start1 });
      await testClient.mine({ blocks: 1 });

      // Verify Bob cannot withdraw during locked period
      await expect(
        veLike.write.withdraw(
          [100n * 10n ** 6n, bob.account.address, bob.account.address],
          { account: bob.account.address },
        ),
      ).to.be.rejectedWith("ErrWithdrawLocked");

      // Advance past period 1 (lock also expires since lockTime == end1)
      await testClient.setNextBlockTimestamp({ timestamp: end1 + 1n });
      await testClient.mine({ blocks: 1 });

      // Bob earned all 10000 LIKE from period 1
      const pendingP1 = await veLike.read.getPendingReward([
        bob.account.address,
      ]);
      expect(pendingP1).to.equal(10000n * 10n ** 6n);

      // --- Rotation: deploy veLikeRewardNoLock, replace reward1 ---
      const reward2Impl = await viem.deployContract("veLikeRewardNoLock");
      const reward2InitData = encodeFunctionData({
        abi: reward2Impl.abi,
        functionName: "initialize",
        args: [deployer.account.address],
      });
      const reward2Proxy = await viem.deployContract("ERC1967Proxy", [
        reward2Impl.address,
        reward2InitData,
      ]);
      const reward2 = await viem.getContractAt(
        "veLikeRewardNoLock",
        reward2Proxy.address,
      );
      await reward2.write.setVault([veLike.address], {
        account: deployer.account.address,
      });
      await reward2.write.setLikecoin([likecoin.address], {
        account: deployer.account.address,
      });

      // Remove lock for period 2 (no-lock model going forward)
      await veLike.write.setLockTime([0n], {
        account: deployer.account.address,
      });

      // Pre-initialize totalStaked on reward2 (captures Bob's 100 LIKE)
      await reward2.write.initTotalStaked({
        account: deployer.account.address,
      });

      // Switch active reward to reward2, mark reward1 as legacy
      await veLike.write.setRewardContract([reward2.address], {
        account: deployer.account.address,
      });
      await veLike.write.setLegacyRewardContract([reward1.address, true], {
        account: deployer.account.address,
      });

      // --- Period 2 setup: 5000 LIKE over 500 seconds ---
      await likecoin.write.approve([reward2.address, 5000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      const block2 = await publicClient.getBlock();
      const start2 = block2.timestamp + 100n;
      const end2 = start2 + 500n;
      await reward2.write.addReward(
        [deployer.account.address, 5000n * 10n ** 6n, start2, end2],
        { account: deployer.account.address },
      );

      // Approve Rick's deposit before advancing time
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });

      // Advance to period 2 start
      await testClient.setNextBlockTimestamp({ timestamp: start2 });
      await testClient.mine({ blocks: 1 });

      // Rick deposits 100 LIKE at start2 + 1
      await testClient.setNextBlockTimestamp({ timestamp: start2 + 1n });
      await veLike.write.deposit([100n * 10n ** 6n, rick.account.address], {
        account: rick.account.address,
      });

      // Advance past period 2
      await testClient.setNextBlockTimestamp({ timestamp: end2 + 1n });
      await testClient.mine({ blocks: 1 });

      // --- Verify: Bob auto-enrolled in period 2, gets retroactive rewards ---
      // Same math as rotation integration test:
      // 5000 LIKE over 500s = 10 LIKE/s
      // 1s with Bob only (totalStaked=100): Bob gets 10 LIKE
      // 499s with Bob+Rick (totalStaked=200): Bob 2495, Rick 2495
      // Bob total: 2505, Rick total: 2495
      const bobP2 = await veLike.read.getPendingReward([bob.account.address]);
      const rickP2 = await veLike.read.getPendingReward([rick.account.address]);
      expect(bobP2).to.equal(2505n * 10n ** 6n);
      expect(rickP2).to.equal(2495n * 10n ** 6n);

      // --- Bob claims legacy reward from period 1 (veLikeReward with lock) ---
      const bobBalanceBefore = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      await veLike.write.claimLegacyReward(
        [reward1.address, bob.account.address],
        { account: bob.account.address },
      );
      const bobBalanceAfter = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      expect(bobBalanceAfter - bobBalanceBefore).to.equal(10000n * 10n ** 6n);

      // --- Bob claims current reward from period 2 ---
      const bobBalanceBefore2 = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      await veLike.write.claimReward([bob.account.address], {
        account: bob.account.address,
      });
      const bobBalanceAfter2 = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      expect(bobBalanceAfter2 - bobBalanceBefore2).to.equal(2505n * 10n ** 6n);

      // --- Rick claims current reward from period 2 ---
      const rickBalanceBefore = await likecoin.read.balanceOf([
        rick.account.address,
      ]);
      await veLike.write.claimReward([rick.account.address], {
        account: rick.account.address,
      });
      const rickBalanceAfter = await likecoin.read.balanceOf([
        rick.account.address,
      ]);
      expect(rickBalanceAfter - rickBalanceBefore).to.equal(2495n * 10n ** 6n);

      // --- Rick has no legacy reward (wasn't in period 1) ---
      await expect(
        veLike.write.claimLegacyReward(
          [reward1.address, rick.account.address],
          { account: rick.account.address },
        ),
      ).to.be.rejectedWith("ErrNoRewardToClaim");
    });
  });

  describe("reward rotation integration", async function () {
    async function deployNewVeLikeRewardNoLock(ownerAddress: `0x${string}`) {
      const impl = await viem.deployContract("veLikeRewardNoLock");
      const initData = encodeFunctionData({
        abi: impl.abi,
        functionName: "initialize",
        args: [ownerAddress],
      });
      const proxy = await viem.deployContract("ERC1967Proxy", [
        impl.address,
        initData,
      ]);
      return await viem.getContractAt("veLikeRewardNoLock", proxy.address);
    }

    /**
     * Flow: deploy reward1 → Bob stakes → period 1 ends →
     * rotate to reward2 (reward1 becomes legacy) → initTotalStaked →
     * Rick stakes → period 2 ends →
     * Bob auto-earns period 2 rewards (lazy sync), Rick also earns.
     * Bob claims legacy for period 1.
     */
    it("should support full rotation with auto-enrollment of existing stakers", async function () {
      const {
        veLike,
        likecoin,
        deployer,
        bob,
        rick,
        publicClient,
        testClient,
      } = await loadFixture(initialMintNoLock);

      // --- Period 1 setup ---
      const reward1 = await deployNewVeLikeRewardNoLock(
        deployer.account.address,
      );
      await reward1.write.setVault([veLike.address], {
        account: deployer.account.address,
      });
      await reward1.write.setLikecoin([likecoin.address], {
        account: deployer.account.address,
      });
      await veLike.write.setRewardContract([reward1.address], {
        account: deployer.account.address,
      });

      // Bob deposits 100 LIKE into the vault
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: bob.account.address,
      });
      await veLike.write.deposit([100n * 10n ** 6n, bob.account.address], {
        account: bob.account.address,
      });

      // Fund and start period 1: 10000 LIKE over 1000 seconds
      await likecoin.write.approve([reward1.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      const block1 = await publicClient.getBlock();
      const start1 = block1.timestamp + 100n;
      const end1 = start1 + 1000n;
      await reward1.write.addReward(
        [deployer.account.address, 10000n * 10n ** 6n, start1, end1],
        { account: deployer.account.address },
      );

      // Advance past period 1
      await testClient.setNextBlockTimestamp({ timestamp: end1 + 1n });
      await testClient.mine({ blocks: 1 });

      const pendingP1 = await veLike.read.getPendingReward([
        bob.account.address,
      ]);
      expect(pendingP1).to.equal(10000n * 10n ** 6n);

      // --- Rotate: reward2 replaces reward1, reward1 becomes legacy ---
      const reward2 = await deployNewVeLikeRewardNoLock(
        deployer.account.address,
      );
      await reward2.write.setVault([veLike.address], {
        account: deployer.account.address,
      });
      await reward2.write.setLikecoin([likecoin.address], {
        account: deployer.account.address,
      });

      // Pre-initialize totalStaked on reward2 BEFORE setting as active reward.
      // This captures all existing vault holders (Bob's 100 LIKE).
      await reward2.write.initTotalStaked({
        account: deployer.account.address,
      });

      await veLike.write.setRewardContract([reward2.address], {
        account: deployer.account.address,
      });
      await veLike.write.setLegacyRewardContract([reward1.address, true], {
        account: deployer.account.address,
      });

      // getPendingReward now queries reward2, which knows nothing about Bob yet
      // but _effectiveStakedAmount returns vault balance for un-synced users
      expect(
        await veLike.read.getPendingReward([bob.account.address]),
      ).to.equal(0n);

      // --- Period 2 setup: 5000 LIKE over 500 seconds ---
      await likecoin.write.approve([reward2.address, 5000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      const block2 = await publicClient.getBlock();
      const start2 = block2.timestamp + 100n;
      const end2 = start2 + 500n;
      await reward2.write.addReward(
        [deployer.account.address, 5000n * 10n ** 6n, start2, end2],
        { account: deployer.account.address },
      );

      // Approve Rick's deposit before advancing time
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });

      // Advance to period start
      await testClient.setNextBlockTimestamp({ timestamp: start2 });
      await testClient.mine({ blocks: 1 });

      // Rick deposits 100 LIKE (same amount as Bob) at start2 + 1
      await testClient.setNextBlockTimestamp({ timestamp: start2 + 1n });
      await veLike.write.deposit([100n * 10n ** 6n, rick.account.address], {
        account: rick.account.address,
      });

      // Advance past period 2
      await testClient.setNextBlockTimestamp({ timestamp: end2 + 1n });
      await testClient.mine({ blocks: 1 });

      // --- Verify auto-enrollment: Bob earns period 2 rewards without re-depositing ---
      // Bob (100 LIKE) and Rick (100 LIKE) share the 5000 LIKE reward.
      // Bob's share is retroactive from period start (rewardIndex stays 0).
      //
      // Total reward = 5000 LIKE over 500 seconds = 10 LIKE/s
      // For 1 second before Rick deposits: totalStaked = 100 (Bob only)
      //   Bob gets 100% of 10 LIKE/s = 10 LIKE
      //
      // For remaining 499 seconds: totalStaked = 200 (Bob + Rick)
      //   Bob gets 50% of 10 LIKE/s * 499 = 2495 LIKE
      //   Rick gets 50% of 10 LIKE/s * 499 = 2495 LIKE
      //
      // Bob total: 10 + 2495 = 2505 LIKE
      // Rick total: 2495 LIKE
      // Grand total: 2505 + 2495 = 5000 LIKE ✓
      const bobP2 = await veLike.read.getPendingReward([bob.account.address]);
      const rickP2 = await veLike.read.getPendingReward([rick.account.address]);
      expect(bobP2).to.equal(2505n * 10n ** 6n);
      expect(rickP2).to.equal(2495n * 10n ** 6n);

      // --- Bob claims legacy reward from period 1 ---
      const bobBalanceBefore = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      await veLike.write.claimLegacyReward(
        [reward1.address, bob.account.address],
        {
          account: bob.account.address,
        },
      );
      const bobBalanceAfter = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      expect(bobBalanceAfter - bobBalanceBefore).to.equal(10000n * 10n ** 6n);

      // --- Rick claims current reward from period 2 ---
      const rickBalanceBefore = await likecoin.read.balanceOf([
        rick.account.address,
      ]);
      await veLike.write.claimReward([rick.account.address], {
        account: rick.account.address,
      });
      const rickBalanceAfter = await likecoin.read.balanceOf([
        rick.account.address,
      ]);
      expect(rickBalanceAfter - rickBalanceBefore).to.equal(2495n * 10n ** 6n);

      // --- Bob claims current reward from period 2 (auto-enrolled) ---
      const bobBalanceBefore2 = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      await veLike.write.claimReward([bob.account.address], {
        account: bob.account.address,
      });
      const bobBalanceAfter2 = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      expect(bobBalanceAfter2 - bobBalanceBefore2).to.equal(2505n * 10n ** 6n);

      // Rick has no legacy reward (wasn't in period 1)
      await expect(
        veLike.write.claimLegacyReward(
          [reward1.address, rick.account.address],
          {
            account: rick.account.address,
          },
        ),
      ).to.be.rejectedWith("ErrNoRewardToClaim");
    });
  });

  describe("syncStakers and lazy sync", async function () {
    async function deployNewVeLikeRewardNoLock(ownerAddress: `0x${string}`) {
      const impl = await viem.deployContract("veLikeRewardNoLock");
      const initData = encodeFunctionData({
        abi: impl.abi,
        functionName: "initialize",
        args: [ownerAddress],
      });
      const proxy = await viem.deployContract("ERC1967Proxy", [
        impl.address,
        initData,
      ]);
      return await viem.getContractAt("veLikeRewardNoLock", proxy.address);
    }

    /**
     * Shared fixture for syncStakers tests:
     * - Bob deposits 100 LIKE via bootstrap reward contract
     * - reward1 deployed with initTotalStaked (captures Bob's 100)
     * - reward1 set as active, period 1 funded (10000 LIKE / 1000s)
     * - Bob does NO operation during period 1 (never synced into reward1)
     */
    async function syncStakersFixture() {
      const {
        veLike,
        likecoin,
        deployer,
        bob,
        rick,
        publicClient,
        testClient,
      } = await loadFixture(initialMintNoLock);

      // Bob deposits 100 LIKE via bootstrap reward contract
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: bob.account.address,
      });
      await veLike.write.deposit([100n * 10n ** 6n, bob.account.address], {
        account: bob.account.address,
      });

      // Deploy reward1 with initTotalStaked
      const reward1 = await deployNewVeLikeRewardNoLock(
        deployer.account.address,
      );
      await reward1.write.setVault([veLike.address], {
        account: deployer.account.address,
      });
      await reward1.write.setLikecoin([likecoin.address], {
        account: deployer.account.address,
      });
      await reward1.write.initTotalStaked({
        account: deployer.account.address,
      });
      await veLike.write.setRewardContract([reward1.address], {
        account: deployer.account.address,
      });

      // Fund period 1: 10000 LIKE over 1000 seconds
      await likecoin.write.approve([reward1.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      const block1 = await publicClient.getBlock();
      const start1 = block1.timestamp + 100n;
      const end1 = start1 + 1000n;
      await reward1.write.addReward(
        [deployer.account.address, 10000n * 10n ** 6n, start1, end1],
        { account: deployer.account.address },
      );

      return {
        veLike,
        likecoin,
        deployer,
        bob,
        rick,
        publicClient,
        testClient,
        reward1,
        start1,
        end1,
      };
    }

    it("should revert with ErrNotActive before period starts", async function () {
      const { reward1, deployer, bob } = await loadFixture(syncStakersFixture);

      // Period hasn't started yet — block.timestamp < start1
      await expect(
        reward1.write.syncStakers([[bob.account.address]], {
          account: deployer.account.address,
        }),
      ).to.be.rejectedWith("ErrNotActive");
    });

    it("should revert with ErrNotActive after period ends", async function () {
      const { reward1, deployer, bob, testClient, end1 } =
        await loadFixture(syncStakersFixture);

      // Advance past end of period
      await testClient.setNextBlockTimestamp({ timestamp: end1 + 1n });
      await testClient.mine({ blocks: 1 });

      await expect(
        reward1.write.syncStakers([[bob.account.address]], {
          account: deployer.account.address,
        }),
      ).to.be.rejectedWith("ErrNotActive");
    });

    it("should sync un-synced account during active period", async function () {
      const { reward1, deployer, bob, testClient, start1 } =
        await loadFixture(syncStakersFixture);

      // Advance to active period
      await testClient.setNextBlockTimestamp({ timestamp: start1 + 1n });
      await testClient.mine({ blocks: 1 });

      // Bob is un-synced (stakedAmount == 0 in reward1)
      await reward1.write.syncStakers([[bob.account.address]], {
        account: deployer.account.address,
      });

      // Verify Bob's pending reward is non-zero (he is now synced and earning)
      const pending = await reward1.read.getPendingReward([
        bob.account.address,
      ]);
      expect(pending > 0n).to.be.true;
    });

    it("should revert with ErrAlreadySynced when account is synced with matching balance", async function () {
      const { reward1, deployer, bob, testClient, start1 } =
        await loadFixture(syncStakersFixture);

      // Advance to active period and sync Bob
      await testClient.setNextBlockTimestamp({ timestamp: start1 + 1n });
      await testClient.mine({ blocks: 1 });

      await reward1.write.syncStakers([[bob.account.address]], {
        account: deployer.account.address,
      });

      // Try to sync again — vault balance still matches stakedAmount
      await expect(
        reward1.write.syncStakers([[bob.account.address]], {
          account: deployer.account.address,
        }),
      ).to.be.rejectedWith("ErrAlreadySynced");
    });

    it("should revert with ErrMismatchSync when account is synced but balance changed", async function () {
      const { veLike, reward1, likecoin, deployer, bob, testClient, start1 } =
        await loadFixture(syncStakersFixture);

      // Advance to active period and sync Bob (stakedAmount = 100)
      await testClient.setNextBlockTimestamp({ timestamp: start1 + 1n });
      await testClient.mine({ blocks: 1 });

      await reward1.write.syncStakers([[bob.account.address]], {
        account: deployer.account.address,
      });

      // Switch active reward contract away from reward1 so that
      // veLike.deposit goes to a different contract, leaving reward1's
      // stakedAmount unchanged while vault balance grows.
      const reward2 = await deployNewVeLikeRewardNoLock(
        deployer.account.address,
      );
      await reward2.write.setVault([veLike.address], {
        account: deployer.account.address,
      });
      await reward2.write.setLikecoin([likecoin.address], {
        account: deployer.account.address,
      });
      await reward2.write.initTotalStaked({
        account: deployer.account.address,
      });
      await veLike.write.setRewardContract([reward2.address], {
        account: deployer.account.address,
      });

      // Bob deposits 50 more LIKE → vault balance = 150
      // This goes through reward2.deposit(), so reward1's stakedAmount stays 100
      await likecoin.write.approve([veLike.address, 50n * 10n ** 6n], {
        account: bob.account.address,
      });
      await veLike.write.deposit([50n * 10n ** 6n, bob.account.address], {
        account: bob.account.address,
      });

      // Try to sync on reward1 — stakedAmount (100) != vaultBalance (150)
      await expect(
        reward1.write.syncStakers([[bob.account.address]], {
          account: deployer.account.address,
        }),
      ).to.be.rejectedWith("ErrMismatchSync");
    });

    it("should revert whole tx on first bad account in batch", async function () {
      const { reward1, deployer, bob, rick, testClient, start1 } =
        await loadFixture(syncStakersFixture);

      // Advance to active period and sync Bob first
      await testClient.setNextBlockTimestamp({ timestamp: start1 + 1n });
      await testClient.mine({ blocks: 1 });

      await reward1.write.syncStakers([[bob.account.address]], {
        account: deployer.account.address,
      });

      // Batch with [bob (already synced), rick (un-synced)] — reverts on bob
      await expect(
        reward1.write.syncStakers(
          [[bob.account.address, rick.account.address]],
          { account: deployer.account.address },
        ),
      ).to.be.rejectedWith("ErrAlreadySynced");
    });

    it("should revert when called by non-owner", async function () {
      const { reward1, bob, testClient, start1 } =
        await loadFixture(syncStakersFixture);

      // Advance to active period
      await testClient.setNextBlockTimestamp({ timestamp: start1 + 1n });
      await testClient.mine({ blocks: 1 });

      await expect(
        reward1.write.syncStakers([[bob.account.address]], {
          account: bob.account.address,
        }),
      ).to.be.rejectedWith("OwnableUnauthorizedAccount");
    });

    /**
     * Full rotation flow with syncStakers fixing the stale-balance problem:
     *
     * Period 0 (bootstrap): Bob deposits 100 LIKE via initial reward contract.
     *
     * Period 1 (veLikeRewardNoLock #1):
     *   - Deployed with initTotalStaked() → totalStaked = 100
     *   - Bob does NO operation on veLike during period 1
     *   - Owner calls syncStakers([bob]) before rotation → freezes Bob at 100
     *   - 10000 LIKE reward accrues over 1000 seconds
     *
     * Rotation to period 2 (veLikeRewardNoLock #2):
     *   - Bob deposits 50 more LIKE → vault balance goes from 100 → 150
     *   - 5000 LIKE reward accrues over 500 seconds
     *
     * Bob claims:
     *   (a) legacy reward from reward1 — based on frozen 100 LIKE → 10000 LIKE
     *   (b) current reward from reward2 — based on 150 LIKE → 5000 LIKE
     */
    it("should claim legacy and current rewards based on respective period balances", async function () {
      const {
        veLike,
        likecoin,
        deployer,
        bob,
        publicClient,
        testClient,
        reward1,
        start1,
        end1,
      } = await loadFixture(syncStakersFixture);

      // Advance to active period, then sync Bob's stake into reward1
      await testClient.setNextBlockTimestamp({ timestamp: start1 + 1n });
      await testClient.mine({ blocks: 1 });

      // Owner syncs Bob before rotation — freezes stakedAmount at 100
      await reward1.write.syncStakers([[bob.account.address]], {
        account: deployer.account.address,
      });

      // Advance past period 1
      await testClient.setNextBlockTimestamp({ timestamp: end1 + 1n });
      await testClient.mine({ blocks: 1 });

      // --- Rotate: reward2 replaces reward1, reward1 becomes legacy ---
      const reward2 = await deployNewVeLikeRewardNoLock(
        deployer.account.address,
      );
      await reward2.write.setVault([veLike.address], {
        account: deployer.account.address,
      });
      await reward2.write.setLikecoin([likecoin.address], {
        account: deployer.account.address,
      });
      await reward2.write.initTotalStaked({
        account: deployer.account.address,
      });
      await veLike.write.setRewardContract([reward2.address], {
        account: deployer.account.address,
      });
      await veLike.write.setLegacyRewardContract([reward1.address, true], {
        account: deployer.account.address,
      });

      // --- Bob deposits 50 more LIKE → vault balance: 100 → 150 ---
      // reward1 has stakedAmount = 100 (frozen by syncStakers), unaffected.
      await likecoin.write.approve([veLike.address, 50n * 10n ** 6n], {
        account: bob.account.address,
      });
      await veLike.write.deposit([50n * 10n ** 6n, bob.account.address], {
        account: bob.account.address,
      });

      // Fund period 2: 5000 LIKE over 500 seconds
      await likecoin.write.approve([reward2.address, 5000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      const block2 = await publicClient.getBlock();
      const start2 = block2.timestamp + 100n;
      const end2 = start2 + 500n;
      await reward2.write.addReward(
        [deployer.account.address, 5000n * 10n ** 6n, start2, end2],
        { account: deployer.account.address },
      );

      // Advance past period 2
      await testClient.setNextBlockTimestamp({ timestamp: end2 + 1n });
      await testClient.mine({ blocks: 1 });

      // --- Claim legacy reward from period 1 ---
      // Bob's frozen stakedAmount = 100, totalStaked = 100 → reward = 10000 LIKE
      const bobBefore = await likecoin.read.balanceOf([bob.account.address]);
      await veLike.write.claimLegacyReward(
        [reward1.address, bob.account.address],
        { account: bob.account.address },
      );
      const bobAfter = await likecoin.read.balanceOf([bob.account.address]);
      const legacyReward = bobAfter - bobBefore;
      expect(legacyReward).to.equal(10000n * 10n ** 6n);

      // --- Claim current reward from period 2 ---
      // Bob has 150 LIKE staked (sole staker) → gets all 5000 LIKE
      const bobBefore2 = await likecoin.read.balanceOf([bob.account.address]);
      await veLike.write.claimReward([bob.account.address], {
        account: bob.account.address,
      });
      const bobAfter2 = await likecoin.read.balanceOf([bob.account.address]);
      const currentReward = bobAfter2 - bobBefore2;
      // Allow 1-unit rounding tolerance from accumulator integer division
      const expected2 = 5000n * 10n ** 6n;
      expect(currentReward >= expected2 - 1n && currentReward <= expected2).to
        .be.true;
    });
  });
});
