import { expect } from "chai";
import { EventLog, BaseContract } from "ethers";
import { ethers, upgrades } from "hardhat";
import { BookConfigLoader } from "./BookConfigLoader";

describe("BookNFTClass", () => {
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
    const likeProtocol = await upgrades.deployProxy(
      this.LikeProtocol,
      [this.protocolOwner.address],
      {
        initializer: "initialize",
      },
    );
    deployment = await likeProtocol.waitForDeployment();
    contractAddress = await deployment.getAddress();
    protocolContract = await ethers.getContractAt(
      "LikeProtocol",
      contractAddress,
    );

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
        .mint(this.classOwner.address, [
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
        ])
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
        .mint(this.classOwner.address, [
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
        ])
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
        .mint(this.classOwner.address, [
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
        ])
        .then((tx) => tx.wait());
    };

    const safeMintWithTokenId = async (fromTokenId: number) => {
      await likeClassOwnerSigner
        .safeMintWithTokenId(fromTokenId, this.classOwner.address, [
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
        ])
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
    const likeProtocol = await upgrades.deployProxy(
      this.LikeProtocol,
      [this.protocolOwner.address],
      {
        initializer: "initialize",
      },
    );
    deployment = await likeProtocol.waitForDeployment();
    contractAddress = await deployment.getAddress();
    protocolContract = await ethers.getContractAt(
      "LikeProtocol",
      contractAddress,
    );
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

  it("should not allow protocol owner to mint NFT", async function () {
    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );

    const mintNFT = async () => {
      await likeProtocolOwnerSigner
        .mintNFT({
          to: this.classCreatorSigner.address,
          classId: this.classId,
          input: {
            metadata: JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "#0001 Description",
              name: "#0001",
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
          },
        })
        .then((tx) => tx.wait());
    };

    await expect(mintNFT()).to.be.rejected;
  });

  it("should allow class owner to mint NFT", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);

    const mintNFT = async () => {
      await likeClassOwnerSigner
        .mint(this.likerLand.address, [
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
        ])
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
    const likeProtocol = await upgrades.deployProxy(
      this.LikeProtocol,
      [this.protocolOwner.address],
      {
        initializer: "initialize",
      },
    );
    deployment = await likeProtocol.waitForDeployment();
    contractAddress = await deployment.getAddress();
    protocolContract = await ethers.getContractAt(
      "LikeProtocol",
      contractAddress,
    );

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

  it("should not allow protocol owner to mint NFT", async function () {
    const likeProtocolOwnerSigner = protocolContract.connect(
      this.protocolOwner,
    );

    const mintNFT = async () => {
      await likeProtocolOwnerSigner
        .mintNFT({
          creator: this.classCreatorSigner.address,
          classId: this.classId,
          input: {
            metadata: JSON.stringify({
              image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
              image_data: "",
              external_url: "https://www.google.com",
              description: "#0001 Description",
              name: "#0001",
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
          },
        })
        .then((tx) => tx.wait());
    };

    await expect(mintNFT()).to.be.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(0n);
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
        .mint(this.classOwner.address, [
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
        ])
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
        .mint(this.randomSigner.address, [
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
        ])
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
    const likeProtocol = await upgrades.deployProxy(
      this.LikeProtocol,
      [this.protocolOwner.address],
      {
        initializer: "initialize",
      },
    );
    deployment = await likeProtocol.waitForDeployment();
    contractAddress = await deployment.getAddress();
    protocolContract = await ethers.getContractAt(
      "LikeProtocol",
      contractAddress,
    );

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
