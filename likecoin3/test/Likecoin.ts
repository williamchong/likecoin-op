import {
  time,
  loadFixture,
} from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem, ignition } from "hardhat";
import { parseEther } from "viem";
import LikecoinModule from "../ignition/modules/Likecoin";

describe("Likecoin", async function () {
  async function deployToken() {
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

    return {
      likecoin,
      likecoinImpl,
      likecoinProxy,
      deployer,
      rick,
      kin,
      bob,
      publicClient,
    };
  }

  describe("Basic ERC20 functionality", async function () {
    it("should have 18 decimals", async function () {
      const { likecoin } = await loadFixture(deployToken);
      expect(await likecoin.read.decimals()).to.equal(18);
    });

    it("should belong to the deployer", async function () {
      const { likecoin, deployer } = await loadFixture(deployToken);
      const owner = await likecoin.read.owner();
      expect(owner.toLowerCase()).to.equal(
        deployer.account.address.toLowerCase(),
      );
    });

    it("Should have correct name and symbol", async function () {
      const { likecoin } = await loadFixture(deployToken);
      expect(await likecoin.read.name()).to.equal("Likecoin");
      expect(await likecoin.read.symbol()).to.equal("LIKE");
    });

    it("Should mint tokens correctly", async function () {
      const { likecoin, bob } = await loadFixture(deployToken);
      const amount = parseEther("100");

      await likecoin.write.mint([bob.account.address, amount]);
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        amount,
      );
    });

    it("Should fail to mint tokens if not owner", async function () {
      const { likecoin, rick } = await loadFixture(deployToken);
      const amount = parseEther("100");
      await expect(
        likecoin.write.mint([rick.account.address, amount], {
          account: rick.account,
        }),
      ).to.be.rejectedWith("OwnableUnauthorizedAccount");
    });

    it("Should transfer tokens between accounts", async function () {
      const { likecoin, rick, bob } = await loadFixture(deployToken);
      const amount = parseEther("100");

      await likecoin.write.mint([rick.account.address, amount]);

      await likecoin.write.transfer([bob.account.address, parseEther("50")], {
        account: rick.account,
      });

      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        parseEther("50"),
      );
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        parseEther("50"),
      );
    });

    it("Should handle approvals and transferFrom", async function () {
      const { likecoin, rick, bob } = await loadFixture(deployToken);
      const amount = parseEther("100");

      await likecoin.write.mint([rick.account.address, amount]);

      await likecoin.write.approve([bob.account.address, parseEther("50")], {
        account: rick.account,
      });

      expect(
        await likecoin.read.allowance([
          rick.account.address,
          bob.account.address,
        ]),
      ).to.equal(parseEther("50"));

      await likecoin.write.transferFrom(
        [rick.account.address, bob.account.address, parseEther("30")],
        { account: bob.account },
      );

      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        parseEther("70"),
      );
      expect(await likecoin.read.balanceOf([bob.account.address])).to.equal(
        parseEther("30"),
      );
      expect(
        await likecoin.read.allowance([
          rick.account.address,
          bob.account.address,
        ]),
      ).to.equal(parseEther("20"));
    });

    it("Should burn tokens correctly", async function () {
      const { likecoin, rick } = await loadFixture(deployToken);
      const amount = parseEther("100");

      await likecoin.write.mint([rick.account.address, amount]);

      await likecoin.write.burn([parseEther("30")], {
        account: rick.account,
      });

      expect(await likecoin.read.balanceOf([rick.account.address])).to.equal(
        parseEther("70"),
      );
    });

    it("Should fail when transferring more than balance", async function () {
      const { likecoin, rick, bob } = await loadFixture(deployToken);
      const amount = parseEther("100");

      await likecoin.write.mint([rick.account.address, amount]);

      await expect(
        likecoin.write.transfer([bob.account.address, parseEther("150")], {
          account: rick.account,
        }),
      );
    });

    it("Should fail when transferring with insufficient allowance", async function () {
      const { likecoin, rick, bob } = await loadFixture(deployToken);
      const amount = parseEther("100");

      await likecoin.write.mint([rick.account.address, amount]);
      await likecoin.write.approve([bob.account.address, parseEther("20")], {
        account: rick.account,
      });

      await expect(
        likecoin.write.transferFrom(
          [rick.account.address, bob.account.address, parseEther("50")],
          { account: bob.account },
        ),
      );
    });
  });
});
