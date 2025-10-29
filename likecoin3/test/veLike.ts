import {
  time,
  loadFixture,
} from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem, ignition } from "hardhat";

import "./setup";
import veLikeModule from "../ignition/modules/veLike";
import { deployProtocol } from "./factory";

describe("veLike ", async function () {
  async function deployVeLike() {
    const [deployer, rick, kin, bob] = await viem.getWalletClients();
    const publicClient = await viem.getPublicClient();
    const testClient = await viem.getTestClient();
    const { veLike, veLikeImpl, veLikeProxy, likecoin } = await ignition.deploy(
      veLikeModule,
      {
        parameters: {
          LikecoinModule: {
            initOwner: deployer.account.address,
          },
          veLikeV0Module: {
            initOwner: deployer.account.address,
          },
        },
        defaultSender: deployer.account.address,
        strategy: "create2",
      },
    );
    return {
      veLike,
      veLikeImpl,
      veLikeProxy,
      likecoin,
      deployer,
      rick,
      kin,
      bob,
      publicClient,
      testClient,
    };
  }

  async function initialMint() {
    const {
      veLike,
      likecoin,
      deployer,
      rick,
      kin,
      bob,
      publicClient,
      testClient,
    } = await loadFixture(deployVeLike);
    await likecoin.write.mint([deployer.account.address, 50000n * 10n ** 6n], {
      account: deployer.account.address,
    });
    await likecoin.write.mint([rick.account.address, 10000n * 10n ** 6n], {
      account: deployer.account.address,
    });
    await likecoin.write.mint([kin.account.address, 10000n * 10n ** 6n], {
      account: deployer.account.address,
    });
    await likecoin.write.mint([bob.account.address, 10000n * 10n ** 6n], {
      account: deployer.account.address,
    });
    return {
      veLike,
      likecoin,
      deployer,
      rick,
      kin,
      bob,
      publicClient,
      testClient,
    };
  }

  async function initialCondition() {
    const {
      veLike,
      likecoin,
      deployer,
      publicClient,
      rick,
      kin,
      bob,
      testClient,
    } = await loadFixture(initialMint);
    await likecoin.write.approve([veLike.address, 10000n * 10n ** 6n], {
      account: deployer.account.address,
    });
    const block = await publicClient.getBlock();
    const startTime = block.timestamp + 100n;
    const endTime = startTime + 1000n;
    await veLike.write.addReward([10000n * 10n ** 6n, startTime, endTime], {
      account: deployer.account.address,
    });

    await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
      account: bob.account.address,
    });
    await veLike.write.deposit([100n * 10n ** 6n, bob.account.address], {
      account: bob.account.address,
    });

    // Test case assume start of block is the startTime
    await testClient.setNextBlockTimestamp({
      timestamp: startTime,
    });
    await testClient.mine({
      blocks: 1,
    });

    return {
      veLike,
      likecoin,
      deployer,
      publicClient,
      rick,
      kin,
      bob,
      startTime,
      endTime,
      testClient,
    };
  }

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
      await veLike.write.addReward([10000n * 10n ** 6n, startTime, endTime], {
        account: deployer.account.address,
      });
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
      await veLike.write.addReward([10000n * 10n ** 6n, startTime, endTime], {
        account: deployer.account.address,
      });
      expect(await veLike.read.totalAssets()).to.equal(100n * 10n ** 6n);
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

  describe("reward condition setting", async function () {
    it("should return empty reward condition if no reward condition is set", async function () {
      const { veLike } = await loadFixture(initialMint);
      const condition = await veLike.read.getCurrentCondition();
      expect(condition.startTime).to.equal(0n);
      expect(condition.endTime).to.equal(0n);
      expect(condition.rewardAmount).to.equal(0n);
      expect(condition.rewardIndex).to.equal(0n);
    });

    it("should return the correct reward condition if reward condition is set", async function () {
      const { veLike, likecoin, deployer, publicClient } =
        await loadFixture(initialMint);

      await likecoin.write.approve([veLike.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });

      const block = await publicClient.getBlock();
      const startTime = block.timestamp;
      const endTime = block.timestamp + 1000n;
      await veLike.write.addReward([10000n * 10n ** 6n, startTime, endTime], {
        account: deployer.account.address,
      });

      const condition = await veLike.read.getCurrentCondition();
      expect(condition.startTime).to.equal(startTime);
      expect(condition.endTime).to.equal(endTime);
      expect(condition.rewardAmount).to.equal(10000n * 10n ** 6n);
      expect(condition.rewardIndex).to.equal(0n);

      expect(
        await likecoin.read.balanceOf([deployer.account.address]),
      ).to.equal(40000n * 10n ** 6n);
      // Reward does not count toward total assets
      expect(await veLike.read.totalAssets()).to.equal(0n);
    });

    it("should not able to set reward condition with startTime before current lastRewardBlock", async function () {
      const { veLike, likecoin, deployer, publicClient } =
        await loadFixture(initialMint);

      await likecoin.write.approve([veLike.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });

      const block = await publicClient.getBlock();
      const startTime = block.timestamp;
      const endTime = block.timestamp + 1000n;
      await veLike.write.addReward([10000n * 10n ** 6n, startTime, endTime], {
        account: deployer.account.address,
      });
      // The lastRewardBlock is updated to the startTime

      await likecoin.write.approve([veLike.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      await expect(
        veLike.write.addReward(
          [10000n * 10n ** 6n, startTime - 1000n, endTime],
          {
            account: deployer.account.address,
          },
        ),
      ).to.be.rejectedWith("ErrConflictCondition");
    });

    it("should not able to set reward condition with endTime before current block", async function () {
      const { veLike, likecoin, deployer, publicClient } =
        await loadFixture(initialMint);
      const block = await publicClient.getBlock();
      const startTime = block.timestamp + 1000n;
      const endTime = block.timestamp - 2000n;
      await expect(
        veLike.write.addReward([10000n * 10n ** 6n, startTime, endTime], {
          account: deployer.account.address,
        }),
      ).to.be.rejectedWith("ErrConflictCondition");
    });

    it("should not able to set reward condition with startTime after endTime", async function () {
      const { veLike, likecoin, deployer, publicClient } =
        await loadFixture(initialMint);
      const block = await publicClient.getBlock();
      const startTime = block.timestamp + 3000n;
      const endTime = block.timestamp + 2000n;
      await expect(
        veLike.write.addReward([10000n * 10n ** 6n, startTime, endTime], {
          account: deployer.account.address,
        }),
      ).to.be.rejectedWith("ErrConflictCondition");
    });

    it("should not able to set reward condition with startTime before current endTime", async function () {
      const { veLike, likecoin, deployer, publicClient } =
        await loadFixture(initialMint);

      await likecoin.write.approve([veLike.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      const block = await publicClient.getBlock();
      const startTime = block.timestamp;
      const endTime = block.timestamp + 1n;
      await veLike.write.addReward([10000n * 10n ** 6n, startTime, endTime], {
        account: deployer.account.address,
      });

      await likecoin.write.approve([veLike.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      await expect(
        veLike.write.addReward([10000n * 10n ** 6n, startTime, endTime], {
          account: deployer.account.address,
        }),
      ).to.be.rejectedWith("ErrConflictCondition");
    });

    it("should able to set new reward condition after current expire", async function () {
      const { veLike, likecoin, deployer, publicClient } =
        await loadFixture(initialMint);
      await likecoin.write.approve([veLike.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      const block = await publicClient.getBlock();
      const startTime = block.timestamp;
      const endTime = block.timestamp + 1n;
      await veLike.write.addReward([10000n * 10n ** 6n, startTime, endTime], {
        account: deployer.account.address,
      });

      await likecoin.write.approve([veLike.address, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });
      await veLike.write.addReward(
        [10000n * 10n ** 6n, endTime + 1n, endTime + 1000n],
        {
          account: deployer.account.address,
        },
      );

      const condition = await veLike.read.getCurrentCondition();
      expect(condition.startTime).to.equal(endTime + 1n);
      expect(condition.endTime).to.equal(endTime + 1000n);
      expect(condition.rewardAmount).to.equal(10000n * 10n ** 6n);
      expect(condition.rewardIndex).to.equal(0n);
      expect(await veLike.read.totalAssets()).to.equal(0n);
    });
  });

  describe("reward distribution", async function () {
    it("should have correct initial reward condition", async function () {
      const { veLike, bob, testClient } = await loadFixture(initialCondition);
      const condition = await veLike.read.getCurrentCondition();

      // Cross check the top function not changed accidentally.
      expect(condition.endTime).to.equal(condition.startTime + 1000n);
      expect(condition.rewardAmount).to.equal(10000n * 10n ** 6n);
      expect(condition.rewardIndex).to.equal(0n);

      // Bob deposit before the startTime, so the lastRewardTime should be the startTime
      const lastRewardTime = await veLike.read.getLastRewardTime();
      expect(lastRewardTime).to.equal(condition.startTime);

      await testClient.setNextBlockTimestamp({
        timestamp: condition.startTime + 100n,
      });
      await testClient.mine({
        blocks: 1,
      });
      const bobReward = await veLike.read.getPendingReward([
        bob.account.address,
      ]);
      expect(bobReward).to.equal(1000n * 10n ** 6n);
    });

    it("should have return zero on account never deposited", async function () {
      const { veLike, rick } = await loadFixture(initialCondition);
      expect(
        await veLike.read.getPendingReward([rick.account.address]),
      ).to.equal(0n);
    });

    it("should have update totalStaked and lastRewardTime correctly on new deposit", async function () {
      const { veLike, likecoin, deployer, rick, publicClient, testClient } =
        await loadFixture(initialCondition);
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });
      const depositTx = await veLike.write.deposit(
        [100n * 10n ** 6n, rick.account.address],
        {
          account: rick.account.address,
        },
      );
      const block = await publicClient.getBlock({
        blockHash: depositTx.blockHash,
      });
      const lastRewardTime = await veLike.read.getLastRewardTime();
      expect(block.timestamp).to.equal(lastRewardTime);
      expect(await veLike.read.totalSupply()).to.equal(200n * 10n ** 6n);
    });

    it("should automatically claim reward on new deposit", async function () {
      const {
        veLike,
        likecoin,
        deployer,
        rick,
        publicClient,
        startTime,
        testClient,
      } = await loadFixture(initialCondition);

      const block0 = await publicClient.getBlock();
      await likecoin.write.approve([veLike.address, 100n * 10n ** 6n], {
        account: rick.account.address,
      });
      const depositTx = await veLike.write.deposit(
        [100n * 10n ** 6n, rick.account.address],
        {
          account: rick.account.address,
        },
      );
      const block = await publicClient.getBlock({
        blockHash: depositTx.blockHash,
      });

      const timePassed = 3n;
      await testClient.setNextBlockTimestamp({
        timestamp: block.timestamp + timePassed,
      });
      await testClient.mine({
        blocks: 1,
      });

      expect(
        await veLike.read.getPendingReward([rick.account.address]),
      ).to.not.equal(0n);
      // Per second reward is 10000n * 10n ** 6n / 1000n = 10n * 10n ** 6n;
      // rick should have 50% in last past x seconds, i.e. 4LIKE
      expect(
        await veLike.read.getPendingReward([rick.account.address]),
      ).to.equal(timePassed * 5n * 10n ** 6n);
    });
  });
});
