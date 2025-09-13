import { loadFixture } from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem, ignition } from "hardhat";
import LikeStakePositionModule from "../ignition/modules/LikeStakePosition";

describe("LikeStakePosition", async function () {
  async function deployLSP() {
    const [deployer, rick, kin] = await viem.getWalletClients();
    const publicClient = await viem.getPublicClient();

    const { likeStakePosition, likeStakePositionImpl, likeStakePositionProxy } =
      await ignition.deploy(LikeStakePositionModule, {
        parameters: {
          LikeStakePositionModule: {
            initOwner: deployer.account.address,
          },
        },
        defaultSender: deployer.account.address,
      });

    return {
      likeStakePosition,
      likeStakePositionImpl,
      likeStakePositionProxy,
      deployer,
      rick,
      kin,
      publicClient,
    };
  }

  describe("Initialization & Ownership", async function () {
    it("should initialize with correct owner", async function () {
      const { likeStakePosition, deployer } = await loadFixture(deployLSP);
      const owner = await likeStakePosition.read.owner();
      expect(owner.toLowerCase()).to.equal(
        deployer.account.address.toLowerCase(),
      );
    });
  });

  describe("Pause controls", async function () {
    it("owner can pause and unpause", async function () {
      const { likeStakePosition } = await loadFixture(deployLSP);
      expect(await likeStakePosition.read.paused()).to.equal(false);
      await likeStakePosition.write.pause();
      expect(await likeStakePosition.read.paused()).to.equal(true);
      await likeStakePosition.write.unpause();
      expect(await likeStakePosition.read.paused()).to.equal(false);
    });

    it("reverts manager ops when paused", async function () {
      const { likeStakePosition, deployer, rick, kin } =
        await loadFixture(deployLSP);
      // Normally the manager is set to LikeCollective, but any address should be same
      await likeStakePosition.write.setManager([rick.account.address], {
        account: deployer.account,
      });
      await likeStakePosition.write.pause({ account: deployer.account });

      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      await expect(
        likeStakePosition.write.mintPosition(
          [kin.account.address, mockBookNFT, 100n, 0n],
          { account: rick.account },
        ),
      ).to.be.rejectedWith("EnforcedPause");
    });
  });

  describe("Manager controls", async function () {
    it("owner can set manager and emits event", async function () {
      const { likeStakePosition, deployer, publicClient } =
        await loadFixture(deployLSP);
      const mockManagerAddress = "0x1234567890123456789012345678901234567890";
      const hash = await likeStakePosition.write.setManager(
        [mockManagerAddress],
        {
          account: deployer.account,
        },
      );
      const receipt = await publicClient.waitForTransactionReceipt({ hash });
      const logs = await publicClient.getContractEvents({
        address: likeStakePosition.address,
        abi: likeStakePosition.abi,
        eventName: "ManagerUpdated",
        fromBlock: receipt.blockNumber,
        toBlock: receipt.blockNumber,
      });
      expect(logs.length).to.be.greaterThan(0);
      expect((logs[0]!.args as any).manager.toLowerCase()).to.equal(
        mockManagerAddress.toLowerCase(),
      );
      const manager = await likeStakePosition.read.manager();
      expect(manager.toLowerCase()).to.equal(mockManagerAddress.toLowerCase());
    });

    it("non-manager cannot mint/update/burn", async function () {
      const { likeStakePosition, deployer, rick } =
        await loadFixture(deployLSP);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";

      // set manager to deployer
      await likeStakePosition.write.setManager([deployer.account.address], {
        account: deployer.account,
      });

      // non-manager tries to mint
      await expect(
        likeStakePosition.write.mintPosition(
          [rick.account.address, mockBookNFT, 100n, 0n],
          { account: rick.account },
        ),
      ).to.be.rejectedWith("ErrNotManager");

      // manager mints so that we have a token to update/burn
      await likeStakePosition.write.mintPosition(
        [rick.account.address, mockBookNFT, 100n, 0n],
        { account: deployer.account },
      );

      // non-manager tries to update
      await expect(
        likeStakePosition.write.updatePosition([1n, 200n, 1n], {
          account: rick.account,
        }),
      ).to.be.rejectedWith("ErrNotManager");

      // non-manager tries to burn
      await expect(
        likeStakePosition.write.burnPosition([1n], { account: rick.account }),
      ).to.be.rejectedWith("ErrNotManager");

      // manager burns
      await likeStakePosition.write.burnPosition([1n], {
        account: deployer.account,
      });
    });
  });

  describe("Mint / Update / Burn flows", async function () {
    it("manager can mint, update and burn positions", async function () {
      const { likeStakePosition, deployer, rick } =
        await loadFixture(deployLSP);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";

      await likeStakePosition.write.setManager([deployer.account.address], {
        account: deployer.account,
      });

      // Mint
      await likeStakePosition.write.mintPosition(
        [rick.account.address, mockBookNFT, 42n, 0n],
        { account: deployer.account },
      );

      // tokenId should be 1 for first mint; verify owner and position data
      const tokenOwner = await likeStakePosition.read.ownerOf([1n]);
      expect(tokenOwner.toLowerCase()).to.equal(
        rick.account.address.toLowerCase(),
      );

      const pos1 = await likeStakePosition.read.getPosition([1n]);
      expect(pos1.bookNFT.toLowerCase()).to.equal(mockBookNFT.toLowerCase());
      expect(pos1.stakedAmount).to.equal(42n);
      expect(pos1.rewardIndex).to.equal(0n);

      // Update
      await likeStakePosition.write.updatePosition([1n, 84n, 3n], {
        account: deployer.account,
      });
      const pos2 = await likeStakePosition.read.getPosition([1n]);
      expect(pos2.stakedAmount).to.equal(84n);
      expect(pos2.rewardIndex).to.equal(3n);

      // Burn
      await likeStakePosition.write.burnPosition([1n], {
        account: deployer.account,
      });
      await expect(likeStakePosition.read.ownerOf([1n])).to.be.rejected; // token no longer exists
    });
  });

  describe("Token URI behavior", async function () {
    it("returns baseURI + tokenId when tokenURI not set", async function () {
      const { likeStakePosition, deployer, rick } =
        await loadFixture(deployLSP);
      const mockBookNFT = "0xe3DB5cD985d869a18C000Cb4D51180B1Cb450e8C";

      await likeStakePosition.write.setManager([deployer.account.address], {
        account: deployer.account,
      });
      await likeStakePosition.write.setBaseURI(["https://3ook.com/lspcover/"], {
        account: deployer.account,
      });

      await likeStakePosition.write.mintPosition(
        [rick.account.address, mockBookNFT, 7n, 0n],
        { account: deployer.account },
      );

      const uri = await likeStakePosition.read.tokenURI([1n]);

      expect(uri).to.equal(
        "https://3ook.com/lspcover/49tc2YXYaaGMAAy01RGAsctFDowAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAcJl5cMUYEtw6AQx9AbUODRfcecg=",
      );
    });
  });
});
