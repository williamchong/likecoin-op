import { expect } from "chai";
import { ethers, upgrades } from "hardhat";

describe("LikeNFT", () => {
  before(async function () {
    this.LikeNFT = await ethers.getContractFactory("LikeNFT");
    const [ownerSigner, signer1] = await ethers.getSigners();

    this.ownerSigner = ownerSigner;
    this.signer1 = signer1;
  });

  beforeEach(async function () {
    const likeNFT = await upgrades.deployProxy(
      this.LikeNFT,
      [this.ownerSigner.address],
      {
        initializer: "initialize",
      },
    );
    const deployment = await likeNFT.waitForDeployment();
    this.contractAddress = await deployment.getAddress();
  });

  it("should be able to pause", async function () {
    const LikeNFTOwnerSigner = await ethers.getContractFactory("LikeNFT", {
      signer: this.ownerSigner,
    });
    const likeNFTOwnerSigner = LikeNFTOwnerSigner.attach(this.contractAddress);

    const classOperation = async () => {
      await likeNFTOwnerSigner
        .newClass({
          creator: this.ownerSigner,
          parent: {
            type_: 1,
            iscn_id_prefix:
              "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
          },
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
    await expect(likeNFTOwnerSigner.pause()).to.be.not.rejected;
    await expect(classOperation()).to.be.rejectedWith(
      "VM Exception while processing transaction: reverted with custom error 'EnforcedPause()'",
    );
    await expect(likeNFTOwnerSigner.unpause()).to.be.not.rejected;
    await expect(classOperation()).to.be.not.rejected;
  });
  it("should be able to create new class", async function () {
    const LikeNFTOwnerSigner = await ethers.getContractFactory("LikeNFT", {
      signer: this.ownerSigner,
    });
    const likeNFTOwnerSigner = LikeNFTOwnerSigner.attach(this.contractAddress);

    const newClass = async () => {
      await likeNFTOwnerSigner
        .newClass({
          creator: this.ownerSigner,
          parent: {
            type_: 1,
            iscn_id_prefix:
              "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
          },
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

    await expect(newClass()).to.be.not.rejected;
    await expect(newClass()).to.be.not.rejected;
  });
});

describe("LikeNFT class operations", () => {
  before(async function () {
    this.LikeNFT = await ethers.getContractFactory("LikeNFT");
    const [ownerSigner, signer1] = await ethers.getSigners();

    this.ownerSigner = ownerSigner;
    this.signer1 = signer1;
  });

  beforeEach(async function () {
    const likeNFT = await upgrades.deployProxy(
      this.LikeNFT,
      [this.ownerSigner.address],
      {
        initializer: "initialize",
      },
    );
    const deployment = await likeNFT.waitForDeployment();
    this.contractAddress = await deployment.getAddress();

    const LikeNFTOwnerSigner = await ethers.getContractFactory("LikeNFT", {
      signer: this.ownerSigner,
    });
    const likeNFTOwnerSigner = LikeNFTOwnerSigner.attach(this.contractAddress);

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      likeNFTOwnerSigner.on("NewClass", (id, params, event) => {
        event.removeListener();
        resolve({ id });
      });

      setTimeout(() => {
        reject(new Error("timeout"));
      }, 60000);
    });

    likeNFTOwnerSigner
      .newClass({
        creator: this.ownerSigner,
        parent: {
          type_: 1,
          iscn_id_prefix:
            "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
        },
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
    const LikeNFTOwnerSigner = await ethers.getContractFactory("LikeNFT", {
      signer: this.ownerSigner,
    });
    const likeNFTOwnerSigner = LikeNFTOwnerSigner.attach(this.contractAddress);

    const updateClass = async () => {
      await likeNFTOwnerSigner
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
      await likeNFTOwnerSigner
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
    await expect(updateClass()).to.be.rejectedWith(
      "VM Exception while processing transaction: reverted with custom error 'ErrCannotUpdateClassWithMintedTokens()'",
    );
  });

  it("should be able to mint class", async function () {
    const LikeNFTOwnerSigner = await ethers.getContractFactory("LikeNFT", {
      signer: this.ownerSigner,
    });
    const likeNFTOwnerSigner = LikeNFTOwnerSigner.attach(this.contractAddress);
    const mintNFT = async () => {
      await likeNFTOwnerSigner
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
});
