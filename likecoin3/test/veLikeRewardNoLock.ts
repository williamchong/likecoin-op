import { loadFixture } from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem } from "hardhat";
import { encodeFunctionData } from "viem";

import "./setup";
import { initialMintNoLock } from "./factory";

const DECIMALS = 10n ** 6n;

async function deployNewVeLikeRewardNoLock(ownerAddress: `0x${string}`) {
  const impl = await viem.deployContract(
    "contracts/veLikeRewardNoLockV2.sol:veLikeRewardNoLock",
  );
  const initData = encodeFunctionData({
    abi: impl.abi,
    functionName: "initialize",
    args: [ownerAddress],
  });
  const proxy = await viem.deployContract("ERC1967Proxy", [
    impl.address,
    initData,
  ]);
  return await viem.getContractAt(
    "contracts/veLikeRewardNoLockV2.sol:veLikeRewardNoLock",
    proxy.address,
  );
}

describe("veLikeRewardNoLock", async function () {
  // -----------------------------------------------------------------------
  // Basic API calls
  // -----------------------------------------------------------------------
  describe("basic API", async function () {
    it("should have the deployer as owner", async function () {
      const { veLikeReward, deployer } = await loadFixture(initialMintNoLock);
      expect(await veLikeReward.read.owner()).to.equalAddress(
        deployer.account.address,
      );
    });

    it("should expose the configured vault and likecoin", async function () {
      const { veLikeReward, veLike, likecoin } =
        await loadFixture(initialMintNoLock);
      const [vault, likecoinConfig, rewardPool, totalStaked, lastRewardTime] =
        await veLikeReward.read.getConfig();
      expect(vault).to.equalAddress(veLike.address);
      expect(likecoinConfig).to.equalAddress(likecoin.address);
      expect(rewardPool).to.equal(0n);
      expect(totalStaked).to.equal(0n);
      expect(lastRewardTime).to.equal(0n);
    });

    it("should return an empty staking condition before any reward is added", async function () {
      const { veLikeReward } = await loadFixture(initialMintNoLock);
      const condition = await veLikeReward.read.getCurrentCondition();
      expect(condition.startTime).to.equal(0n);
      expect(condition.endTime).to.equal(0n);
      expect(condition.rewardAmount).to.equal(0n);
      expect(condition.rewardIndex).to.equal(0n);
    });

    it("should reject deposit() from a non-vault caller", async function () {
      const { veLikeReward, rick } = await loadFixture(initialMintNoLock);
      await expect(
        veLikeReward.write.deposit([rick.account.address, 100n * DECIMALS], {
          account: rick.account.address,
        }),
      ).to.be.rejectedWith("ErrUnauthorized");
    });

    it("initTotalStaked() should enable auto-sync and be callable only once", async function () {
      const { veLike, deployer } = await loadFixture(initialMintNoLock);
      const reward = await deployNewVeLikeRewardNoLock(
        deployer.account.address,
      );
      await reward.write.setVault([veLike.address], {
        account: deployer.account.address,
      });

      await reward.write.initTotalStaked({ account: deployer.account.address });
      // totalSupply is 0 in initialMintNoLock (no deposits yet).
      const [, , , totalStaked] = await reward.read.getConfig();
      expect(totalStaked).to.equal(0n);

      await expect(
        reward.write.initTotalStaked({ account: deployer.account.address }),
      ).to.be.rejectedWith("Already initialized");
    });

    it("finalizeSync() should revert before initTotalStaked()", async function () {
      const { veLike, deployer } = await loadFixture(initialMintNoLock);
      const reward = await deployNewVeLikeRewardNoLock(
        deployer.account.address,
      );
      await reward.write.setVault([veLike.address], {
        account: deployer.account.address,
      });
      await expect(
        reward.write.finalizeSync({ account: deployer.account.address }),
      ).to.be.rejectedWith("Not initialized or already finalized");
    });

    it("finalizeSync() should disable auto-sync and reject a second call", async function () {
      // The fixture's reward contract already has initTotalStaked() applied.
      const { veLikeReward, deployer } = await loadFixture(initialMintNoLock);
      await veLikeReward.write.finalizeSync({
        account: deployer.account.address,
      });
      await expect(
        veLikeReward.write.finalizeSync({
          account: deployer.account.address,
        }),
      ).to.be.rejectedWith("Not initialized or already finalized");
    });

    it("finalizeSync() should be owner-only", async function () {
      const { veLikeReward, rick } = await loadFixture(initialMintNoLock);
      await expect(
        veLikeReward.write.finalizeSync({
          account: rick.account.address,
        }),
      ).to.be.rejectedWith("OwnableUnauthorizedAccount");
    });

    it("syncStakers() should be callable outside an active reward period", async function () {
      const { veLike, likecoin, deployer, rick } =
        await loadFixture(initialMintNoLock);

      // rick gets a vault balance via the currently-active reward contract.
      await likecoin.write.approve([veLike.address, 100n * DECIMALS], {
        account: rick.account.address,
      });
      await veLike.write.deposit([100n * DECIMALS, rick.account.address], {
        account: rick.account.address,
      });

      // A freshly rotated-in contract: initTotalStaked captures rick's balance
      // but no reward period has been added yet, so it is NOT active.
      const reward2 = await deployNewVeLikeRewardNoLock(
        deployer.account.address,
      );
      await reward2.write.setVault([veLike.address], {
        account: deployer.account.address,
      });
      await reward2.write.initTotalStaked({
        account: deployer.account.address,
      });

      // syncStakers succeeds despite there being no active period, because
      // auto-sync is enabled.
      await reward2.write.syncStakers([[rick.account.address]], {
        account: deployer.account.address,
      });

      // rick is now materialized: a repeat sync reverts ErrAlreadySynced.
      await expect(
        reward2.write.syncStakers([[rick.account.address]], {
          account: deployer.account.address,
        }),
      ).to.be.rejectedWith("ErrAlreadySynced");
    });

    it("syncStakers() should be rejected once auto-sync is finalized", async function () {
      const { veLikeReward, deployer, rick } =
        await loadFixture(initialMintNoLock);
      await veLikeReward.write.finalizeSync({
        account: deployer.account.address,
      });
      await expect(
        veLikeReward.write.syncStakers([[rick.account.address]], {
          account: deployer.account.address,
        }),
      ).to.be.rejectedWith("Not initialized or already finalized");
    });
  });

  // -----------------------------------------------------------------------
  // Full rotation scenario:
  //   - Period 1: rick & kin stake (stakerInfos synced through the vault)
  //   - Period 2: a fresh veLikeRewardNoLock is rotated in with initTotalStaked
  //   - Period 2: bob (a late joiner) deposits, so his vault balance is non-zero
  //   - At rotation, reward1's stakers are materialized via syncStakers() (or
  //     already recorded via deposits) and then finalizeSync() disables
  //     auto-sync.
  //   - bob claims the period-1 legacy reward via veLike.claimLegacyReward
  //   - Expectation: bob earns ZERO from period 1 (he never staked in it).
  //     reward1 has auto-sync enabled (initTotalStaked runs in the fixture),
  //     but finalizeSync() turns it off before reward1 becomes legacy, so
  //     bob's period-2 vault balance must NOT leak into period-1 rewards.
  // -----------------------------------------------------------------------
  describe("rotation with a late joiner", async function () {
    it("bob's period-1 legacy reward is zero while rick & kin earn their share", async function () {
      // Step 1: setup of veLike with rick, kin, bob. reward1 is the period-1
      // reward contract already wired into the vault by the fixture.
      const {
        veLike,
        veLikeReward: reward1,
        likecoin,
        deployer,
        rick,
        kin,
        bob,
        publicClient,
        testClient,
      } = await loadFixture(initialMintNoLock);

      // --- Step 2: period 1 — rick & kin interact, stakerInfos get synced ---

      // Fund and start period 1: 10000 LIKE over 1000 seconds.
      await likecoin.write.approve([reward1.address, 10000n * DECIMALS], {
        account: deployer.account.address,
      });
      const block1 = await publicClient.getBlock();
      const start1 = block1.timestamp + 100n;
      const end1 = start1 + 1000n;
      await reward1.write.addReward(
        [deployer.account.address, 10000n * DECIMALS, start1, end1],
        { account: deployer.account.address },
      );

      // rick and kin each deposit 100 LIKE before the period starts. The vault
      // forwards each deposit to reward1.deposit(), syncing their stakerInfos.
      for (const user of [rick, kin]) {
        await likecoin.write.approve([veLike.address, 100n * DECIMALS], {
          account: user.account.address,
        });
        await veLike.write.deposit([100n * DECIMALS, user.account.address], {
          account: user.account.address,
        });
      }

      // stakerInfos synced: totalStaked reflects both deposits.
      const [, , , totalStakedP1] = await reward1.read.getConfig();
      expect(totalStakedP1).to.equal(200n * DECIMALS);

      // Halfway through period 1 each holds 50% of a 10000 LIKE pool.
      await testClient.setNextBlockTimestamp({ timestamp: start1 + 500n });
      await testClient.mine({ blocks: 1 });
      expect(
        await reward1.read.getPendingReward([rick.account.address]),
      ).to.equal(2500n * DECIMALS);
      expect(
        await reward1.read.getPendingReward([kin.account.address]),
      ).to.equal(2500n * DECIMALS);
      // bob never staked in period 1.
      expect(
        await reward1.read.getPendingReward([bob.account.address]),
      ).to.equal(0n);

      // Advance past period 1. Each accrues the full 5000 LIKE.
      await testClient.setNextBlockTimestamp({ timestamp: end1 + 1n });
      await testClient.mine({ blocks: 1 });
      expect(
        await reward1.read.getPendingReward([rick.account.address]),
      ).to.equal(5000n * DECIMALS);
      expect(
        await reward1.read.getPendingReward([kin.account.address]),
      ).to.equal(5000n * DECIMALS);

      // --- Step 3: period 2 — rotate in a fresh contract with initTotalStaked ---
      const reward2 = await deployNewVeLikeRewardNoLock(
        deployer.account.address,
      );
      await reward2.write.setVault([veLike.address], {
        account: deployer.account.address,
      });
      await reward2.write.setLikecoin([likecoin.address], {
        account: deployer.account.address,
      });

      // initTotalStaked captures all existing vault holders (rick + kin = 200)
      // and enables auto-sync for them in period 2.
      await reward2.write.initTotalStaked({
        account: deployer.account.address,
      });
      const [, , , totalStakedP2Init] = await reward2.read.getConfig();
      expect(totalStakedP2Init).to.equal(200n * DECIMALS);

      // Close out reward1 before it becomes legacy. Its stakers (rick & kin)
      // are already recorded via their period-1 deposits; a larger holder set
      // would be materialized across syncStakers() batches first. finalizeSync()
      // then disables auto-sync, so reward1 will no longer credit anyone from
      // their current vault balance.
      await reward1.write.finalizeSync({ account: deployer.account.address });

      // Switch the vault's active reward contract to reward2, and allowlist
      // reward1 so users can still claim period-1 rewards.
      await veLike.write.setRewardContract([reward2.address], {
        account: deployer.account.address,
      });
      await veLike.write.setLegacyRewardContract([reward1.address, true], {
        account: deployer.account.address,
      });

      // Fund and start period 2: 5000 LIKE over 500 seconds.
      await likecoin.write.approve([reward2.address, 5000n * DECIMALS], {
        account: deployer.account.address,
      });
      const block2 = await publicClient.getBlock();
      const start2 = block2.timestamp + 100n;
      const end2 = start2 + 500n;
      await reward2.write.addReward(
        [deployer.account.address, 5000n * DECIMALS, start2, end2],
        { account: deployer.account.address },
      );

      // --- Step 4: period 2 — bob deposits, so his vault balance is non-zero ---
      await testClient.setNextBlockTimestamp({ timestamp: start2 });
      await testClient.mine({ blocks: 1 });

      await likecoin.write.approve([veLike.address, 100n * DECIMALS], {
        account: bob.account.address,
      });
      await veLike.write.deposit([100n * DECIMALS, bob.account.address], {
        account: bob.account.address,
      });

      // bob now holds veLIKE (his vault balance is non-zero) and is staked in
      // reward2 only.
      expect(await veLike.read.balanceOf([bob.account.address])).to.equal(
        100n * DECIMALS,
      );
      const [, , , totalStakedP2After] = await reward2.read.getConfig();
      expect(totalStakedP2After).to.equal(300n * DECIMALS);

      // --- Step 5 & 6: bob claims the period-1 legacy reward → must be ZERO ---

      // reward1 was closed out with finalizeSync() at rotation, so auto-sync
      // is off. bob's period-2 vault balance therefore does NOT make him
      // eligible for period-1 rewards. His effective period-1 stake is 0,
      // hence pending reward is 0.
      expect(
        await reward1.read.getPendingReward([bob.account.address]),
      ).to.equal(0n);

      // A zero legacy reward surfaces as ErrNoRewardToClaim, and bob's LIKE
      // balance is unchanged.
      const bobBalanceBefore = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      await expect(
        veLike.write.claimLegacyReward([reward1.address, bob.account.address], {
          account: bob.account.address,
        }),
      ).to.be.rejectedWith("ErrNoRewardToClaim");
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        bobBalanceBefore,
      );

      // Sanity contrast: rick (a genuine period-1 staker) DOES get his 5000
      // LIKE legacy reward, proving the zero result for bob is meaningful.
      const rickBalanceBefore = await likecoin.read.balanceOf([
        rick.account.address,
      ]);
      await veLike.write.claimLegacyReward(
        [reward1.address, rick.account.address],
        { account: rick.account.address },
      );
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        rickBalanceBefore + 5000n * DECIMALS,
      );
    });

    // -------------------------------------------------------------------
    // Scenario:
    //   1. veLike with rick, kin, bob, alice.
    //   2. Period 1: rick & kin stake (stakerInfos synced).
    //   3. Period 1 expires, then bob deposits while reward1 is still the
    //      active reward contract — he joins too late to earn anything from
    //      period 1 (his rewardIndex is set to the frozen final index).
    //   4. Rotate: reward1 is closed out with finalizeSync(); reward2 is set
    //      up with initTotalStaked() (capturing rick + kin + bob = 300).
    //   5. Period 2: alice (a brand-new staker) deposits, earns her correct
    //      period-2 share, and her period-1 legacy claim is zero.
    //   6. bob's period-1 reward is zero (joined after the period ended).
    // -------------------------------------------------------------------
    it("late depositor (bob) and new joiner (alice) both earn zero from period 1", async function () {
      // --- Step 1: setup with rick, kin, bob, alice ---
      const {
        veLike,
        veLikeReward: reward1,
        likecoin,
        deployer,
        rick,
        kin,
        bob,
        publicClient,
        testClient,
      } = await loadFixture(initialMintNoLock);
      // alice is not part of the fixture wallet set; fund her explicitly.
      const alice = (await viem.getWalletClients())[4];
      await likecoin.write.mint([alice.account.address, 10000n * DECIMALS], {
        account: deployer.account.address,
      });

      // --- Step 2: period 1 — rick & kin stake; stakerInfos synced ---
      await likecoin.write.approve([reward1.address, 10000n * DECIMALS], {
        account: deployer.account.address,
      });
      const block1 = await publicClient.getBlock();
      const start1 = block1.timestamp + 100n;
      const end1 = start1 + 1000n;
      await reward1.write.addReward(
        [deployer.account.address, 10000n * DECIMALS, start1, end1],
        { account: deployer.account.address },
      );

      for (const user of [rick, kin]) {
        await likecoin.write.approve([veLike.address, 100n * DECIMALS], {
          account: user.account.address,
        });
        await veLike.write.deposit([100n * DECIMALS, user.account.address], {
          account: user.account.address,
        });
      }

      const [, , , totalStakedP1] = await reward1.read.getConfig();
      expect(totalStakedP1).to.equal(200n * DECIMALS);

      // --- Step 3: period 1 expires, then bob deposits (too late) ---
      await testClient.setNextBlockTimestamp({ timestamp: end1 + 1n });
      await testClient.mine({ blocks: 1 });

      // rick & kin earned their full 5000 each over the period.
      expect(
        await reward1.read.getPendingReward([rick.account.address]),
      ).to.equal(5000n * DECIMALS);
      expect(
        await reward1.read.getPendingReward([kin.account.address]),
      ).to.equal(5000n * DECIMALS);

      // bob deposits after the period ended. reward1 is still the active
      // reward contract, so the deposit is routed here and bob's rewardIndex
      // is set to the frozen final index — he earns nothing from period 1.
      await likecoin.write.approve([veLike.address, 100n * DECIMALS], {
        account: bob.account.address,
      });
      await veLike.write.deposit([100n * DECIMALS, bob.account.address], {
        account: bob.account.address,
      });
      expect(
        await reward1.read.getPendingReward([bob.account.address]),
      ).to.equal(0n);
      const [, , , totalStakedP1AfterBob] = await reward1.read.getConfig();
      expect(totalStakedP1AfterBob).to.equal(300n * DECIMALS);

      // --- Step 4: rotate — finalizeSync(reward1), set up reward2 ---
      // bob, rick, kin are all materialized (via deposits), so no syncStakers
      // batch is needed before finalizing.
      await reward1.write.finalizeSync({ account: deployer.account.address });

      const reward2 = await deployNewVeLikeRewardNoLock(
        deployer.account.address,
      );
      await reward2.write.setVault([veLike.address], {
        account: deployer.account.address,
      });
      await reward2.write.setLikecoin([likecoin.address], {
        account: deployer.account.address,
      });
      // Capture all existing vault holders (rick + kin + bob = 300).
      await reward2.write.initTotalStaked({
        account: deployer.account.address,
      });
      const [, , , totalStakedP2Init] = await reward2.read.getConfig();
      expect(totalStakedP2Init).to.equal(300n * DECIMALS);

      await veLike.write.setRewardContract([reward2.address], {
        account: deployer.account.address,
      });
      await veLike.write.setLegacyRewardContract([reward1.address, true], {
        account: deployer.account.address,
      });

      // Fund and start period 2: 5000 LIKE over 500 seconds.
      await likecoin.write.approve([reward2.address, 5000n * DECIMALS], {
        account: deployer.account.address,
      });
      const block2 = await publicClient.getBlock();
      const start2 = block2.timestamp + 100n;
      const end2 = start2 + 500n;
      await reward2.write.addReward(
        [deployer.account.address, 5000n * DECIMALS, start2, end2],
        { account: deployer.account.address },
      );

      // --- Step 5: period 2 — alice deposits, earns her correct share ---
      // alice deposits before the period starts, so she stakes for the whole
      // period alongside the 3 auto-enrolled holders (4 x 100 = 400 total).
      await likecoin.write.approve([veLike.address, 100n * DECIMALS], {
        account: alice.account.address,
      });
      await veLike.write.deposit([100n * DECIMALS, alice.account.address], {
        account: alice.account.address,
      });
      const [, , , totalStakedP2AfterAlice] = await reward2.read.getConfig();
      expect(totalStakedP2AfterAlice).to.equal(400n * DECIMALS);

      // Run the full period 2.
      await testClient.setNextBlockTimestamp({ timestamp: end2 + 1n });
      await testClient.mine({ blocks: 1 });

      // 5000 LIKE over 4 equal stakers => 1250 LIKE each.
      expect(
        await veLike.read.getPendingReward([alice.account.address]),
      ).to.equal(1250n * DECIMALS);

      // alice's period-1 legacy claim is zero: reward1 is finalized, so her
      // period-2 vault balance is not credited there, and she never staked in
      // period 1.
      expect(
        await reward1.read.getPendingReward([alice.account.address]),
      ).to.equal(0n);
      await expect(
        veLike.write.claimLegacyReward(
          [reward1.address, alice.account.address],
          { account: alice.account.address },
        ),
      ).to.be.rejectedWith("ErrNoRewardToClaim");

      // alice can still claim her period-2 reward properly.
      const aliceBalanceBefore = await likecoin.read.balanceOf([
        alice.account.address,
      ]);
      await veLike.write.claimReward([alice.account.address], {
        account: alice.account.address,
      });
      expect(await likecoin.read.balanceOf([alice.account.address])).to.equal(
        aliceBalanceBefore + 1250n * DECIMALS,
      );

      // --- Step 6: bob's period-1 reward is zero ---
      expect(
        await reward1.read.getPendingReward([bob.account.address]),
      ).to.equal(0n);
      const bobBalanceBefore = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      await expect(
        veLike.write.claimLegacyReward([reward1.address, bob.account.address], {
          account: bob.account.address,
        }),
      ).to.be.rejectedWith("ErrNoRewardToClaim");
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        bobBalanceBefore,
      );

      // Sanity contrast: rick still claims his genuine 5000 period-1 reward.
      const rickBalanceBefore = await likecoin.read.balanceOf([
        rick.account.address,
      ]);
      await veLike.write.claimLegacyReward(
        [reward1.address, rick.account.address],
        { account: rick.account.address },
      );
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        rickBalanceBefore + 5000n * DECIMALS,
      );
    });

    // -------------------------------------------------------------------
    // Scenario (existing staker tops up after each period ends):
    //   1. veLike with rick, kin, bob.
    //   2. Period 1: rick & kin stake (stakerInfos synced).
    //   3. Period 1 expires, then kin tops up while reward1 is still active.
    //      The deposit auto-claims kin's earned period-1 reward, and the newly
    //      added stake earns nothing more.
    //   4. Rotate: finalizeSync(reward1); reward2 with initTotalStaked().
    //   5. Period 2: bob (brand-new) deposits, earns his correct share, and
    //      his period-1 legacy claim is zero.
    //   6. kin's period-1 reward is correct (the 5000 he auto-claimed).
    //   7. Period 2 expires; reward2 is still active. kin tops up — the
    //      deposit auto-claims his correct period-2 reward, and the new stake
    //      earns nothing more.
    // -------------------------------------------------------------------
    it("existing staker (kin) topping up after a period ends earns nothing extra", async function () {
      const {
        veLike,
        veLikeReward: reward1,
        likecoin,
        deployer,
        rick,
        kin,
        bob,
        publicClient,
        testClient,
      } = await loadFixture(initialMintNoLock);

      // --- Step 2: period 1 — rick & kin stake ---
      await likecoin.write.approve([reward1.address, 10000n * DECIMALS], {
        account: deployer.account.address,
      });
      const block1 = await publicClient.getBlock();
      const start1 = block1.timestamp + 100n;
      const end1 = start1 + 1000n;
      await reward1.write.addReward(
        [deployer.account.address, 10000n * DECIMALS, start1, end1],
        { account: deployer.account.address },
      );

      for (const user of [rick, kin]) {
        await likecoin.write.approve([veLike.address, 100n * DECIMALS], {
          account: user.account.address,
        });
        await veLike.write.deposit([100n * DECIMALS, user.account.address], {
          account: user.account.address,
        });
      }
      expect((await reward1.read.getConfig())[3]).to.equal(200n * DECIMALS);

      // --- Step 3: period 1 expires, then kin tops up ---
      await testClient.setNextBlockTimestamp({ timestamp: end1 + 1n });
      await testClient.mine({ blocks: 1 });

      // Both earned their full 5000 over the period.
      expect(
        await reward1.read.getPendingReward([rick.account.address]),
      ).to.equal(5000n * DECIMALS);
      expect(
        await reward1.read.getPendingReward([kin.account.address]),
      ).to.equal(5000n * DECIMALS);

      // kin tops up 100 more. The deposit auto-claims his 5000 period-1 reward;
      // the new stake is pinned to the frozen final index, so it earns nothing.
      const kinBalanceBeforeTopup = await likecoin.read.balanceOf([
        kin.account.address,
      ]);
      await likecoin.write.approve([veLike.address, 100n * DECIMALS], {
        account: kin.account.address,
      });
      await veLike.write.deposit([100n * DECIMALS, kin.account.address], {
        account: kin.account.address,
      });

      // kin received his 5000 reward minus the 100 he just staked.
      expect(await likecoin.read.balanceOf([kin.account.address])).to.equal(
        kinBalanceBeforeTopup - 100n * DECIMALS + 5000n * DECIMALS,
      );
      // No further period-1 reward accrues to the new stake.
      expect(
        await reward1.read.getPendingReward([kin.account.address]),
      ).to.equal(0n);
      expect((await reward1.read.getConfig())[3]).to.equal(300n * DECIMALS);

      // --- Step 4: rotate — finalizeSync(reward1), set up reward2 ---
      await reward1.write.finalizeSync({ account: deployer.account.address });

      const reward2 = await deployNewVeLikeRewardNoLock(
        deployer.account.address,
      );
      await reward2.write.setVault([veLike.address], {
        account: deployer.account.address,
      });
      await reward2.write.setLikecoin([likecoin.address], {
        account: deployer.account.address,
      });
      // Captures rick (100) + kin (200) = 300.
      await reward2.write.initTotalStaked({
        account: deployer.account.address,
      });
      expect((await reward2.read.getConfig())[3]).to.equal(300n * DECIMALS);

      await veLike.write.setRewardContract([reward2.address], {
        account: deployer.account.address,
      });
      await veLike.write.setLegacyRewardContract([reward1.address, true], {
        account: deployer.account.address,
      });

      // Fund and start period 2: 5000 LIKE over 500 seconds.
      await likecoin.write.approve([reward2.address, 5000n * DECIMALS], {
        account: deployer.account.address,
      });
      const block2 = await publicClient.getBlock();
      const start2 = block2.timestamp + 100n;
      const end2 = start2 + 500n;
      await reward2.write.addReward(
        [deployer.account.address, 5000n * DECIMALS, start2, end2],
        { account: deployer.account.address },
      );

      // --- Step 5: period 2 — bob (brand-new) deposits ---
      // Stakes: rick 100, kin 200 (auto-enrolled), bob 100 => 400 total.
      await likecoin.write.approve([veLike.address, 100n * DECIMALS], {
        account: bob.account.address,
      });
      await veLike.write.deposit([100n * DECIMALS, bob.account.address], {
        account: bob.account.address,
      });
      expect((await reward2.read.getConfig())[3]).to.equal(400n * DECIMALS);

      await testClient.setNextBlockTimestamp({ timestamp: end2 + 1n });
      await testClient.mine({ blocks: 1 });

      // bob's share: 100/400 of 5000 = 1250.
      expect(
        await veLike.read.getPendingReward([bob.account.address]),
      ).to.equal(1250n * DECIMALS);

      // bob's period-1 legacy claim is zero (never staked in period 1).
      expect(
        await reward1.read.getPendingReward([bob.account.address]),
      ).to.equal(0n);
      await expect(
        veLike.write.claimLegacyReward([reward1.address, bob.account.address], {
          account: bob.account.address,
        }),
      ).to.be.rejectedWith("ErrNoRewardToClaim");

      // bob claims his correct period-2 reward.
      const bobBalanceBefore = await likecoin.read.balanceOf([
        bob.account.address,
      ]);
      await veLike.write.claimReward([bob.account.address], {
        account: bob.account.address,
      });
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        bobBalanceBefore + 1250n * DECIMALS,
      );

      // --- Step 6: kin's period-1 reward is correct (5000, auto-claimed) ---
      expect(
        await reward1.read.getClaimedReward([kin.account.address]),
      ).to.equal(5000n * DECIMALS);
      expect(
        await reward1.read.getPendingReward([kin.account.address]),
      ).to.equal(0n);

      // --- Step 7: period 2 ended, reward2 still active, kin tops up ---
      // kin's pending period-2 reward is his 200/400 share = 2500.
      expect(
        await reward2.read.getPendingReward([kin.account.address]),
      ).to.equal(2500n * DECIMALS);

      const kinBalanceBeforeP2Topup = await likecoin.read.balanceOf([
        kin.account.address,
      ]);
      await likecoin.write.approve([veLike.address, 100n * DECIMALS], {
        account: kin.account.address,
      });
      await veLike.write.deposit([100n * DECIMALS, kin.account.address], {
        account: kin.account.address,
      });

      // The deposit auto-claims kin's 2500 period-2 reward; the new 100 stake
      // earns nothing more (period already ended).
      expect(await likecoin.read.balanceOf([kin.account.address])).to.equal(
        kinBalanceBeforeP2Topup - 100n * DECIMALS + 2500n * DECIMALS,
      );
      expect(
        await reward2.read.getClaimedReward([kin.account.address]),
      ).to.equal(2500n * DECIMALS);
      expect(
        await reward2.read.getPendingReward([kin.account.address]),
      ).to.equal(0n);
    });
  });
});
