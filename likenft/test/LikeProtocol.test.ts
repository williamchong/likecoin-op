import { expect } from "chai";
import { BaseContract, EventLog } from "ethers";
import { ethers, upgrades } from "hardhat";

describe("LikeProtocol", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    const [ownerSigner, randomSigner] = await ethers.getSigners();

    this.ownerSigner = ownerSigner;
    this.randomSigner = randomSigner;
  });

  let deployment: BaseContract;
  let contractAddress: string;
  let contract: any;
  beforeEach(async function () {
    const likeProtocol = await upgrades.deployProxy(
      this.LikeProtocol,
      [this.ownerSigner.address],
      {
        initializer: "initialize",
      },
    );
    deployment = await likeProtocol.waitForDeployment();
    contractAddress = await deployment.getAddress();
    contract = await ethers.getContractAt("LikeProtocol", contractAddress);
  });

  it("should set the right owner", async function () {
    expect(await contract.owner()).to.equal(this.ownerSigner.address);
  });

  it("should allow ownership transfer", async function () {
    await contract.transferOwnership(this.randomSigner.address);
    expect(await contract.owner()).to.equal(this.randomSigner.address);
  });

  it("should not allow random ownership transfer", async function () {
    const likeProtocolSigner = contract.connect(this.randomSigner);
    try {
      await likeProtocolSigner.transferOwnership(this.randomSigner.address);
    } catch (error) {
      expect(error).to.be.instanceOf(Error);
    }
    expect(await contract.owner()).to.equal(this.ownerSigner.address);
  });

  it("should not paused by random address", async function () {
    const likeProtocolSigner = contract.connect(this.randomSigner);
    await expect(likeProtocolSigner.pause()).to.be.rejected;
  });

  it("should be paused by owner address", async function () {
    const likeProtocolSigner = contract.connect(this.ownerSigner);
    await expect(likeProtocolSigner.pause()).to.be.not.rejected;
  });

  it("should not unpaused by random address", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);
    await expect(likeProtocolOwnerSigner.pause()).to.be.not.rejected;
    const likeProtocolRandomSigner = contract.connect(this.randomSigner);
    await expect(likeProtocolRandomSigner.unpause()).to.be.rejected;
  });

  it("should not operatable when paused", async function () {
    const likeProtocolSigner = contract.connect(this.ownerSigner);

    const classOperation = async () => {
      await likeProtocolSigner
        .newClass({
          creator: this.ownerSigner,
          input: {
            name: "My Book",
            symbol: "KOOB",
            metadata: JSON.stringify({
              name: "Collection Name",
              symbol: "Collection SYMB",
              description: "Collection Description",
              image:
                "ipfs://bafybeiezq4yqosc2u4saanove5bsa3yciufwhfduemy5z6vvf6q3c5lnbi",
              banner_image: "",
              featured_image: "",
              external_link: "https://www.example.com",
              collaborators: [],
            }),
            config: {
              max_supply: 10,
            },
          },
        })
        .then((tx) => tx.wait());
    };

    await expect(classOperation()).to.be.not.rejected;
    await expect(likeProtocolSigner.pause()).to.be.not.rejected;
    await expect(classOperation()).to.be.rejectedWith(
      "VM Exception while processing transaction: reverted with custom error 'EnforcedPause()'",
    );
    await expect(likeProtocolSigner.unpause()).to.be.not.rejected;
    await expect(classOperation()).to.be.not.rejected;
  });

  it("should be able to create new class", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);

    const newClass = async () => {
      await likeProtocolOwnerSigner
        .newClass({
          creator: this.ownerSigner,
          input: {
            name: "My Book",
            symbol: "KOOB",
            metadata: JSON.stringify({
              name: "Collection Name",
              symbol: "Collection SYMB",
              description: "Collection Description",
              image:
                "ipfs://bafybeiezq4yqosc2u4saanove5bsa3yciufwhfduemy5z6vvf6q3c5lnbi",
              banner_image: "",
              featured_image: "",
              external_link: "https://www.example.com",
              collaborators: [],
            }),
            config: {
              max_supply: 10,
            },
          },
        })
        .then((tx) => tx.wait());
    };

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeProtocolOwnerSigner.on("NewClass", (id, params, event) => {
        event.removeListener();
        resolve({ id });
      });

      setTimeout(() => {
        reject(new Error("timeout"));
      }, 20000);
    });

    await expect(newClass()).to.be.not.rejected;
    const newClassEvent = await NewClassEvent;
    const classId = newClassEvent.id;

    const _newNFTClass = await ethers.getContractAt("LikeNFTClass", classId);
    expect(await _newNFTClass.name()).to.equal("My Book");
    expect(await _newNFTClass.symbol()).to.equal("KOOB");
  });

  it("should be allow everyone to create new class", async function () {
    const likeNFTSigner = contract.connect(this.randomSigner);

    const newClass = async () => {
      await likeNFTSigner
        .newClass({
          creator: this.randomSigner,
          input: {
            name: "My Book",
            symbol: "KOOB",
            metadata: JSON.stringify({
              name: "Random by somone",
              symbol: "No data",
            }),
            config: {
              max_supply: 10,
            },
          },
        })
        .then((tx) => tx.wait());
    };

    await expect(newClass()).to.be.not.rejected;
  });
});
