import { loadFixture } from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem, ignition } from "hardhat";
import { encodeFunctionData } from "viem";

import "./setup";
import { BookConfigLoader } from "./BookConfigLoader";
import LikeProtocolV1Module from "../ignition/modules/LikeProtocolV1";
import LikeProtocol from "../ignition/modules/LikeProtocol";

describe("LikeProtocol as create2 factory for BookNFT", () => {
  async function initLikeProtocolCreate3Factory() {
    const [deployer] = await viem.getWalletClients();
    const { likeProtocol } = await ignition.deploy(LikeProtocolV1Module, {
      parameters: {
        LikeProtocolV0Module: {
          initOwner: deployer.account.address,
        },
      },
      defaultSender: deployer.account.address,
    });

    return { likeProtocol, deployer };
  }

  it("should be able to create new BookNFT", async function () {
    const { likeProtocol, deployer } = await loadFixture(
      initLikeProtocolCreate3Factory,
    );
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    await likeProtocol.write.newBookNFT([
      {
        creator: deployer.account.address,
        updaters: [deployer.account.address],
        minters: [deployer.account.address],
        config: bookConfig,
      },
    ]);
    const logs = await likeProtocol.getEvents.NewBookNFT();
    expect(logs.length).to.equal(1);
    const bookNFTAddress = logs[0].args.bookNFT;
    expect(await likeProtocol.read.isBookNFT([bookNFTAddress])).to.be.true;
  });

  it("should be able to precompute address", async function () {
    const { likeProtocol, deployer } = await loadFixture(
      initLikeProtocolCreate3Factory,
    );
    // Any upgrade to openzeppelin would cause this test to fail, event comments updates
    const salt =
      "0x0000000000000000000000000000000000000000000000000000000000000000";
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    const deployerAddress = deployer.account.address;
    expect(deployerAddress).to.equalAddress(
      "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
    );
    expect(likeProtocol.address).to.equalAddress(
      "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9",
    );
    const msgNewBookNFT = {
      creator: deployer.account.address,
      updaters: [deployer.account.address],
      minters: [deployer.account.address],
      config: bookConfig,
    };
    const targetAddress = await likeProtocol.read.precomputeAddress([
      salt,
      msgNewBookNFT,
    ]);
    expect(targetAddress).to.equal(
      "0x49076914E89608B826b7dF00ca0ECe5520A84019",
    );
    const bookConfig1 = BookConfigLoader.load(
      "./test/fixtures/BookConfig1.json",
    );
    const msgNewBookNFT1 = {
      creator: deployer.account.address,
      updaters: [deployer.account.address],
      minters: [deployer.account.address],
      config: bookConfig1,
    };
    const targetAddress1 = await likeProtocol.read.precomputeAddress([
      salt,
      msgNewBookNFT1,
    ]);
    expect(targetAddress1).to.equalAddress(
      "0x69FE487677C8B98E5e0a485C60C55172aaF0e031",
    );
  });

  it("should be able to create new BookNFT with precomputed address", async function () {
    const { likeProtocol, deployer } = await loadFixture(
      initLikeProtocolCreate3Factory,
    );
    const salt = ("0x" +
      deployer.account.address.slice(2) +
      "0".repeat(24)) as `0x${string}`;
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    const msgNewBookNFT = {
      creator: deployer.account.address,
      updaters: [deployer.account.address],
      minters: [deployer.account.address],
      config: bookConfig,
    };
    const precomputedAddress = await likeProtocol.read.precomputeAddress([
      salt,
      msgNewBookNFT,
    ]);

    await likeProtocol.write.create2BookNFT([salt, msgNewBookNFT], {
      account: deployer.account,
    });
    const logs = await likeProtocol.getEvents.NewBookNFT();
    expect(logs.length).to.equal(1);
    const actualAddress = logs[0].args.bookNFT;
    expect(await likeProtocol.read.isBookNFT([actualAddress])).to.be.true;

    const book0NFT = await viem.getContractAt("BookNFT", actualAddress);
    expect(await book0NFT.read.name()).to.equal("My Book");
    expect(await book0NFT.read.symbol()).to.equal("KOOB");
    const beacon = await book0NFT.read.getProtocolBeacon();
    expect(beacon.toLowerCase()).to.equal(likeProtocol.address.toLowerCase());
    expect(await book0NFT.read.contractURI()).to.equal(
      "data:application/json;base64,eyJuYW1lIjoiQ29sbGVjdGlvbiBOYW1lIiwic3ltYm9sIjoiQ29sbGVjdGlvbiBTWU1CIiwiZGVzY3JpcHRpb24iOiJDb2xsZWN0aW9uIERlc2NyaXB0aW9uIiwiaW1hZ2UiOiJpcGZzOi8vYmFmeWJlaWV6cTR5cW9zYzJ1NHNhYW5vdmU1YnNhM3ljaXVmd2hmZHVlbXk1ejZ2dmY2cTNjNWxuYmkiLCJiYW5uZXJfaW1hZ2UiOiIiLCJmZWF0dXJlZF9pbWFnZSI6IiIsImV4dGVybmFsX2xpbmsiOiJodHRwczovL3d3dy5leGFtcGxlLmNvbSIsImNvbGxhYm9yYXRvcnMiOltdfQ==",
    );
    const royalty = await book0NFT.read.royaltyInfo([0n, 1000n]);
    expect(royalty[0].toLowerCase()).to.equal(
      deployer.account.address.toLowerCase(),
    );
    expect(royalty[1]).to.equal(0n);

    expect(actualAddress).to.equalAddress(precomputedAddress);
  });

  it("should be not able to create same BookNFT with same salt", async function () {
    const { likeProtocol, deployer } = await loadFixture(
      initLikeProtocolCreate3Factory,
    );
    const salt = ("0x" +
      deployer.account.address.slice(2) +
      "0".repeat(24)) as `0x${string}`;
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    const msgNewBookNFT = {
      creator: deployer.account.address,
      updaters: [deployer.account.address],
      minters: [deployer.account.address],
      config: bookConfig,
    };
    const precomputedAddress = await likeProtocol.read.precomputeAddress([
      salt,
      msgNewBookNFT,
    ]);

    await likeProtocol.write.create2BookNFT([salt, msgNewBookNFT], {
      account: deployer.account,
    });
    const logs = await likeProtocol.getEvents.NewBookNFT();
    expect(logs.length).to.equal(1);
    const actualAddress = logs[0].args.bookNFT;
    expect(await likeProtocol.read.isBookNFT([actualAddress])).to.be.true;

    await expect(
      likeProtocol.write.create2BookNFT([salt, msgNewBookNFT], {
        account: deployer.account,
      }),
    ).to.be.rejectedWith("FailedDeployment()");
  });

  it("should not alter computed address after BookNFT upgrade", async function () {
    const { likeProtocol, deployer } = await loadFixture(
      initLikeProtocolCreate3Factory,
    );
    const salt = ("0x" +
      deployer.account.address.slice(2) +
      "0".repeat(24)) as `0x${string}`;
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    const msgNewBookNFT = {
      creator: deployer.account.address,
      updaters: [deployer.account.address],
      minters: [deployer.account.address],
      config: bookConfig,
    };
    const precomputedAddress = await likeProtocol.read.precomputeAddress([
      salt,
      msgNewBookNFT,
    ]);

    const bookNFTMock = await viem.deployContract("BookNFTMock");
    await likeProtocol.write.upgradeTo([bookNFTMock.address], {
      account: deployer.account,
    });
    const newAddress = await likeProtocol.read.precomputeAddress([
      salt,
      msgNewBookNFT,
    ]);
    expect(newAddress).to.equal(precomputedAddress);
  });

  it("should not alter computed address after LikeProtocol upgrade", async function () {
    const { likeProtocol, deployer } = await loadFixture(
      initLikeProtocolCreate3Factory,
    );
    const salt = ("0x" +
      deployer.account.address.slice(2) +
      "0".repeat(24)) as `0x${string}`;
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    const msgNewBookNFT = {
      creator: deployer.account.address,
      updaters: [deployer.account.address],
      minters: [deployer.account.address],
      config: bookConfig,
    };
    const precomputedAddress = await likeProtocol.read.precomputeAddress([
      salt,
      msgNewBookNFT,
    ]);

    const likeProtocolMock = await viem.deployContract("LikeProtocolMock");
    await likeProtocol.write.upgradeToAndCall(
      [likeProtocolMock.address, "0x"],
      {
        account: deployer.account,
      },
    );
    const newAddress = await likeProtocol.read.precomputeAddress([
      salt,
      msgNewBookNFT,
    ]);
    expect(newAddress).to.equal(precomputedAddress);
  });

  it("should not alter computed address after LikeProtocol upgrade with BookNFT upgrade", async function () {
    const { likeProtocol, deployer } = await loadFixture(
      initLikeProtocolCreate3Factory,
    );
    const salt = ("0x" +
      deployer.account.address.slice(2) +
      "0".repeat(24)) as `0x${string}`;
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    const msgNewBookNFT = {
      creator: deployer.account.address,
      updaters: [deployer.account.address],
      minters: [deployer.account.address],
      config: bookConfig,
    };
    const precomputedAddress = await likeProtocol.read.precomputeAddress([
      salt,
      msgNewBookNFT,
    ]);

    const likeProtocolMock = await viem.deployContract("LikeProtocolMock");
    const bookNFTMock = await viem.deployContract("BookNFTMock");
    const upgradeToAndCallData = await encodeFunctionData({
      abi: likeProtocol.abi,
      functionName: "upgradeTo",
      args: [bookNFTMock.address],
    });
    await likeProtocol.write.upgradeToAndCall(
      [likeProtocolMock.address, upgradeToAndCallData],
      {
        account: deployer.account,
      },
    );
    const newAddress = await likeProtocol.read.precomputeAddress([
      salt,
      msgNewBookNFT,
    ]);
    expect(newAddress).to.equal(precomputedAddress);
  });
});
