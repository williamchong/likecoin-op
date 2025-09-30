import { loadFixture } from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { viem, ignition } from "hardhat";

import "./setup";
import { BookConfigLoader } from "./BookConfigLoader";
import { deployProtocol } from "./factory";

describe("LikeProtocol", () => {
  it("should have expected proxy address", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);
    const deployerAddress = deployer.account.address;
    expect(deployerAddress).to.equalAddress(
      "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
    );
    expect(likeProtocol.address).to.equalAddress(
      "0x44322611140BC362972878FdEcEF335315E2c364",
    );
  });

  it("should have the correct STORAGE_SLOT", async function () {
    const likeProtocolMock = await viem.deployContract("LikeProtocolMock");
    expect(await likeProtocolMock.read.protocolDataStorage()).to.equal(
      "0xe3ffde652b1592025b57f85d2c64876717f9cdf4e44b57422a295c18d0719a00",
    );
  });

  it("should have the correct bookNFTImplementation", async function () {
    const { likeProtocol, bookNFTImpl } = await loadFixture(deployProtocol);
    const impl = await likeProtocol.read.implementation();
    expect(impl.toLowerCase()).to.equal(bookNFTImpl.address.toLowerCase());
  });

  it("should be upgradable", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);
    const likeProtocolMock = await viem.deployContract("LikeProtocolMock");
    await likeProtocol.write.upgradeToAndCall(
      [likeProtocolMock.address, "0x"],
      {
        account: deployer.account,
      },
    );
    const proxyAsMock = await viem.getContractAt(
      "LikeProtocolMock",
      likeProtocol.address,
    );
    const owner = await proxyAsMock.read.owner();
    expect(owner.toLowerCase()).to.equal(
      deployer.account.address.toLowerCase(),
    );
    expect(await proxyAsMock.read.version()).to.equal(2n);
  });

  it("should not be upgradable by random address", async function () {
    const { likeProtocol, randomSigner, deployer } =
      await loadFixture(deployProtocol);
    const likeProtocolMock = await viem.deployContract("LikeProtocolMock");
    await expect(
      likeProtocol.write.upgradeToAndCall([likeProtocolMock.address, "0x"], {
        account: randomSigner.account,
      }),
    ).to.be.rejected;
    const owner = await likeProtocol.read.owner();
    expect(owner.toLowerCase()).to.equal(
      deployer.account.address.toLowerCase(),
    );
  });

  it("should set the right owner", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);
    const owner = await likeProtocol.read.owner();
    expect(owner).to.equalAddress(deployer.account.address);
  });

  it("should allow ownership transfer", async function () {
    const { likeProtocol, deployer, randomSigner } =
      await loadFixture(deployProtocol);
    await likeProtocol.write.transferOwnership([randomSigner.account.address], {
      account: deployer.account,
    });
    const owner = await likeProtocol.read.owner();
    expect(owner).to.equalAddress(randomSigner.account.address);
  });

  it("should not allow random ownership transfer", async function () {
    const { likeProtocol, randomSigner, deployer } =
      await loadFixture(deployProtocol);
    await expect(
      likeProtocol.write.transferOwnership([randomSigner.account.address], {
        account: randomSigner.account,
      }),
    ).to.be.rejected;
    const owner = await likeProtocol.read.owner();
    expect(owner).to.equalAddress(deployer.account.address);
  });

  it("should not paused by random address", async function () {
    const { likeProtocol, randomSigner } = await loadFixture(deployProtocol);
    await expect(likeProtocol.write.pause({ account: randomSigner.account })).to
      .be.rejected;
  });

  it("should be paused by owner address", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);
    await expect(likeProtocol.write.pause({ account: deployer.account })).to.be
      .not.rejected;
  });

  it("should not unpaused by random address", async function () {
    const { likeProtocol, deployer, randomSigner } =
      await loadFixture(deployProtocol);
    await expect(likeProtocol.write.pause({ account: deployer.account })).to.be
      .not.rejected;
    await expect(likeProtocol.write.unpause({ account: randomSigner.account }))
      .to.be.rejected;
  });

  it("should be able to create new BookNFT", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    const NewClassEvent = new Promise<{ id: `0x${string}` }>((resolve) => {
      const unwatch = likeProtocol.watchEvent.NewBookNFT({
        onLogs: (logs) => {
          const id = logs[0].args.bookNFT as `0x${string}`;
          unwatch();
          resolve({ id });
        },
      });
    });

    await likeProtocol.write.newBookNFT([
      {
        creator: deployer.account.address,
        updaters: [deployer.account.address],
        minters: [deployer.account.address],
        config: bookConfig,
      },
    ]);

    const { id: classId } = await NewClassEvent;
    const newNFTClass = await viem.getContractAt("BookNFT", classId);
    expect(await newNFTClass.read.name()).to.equal("My Book");
    expect(await newNFTClass.read.symbol()).to.equal("KOOB");
    const beacon = await newNFTClass.read.getProtocolBeacon();
    expect(beacon.toLowerCase()).to.equal(likeProtocol.address.toLowerCase());
    expect(await newNFTClass.read.contractURI()).to.equal(
      "data:application/json;base64,eyJuYW1lIjoiQ29sbGVjdGlvbiBOYW1lIiwic3ltYm9sIjoiQ29sbGVjdGlvbiBTWU1CIiwiZGVzY3JpcHRpb24iOiJDb2xsZWN0aW9uIERlc2NyaXB0aW9uIiwiaW1hZ2UiOiJpcGZzOi8vYmFmeWJlaWV6cTR5cW9zYzJ1NHNhYW5vdmU1YnNhM3ljaXVmd2hmZHVlbXk1ejZ2dmY2cTNjNWxuYmkiLCJiYW5uZXJfaW1hZ2UiOiIiLCJmZWF0dXJlZF9pbWFnZSI6IiIsImV4dGVybmFsX2xpbmsiOiJodHRwczovL3d3dy5leGFtcGxlLmNvbSIsImNvbGxhYm9yYXRvcnMiOltdfQ==",
    );
    const royalty = await newNFTClass.read.royaltyInfo([0n, 1000n]);
    expect(royalty[0].toLowerCase()).to.equal(
      deployer.account.address.toLowerCase(),
    );
    expect(royalty[1]).to.equal(0n);
  });

  it("should be able to create new BookNFT with royalty", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    const NewClassEvent = new Promise<{ id: `0x${string}` }>((resolve) => {
      const unwatch = likeProtocol.watchEvent.NewBookNFT({
        onLogs: (logs) => {
          const id = logs[0].args.bookNFT as `0x${string}`;
          unwatch();
          resolve({ id });
        },
      });
    });

    await likeProtocol.write.newBookNFTWithRoyalty([
      {
        creator: deployer.account.address,
        updaters: [deployer.account.address],
        minters: [deployer.account.address],
        config: bookConfig,
      },
      100n,
    ]);

    const { id: classId } = await NewClassEvent;
    const newNFTClass = await viem.getContractAt("BookNFT", classId);
    const [receiver, royaltyAmount] = await newNFTClass.read.royaltyInfo([
      0n,
      1000n,
    ]);
    expect(receiver.toLowerCase()).to.equal(
      deployer.account.address.toLowerCase(),
    );
    expect(royaltyAmount).to.equal(10n);
  });

  it("should set the right royalty receiver on initialization", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);
    const receiver = await likeProtocol.read.getRoyaltyReceiver();
    expect(receiver.toLowerCase()).to.equal(
      deployer.account.address.toLowerCase(),
    );
  });

  it("should allow owner to set the royalty receiver", async function () {
    const { likeProtocol, deployer, randomSigner } =
      await loadFixture(deployProtocol);
    await likeProtocol.write.setRoyaltyReceiver(
      [randomSigner.account.address],
      {
        account: deployer.account,
      },
    );
    const receiver = await likeProtocol.read.getRoyaltyReceiver();
    expect(receiver.toLowerCase()).to.equal(
      randomSigner.account.address.toLowerCase(),
    );
  });

  it("should not allow random address to set the royalty receiver", async function () {
    const { likeProtocol, randomSigner } = await loadFixture(deployProtocol);
    await expect(
      likeProtocol.write.setRoyaltyReceiver([randomSigner.account.address], {
        account: randomSigner.account,
      }),
    ).to.be.rejected;
  });

  it("should mint a already initialized BookNFT which cant be initialized again", async function () {
    const { likeProtocol, deployer, randomSigner } =
      await loadFixture(deployProtocol);
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    const NewClassEvent = new Promise<{ id: `0x${string}` }>((resolve) => {
      const unwatch = likeProtocol.watchEvent.NewBookNFT({
        onLogs: (logs) => {
          const id = logs[0].args.bookNFT as `0x${string}`;
          unwatch();
          resolve({ id });
        },
      });
    });
    await likeProtocol.write.newBookNFT([
      {
        creator: deployer.account.address,
        updaters: [deployer.account.address],
        minters: [deployer.account.address],
        config: bookConfig,
      },
    ]);
    const { id: classId } = await NewClassEvent;

    const newNFTClass = await viem.getContractAt("BookNFT", classId);
    await expect(
      newNFTClass.write.initialize([bookConfig.name, bookConfig.symbol], {
        account: randomSigner.account,
      }),
    ).to.be.rejectedWith("InvalidInitialization()");
  });

  it("should not allow to create new BookNFT when paused", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);
    const classOperation = async () => {
      await likeProtocol.write.newBookNFT([
        {
          creator: deployer.account.address,
          updaters: [deployer.account.address],
          minters: [deployer.account.address],
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
            max_supply: 10n,
          },
        },
      ]);
    };

    await expect(classOperation()).to.be.not.rejected;
    await expect(likeProtocol.write.pause({ account: deployer.account })).to.be
      .not.rejected;
    await expect(classOperation()).to.be.rejectedWith("EnforcedPause");
    await expect(likeProtocol.write.unpause({ account: deployer.account })).to
      .be.not.rejected;
    await expect(classOperation()).to.be.not.rejected;
  });

  it("should be allow everyone to create new BookNFT", async function () {
    const { likeProtocol, randomSigner } = await loadFixture(deployProtocol);
    const newClass = async () => {
      await likeProtocol.write.newBookNFT(
        [
          {
            creator: randomSigner.account.address,
            updaters: [randomSigner.account.address],
            minters: [randomSigner.account.address],
            config: {
              name: "My Book",
              symbol: "KOOB",
              metadata: JSON.stringify({
                name: "Random by somone",
                symbol: "No data",
              }),
              max_supply: 10n,
            },
          },
        ],
        { account: randomSigner.account },
      );
    };

    await expect(newClass()).to.be.not.rejected;
  });

  it("should not allow to create new BookNFT when max supply is 0", async function () {
    const { likeProtocol, randomSigner } = await loadFixture(deployProtocol);
    const newClass = async () => {
      await likeProtocol.write.newBookNFT(
        [
          {
            creator: randomSigner.account.address,
            updaters: [randomSigner.account.address],
            minters: [randomSigner.account.address],
            config: {
              name: "My Book",
              symbol: "KOOB",
              metadata: JSON.stringify({
                name: "Random by somone",
                symbol: "No data",
              }),
              max_supply: 0n,
            },
          },
        ],
        { account: randomSigner.account },
      );
    };

    await expect(newClass()).to.be.rejectedWith("ErrMaxSupplyZero()");
  });

  it("should not allow everyone to create new BookNFT when paused", async function () {
    const { likeProtocol, deployer, randomSigner } =
      await loadFixture(deployProtocol);
    const classOperation = async () => {
      await likeProtocol.write.newBookNFT(
        [
          {
            creator: randomSigner.account.address,
            updaters: [randomSigner.account.address],
            minters: [randomSigner.account.address],
            config: {
              name: "My Book",
              symbol: "KOOB",
              metadata: JSON.stringify({
                name: "Random by somone",
                symbol: "No data",
              }),
              max_supply: 10n,
            },
          },
        ],
        { account: randomSigner.account },
      );
    };

    await expect(classOperation()).to.be.not.rejected;
    await expect(likeProtocol.write.pause({ account: deployer.account })).to.be
      .not.rejected;
    await expect(classOperation()).to.be.rejectedWith("EnforcedPause");
    await expect(likeProtocol.write.unpause({ account: deployer.account })).to
      .be.not.rejected;
    await expect(classOperation()).to.be.not.rejected;
  });

  it("should retain the BookNFT paused state after upgrade", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);
    expect(await likeProtocol.read.paused()).to.be.false;
    await likeProtocol.write.pause({ account: deployer.account });
    expect(await likeProtocol.read.paused()).to.be.true;

    const likeProtocolMock = await viem.deployContract("LikeProtocolMock");
    await likeProtocol.write.upgradeToAndCall(
      [likeProtocolMock.address, "0x"],
      {
        account: deployer.account,
      },
    );
    const proxyAsMock = await viem.getContractAt(
      "LikeProtocolMock",
      likeProtocol.address,
    );
    expect(await proxyAsMock.read.version()).to.equal(2n);

    expect(await likeProtocol.read.paused()).to.be.true;
    await likeProtocol.write.unpause({ account: deployer.account });
    expect(await likeProtocol.read.paused()).to.be.false;
  });

  it("should retain the BookNFT mapping after upgrade", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);
    const NewClassEvent = new Promise<{ id: `0x${string}` }>((resolve) => {
      const unwatch = likeProtocol.watchEvent.NewBookNFT({
        onLogs: (logs) => {
          const id = logs[0].args.bookNFT as `0x${string}`;
          unwatch();
          resolve({ id });
        },
      });
    });
    await likeProtocol.write.newBookNFT([
      {
        creator: deployer.account.address,
        updaters: [deployer.account.address],
        minters: [deployer.account.address],
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
          max_supply: 10n,
        },
      },
    ]);

    const { id: classId } = await NewClassEvent;
    expect(await likeProtocol.read.isBookNFT([classId])).to.be.true;

    const likeProtocolMock = await viem.deployContract("LikeProtocolMock");
    await likeProtocol.write.upgradeToAndCall(
      [likeProtocolMock.address, "0x"],
      {
        account: deployer.account,
      },
    );

    const proxyAsMock = await viem.getContractAt(
      "LikeProtocolMock",
      likeProtocol.address,
    );
    expect(await proxyAsMock.read.version()).to.equal(2n);
    expect(await proxyAsMock.read.isBookNFT([classId])).to.be.true;
  });
});

describe("LikeProtocol events", () => {
  it("should emit NewBookNFT event calling newBookNFT", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    expect(bookConfig.name).to.equal("My Book");
    expect(bookConfig.symbol).to.equal("KOOB");

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
    const book0 = await viem.getContractAt("BookNFT", bookNFTAddress);
    expect(await book0.read.symbol()).to.equal(bookConfig.symbol);
    expect(await book0.read.name()).to.equal(bookConfig.name);
  });

  it("should emit multiple NewBookNFT events calling newBookNFTs", async function () {
    const { likeProtocol, deployer } = await loadFixture(deployProtocol);

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    const bookConfig1 = BookConfigLoader.load(
      "./test/fixtures/BookConfig1.json",
    );

    await likeProtocol.write.newBookNFTs([
      [
        {
          creator: deployer.account.address,
          updaters: [deployer.account.address],
          minters: [deployer.account.address],
          config: bookConfig,
        },
        {
          creator: deployer.account.address,
          updaters: [deployer.account.address],
          minters: [deployer.account.address],
          config: bookConfig1,
        },
      ],
    ]);

    const logs = await likeProtocol.getEvents.NewBookNFT();
    expect(logs.length).to.equal(2);
    expect(await likeProtocol.read.isBookNFT([logs[0].args.bookNFT])).to.be
      .true;
    expect(await likeProtocol.read.isBookNFT([logs[1].args.bookNFT])).to.be
      .true;

    const book0 = await viem.getContractAt("BookNFT", logs[0].args.bookNFT);
    const book1 = await viem.getContractAt("BookNFT", logs[1].args.bookNFT);
    expect(await book0.read.symbol()).to.equal(bookConfig.symbol);
    expect(await book1.read.symbol()).to.equal(bookConfig1.symbol);
  });
});

describe("LikeProtocol as Beacon", () => {
  it("should only owner can upgrade the implementation", async function () {
    const { likeProtocol, deployer, randomSigner } =
      await loadFixture(deployProtocol);

    const bookNFTMock = await viem.deployContract("BookNFTMock");
    await expect(
      likeProtocol.write.upgradeTo([bookNFTMock.address], {
        account: deployer.account,
      }),
    ).to.be.not.rejected;
    await expect(
      likeProtocol.write.upgradeTo([bookNFTMock.address], {
        account: randomSigner.account,
      }),
    ).to.be.rejected;
    expect(await likeProtocol.read.implementation()).to.equalAddress(
      bookNFTMock.address,
    );
  });
});
