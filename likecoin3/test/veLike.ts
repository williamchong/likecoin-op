import {
  time,
  loadFixture,
} from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem, ignition } from "hardhat";

import "./setup";
import veLikeModule from "../ignition/modules/veLike";

describe("veLike", async function () {
  async function deployVeLike() {
    const [deployer, rick, kin, bob] = await viem.getWalletClients();
    const publicClient = await viem.getPublicClient();
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
    };
  }

  it("should have the correct owner", async function () {
    const { veLike, deployer } = await loadFixture(deployVeLike);
    expect(await veLike.read.owner()).to.equalAddress(deployer.account.address);
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
    expect(await newVeLike.read.dataStorage()).to.equal(
      "0xb9e14b2a89d227541697d62a06ecbf5ccc9ad849800745b40b2826662a177600",
    );
  });
});
