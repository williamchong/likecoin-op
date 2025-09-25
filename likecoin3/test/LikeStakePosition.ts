import { loadFixture } from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem, ignition } from "hardhat";
import LikeStakePositionModule from "../ignition/modules/LikeStakePosition";

describe("LikeStakePosition", async function () {
  async function deployLSP() {
    const [deployer, rick, kin] = await viem.getWalletClients();
    const publicClient = await viem.getPublicClient();

    const { likeStakePosition, likeStakePositionImpl } = await ignition.deploy(
      LikeStakePositionModule,
      {
        parameters: {
          LikecoinModule: {
            initOwner: deployer.account.address,
          },
          LikeCollectiveV0Module: {
            initOwner: deployer.account.address,
          },
          LikeStakePositionV0Module: {
            initOwner: deployer.account.address,
          },
        },
        defaultSender: deployer.account.address,
      },
    );

    return {
      likeStakePosition,
      likeStakePositionImpl,
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

      const nextTokenId = await likeStakePosition.read.getNextTokenId();
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
        likeStakePosition.write.updatePosition([1n, 200n, nextTokenId], {
          account: rick.account,
        }),
      ).to.be.rejectedWith("ErrNotManager");

      // non-manager tries to burn
      await expect(
        likeStakePosition.write.burnPosition([1n], { account: rick.account }),
      ).to.be.rejectedWith("ErrNotManager");

      // manager burns
      await likeStakePosition.write.burnPosition([nextTokenId], {
        account: deployer.account,
      });
    });
  });

  describe("Mint / Update / Burn flows", async function () {
    it("manager can mint, update and burn positions", async function () {
      const { likeStakePosition, deployer, rick } =
        await loadFixture(deployLSP);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const nextTokenId = await likeStakePosition.read.getNextTokenId();

      await likeStakePosition.write.setManager([deployer.account.address], {
        account: deployer.account,
      });

      // Mint
      await likeStakePosition.write.mintPosition(
        [rick.account.address, mockBookNFT, 42n, 0n],
        { account: deployer.account },
      );

      // tokenId should be 0 for first mint; verify owner and position data
      const tokenOwner = await likeStakePosition.read.ownerOf([nextTokenId]);
      expect(tokenOwner.toLowerCase()).to.equal(
        rick.account.address.toLowerCase(),
      );

      const pos1 = await likeStakePosition.read.getPosition([nextTokenId]);
      expect(pos1.bookNFT.toLowerCase()).to.equal(mockBookNFT.toLowerCase());
      expect(pos1.stakedAmount).to.equal(42n);
      expect(pos1.rewardIndex).to.equal(0n);

      // Update
      await likeStakePosition.write.updatePosition([nextTokenId, 84n, 3n], {
        account: deployer.account,
      });
      const pos2 = await likeStakePosition.read.getPosition([nextTokenId]);
      expect(pos2.stakedAmount).to.equal(84n);
      expect(pos2.rewardIndex).to.equal(3n);

      // Burn
      await likeStakePosition.write.burnPosition([nextTokenId], {
        account: deployer.account,
      });
      await expect(likeStakePosition.read.ownerOf([nextTokenId])).to.be
        .rejected; // token no longer exists
    });

    it("should track user positions correctly", async function () {
      const { likeStakePosition, deployer, rick } =
        await loadFixture(deployLSP);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";
      const nextTokenId = await likeStakePosition.read.getNextTokenId();
      await likeStakePosition.write.setManager([deployer.account.address], {
        account: deployer.account,
      });

      await likeStakePosition.write.mintPosition(
        [rick.account.address, mockBookNFT, 42n, 0n],
        { account: deployer.account },
      );
      await likeStakePosition.write.mintPosition(
        [rick.account.address, mockBookNFT, 42n, 0n],
        { account: deployer.account },
      );
      await likeStakePosition.write.mintPosition(
        [rick.account.address, mockBookNFT, 42n, 0n],
        { account: deployer.account },
      );

      const balanceOf = await likeStakePosition.read.balanceOf([
        rick.account.address,
      ]);
      expect(balanceOf).to.equal(3n);
      const firsttokenId = await likeStakePosition.read.tokenOfOwnerByIndex([
        rick.account.address,
        0n,
      ]);
      expect(firsttokenId).to.equal(0n);
      const secondtokenId = await likeStakePosition.read.tokenOfOwnerByIndex([
        rick.account.address,
        1n,
      ]);
      expect(secondtokenId).to.equal(1n);
      const thirdtokenId = await likeStakePosition.read.tokenOfOwnerByIndex([
        rick.account.address,
        2n,
      ]);
      expect(thirdtokenId).to.equal(2n);

      const positions = await likeStakePosition.read.getUserPositions([
        rick.account.address,
      ]);
      expect(positions.length).to.equal(3);
      expect(positions[0]).to.equal(0n);
      expect(positions[1]).to.equal(1n);
      expect(positions[2]).to.equal(2n);
    });

    it("should track user positions correct even after transfer", async function () {
      const { likeStakePosition, deployer, rick, kin } =
        await loadFixture(deployLSP);
      const mockBookNFT = "0x1234567890123456789012345678901234567890";

      await likeStakePosition.write.setManager([deployer.account.address], {
        account: deployer.account,
      });
      await likeStakePosition.write.mintPosition(
        [rick.account.address, mockBookNFT, 42n, 0n],
        { account: deployer.account },
      );
      await likeStakePosition.write.mintPosition(
        [rick.account.address, mockBookNFT, 42n, 0n],
        { account: deployer.account },
      );
      await likeStakePosition.write.mintPosition(
        [rick.account.address, mockBookNFT, 42n, 0n],
        { account: deployer.account },
      );

      expect(
        await likeStakePosition.read.balanceOf([rick.account.address]),
      ).to.equal(3n);
      await likeStakePosition.write.transferFrom(
        [rick.account.address, kin.account.address, 1n],
        { account: rick.account },
      );
      expect(
        await likeStakePosition.read.balanceOf([rick.account.address]),
      ).to.equal(2n);
      expect(
        await likeStakePosition.read.balanceOf([kin.account.address]),
      ).to.equal(1n);

      const positions = await likeStakePosition.read.getUserPositions([
        rick.account.address,
      ]);
      expect(positions.length).to.equal(2);

      const positionsBob = await likeStakePosition.read.getUserPositions([
        kin.account.address,
      ]);
      expect(positionsBob.length).to.equal(1);
      expect(positionsBob[0]).to.equal(1n);
    });
  });

  describe("position info view functions", async function () {
    async function preparePositionInfoData() {
      const { likeStakePosition, deployer, rick, kin } = await deployLSP();
      const bookNFT1 = "0x1111111111111111111111111111111111111111";
      const bookNFT2 = "0x2222222222222222222222222222222222222222";

      await likeStakePosition.write.setManager([deployer.account.address], {
        account: deployer.account,
      });

      const mintArray = [
        {
          bookNFT: bookNFT1,
          stakedAmount: 10n,
          rewardIndex: 1n,
          initialStaker: rick,
        },
        {
          bookNFT: bookNFT1,
          stakedAmount: 20n,
          rewardIndex: 2n,
          initialStaker: rick,
        },
        {
          bookNFT: bookNFT2,
          stakedAmount: 30n,
          rewardIndex: 3n,
          initialStaker: rick,
        },
        {
          bookNFT: bookNFT1,
          stakedAmount: 40n,
          rewardIndex: 4n,
          initialStaker: kin,
        },
      ];

      for (const info of mintArray) {
        await likeStakePosition.write.mintPosition(
          [
            info.initialStaker.account.address,
            info.bookNFT,
            info.stakedAmount,
            info.rewardIndex,
          ],
          { account: deployer.account },
        );
      }

      return {
        likeStakePosition,
        deployer,
        rick,
        kin,
        mintArray,
        bookNFT1,
        bookNFT2,
      };
    }

    it("should return correct position info", async function () {
      const { likeStakePosition, deployer, rick, kin, mintArray } =
        await loadFixture(preparePositionInfoData);
      const infoArray = [];
      for (let i = 0; i < mintArray.length; i++) {
        infoArray.push(await likeStakePosition.read.positionInfo([i]));
      }
      for (let i = 0; i < infoArray.length; i++) {
        expect(
          infoArray[i].initialStaker.toLowerCase(),
          "initialStaker is same",
        ).to.equal(mintArray[i].initialStaker.account.address.toLowerCase());
        expect(infoArray[i].bookNFT, "bookNFT is same").to.equal(
          mintArray[i].bookNFT,
        );
        expect(infoArray[i].stakedAmount, "stakedAmount is same").to.equal(
          mintArray[i].stakedAmount,
        );
        expect(infoArray[i].rewardIndex, "rewardIndex is same").to.equal(
          mintArray[i].rewardIndex,
        );
      }
    });

    it("should return correct position info by user", async function () {
      const { likeStakePosition, rick, bookNFT1, bookNFT2 } = await loadFixture(
        preparePositionInfoData,
      );
      const positionArray = await likeStakePosition.read.getUserPositions([
        rick.account.address,
      ]);
      expect(positionArray.length, "Rick have 3 positions").to.equal(3);
      const pos1 = await likeStakePosition.read.getPosition([positionArray[0]]);
      const pos2 = await likeStakePosition.read.getPosition([positionArray[1]]);
      const pos3 = await likeStakePosition.read.getPosition([positionArray[2]]);
      expect(pos1.bookNFT).to.equal(bookNFT1);
      expect(pos2.bookNFT).to.equal(bookNFT1);
      expect(pos3.bookNFT).to.equal(bookNFT2);
    });

    it("should return correct position info by user and bookNFT", async function () {
      const { likeStakePosition, rick, bookNFT1, bookNFT2 } = await loadFixture(
        preparePositionInfoData,
      );
      const positionArray =
        await likeStakePosition.read.getUserPositionByBookNFT([
          rick.account.address,
          bookNFT1,
        ]);
      expect(
        positionArray.length,
        "Rick have 2 positions on BookNFT1",
      ).to.equal(2);
      const pos1 = await likeStakePosition.read.getPosition([positionArray[0]]);
      const pos2 = await likeStakePosition.read.getPosition([positionArray[1]]);
      expect(pos1.bookNFT).to.equal(bookNFT1);
      expect(pos2.bookNFT).to.equal(bookNFT1);

      const positionArrayByBookNFT2 =
        await likeStakePosition.read.getUserPositionByBookNFT([
          rick.account.address,
          bookNFT2,
        ]);
      expect(
        positionArrayByBookNFT2.length,
        "Rick have 1 positions on BookNFT2",
      ).to.equal(1);
      const pos3 = await likeStakePosition.read.getPosition([
        positionArrayByBookNFT2[0],
      ]);
      expect(pos3.bookNFT).to.equal(bookNFT2);
    });

    it("should return correct position after ERC721 transfer", async function () {
      const { likeStakePosition, rick, kin, bookNFT1, bookNFT2 } =
        await loadFixture(preparePositionInfoData);
      await likeStakePosition.write.transferFrom(
        [rick.account.address, kin.account.address, 1n],
        { account: rick.account },
      );
      const positionArray = await likeStakePosition.read.getUserPositions([
        rick.account.address,
      ]);
      expect(
        positionArray.length,
        "Rick have 2 positions after transfer",
      ).to.equal(2);
      const pos1 = await likeStakePosition.read.getPosition([positionArray[0]]);
      const pos2 = await likeStakePosition.read.getPosition([positionArray[1]]);
      expect(pos1.bookNFT).to.equal(bookNFT1);
      expect(pos2.bookNFT).to.equal(bookNFT2);
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
        [rick.account.address, mockBookNFT, 14n, 0n],
        { account: deployer.account },
      );
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
