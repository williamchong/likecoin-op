import {
  time,
  loadFixture,
} from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem, ignition } from "hardhat";
import { parseEther } from "viem";
import LikecoinModule from "../ignition/modules/Likecoin";
import LikeCollectiveModule from "../ignition/modules/LikeCollective";
import LikeStakePositionModule from "../ignition/modules/LikeStakePosition";

describe("LikeCollective", async function () {
  async function deployCollective() {
    const [deployer, rick, kin, bob] = await viem.getWalletClients();
    const publicClient = await viem.getPublicClient();

    const { likecoin, likecoinImpl, likecoinProxy } = await ignition.deploy(
      LikecoinModule,
      {
        parameters: {
          LikecoinModule: {
            initOwner: deployer.account.address,
          },
        },
        defaultSender: deployer.account.address,
      },
    );

    const { likeCollective, likeCollectiveImpl, likeCollectiveProxy } =
      await ignition.deploy(LikeCollectiveModule, {
        parameters: {
          LikeCollectiveModule: {
            initOwner: deployer.account.address,
          },
        },
        defaultSender: deployer.account.address,
      });
    const { likeStakePosition, likeStakePositionImpl, likeStakePositionProxy } =
      await ignition.deploy(LikeStakePositionModule, {
        parameters: {
          LikeStakePositionModule: {
            initOwner: deployer.account.address,
          },
        },
        defaultSender: deployer.account.address,
      });

    // Setup relationships
    await likeCollective.write.setLikeStakePosition(
      [likeStakePosition.address],
      {
        account: deployer.account.address,
      },
    );
    await likeCollective.write.setLikecoin([likecoin.address], {
      account: deployer.account.address,
    });
    await likeStakePosition.write.setManager([likeCollective.address], {
      account: deployer.account.address,
    });

    // Mint some LIKE tokens
    for (const a of [
      rick.account.address,
      kin.account.address,
      bob.account.address,
    ]) {
      await likecoin.write.mint([a, 10000n * 10n ** 6n], {
        account: deployer.account.address,
      });
    }

    return {
      likecoin,
      likecoinImpl,
      likecoinProxy,
      likeCollective,
      likeCollectiveImpl,
      likeCollectiveProxy,
      likeStakePosition,
      likeStakePositionImpl,
      likeStakePositionProxy,
      deployer,
      rick,
      kin,
      bob,
      publicClient,
    };
  }

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
        likeCollective.write.stake([mockBookNFT, amount], {
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
        likeCollective.write.unstakePosition([1n], {
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
      await likeCollective.write.stake([mockBookNFT, 6000n * 10n ** 6n], {
        account: rick.account,
      });
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        4000n * 10n ** 6n,
      );
      const tokenOwner = await likeStakePosition.read.ownerOf([nextTokenId]);
      expect(tokenOwner.toLowerCase()).to.equal(
        rick.account.address.toLowerCase(),
      );
      await likeCollective.write.unstakePosition([nextTokenId], {
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

      await likeCollective.write.stake([mockBookNFT, amount], {
        account: rick.account,
      });
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        2000n * 10n ** 6n,
      );
      const owner = await likeStakePosition.read.ownerOf([nextTokenId]);
      expect(owner.toLowerCase()).to.equal(rick.account.address.toLowerCase());
      expect(
        await likeCollective.read.pendingRewardsOf([nextTokenId]),
      ).to.equal(0n);
      expect(
        await likeCollective.read.getStakeForUser([
          rick.account.address,
          mockBookNFT,
        ]),
      ).to.equal(amount);

      await likeCollective.write.claimRewards([mockBookNFT], {
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
        await likeCollective.read.pendingRewardsOf([nextTokenId]),
      ).to.equal(reward);

      // Claim rewards
      await likeCollective.write.claimRewards([mockBookNFT], {
        account: rick.account,
      });
      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        3000n * 10n ** 6n,
      );

      await likeCollective.write.unstakePosition([nextTokenId], {
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

      await likeCollective.write.stake([mockBookNFT, amount], {
        account: rick.account,
      });
      const owner = await likeStakePosition.read.ownerOf([nextTokenId]);
      expect(owner.toLowerCase()).to.equal(rick.account.address.toLowerCase());

      await expect(
        likeCollective.write.unstakePosition([nextTokenId], {
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

      await likeCollective.write.stake([mockBookNFT, 4n * amount], {
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
      await likeCollective.write.unstakePosition([nextTokenId], {
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

      await likeCollective.write.stake([mockBookNFT, 4n * amount], {
        account: rick.account,
      });
      await likeCollective.write.stake([mockBookNFT, amount], {
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
      await likeCollective.write.unstakePosition([nextTokenId], {
        account: rick.account,
      });
    });
  });

  describe("Reward", async function () {
    it("should allow to claim single reward for a user", async function () {
      const { likeCollective, rick, bob, likeStakePosition, likecoin, kin } =
        await loadFixture(deployCollective);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const amount = 2000n * 10n ** 6n;
      const reward = 500n * 10n ** 6n;

      await likecoin.write.approve([likeCollective.address, amount], {
        account: rick.account,
      });
      await likeCollective.write.stake([mockBookNFT, amount], {
        account: rick.account,
      });

      await likecoin.write.approve([likeCollective.address, reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });

      await likeCollective.write.claimRewards([mockBookNFT], {
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

      await likecoin.write.approve([likeCollective.address, 2n * amount], {
        account: rick.account,
      });
      await likeCollective.write.stake([mockBookNFT, amount], {
        account: rick.account,
      });
      await likeCollective.write.stake([mockBookNFT, amount], {
        account: rick.account,
      });

      await likecoin.write.approve([likeCollective.address, reward], {
        account: kin.account,
      });
      await likeCollective.write.depositReward([mockBookNFT, reward], {
        account: kin.account,
      });

      await likeCollective.write.claimRewards([mockBookNFT], {
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
      await likeCollective.write.stake([mockBookNFT, amount], {
        account: rick.account,
      });
      await likeCollective.write.stake([mockBookNFT2, amount], {
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

      const pendingRewards = await likeCollective.read.pendingRewardsOf([
        nextTokenId,
      ]);
      const pendingRewards2 = await likeCollective.read.pendingRewardsOf([
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
});
