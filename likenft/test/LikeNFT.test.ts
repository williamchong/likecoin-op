import { expect } from "chai";
import { EventLog } from "ethers";
import { ethers, upgrades } from "hardhat";


describe("LikeNFT token operations", () => {
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

    await likeProtocolOwnerSigner
      .mintNFT({
        creator: this.ownerSigner,
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
  });

  it("should be able to send", async function () {
    const ClassOwnerSigner = await ethers.getContractFactory("Class", {
      signer: this.ownerSigner,
    });
    const classOwnerSigner = ClassOwnerSigner.attach(this.classId);
    await expect(
      classOwnerSigner
        .transferWithMemo(this.ownerSigner, this.signer1, 0, "memo1")
        .then((tx) => tx.wait()),
    ).to.be.not.rejected;
    await expect(
      classOwnerSigner
        .transferWithMemo(this.ownerSigner, this.signer1, 0, "memo1")
        .then((tx) => tx.wait()),
    ).to.be.rejectedWith(
      "VM Exception while processing transaction: reverted with custom error 'TransferFromIncorrectOwner()'",
    );

    const filters = classOwnerSigner.filters.TransferWithMemo(null, null, 0);
    const logs1 = await classOwnerSigner.queryFilter(filters);
    expect((logs1[0] as EventLog).args[3]).to.equal("memo1");

    const ClassSigner1 = await ethers.getContractFactory("Class", {
      signer: this.signer1,
    });
    const classSigner1 = ClassSigner1.attach(this.classId);
    await expect(
      classSigner1
        .transferWithMemo(this.signer1, this.ownerSigner, 0, "memo2")
        .then((tx) => tx.wait()),
    ).to.be.not.rejected;
    await expect(
      classSigner1
        .transferWithMemo(this.signer1, this.ownerSigner, 0, "memo2")
        .then((tx) => tx.wait()),
    ).to.be.rejectedWith(
      "VM Exception while processing transaction: reverted with custom error 'TransferFromIncorrectOwner()'",
    );

    const logs2 = await classOwnerSigner.queryFilter(filters);
    expect((logs2[0] as EventLog).args[0]).to.equal(this.ownerSigner.address);
    expect((logs2[0] as EventLog).args[1]).to.equal(this.signer1.address);
    expect((logs2[0] as EventLog).args[2]).to.equal(0n);
    expect((logs2[0] as EventLog).args[3]).to.equal("memo1");
    expect((logs2[1] as EventLog).args[0]).to.equal(this.signer1.address);
    expect((logs2[1] as EventLog).args[1]).to.equal(this.ownerSigner.address);
    expect((logs2[1] as EventLog).args[2]).to.equal(0n);
    expect((logs2[1] as EventLog).args[3]).to.equal("memo2");
  });
});
