import { expect } from "chai";
import { BaseContract, EventLog } from "ethers";
import { ethers, upgrades } from "hardhat";

import { BookConfigLoader } from "./BookConfigLoader";
import { createProtocol } from "./ProtocolFactory";

describe("LikeProtocol", () => {
  before(async function () {
    this.BookNFT = await ethers.getContractFactory("BookNFT");
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    this.LikeProtocolMock = await ethers.getContractFactory("LikeProtocolMock");
    const [ownerSigner, randomSigner] = await ethers.getSigners();

    this.ownerSigner = ownerSigner;
    this.randomSigner = randomSigner;
  });

  let deployment: BaseContract;
  let contractAddress: string;
  let contract: any;
  let bookNFTContractAddress: string;
  beforeEach(async function () {
    const {
      likeProtocol,
      likeProtocolDeployment,
      likeProtocolAddress,
      likeProtocolContract,
      bookNFT,
      bookNFTDeployment,
      bookNFTAddress,
      bookNFTContract,
    } = await createProtocol(this.ownerSigner);

    deployment = likeProtocolDeployment;
    contractAddress = likeProtocolAddress;
    contract = likeProtocolContract;
    bookNFTContractAddress = bookNFTAddress;
  });

  it("should have the correct STORAGE_SLOT", async function () {
    const likeProtocolMockOwnerSigner = this.LikeProtocolMock.connect(
      this.ownerSigner,
    );
    const newLikeProtocol = await upgrades.upgradeProxy(
      contractAddress,
      likeProtocolMockOwnerSigner,
    );
    expect(await newLikeProtocol.protocolDataStorage()).to.equal(
      "0xe3ffde652b1592025b57f85d2c64876717f9cdf4e44b57422a295c18d0719a00",
    );
  });

  it("should have the correct bookNFTImplementation", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);
    expect(await likeProtocolOwnerSigner.implementation()).to.equal(
      bookNFTContractAddress,
    );
  });

  it("should be upgradable", async function () {
    const likeProtocolMockOwnerSigner = this.LikeProtocolMock.connect(
      this.ownerSigner,
    );
    const newLikeProtocol = await upgrades.upgradeProxy(
      contractAddress,
      likeProtocolMockOwnerSigner,
    );
    expect(await newLikeProtocol.getAddress()).to.equal(contractAddress);
    expect(await newLikeProtocol.owner()).to.equal(this.ownerSigner.address);
    expect(await newLikeProtocol.version()).to.equal(2n);
  });

  it("should not be upgradable by random address", async function () {
    const likeProtocolMockRandomSigner = this.LikeProtocolMock.connect(
      this.randomSigner,
    );
    await expect(
      upgrades.upgradeProxy(contractAddress, likeProtocolMockRandomSigner),
    ).to.be.rejected;
    expect(await contract.owner()).to.equal(this.ownerSigner.address);
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

  it("should be able to create new BookNFT", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    const newClass = async () => {
      await likeProtocolOwnerSigner
        .newBookNFT({
          creator: this.ownerSigner,
          updaters: [this.ownerSigner],
          minters: [this.ownerSigner],
          config: bookConfig,
        })
        .then((tx) => tx.wait());
    };

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeProtocolOwnerSigner.on("NewBookNFT", (id, params, event) => {
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

    const _newNFTClass = await ethers.getContractAt("BookNFT", classId);
    expect(await _newNFTClass.name()).to.equal("My Book");
    expect(await _newNFTClass.symbol()).to.equal("KOOB");
    expect(await _newNFTClass.getProtocolBeacon()).to.equal(contractAddress);
    const [receiver, royaltyAmount] = await _newNFTClass.royaltyInfo(0, 1000);
    expect(receiver).to.equal(this.ownerSigner.address);
    expect(royaltyAmount).to.equal(0n);
  });

  it("should be able to create new BookNFT with royalty", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    const newClass = async () => {
      await likeProtocolOwnerSigner
        .newBookNFTWithRoyalty(
          {
            creator: this.ownerSigner,
            updaters: [this.ownerSigner],
            minters: [this.ownerSigner],
            config: bookConfig,
          },
          100,
        )
        .then((tx) => tx.wait());
    };

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeProtocolOwnerSigner.on("NewBookNFT", (id, params, event) => {
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

    expect(classId).to.not.equal(bookNFTContractAddress);
    const _newNFTClass = await ethers.getContractAt("BookNFT", classId);
    const [receiver, royaltyAmount] = await _newNFTClass.royaltyInfo(0, 1000);
    expect(receiver).to.equal(this.ownerSigner.address);
    expect(royaltyAmount).to.equal(10n);
  });

  it("should set the right royalty receiver on initialization", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);
    expect(await likeProtocolOwnerSigner.getRoyaltyReceiver()).to.equal(
      this.ownerSigner.address,
    );
  });

  it("should allow owner to set the royalty receiver", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);
    await likeProtocolOwnerSigner.setRoyaltyReceiver(this.randomSigner.address);
    expect(await likeProtocolOwnerSigner.getRoyaltyReceiver()).to.equal(
      this.randomSigner.address,
    );
  });

  it("should not allow random address to set the royalty receiver", async function () {
    const likeProtocolRandomSigner = contract.connect(this.randomSigner);
    await expect(
      likeProtocolRandomSigner.setRoyaltyReceiver(this.randomSigner.address),
    ).to.be.rejected;
  });

  it("should mint a already initialized BookNFT which cant be initialized again", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    const newClass = async () => {
      await likeProtocolOwnerSigner
        .newBookNFT({
          creator: this.ownerSigner,
          updaters: [this.ownerSigner],
          minters: [this.ownerSigner],
          config: bookConfig,
        })
        .then((tx) => tx.wait());
    };

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeProtocolOwnerSigner.on("NewBookNFT", (id, params, event) => {
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

    const _newNFTClass = await ethers.getContractAt("BookNFT", classId);
    const bookNFTRandomSigner = _newNFTClass.connect(this.randomSigner);
    await expect(
      bookNFTRandomSigner.initialize({
        creator: this.randomSigner,
        updaters: [this.randomSigner, this.randomSigner],
        minters: [this.randomSigner, this.randomSigner],
        config: bookConfig,
      }),
    ).to.be.rejectedWith("InvalidInitialization()");
  });

  it("should not allow to create new BookNFT when paused", async function () {
    const likeProtocolSigner = contract.connect(this.ownerSigner);

    const classOperation = async () => {
      await likeProtocolSigner
        .newBookNFT({
          creator: this.ownerSigner,
          updaters: [this.ownerSigner],
          minters: [this.ownerSigner],
          config: {
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
            max_supply: 10,
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

  it("should be allow everyone to create new BookNFT", async function () {
    const likeNFTSigner = contract.connect(this.randomSigner);

    const newClass = async () => {
      await likeNFTSigner
        .newBookNFT({
          creator: this.randomSigner,
          updaters: [this.randomSigner],
          minters: [this.randomSigner],
          config: {
            name: "My Book",
            symbol: "KOOB",
            metadata: JSON.stringify({
              name: "Random by somone",
              symbol: "No data",
            }),
            max_supply: 10,
          },
        })
        .then((tx) => tx.wait());
    };

    await expect(newClass()).to.be.not.rejected;
  });

  it("should not allow to create new BookNFT when max supply is 0", async function () {
    const likeNFTSigner = contract.connect(this.randomSigner);

    const newClass = async () => {
      await likeNFTSigner
        .newBookNFT({
          creator: this.randomSigner,
          updaters: [this.randomSigner],
          minters: [this.randomSigner],
          config: {
            name: "My Book",
            symbol: "KOOB",
            metadata: JSON.stringify({
              name: "Random by somone",
              symbol: "No data",
            }),
            max_supply: 0,
          },
        })
        .then((tx) => tx.wait());
    };

    await expect(newClass()).to.be.rejectedWith(
      "VM Exception while processing transaction: reverted with custom error 'ErrMaxSupplyZero()'",
    );
  });

  it("should not allow everyone to create new BookNFT when paused", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);
    const likeProtocolRandomSigner = contract.connect(this.randomSigner);

    const classOperation = async () => {
      await likeProtocolRandomSigner
        .newBookNFT({
          creator: this.randomSigner,
          updaters: [this.randomSigner],
          minters: [this.randomSigner],
          config: {
            name: "My Book",
            symbol: "KOOB",
            metadata: JSON.stringify({
              name: "Random by somone",
              symbol: "No data",
            }),
            max_supply: 10,
          },
        })
        .then((tx) => tx.wait());
    };

    await expect(classOperation()).to.be.not.rejected;
    await expect(likeProtocolOwnerSigner.pause()).to.be.not.rejected;
    await expect(classOperation()).to.be.rejectedWith(
      "VM Exception while processing transaction: reverted with custom error 'EnforcedPause()'",
    );
    await expect(likeProtocolOwnerSigner.unpause()).to.be.not.rejected;
    await expect(classOperation()).to.be.not.rejected;
  });

  it("should retain the BookNFT paused state after upgrade", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);

    expect(await likeProtocolOwnerSigner.paused()).to.be.false;
    await likeProtocolOwnerSigner.pause();
    expect(await likeProtocolOwnerSigner.paused()).to.be.true;

    const likeProtocolMockOwnerSigner = this.LikeProtocolMock.connect(
      this.ownerSigner,
    );
    const newLikeProtocol = await upgrades.upgradeProxy(
      contractAddress,
      likeProtocolMockOwnerSigner,
    );
    expect(await newLikeProtocol.owner()).to.equal(this.ownerSigner.address);
    expect(await newLikeProtocol.version()).to.equal(2n);

    const proxyContract = await ethers.getContractAt(
      "LikeProtocolMock",
      contractAddress,
    );
    expect(await proxyContract.version()).to.equal(2n);

    expect(await likeProtocolOwnerSigner.paused()).to.be.true;
    await likeProtocolOwnerSigner.unpause();
    expect(await likeProtocolOwnerSigner.paused()).to.be.false;
  });

  it("should retain the BookNFT mapping after upgrade", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);
    const newClass = async () => {
      await likeProtocolOwnerSigner
        .newBookNFT({
          creator: this.ownerSigner,
          updaters: [this.ownerSigner],
          minters: [this.ownerSigner],
          config: {
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
            max_supply: 10,
          },
        })
        .then((tx) => tx.wait());
    };

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeProtocolOwnerSigner.on("NewBookNFT", (id, params, event) => {
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
    expect(await likeProtocolOwnerSigner.isBookNFT(classId)).to.be.true;

    const likeProtocolMockOwnerSigner = this.LikeProtocolMock.connect(
      this.ownerSigner,
    );
    const newLikeProtocol = await upgrades.upgradeProxy(
      contractAddress,
      likeProtocolMockOwnerSigner,
    );
    expect(await newLikeProtocol.owner()).to.equal(this.ownerSigner.address);
    expect(await newLikeProtocol.version()).to.equal(2n);
    expect(await newLikeProtocol.isBookNFT(classId)).to.be.true;

    const proxyContract = await ethers.getContractAt(
      "LikeProtocolMock",
      contractAddress,
    );
    expect(await proxyContract.isBookNFT(classId)).to.be.true;
  });
});

describe("LikeProtocol events", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    this.LikeProtocolMock = await ethers.getContractFactory("LikeProtocolMock");
    this.BookNFT = await ethers.getContractFactory("BookNFT");
    const [ownerSigner, randomSigner] = await ethers.getSigners();

    this.ownerSigner = ownerSigner;
    this.randomSigner = randomSigner;
  });

  let deployment: BaseContract;
  let contractAddress: string;
  let contract: any;
  beforeEach(async function () {
    const {
      likeProtocol,
      likeProtocolDeployment,
      likeProtocolAddress,
      likeProtocolContract,
      bookNFT,
      bookNFTDeployment,
      bookNFTAddress,
      bookNFTContract,
    } = await createProtocol(this.ownerSigner);

    deployment = likeProtocolDeployment;
    contractAddress = likeProtocolAddress;
    contract = likeProtocolContract;
  });

  it("should emit NewBookNFT event calling newBookNFT", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeProtocolOwnerSigner.on("NewBookNFT", (id, params, event) => {
        event.removeListener();
        resolve({ id });
      });

      setTimeout(() => {
        reject(new Error("timeout"));
      }, 60000);
    });

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    expect(bookConfig.name).to.equal("My Book");
    expect(bookConfig.symbol).to.equal("KOOB");

    const newClass = async () => {
      await likeProtocolOwnerSigner
        .newBookNFT({
          creator: this.ownerSigner,
          updaters: [this.ownerSigner],
          minters: [this.ownerSigner],
          config: bookConfig,
        })
        .then((tx) => tx.wait());
    };

    await expect(newClass()).to.be.not.rejected;
    const newClassEvent = await NewClassEvent;
    const classId = newClassEvent.id;
    expect(await likeProtocolOwnerSigner.isBookNFT(classId)).to.be.true;
  });

  it("should emit multiple NewBookNFT events calling newBookNFTs", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);

    const NewClassEvent = new Promise<string[]>((resolve, reject) => {
      const ids: string[] = [];
      const listener = (id: string, params: any, event: any) => {
        ids.push(id);
        if (ids.length === 2) {
          // We expect 2 events
          likeProtocolOwnerSigner.off("NewBookNFT", listener);
          resolve(ids);
        }
      };

      likeProtocolOwnerSigner.on("NewBookNFT", listener);

      setTimeout(() => {
        reject(new Error("timeout waiting for NewBookNFT events"));
      }, 60000); // Longer timeout just in case, but won't wait unless there's an issue
    });

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    const bookConfig1 = BookConfigLoader.load(
      "./test/fixtures/BookConfig1.json",
    );
    expect(bookConfig.name).to.equal("My Book");
    expect(bookConfig.symbol).to.equal("KOOB");
    expect(bookConfig1.name).to.equal("My Book 1");
    expect(bookConfig1.symbol).to.equal("KOOB1");

    const newClasses = async () => {
      await likeProtocolOwnerSigner
        .newBookNFTs([
          {
            creator: this.ownerSigner,
            updaters: [this.ownerSigner],
            minters: [this.ownerSigner],
            config: bookConfig,
          },
          {
            creator: this.ownerSigner,
            updaters: [this.ownerSigner],
            minters: [this.ownerSigner],
            config: bookConfig1,
          },
        ])
        .then((tx) => tx.wait());
    };

    await expect(newClasses()).to.be.not.rejected;
    const classIds = await NewClassEvent;

    expect(classIds.length).to.equal(2);
    expect(await likeProtocolOwnerSigner.isBookNFT(classIds[0])).to.be.true;
    expect(await likeProtocolOwnerSigner.isBookNFT(classIds[1])).to.be.true;

    let bookNFTOwnerSigner = this.BookNFT.connect(this.ownerSigner);
    bookNFTOwnerSigner = bookNFTOwnerSigner.attach(classIds[0]);
    expect(await bookNFTOwnerSigner.symbol()).to.equal(bookConfig.symbol);

    bookNFTOwnerSigner = bookNFTOwnerSigner.attach(classIds[1]);
    expect(await bookNFTOwnerSigner.symbol()).to.equal(bookConfig1.symbol);
  });
});

describe("LikeProtocol as Beacon", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    this.LikeProtocolMock = await ethers.getContractFactory("LikeProtocolMock");
    this.BookNFT = await ethers.getContractFactory("BookNFT");
    this.BookNFTMock = await ethers.getContractFactory("BookNFTMock");
    const [ownerSigner, randomSigner] = await ethers.getSigners();

    this.ownerSigner = ownerSigner;
    this.randomSigner = randomSigner;
  });

  let deployment: BaseContract;
  let contractAddress: string;
  let contract: any;
  let bookNFTContractAddress: string;
  let bookNFTMockContractAddress: string;
  beforeEach(async function () {
    const {
      likeProtocol,
      likeProtocolDeployment,
      likeProtocolAddress,
      likeProtocolContract,
      bookNFTDeployment,
      bookNFTAddress,
    } = await createProtocol(this.ownerSigner);

    deployment = likeProtocolDeployment;
    contractAddress = likeProtocolAddress;
    contract = likeProtocolContract;
    bookNFTContractAddress = bookNFTAddress;

    const bookNFTMockDeployment = await this.BookNFTMock.deploy();
    bookNFTMockContractAddress = await bookNFTMockDeployment.getAddress();
  });

  it("should only owner can upgrade the implementation", async function () {
    const likeProtocolOwnerSigner = contract.connect(this.ownerSigner);
    const likeProtocolRandomSigner = contract.connect(this.randomSigner);

    await expect(likeProtocolOwnerSigner.upgradeTo(bookNFTMockContractAddress))
      .to.be.not.rejected;
    await expect(likeProtocolRandomSigner.upgradeTo(bookNFTMockContractAddress))
      .to.be.rejected;
  });
});
