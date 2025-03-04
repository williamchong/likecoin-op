import { expect } from "chai";
import { EventLog, BaseContract } from "ethers";
import { ethers, upgrades } from "hardhat";

describe("LikeNFTClass", () => {
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
        "NewClass",
        (id: string, params: any, event: any) => {
          event.removeListener();
          resolve({ id });
        },
      );
      setTimeout(() => {
        reject(new Error("timeout"));
      }, 20000);
    });

    await likeProtocolOwnerSigner.newClass({
      creator: this.classOwner,
      updaters: [this.classOwner, this.likerLand],
      minters: [this.classOwner, this.likerLand],
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
    });

    const newClassEvent = await NewClassEvent;
    nftClassId = newClassEvent.id;
    nftClassContract = await ethers.getContractAt("LikeNFTClass", nftClassId);
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

  it("should allow class owner to update class and mint NFT", async function () {
    const likeClassOwnerSigner = nftClassContract.connect(this.classOwner);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);

    const updateClass = async () => {
      await likeClassOwnerSigner
        .update({
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
    await expect(await nftClassContract.totalSupply()).to.equal(0n);
    await expect(mintNFT()).to.be.not.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(1n);
    await expect(updateClass()).to.be.not.rejected;
  });

  it("should allow class owner to update class and mint NFT via protocol contract", async function () {
    const likeProtocolClassOwnerSigner = protocolContract.connect(
      this.classOwner,
    );
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);

    const updateClass = async () => {
      await likeProtocolClassOwnerSigner
        .updateClass({
          classId: nftClassId,
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

    const mintNFT = async () => {
      await likeProtocolClassOwnerSigner
        .mintNFT({
          to: this.classOwner.address,
          classId: nftClassId,
          input: {
            metadata: JSON.stringify({
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
            }),
          },
        })
        .then((tx) => tx.wait());
    };

    await expect(updateClass()).to.be.not.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(0n);
    await expect(mintNFT()).to.be.not.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(1n);
    await expect(updateClass()).to.be.not.rejected;
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

  it("should not allow random address update class via protocol contract", async function () {
    const likeProtocolRandomSigner = protocolContract.connect(
      this.randomSigner,
    );

    await expect(
      likeProtocolRandomSigner.updateClass({
        classId: nftClassId,
        input: {
          name: "Hi Jack",
          symbol: "HIJACK",
          metadata: JSON.stringify({}),
          config: {
            max_supply: 0,
          },
        },
      }),
    ).to.be.rejected;

    await expect(await nftClassContract.owner()).to.equal(
      this.classOwner.address,
    );
    await expect(await nftClassContract.symbol()).to.equal("KOOB");
  });

  it("should allow class owner to mint NFT via protocol contract", async function () {
    const likeProtocolClassOwnerSigner = protocolContract.connect(
      this.classOwner,
    );
    const mintNFT = async () => {
      await likeProtocolClassOwnerSigner
        .mintNFT({
          to: this.classOwner,
          classId: nftClassId,
          input: {
            metadata: JSON.stringify({
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
          },
        })
        .then((tx) => tx.wait());
    };
    await expect(mintNFT()).to.be.not.rejected;
  });

  it("should allow class owner to mintNFTs in batch", async function () {
    const likeClassOwnerSigner = protocolContract.connect(this.classOwner);
    const mintNFT = async () => {
      await likeClassOwnerSigner
        .mintNFTs({
          to: this.classOwner,
          classId: nftClassId,
          inputs: [
            {
              metadata: JSON.stringify({
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
            },
            {
              metadata: JSON.stringify({
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
            },
          ],
        })
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

  it("should check token id when safe mint with token id via protocol contract", async function () {
    const likeProtocolOwnerSigner = protocolContract.connect(this.classOwner);
    expect(await nftClassContract.owner()).to.equal(this.classOwner.address);

    const mintNFT = async () => {
      await likeProtocolOwnerSigner
        .mintNFT({
          to: this.classOwner.address,
          classId: nftClassId,
          input: {
            metadata: JSON.stringify({
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
            }),
          },
        })
        .then((tx) => tx.wait());
    };

    const safeMintNFTsWithTokenId = async (fromTokenId: number) => {
      await likeProtocolOwnerSigner
        .safeMintNFTsWithTokenId({
          to: this.classOwner.address,
          classId: nftClassId,
          fromTokenId,
          inputs: [
            {
              metadata: JSON.stringify({
                image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
                image_data: "",
                external_url: "https://www.google.com",
                description: "202412191729 #0001 Description",
                name: "202412191729 #0002",
                attributes: [
                  {
                    trait_type: "ISCN ID",
                    value:
                      "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                  },
                ],
              }),
            },
            {
              metadata: JSON.stringify({
                image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
                image_data: "",
                external_url: "https://www.google.com",
                description: "202412191729 #0001 Description",
                name: "202412191729 #0003",
                attributes: [
                  {
                    trait_type: "ISCN ID",
                    value:
                      "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
                  },
                ],
              }),
            },
          ],
        })
        .then((tx) => tx.wait());
    };

    await expect(mintNFT()).to.be.not.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(1n);
    await expect(safeMintNFTsWithTokenId(0)).to.be.rejectedWith(
      "ErrTokenIdMintFails(1)",
    );
    await expect(safeMintNFTsWithTokenId(1)).to.be.not.rejected;
    await expect(await nftClassContract.totalSupply()).to.equal(3n);
    await expect(safeMintNFTsWithTokenId(1)).to.be.rejectedWith(
      "ErrTokenIdMintFails(3)",
    );
  });
});

describe("LikeNFTClass permission control", () => {
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
      likeProtocolOwnerSigner.on("NewClass", (id, params, event) => {
        event.removeListener();
        resolve({ id });
      });

      setTimeout(() => {
        reject(new Error("timeout"));
      }, 60000);
    });

    likeProtocolOwnerSigner
      .newClass({
        creator: this.classOwner,
        updaters: [this.classOwner, this.likerLand],
        minters: [this.classOwner, this.likerLand],
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

    const newClassEvent = await NewClassEvent;
    nftClassId = newClassEvent.id;
    nftClassContract = await ethers.getContractAt("LikeNFTClass", nftClassId);
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

describe("LikeNFTClass ownership transfer", () => {
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
      likeProtocolOwnerSigner.on("NewClass", (id, params, event) => {
        event.removeListener();
        resolve({ id });
      });

      setTimeout(() => {
        reject(new Error("timeout"));
      }, 60000);
    });

    likeProtocolOwnerSigner
      .newClass({
        creator: this.classOwner,
        updaters: [this.classOwner, this.likerLand],
        minters: [this.classOwner, this.likerLand],
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

    const newClassEvent = await NewClassEvent;
    nftClassId = newClassEvent.id;
    nftClassContract = await ethers.getContractAt("LikeNFTClass", nftClassId);
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
