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
      creator: this.classOwner.address,
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

  it("should have the right minter", async function () {
    // FIXME: This is not working
    return;
    const MINTER_ROLE = await nftClassContract.MINTER_ROLE();
    expect(
      await nftClassContract.hasRole(MINTER_ROLE, this.classOwner.address),
    ).to.equal(true);
    expect(
      await nftClassContract.hasRole(MINTER_ROLE, this.protocolOwner.address),
    ).to.equal(true);
    // expect(await nftClassContract.hasRole(MINTER_ROLE, contractAddress)).to.equal(true);
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
          creator: this.classOwner.address,
          class_id: nftClassId,
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
          creator: this.classOwner.address,
          class_id: nftClassId,
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
    // FIXME: This is not working
    return;
    const likeProtocolRandomSigner = protocolContract.connect(
      this.randomSigner,
    );

    await likeProtocolRandomSigner.updateClass({
      creator: this.randomSigner,
      class_id: nftClassId,
      input: {
        name: "Hi Jack",
        symbol: "HIJACK",
        metadata: JSON.stringify({}),
        config: {
          max_supply: 0,
        },
      },
    });

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
          creator: this.classOwner,
          class_id: nftClassId,
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
          creator: this.classOwner,
          class_id: nftClassId,
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
  });
});

describe("LikeNFTClass permission control", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    const [likeProtocolOwnerSigner, classCreatorSigner, randomSigner] =
      await ethers.getSigners();

    this.likeProtocolOwnerSigner = likeProtocolOwnerSigner;
    this.classCreatorSigner = classCreatorSigner;
    this.randomSigner = randomSigner;
  });

  beforeEach(async function () {
    const likeProtocol = await upgrades.deployProxy(
      this.LikeProtocol,
      [this.likeProtocolOwnerSigner.address],
      {
        initializer: "initialize",
      },
    );
    const deployment = await likeProtocol.waitForDeployment();
    this.contractAddress = await deployment.getAddress();

    const LikeProtocolOwnerSigner = await ethers.getContractFactory(
      "LikeProtocol",
      {
        signer: this.likeProtocolOwnerSigner,
      },
    );
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(
      this.contractAddress,
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
        creator: this.classCreatorSigner.address,
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
    this.classId = newClassEvent.id;
  });

  it("should allow protocol owner to mint NFT", async function () {
    const LikeProtocolOwnerSigner = await ethers.getContractFactory(
      "LikeProtocol",
      {
        signer: this.likeProtocolOwnerSigner,
      },
    );
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(
      this.contractAddress,
    );

    const mintNFT = async () => {
      await likeProtocolOwnerSigner
        .mintNFT({
          creator: this.classCreatorSigner.address,
          class_id: this.classId,
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

    await expect(mintNFT()).to.be.not.rejected;
  });

  it("should allow creator to mint NFT", async function () {
    const ClassContractWithCreatorSigner = await ethers.getContractFactory(
      "LikeNFTClass",
      {
        signer: this.classCreatorSigner,
      },
    );
    const classContractWithCreatorSigner =
      ClassContractWithCreatorSigner.attach(this.classId);

    const mintNFT = async () => {
      await classContractWithCreatorSigner
        .mint(this.classCreatorSigner.address, [
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
  });
});

describe("LikeNFTClass ownership transfer", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    const [likeProtocolOwnerSigner, classCreatorSigner, nextClassOwnerSigner] =
      await ethers.getSigners();

    this.likeProtocolOwnerSigner = likeProtocolOwnerSigner;
    this.classCreatorSigner = classCreatorSigner;
    this.nextClassOwnerSigner = nextClassOwnerSigner;
  });

  beforeEach(async function () {
    const likeNFT = await upgrades.deployProxy(
      this.LikeProtocol,
      [this.likeProtocolOwnerSigner.address],
      {
        initializer: "initialize",
      },
    );
    const deployment = await likeNFT.waitForDeployment();
    this.contractAddress = await deployment.getAddress();

    const LikeProtocolOwnerSigner = await ethers.getContractFactory(
      "LikeProtocol",
      {
        signer: this.likeProtocolOwnerSigner,
      },
    );
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(
      this.contractAddress,
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
        creator: this.classCreatorSigner.address,
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
    this.classId = newClassEvent.id;

    const ClassContractWithCreatorSigner = await ethers.getContractFactory(
      "LikeNFTClass",
      {
        signer: this.classCreatorSigner,
      },
    );
    const classContractWithCreatorSigner =
      ClassContractWithCreatorSigner.attach(this.classId);

    const transferOwnership = async () => {
      await classContractWithCreatorSigner
        .transferOwnership(this.nextClassOwnerSigner)
        .then((tx) => tx.wait());
    };

    expect(transferOwnership()).to.not.be.rejected;
  });

  it("should allow protocol owner to mint NFT", async function () {
    const LikeProtocolOwnerSigner = await ethers.getContractFactory(
      "LikeProtocol",
      {
        signer: this.likeProtocolOwnerSigner,
      },
    );
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(
      this.contractAddress,
    );

    const mintNFT = async () => {
      await likeProtocolOwnerSigner
        .mintNFT({
          creator: this.classCreatorSigner.address,
          class_id: this.classId,
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

    await expect(mintNFT()).to.be.not.rejected;
  });

  it("should not allow creator to mint NFT", async function () {
    const ClassContractWithCreatorSigner = await ethers.getContractFactory(
      "LikeNFTClass",
      {
        signer: this.classCreatorSigner,
      },
    );
    const classContractWithCreatorSigner =
      ClassContractWithCreatorSigner.attach(this.classId);

    const mintNFT = async () => {
      await classContractWithCreatorSigner
        .mint(this.classCreatorSigner.address, [
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

    await expect(mintNFT()).to.be.rejectedWith(
      "VM Exception while processing transaction: reverted with custom error 'ErrUnauthorized()'",
    );
  });

  it("should allow next owner to mint NFT", async function () {
    const ClassContractWithNextOwnerSigner = await ethers.getContractFactory(
      "LikeNFTClass",
      {
        signer: this.nextClassOwnerSigner,
      },
    );
    const classContractWithNextOwnerSigner =
      ClassContractWithNextOwnerSigner.attach(this.classId);

    const mintNFT = async () => {
      await classContractWithNextOwnerSigner
        .mint(this.nextClassOwnerSigner.address, [
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

    await expect(mintNFT()).to.not.be.rejected;
  });
});
