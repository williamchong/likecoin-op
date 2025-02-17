import { expect } from "chai";
import { EventLog } from "ethers";
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
    this.deployment = deployment;
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

  it("should be allow everyone to create new class", async function () {
    const likeNFTSigner = this.deployment.attach(this.signer1.address);

    const newClass = async () => {
      await likeNFTSigner
        .newClass({
        creator: this.signer1,
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
