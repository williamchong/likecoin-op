import { expect } from "chai";
import { EventLog } from "ethers";
import { ethers, upgrades } from "hardhat";

describe("LikeNFT class operations", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    const [ownerSigner, signer1] = await ethers.getSigners();

    this.ownerSigner = ownerSigner;
    this.signer1 = signer1;
  });

  beforeEach(async function () {
    const likeProtocol = await upgrades.deployProxy(
      this.LikeProtocol,
      [this.ownerSigner.address],
      {
        initializer: "initialize",
      },
    );
    const deployment = await likeProtocol.waitForDeployment();
    this.contractAddress = await deployment.getAddress();

    const LikeProtocolOwnerSigner = await ethers.getContractFactory("LikeProtocol", {
      signer: this.ownerSigner,
    });
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(this.contractAddress);

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

    const newClassEvent = await NewClassEvent;
    this.classId = newClassEvent.id;
  });

  it("should be able to update class", async function () {
    const LikeProtocolOwnerSigner = await ethers.getContractFactory("LikeProtocol", {
      signer: this.ownerSigner,
    });
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(this.contractAddress);

    const updateClass = async () => {
      await likeProtocolOwnerSigner
        .updateClass({
          creator: this.ownerSigner,
          class_id: this.classId,
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
      await likeProtocolOwnerSigner
        .mintNFT({
          creator: this.ownerSigner,
          class_id: this.classId,
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

    await expect(updateClass()).to.be.not.rejected;
    await expect(mintNFT()).to.be.not.rejected;
    await expect(updateClass()).to.be.not.rejected;
  });

  it("should be able to mint class", async function () {
    const LikeProtocolOwnerSigner = await ethers.getContractFactory("LikeProtocol", {
      signer: this.ownerSigner,
    });
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(this.contractAddress);
    const mintNFT = async () => {
      await likeProtocolOwnerSigner
        .mintNFT({
          creator: this.ownerSigner,
          class_id: this.classId,
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

  it("should be able to multiple mint", async function () {
    const LikeProtocolOwnerSigner = await ethers.getContractFactory("LikeProtocol", {
      signer: this.ownerSigner,
    });
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(this.contractAddress);
    const mintNFT = async () => {
      await likeProtocolOwnerSigner
        .mintNFTs({
          creator: this.ownerSigner,
          class_id: this.classId,
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

describe("LikeNFT mint by likenft and creator", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    const [likeProtocolOwnerSigner, classCreatorSigner] = await ethers.getSigners();

    this.likeProtocolOwnerSigner = likeProtocolOwnerSigner;
    this.classCreatorSigner = classCreatorSigner;
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
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(this.contractAddress);

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

  it("should be able to mint by likenft", async function () {
    const LikeProtocolOwnerSigner = await ethers.getContractFactory(
      "LikeProtocol",
      {
        signer: this.likeProtocolOwnerSigner,
      },
    );
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(this.contractAddress);

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

  it("should be able to mint by creator", async function () {
    const ClassContractWithCreatorSigner = await ethers.getContractFactory(
      "Class",
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

describe("LikeNFT mint by likenft and creator after transfer", () => {
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
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(this.contractAddress);

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
      "Class",
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

  it("should be able to mint by likenft", async function () {
    const LikeProtocolOwnerSigner = await ethers.getContractFactory(
      "LikeProtocol",
      {
        signer: this.likeProtocolOwnerSigner,
      },
    );
    const likeProtocolOwnerSigner = LikeProtocolOwnerSigner.attach(this.contractAddress);

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

  it("should not be able to mint by creator", async function () {
    const ClassContractWithCreatorSigner = await ethers.getContractFactory(
      "Class",
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

  it("should be able to mint by next owner", async function () {
    const ClassContractWithNextOwnerSigner = await ethers.getContractFactory(
      "Class",
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