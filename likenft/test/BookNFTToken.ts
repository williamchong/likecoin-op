import { expect } from "chai";
import { EventLog } from "ethers";
import { ethers, upgrades } from "hardhat";

describe("BookNFTToken", () => {
  before(async function () {
    this.LikeProtocol = await ethers.getContractFactory("LikeProtocol");
    const [protocolOwner, classOwner, likerLand, randomSigner, randomSigner2] =
      await ethers.getSigners();

    this.protocolOwner = protocolOwner;
    this.classOwner = classOwner;
    this.likerLand = likerLand;
    this.randomSigner = randomSigner;
    this.randomSigner2 = randomSigner2;
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
      }, 20000);
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

    const likeNFTClassOwnerSigner = nftClassContract.connect(this.classOwner);
    await likeNFTClassOwnerSigner
      .mint(this.classOwner, [
        JSON.stringify({
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
      ])
      .then((tx) => tx.wait());
  });

  it("owner should be able to send once", async function () {
    const likeNFTClassOwnerSigner = nftClassContract.connect(this.classOwner);
    await expect(
      likeNFTClassOwnerSigner
        .transferWithMemo(this.classOwner, this.randomSigner, 0, "memo1")
        .then((tx) => tx.wait()),
    ).to.be.not.rejected;
    await expect(
      likeNFTClassOwnerSigner
        .transferWithMemo(this.classOwner, this.randomSigner2, 0, "memo1fails")
        .then((tx) => tx.wait()),
    ).to.be.rejectedWith(/ERC721InsufficientApproval/);

    const filters = likeNFTClassOwnerSigner.filters.TransferWithMemo(
      null,
      null,
      0,
    );
    const logs1 = await likeNFTClassOwnerSigner.queryFilter(filters);
    expect((logs1[0] as EventLog).args[3]).to.equal("memo1");
  });

  it("should not able to send with random signer", async function () {
    const likeNFTClassRandomSigner = nftClassContract.connect(
      this.randomSigner,
    );
    await expect(
      likeNFTClassRandomSigner
        .transferWithMemo(this.classOwner, this.randomSigner, 0, "memo1")
        .then((tx) => tx.wait()),
    ).to.be.rejectedWith(/ERC721InsufficientApproval/);

    expect(await likeNFTClassRandomSigner.ownerOf(0)).to.equal(
      this.classOwner.address,
    );
  });

  it("should be able to send with memo", async function () {
    const likeNFTClassOwnerSigner = nftClassContract.connect(this.classOwner);
    const likeNFTClassRandomSigner = nftClassContract.connect(
      this.randomSigner,
    );
    await expect(
      likeNFTClassOwnerSigner
        .transferWithMemo(this.classOwner, this.randomSigner, 0, "memo1")
        .then((tx) => tx.wait()),
    ).to.be.not.rejected;
    await expect(
      likeNFTClassRandomSigner
        .transferWithMemo(this.randomSigner, this.classOwner, 0, "memo2")
        .then((tx) => tx.wait()),
    ).to.be.not.rejected;
    await expect(
      likeNFTClassRandomSigner
        .transferWithMemo(this.randomSigner, this.classOwner, 0, "memo2fails")
        .then((tx) => tx.wait()),
    ).to.be.rejectedWith(/ERC721InsufficientApproval/);

    const filters2 = likeNFTClassOwnerSigner.filters.TransferWithMemo(
      null,
      null,
      0,
    );
    const logs2 = await likeNFTClassOwnerSigner.queryFilter(filters2);
    expect((logs2[0] as EventLog).args[0]).to.equal(this.classOwner.address);
    expect((logs2[0] as EventLog).args[1]).to.equal(this.randomSigner.address);
    expect((logs2[0] as EventLog).args[2]).to.equal(0n);
    expect((logs2[0] as EventLog).args[3]).to.equal("memo1");
    expect((logs2[1] as EventLog).args[0]).to.equal(this.randomSigner.address);
    expect((logs2[1] as EventLog).args[1]).to.equal(this.classOwner.address);
    expect((logs2[1] as EventLog).args[2]).to.equal(0n);
    expect((logs2[1] as EventLog).args[3]).to.equal("memo2");
  });
});
