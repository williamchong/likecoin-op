import { expect } from "chai";
import { EventLog, BaseContract } from "ethers";
import { ethers, upgrades } from "hardhat";

import { BookConfigLoader } from "./BookConfigLoader";
import { createProtocol } from "./ProtocolFactory";
describe("BookNFTClass", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    this.BookNFTMock = await ethers.getContractFactory("BookNFTMock");
    const [protocolOwner, classOwner, likerLand, randomSigner] =
      await ethers.getSigners();

    this.protocolOwner = protocolOwner;
    this.classOwner = classOwner;
    this.likerLand = likerLand;
    this.randomSigner = randomSigner;
  });

  let deployment: BaseContract;
  let contractAddress: string;
  let protocolContract: BaseContract;
  let bookNFTImplementation: BaseContract;
  let nftClassId: string;
  let nftClassContract: BaseContract;
  beforeEach(async function () {
    const {
      likeProtocol,
      likeProtocolDeployment,
      likeProtocolAddress,
      likeProtocolContract,
      bookNFTDeployment,
      bookNFTAddress,
    } = await createProtocol(this.protocolOwner);

    deployment = likeProtocolDeployment;
    contractAddress = likeProtocolAddress;
    protocolContract = likeProtocolContract;
    bookNFTImplementation = bookNFTDeployment;

    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeProtocolOwnerSigner.on(
        "NewBookNFT",
        (id: string, params: any, event: any) => {
          event.removeListener();
          resolve({ id });
        },
      );
      setTimeout(() => {
        reject(new Error("timeout"));
      }, 20000);
    });

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    await likeProtocolOwnerSigner.newBookNFT({
      creator: this.classOwner,
      updaters: [this.classOwner, this.likerLand],
      minters: [this.classOwner, this.likerLand],
      config: bookConfig,
    });

    const newClassEvent = await NewClassEvent;
    nftClassId = newClassEvent.id;
    nftClassContract = await ethers.getContractAt("BookNFT", nftClassId);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);
  });

  it("should have the correct STORAGE_SLOT", async function () {
    const bookNFTMockOwnerSigner = this.BookNFTMock.connect(this.protocolOwner);
    const newBookNFT = await bookNFTMockOwnerSigner.deploy();
    expect(await newBookNFT.bookNFTStorage()).to.equal(
      "0x8303e9d27d04c843c8d4a08966b1e1be0214fc0b3375d79db0a8252068c41f00",
    );
  });

  it("should not able to re-initialize", async function () {
    const bookNFTRandomSigner = bookNFTImplementation.connect(this.randomSigner);
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    const owner = await bookNFTRandomSigner.owner();
    expect(owner).is.not.equal(this.randomSigner.address);
    
    await expect(bookNFTRandomSigner.initialize({
      creator: this.randomSigner,
      updaters: [this.randomSigner, this.randomSigner],
      minters: [this.randomSigner, this.randomSigner],
      config: bookConfig,
    })).to.be.rejectedWith(
      "InvalidInitialization()",
    );
  });

  it("should have the right roles assigned", async function () {
    const MINTER_ROLE = await nftClassContract.MINTER_ROLE();
    const UPDATER_ROLE = await nftClassContract.UPDATER_ROLE();
    expect(
      await nftClassContract.hasRole(MINTER_ROLE, this.protocolOwner.address),
    ).to.equal(false);
    expect(
      await nftClassContract.hasRole(MINTER_ROLE, this.classOwner.address),
    ).to.equal(true);
    expect(
      await nftClassContract.hasRole(MINTER_ROLE, this.likerLand.address),
    ).to.equal(true);
    expect(
      await nftClassContract.hasRole(MINTER_ROLE, this.randomSigner.address),
    ).to.equal(false);
    expect(
      await nftClassContract.hasRole(UPDATER_ROLE, this.protocolOwner.address),
    ).to.equal(false);
    expect(
      await nftClassContract.hasRole(UPDATER_ROLE, this.classOwner.address),
    ).to.equal(true);
    expect(
      await nftClassContract.hasRole(UPDATER_ROLE, this.likerLand.address),
    ).to.equal(true);
    expect(
      await nftClassContract.hasRole(UPDATER_ROLE, this.randomSigner.address),
    ).to.equal(false);
  });

  it("should return the right current index", async function () {
    expect(await nftClassContract.getCurrentIndex()).to.equal(0n);

    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);
    const mintNFT = async () => {
      await likeClassOwnerSigner
        .mint(
          this.classOwner.address,
          ["_mint"],
          [
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0001 Description",
              name: "202412191729 #0001",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
          ],
        )
        .then((tx) => tx.wait());
    };

    await expect(mintNFT()).to.be.not.rejected;
    expect(await nftClassContract.getCurrentIndex()).to.equal(1n);
  });

  it("should allow class owner to update class and mint NFT", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);

    const updateClass = async () => {
      await likeClassOwnerSigner
        .update({
          name: "My Book",
          symbol: "NEWBOOK",
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
          max_supply: 20,
        })
        .then((tx) => tx.wait());
    };

    const mintNFT = async () => {
      await likeClassOwnerSigner
        .mint(
          this.classOwner.address,
          ["_mint"],
          [
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0001 Description",
              name: "202412191729 #0001",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
          ],
        )
        .then((tx) => tx.wait());
    };

    await expect(updateClass()).to.be.not.rejected;
    await expect(await nftClassContract.symbol()).to.equal("NEWBOOK");
    await expect(await nftClassContract.totalSupply()).to.equal(0n);
    await expect(mintNFT()).to.be.not.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(1n);
    await expect(updateClass()).to.be.not.rejected;
  });

  it("should reject update class with decreasing max supply", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);
    await expect(await nftClassContract.owner()).to.equal(
      this.classOwner.address,
    );

    await expect(
      likeClassOwnerSigner.update({
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
        max_supply: 5,
      }),
    ).to.be.rejectedWith("ErrSupplyDecrease");

    expect(await nftClassContract.symbol()).to.equal("KOOB");
  });

  it("should not allow random address update class", async function () {
    const likeClassSigner = nftClassContract.connect(this.randomSigner);
    await expect(
      likeClassSigner.update({
        name: "Hi Jack",
        symbol: "HIJACK",
        metadata: JSON.stringify({}),
        config: {
          max_supply: 0,
        },
      }),
    ).to.be.rejected;
    await expect(await nftClassContract.owner()).to.equal(
      this.classOwner.address,
    );
    await expect(await nftClassContract.symbol()).to.equal("KOOB");
  });

  it("should allow class owner to mintNFTs in batch", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);
    await expect(await nftClassContract.totalSupply()).to.equal(0n);
    const mintNFT = async () => {
      await likeClassOwnerSigner
        .batchMint(
          [this.classOwner.address, this.classOwner.address],
          ["_mint1", "_mint2"],
          [
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0001 Description",
              name: "202412191729 #0001",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0001 Description",
              name: "202412191729 #0001",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
          ],
        )
        .then((tx) => tx.wait());
    };
    await expect(mintNFT()).to.be.not.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(2n);
  });

  it("should check token id when safe mint with token id", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);

    const mintNFT = async () => {
      await likeClassOwnerSigner
        .mint(
          this.classOwner.address,
          ["_mint1"],
          [
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0001 Description",
              name: "202412191729 #0001",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
          ],
        )
        .then((tx) => tx.wait());
    };

    const safeMintWithTokenId = async (fromTokenId: number) => {
      await likeClassOwnerSigner
        .safeMintWithTokenId(
          fromTokenId,
          [this.classOwner.address, this.classOwner.address],
          ["_mint1", "_mint2"],
          [
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0002 Description",
              name: "202412191729 #0002",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0003 Description",
              name: "202412191729 #0003",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
          ],
        )
        .then((tx) => tx.wait());
    };

    await expect(mintNFT()).to.be.not.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(1n);
    await expect(safeMintWithTokenId(0)).to.be.rejectedWith(
      "ErrTokenIdMintFails(1)",
    );
    await expect(safeMintWithTokenId(1)).to.be.not.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(3n);
    await expect(safeMintWithTokenId(1)).to.be.rejectedWith(
      "ErrTokenIdMintFails(3)",
    );
  });
});

describe("BookNFT permission control", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    const [protocolOwner, classOwner, likerLand, randomSigner] =
      await ethers.getSigners();

    this.protocolOwner = protocolOwner;
    this.classOwner = classOwner;
    this.likerLand = likerLand;
    this.randomSigner = randomSigner;
  });

  let deployment: BaseContract;
  let contractAddress: string;
  let protocolContract: BaseContract;
  let nftClassId: string;
  let nftClassContract: BaseContract;
  beforeEach(async function () {
    const {
      likeProtocol,
      likeProtocolDeployment,
      likeProtocolAddress,
      likeProtocolContract,
      bookNFTDeployment,
      bookNFTAddress,
    } = await createProtocol(this.protocolOwner);

    deployment = likeProtocolDeployment;
    contractAddress = likeProtocolAddress;
    protocolContract = likeProtocolContract;

    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeProtocolOwnerSigner.on("NewBookNFT", (id, params, event) => {
        event.removeListener();
        resolve({ id });
      });

      setTimeout(() => {
        reject(new Error("timeout"));
      }, 60000);
    });

    likeProtocolOwnerSigner
      .newBookNFT({
        creator: this.classOwner,
        updaters: [this.classOwner, this.likerLand],
        minters: [this.classOwner, this.likerLand],
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

    const newClassEvent = await NewClassEvent;
    nftClassId = newClassEvent.id;
    nftClassContract = await ethers.getContractAt("BookNFT", nftClassId);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);
  });

  it("should allow class owner to mint NFT", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);

    const mintNFT = async () => {
      await likeClassOwnerSigner
        .mint(
          this.likerLand.address,
          ["_mint1", "_mint2"],
          [
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0001 Description",
              name: "202412191729 #0001",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0001 Description",
              name: "202412191729 #0001",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
          ],
        )
        .then((tx) => tx.wait());
    };

    await expect(mintNFT()).to.be.not.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(2n);
    await expect(
      await nftClassContract.balanceOf(this.likerLand.address),
    ).to.equal(2n);
  });
});

describe("BookNFT ownership transfer", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    const [protocolOwner, classOwner, likerLand, randomSigner] =
      await ethers.getSigners();

    this.protocolOwner = protocolOwner;
    this.classOwner = classOwner;
    this.likerLand = likerLand;
    this.randomSigner = randomSigner;
  });

  let deployment: BaseContract;
  let contractAddress: string;
  let protocolContract: BaseContract;
  let nftClassId: string;
  let nftClassContract: BaseContract;
  beforeEach(async function () {
    const {
      likeProtocol,
      likeProtocolDeployment,
      likeProtocolAddress,
      likeProtocolContract,
      bookNFTDeployment,
      bookNFTAddress,
    } = await createProtocol(this.protocolOwner);

    deployment = likeProtocolDeployment;
    contractAddress = likeProtocolAddress;
    protocolContract = likeProtocolContract;

    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeProtocolOwnerSigner.on("NewBookNFT", (id, params, event) => {
        event.removeListener();
        resolve({ id });
      });

      setTimeout(() => {
        reject(new Error("timeout"));
      }, 60000);
    });

    likeProtocolOwnerSigner
      .newBookNFT({
        creator: this.classOwner,
        updaters: [this.classOwner, this.likerLand],
        minters: [this.classOwner, this.likerLand],
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

    const newClassEvent = await NewClassEvent;
    nftClassId = newClassEvent.id;
    nftClassContract = await ethers.getContractAt("BookNFT", nftClassId);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);
  });

  it("should allow class owner to transfer ownership", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);

    const transferOwnership = async () => {
      await likeClassOwnerSigner
        .transferOwnership(this.randomSigner.address)
        .then((tx) => tx.wait());
    };

    await expect(transferOwnership()).to.not.be.rejected;
    expect(await nftClassContract.owner()).to.equal(this.randomSigner.address);
  });

  it("should not allow random signer to transfer ownership", async function () {
    const likeClassRandomSigner = nftClassContract.connect(this.randomSigner);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);

    await expect(
      likeClassRandomSigner.transferOwnership(this.randomSigner.address),
    ).to.be.rejected;
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);
  });

  it("should not modify minter permission when transfer ownership", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);

    const transferOwnership = async () => {
      await likeClassOwnerSigner
        .transferOwnership(this.randomSigner.address)
        .then((tx) => tx.wait());
    };

    await expect(transferOwnership()).to.not.be.rejected;
    expect(await nftClassContract.owner()).to.equal(this.randomSigner.address);

    const mintNFT = async () => {
      await likeClassOwnerSigner
        .mint(
          this.classOwner.address,
          ["_mint1", "_mint2"],
          [
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0001 Description",
              name: "202412191729 #0001",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0001 Description",
              name: "202412191729 #0001",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
          ],
        )
        .then((tx) => tx.wait());
    };

    await expect(mintNFT()).to.be.not.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(2n);
    await expect(
      await nftClassContract.balanceOf(this.classOwner.address),
    ).to.equal(2n);
  });

  it("should not allow next owner to mint NFT without minter permission", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);

    const transferOwnership = async () => {
      await likeClassOwnerSigner
        .transferOwnership(this.randomSigner.address)
        .then((tx) => tx.wait());
    };

    await expect(transferOwnership()).to.not.be.rejected;
    expect(await nftClassContract.owner()).to.equal(this.randomSigner.address);

    const likeClassRandomSigner = nftClassContract.connect(this.randomSigner);
    const mintNFT = async () => {
      await likeClassRandomSigner
        .mint(
          this.randomSigner.address,
          ["_mint1", "_mint2"],
          [
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0001 Description",
              name: "202412191729 #0001",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
            JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "202412191729 #0001 Description",
              name: "202412191729 #0001",
              attributes: [
                {
                  trait_type: "ISCN ID",
                  value:
                    "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                },
              ],
              background_color: "",
              animation_url: "",
              youtube_url: "",
            }),
          ],
        )
        .then((tx) => tx.wait());
    };

    await expect(mintNFT()).to.be.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(0n);
    await expect(
      await nftClassContract.balanceOf(this.randomSigner.address),
    ).to.equal(0n);
  });
});

describe("BookNFT config validation", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    const [protocolOwner, classOwner, likerLand, randomSigner] =
      await ethers.getSigners();

    this.protocolOwner = protocolOwner;
    this.classOwner = classOwner;
    this.likerLand = likerLand;
    this.randomSigner = randomSigner;
  });

  let deployment: BaseContract;
  let contractAddress: string;
  let protocolContract: BaseContract;
  let nftClassId: string;
  let nftClassContract: BaseContract;
  beforeEach(async function () {
    const {
      likeProtocol,
      likeProtocolDeployment,
      likeProtocolAddress,
      likeProtocolContract,
      bookNFTDeployment,
      bookNFTAddress,
    } = await createProtocol(this.protocolOwner);

    deployment = likeProtocolDeployment;
    contractAddress = likeProtocolAddress;
    protocolContract = likeProtocolContract;

    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeProtocolOwnerSigner.on(
        "NewBookNFT",
        (id: string, params: any, event: any) => {
          event.removeListener();
          resolve({ id });
        },
      );
      setTimeout(() => {
        reject(new Error("timeout"));
      }, 20000);
    });

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    await likeProtocolOwnerSigner.newBookNFT({
      creator: this.classOwner,
      updaters: [this.classOwner, this.likerLand],
      minters: [this.classOwner, this.likerLand],
      config: bookConfig,
      max_supply: 10,
    });

    const newClassEvent = await NewClassEvent;
    nftClassId = newClassEvent.id;
    nftClassContract = await ethers.getContractAt("BookNFT", nftClassId);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);
  });

  it("should retuen correct config", async function () {
    const config = await nftClassContract.getBookConfig();
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    expect(config.max_supply).to.equal(BigInt(bookConfig.max_supply));
    expect(config.metadata).to.equal(bookConfig.metadata);
    expect(config.name).to.equal(bookConfig.name);
    expect(config.symbol).to.equal(bookConfig.symbol);
  });

  it("should reject empty name in constructor", async function () {
    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );

    await expect(
      likeProtocolOwnerSigner.newBookNFT({
        creator: this.classOwner,
        updaters: [this.classOwner],
        minters: [this.classOwner],
        config: {
          name: "",
          symbol: "TEST",
          metadata: JSON.stringify({
            name: "Test Collection",
            description: "Test Description",
          }),
          max_supply: 10,
        },
      }),
    ).to.be.rejectedWith("ErrEmptyName()");
  });

  it("should reject decreasing max supply in update", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);

    await expect(
      likeClassOwnerSigner.update({
        name: "Valid Name",
        symbol: "TEST",
        metadata: JSON.stringify({
          name: "Test Collection",
          description: "Test Description",
        }),
        max_supply: 5,
      }),
    ).to.be.rejectedWith("ErrSupplyDecrease");
  });

  it("should reject empty symbol in update", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);

    await expect(
      likeClassOwnerSigner.update({
        name: "Valid Name",
        symbol: "",
        metadata: JSON.stringify({
          name: "Test Collection",
          description: "Test Description",
        }),
        max_supply: 10,
      }),
    ).to.be.rejectedWith("ErrEmptySymbol()");
  });

  it("should not allow zero max supply in update", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);

    await expect(
      likeClassOwnerSigner.update({
        name: "Valid Name",
        symbol: "TEST",
        metadata: JSON.stringify({
          name: "Test Collection",
          description: "Test Description",
        }),
        max_supply: 0,
      }),
    ).to.be.rejectedWith("ErrMaxSupplyZero()");
  });

  it("should not verify metadata JSON", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);

    await expect(
      likeClassOwnerSigner.update({
        name: "Valid Name",
        symbol: "TEST",
        metadata: "invalid json",
        max_supply: 10,
      }),
    ).to.be.not.rejected;
  });
});

describe("BookNFT version", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    this.BookNFTMock = await ethers.getContractFactory("BookNFTMock");
    const [protocolOwner, classOwner, likerLand, randomSigner] =
      await ethers.getSigners();

    this.protocolOwner = protocolOwner;
    this.classOwner = classOwner;
    this.likerLand = likerLand;
    this.randomSigner = randomSigner;
  });

  let deployment: BaseContract;
  let contractAddress: string;
  let protocolContract: BaseContract;
  let nftClassId: string;
  let nftClassContract: BaseContract;
  let v2NFTClassContract: BaseContract;
  beforeEach(async function () {
    const {
      likeProtocol,
      likeProtocolDeployment,
      likeProtocolAddress,
      likeProtocolContract,
      bookNFTDeployment,
      bookNFTAddress,
    } = await createProtocol(this.protocolOwner);

    deployment = likeProtocolDeployment;
    contractAddress = likeProtocolAddress;
    protocolContract = likeProtocolContract;

    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeProtocolOwnerSigner.on(
        "NewBookNFT",
        (id: string, params: any, event: any) => {
          event.removeListener();
          resolve({ id });
        },
      );
      setTimeout(() => {
        reject(new Error("timeout"));
      }, 20000);
    });

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    await likeProtocolOwnerSigner.newBookNFT({
      creator: this.classOwner,
      updaters: [this.classOwner, this.likerLand],
      minters: [this.classOwner, this.likerLand],
      config: bookConfig,
    });

    const newClassEvent = await NewClassEvent;
    nftClassId = newClassEvent.id;
    nftClassContract = await ethers.getContractAt("BookNFT", nftClassId);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);

    // Deploy V2 but not upgrade
    const bookNFTMockOwnerSigner = this.BookNFTMock.connect(this.protocolOwner);
    v2NFTClassContract = await bookNFTMockOwnerSigner.deploy();
  });

  it("should have the correct version on protocol replace implementation", async function () {
    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );
    await likeProtocolOwnerSigner.upgradeTo(v2NFTClassContract.getAddress());
    const version = await v2NFTClassContract.version();
    expect(version).to.equal(2n);
    const beaconProxy = await v2NFTClassContract.attach(nftClassId);
    expect(await beaconProxy.version()).to.equal(2n);
  });

  it("should preserve owner on implementation upgrade", async function () {
    const owner = await nftClassContract.owner();

    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );
    await likeProtocolOwnerSigner.upgradeTo(v2NFTClassContract.getAddress());

    const beaconProxy = await v2NFTClassContract.attach(nftClassId);
    expect(await beaconProxy.owner()).to.equal(owner);
  });

  it("should preserve minters on implementation upgrade", async function () {
    const MINTER_ROLE = await nftClassContract.MINTER_ROLE();
    expect(
      await nftClassContract.hasRole(MINTER_ROLE, this.protocolOwner.address),
    ).to.equal(false);
    expect(
      await nftClassContract.hasRole(MINTER_ROLE, this.classOwner.address),
    ).to.equal(true);
    expect(
      await nftClassContract.hasRole(MINTER_ROLE, this.likerLand.address),
    ).to.equal(true);
    expect(
      await nftClassContract.hasRole(MINTER_ROLE, this.randomSigner.address),
    ).to.equal(false);

    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );
    await likeProtocolOwnerSigner.upgradeTo(v2NFTClassContract.getAddress());

    const beaconProxy = await v2NFTClassContract.attach(nftClassId);
    expect(
      await beaconProxy.hasRole(MINTER_ROLE, this.protocolOwner.address),
    ).to.equal(false);
    expect(
      await beaconProxy.hasRole(MINTER_ROLE, this.classOwner.address),
    ).to.equal(true);
    expect(
      await beaconProxy.hasRole(MINTER_ROLE, this.likerLand.address),
    ).to.equal(true);
    expect(
      await beaconProxy.hasRole(MINTER_ROLE, this.randomSigner.address),
    ).to.equal(false);
  });

  it("should preserve updaters on implementation upgrade", async function () {
    const UPDATER_ROLE = await nftClassContract.UPDATER_ROLE();
    expect(
      await nftClassContract.hasRole(UPDATER_ROLE, this.protocolOwner.address),
    ).to.equal(false);
    expect(
      await nftClassContract.hasRole(UPDATER_ROLE, this.classOwner.address),
    ).to.equal(true);
    expect(
      await nftClassContract.hasRole(UPDATER_ROLE, this.likerLand.address),
    ).to.equal(true);
    expect(
      await nftClassContract.hasRole(UPDATER_ROLE, this.randomSigner.address),
    ).to.equal(false);

    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );
    await likeProtocolOwnerSigner.upgradeTo(v2NFTClassContract.getAddress());

    const beaconProxy = await v2NFTClassContract.attach(nftClassId);
    expect(
      await beaconProxy.hasRole(UPDATER_ROLE, this.protocolOwner.address),
    ).to.equal(false);
    expect(
      await beaconProxy.hasRole(UPDATER_ROLE, this.classOwner.address),
    ).to.equal(true);
    expect(
      await beaconProxy.hasRole(UPDATER_ROLE, this.likerLand.address),
    ).to.equal(true);
    expect(
      await beaconProxy.hasRole(UPDATER_ROLE, this.randomSigner.address),
    ).to.equal(false);
  });

  it("should preserve name, symbol and max supply on implementation upgrade", async function () {
    const originalName = await nftClassContract.name();
    const originalSymbol = await nftClassContract.symbol();
    const originalMaxSupply = await nftClassContract.maxSupply();

    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );
    await likeProtocolOwnerSigner.upgradeTo(v2NFTClassContract.getAddress());

    const beaconProxy = await v2NFTClassContract.attach(nftClassId);
    expect(await beaconProxy.name()).to.equal(originalName);
    expect(await beaconProxy.symbol()).to.equal(originalSymbol);
    expect(await beaconProxy.maxSupply()).to.equal(originalMaxSupply);
  });

  it("should preserve metadata on implementation upgrade", async function () {
    const originalMetadata = await nftClassContract.contractURI();

    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );
    await likeProtocolOwnerSigner.upgradeTo(v2NFTClassContract.getAddress());

    const beaconProxy = await v2NFTClassContract.attach(nftClassId);
    expect(await beaconProxy.contractURI()).to.equal(originalMetadata);
  });
});
