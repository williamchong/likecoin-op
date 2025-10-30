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
