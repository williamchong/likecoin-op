import { expect } from "chai";
import { loadFixture } from "@nomicfoundation/hardhat-network-helpers";
import { viem } from "hardhat";
import { deployProtocol } from "./factory";
import "./setup";

import { BookConfigLoader, BookTokenConfigLoader } from "./BookConfigLoader";

describe("BookNFTToken", () => {
  async function initBookNFTToken() {
    const {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      randomSigner2,
      publicClient,
    } = await loadFixture(deployProtocol);

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      const unwatch = likeProtocol.watchEvent.NewBookNFT({
        onLogs: (logs) => {
          const id = logs[0].args.bookNFT;
          if (id) {
            unwatch();
            resolve(id);
          }
        },
      });
    });

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    await likeProtocol.write.newBookNFT([
      {
        creator: classOwner.account.address,
        updaters: [classOwner.account.address, likerLand.account.address],
        minters: [classOwner.account.address, likerLand.account.address],
        config: bookConfig,
      },
    ]);

    const nftClassId = await NewClassEvent;
    const nftClassContract = await viem.getContractAt("BookNFT", nftClassId);

    const tokenConfig0 = BookTokenConfigLoader.load(
      "./test/fixtures/TokenConfig0.json",
    );
    await nftClassContract.write.mint(
      [classOwner.account.address, ["_mint1"], [tokenConfig0]],
      {
        account: classOwner.account,
      },
    );

    return {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      randomSigner2,
      publicClient,
      nftClassId,
      nftClassContract,
    };
  }

  it("should allow updater to update token metadata", async function () {
    const { nftClassContract, classOwner, likerLand } =
      await loadFixture(initBookNFTToken);
    const TokenConfig0 = BookTokenConfigLoader.load(
      "./test/fixtures/TokenConfig0.json",
    );
    const tokenConfig1 = BookTokenConfigLoader.load(
      "./test/fixtures/TokenConfig1.json",
    );
    expect(await nftClassContract.read.tokenURI([0])).to.equal(
      `data:application/json;base64,${Buffer.from(TokenConfig0).toString("base64")}`,
    );
    await nftClassContract.write.updateTokenMetadata([0, tokenConfig1], {
      account: likerLand.account,
    });
    expect(await nftClassContract.read.tokenURI([0])).to.equal(
      `data:application/json;base64,${Buffer.from(tokenConfig1).toString("base64")}`,
    );
  });

  it("should not allow random signer to update token metadata", async function () {
    const { nftClassContract, randomSigner } =
      await loadFixture(initBookNFTToken);
    const TokenConfig0 = BookTokenConfigLoader.load(
      "./test/fixtures/TokenConfig0.json",
    );
    const tokenConfig1 = BookTokenConfigLoader.load(
      "./test/fixtures/TokenConfig1.json",
    );
    expect(await nftClassContract.read.tokenURI([0])).to.equal(
      `data:application/json;base64,${Buffer.from(TokenConfig0).toString("base64")}`,
    );
    await expect(
      nftClassContract.write.updateTokenMetadata([0, tokenConfig1], {
        account: randomSigner.account,
      }),
    ).to.be.rejectedWith(/ErrUnauthorized()/);
    expect(await nftClassContract.read.tokenURI([0])).to.equal(
      `data:application/json;base64,${Buffer.from(TokenConfig0).toString("base64")}`,
    );
  });

  it("owner should be able to send once", async function () {
    const { nftClassContract, classOwner, randomSigner, randomSigner2 } =
      await loadFixture(initBookNFTToken);
    await expect(
      nftClassContract.write.transferWithMemo(
        [classOwner.account.address, randomSigner.account.address, 0, "memo1"],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.not.rejected;
    expect(await nftClassContract.read.ownerOf([0])).to.equalAddress(
      randomSigner.account.address,
    );
    const logs1 = await nftClassContract.getEvents.TransferWithMemo();
    expect(logs1).to.have.lengthOf(1);
    expect(logs1[0].args.memo).to.equal("memo1");

    await expect(
      nftClassContract.write.transferWithMemo(
        [
          classOwner.account.address,
          randomSigner2.account.address,
          0,
          "memo1fails",
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.rejectedWith(/ERC721InsufficientApproval/);
  });

  it("should not able to send with random signer", async function () {
    const { nftClassContract, classOwner, randomSigner } =
      await loadFixture(initBookNFTToken);
    await expect(
      nftClassContract.write.transferWithMemo(
        [classOwner.account.address, randomSigner.account.address, 0, "memo1"],
        {
          account: randomSigner.account,
        },
      ),
    ).to.be.rejectedWith(/ERC721InsufficientApproval/);

    expect(await nftClassContract.read.ownerOf([0])).to.equalAddress(
      classOwner.account.address,
    );
  });

  it("should be able to send with memo", async function () {
    const { nftClassContract, classOwner, randomSigner } =
      await loadFixture(initBookNFTToken);
    await expect(
      nftClassContract.write.transferWithMemo(
        [classOwner.account.address, randomSigner.account.address, 0, "memo1"],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.not.rejected;
    const logs1 = await nftClassContract.getEvents.TransferWithMemo();
    expect(logs1).to.have.lengthOf(1);
    expect(logs1[0].args.from).to.equalAddress(classOwner.account.address);
    expect(logs1[0].args.to).to.equalAddress(randomSigner.account.address);
    expect(logs1[0].args.tokenId).to.equal(0n);
    expect(logs1[0].args.memo).to.equal("memo1");

    await expect(
      nftClassContract.write.transferWithMemo(
        [randomSigner.account.address, classOwner.account.address, 0, "memo2"],
        {
          account: randomSigner.account,
        },
      ),
    ).to.be.not.rejected;
    const logs2 = await nftClassContract.getEvents.TransferWithMemo();
    expect(logs2).to.have.lengthOf(1);
    expect(logs2[0].args.from).to.equalAddress(randomSigner.account.address);
    expect(logs2[0].args.to).to.equalAddress(classOwner.account.address);
    expect(logs2[0].args.tokenId).to.equal(0n);
    expect(logs2[0].args.memo).to.equal("memo2");

    await expect(
      nftClassContract.write.transferWithMemo(
        [
          randomSigner.account.address,
          classOwner.account.address,
          0,
          "memo2fails",
        ],
        {
          account: randomSigner.account,
        },
      ),
    ).to.be.rejectedWith(/ERC721InsufficientApproval/);
  });
});

describe("BookNFTToken batch actions", () => {
  async function initBookNFTTokenBatch() {
    const {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      randomSigner2,
      publicClient,
    } = await loadFixture(deployProtocol);

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      const unwatch = likeProtocol.watchEvent.NewBookNFT({
        onLogs: (logs) => {
          const id = logs[0].args.bookNFT;
          if (id) {
            unwatch();
            resolve(id);
          }
        },
      });
    });

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    await likeProtocol.write.newBookNFT([
      {
        creator: classOwner.account.address,
        updaters: [classOwner.account.address, likerLand.account.address],
        minters: [classOwner.account.address, likerLand.account.address],
        config: bookConfig,
      },
    ]);

    const nftClassId = await NewClassEvent;
    const nftClassContract = await viem.getContractAt("BookNFT", nftClassId);

    await nftClassContract.write.mint(
      [
        classOwner.account.address,
        ["_mint1"],
        [
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
        ],
      ],
      {
        account: classOwner.account,
      },
    );
    await nftClassContract.write.mint(
      [
        classOwner.account.address,
        ["_mint2"],
        [
          JSON.stringify({
            image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
            image_data: "",
            external_url: "https://www.google.com",
            description: "#0002 Description",
            name: "#0002",
            attributes: [
              {
                trait_type: "ISCN ID",
                value:
                  "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/2",
              },
            ],
            background_color: "",
            animation_url: "",
            youtube_url: "",
          }),
        ],
      ],
      {
        account: classOwner.account,
      },
    );
    await nftClassContract.write.mint(
      [
        likerLand.account.address,
        ["_mint3"],
        [
          JSON.stringify({
            image: "ipfs://QmUEV41Hbi7qkxeYSVUtoE5xkfRFnqSd62fa5v8Naya5Ys",
            image_data: "",
            external_url: "https://www.google.com",
            description: "#0003 Description",
            name: "#0003",
            attributes: [
              {
                trait_type: "ISCN ID",
                value:
                  "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/2",
              },
            ],
            background_color: "",
            animation_url: "",
            youtube_url: "",
          }),
        ],
      ],
      {
        account: classOwner.account,
      },
    );

    return {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      randomSigner2,
      publicClient,
      nftClassId,
      nftClassContract,
    };
  }

  it("owner should be able to send in batch", async function () {
    const { nftClassContract, classOwner, randomSigner, randomSigner2 } =
      await loadFixture(initBookNFTTokenBatch);
    await expect(
      nftClassContract.write.batchTransferWithMemo(
        [
          classOwner.account.address,
          [randomSigner.account.address, randomSigner2.account.address],
          [0, 1],
          ["batch memo1", "batch memo2"],
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.not.rejected;

    expect(await nftClassContract.read.ownerOf([0])).to.equalAddress(
      randomSigner.account.address,
    );
    expect(await nftClassContract.read.ownerOf([1])).to.equalAddress(
      randomSigner2.account.address,
    );

    const logs = await nftClassContract.getEvents.TransferWithMemo();
    expect(logs).to.have.lengthOf(2);

    // Find the logs for each token
    const token0Log = logs.find((log) => log.args.tokenId === 0n);
    const token1Log = logs.find((log) => log.args.tokenId === 1n);

    expect(token0Log).to.not.be.undefined;
    expect(token0Log!.args.from).to.equalAddress(classOwner.account.address);
    expect(token0Log!.args.to).to.equalAddress(randomSigner.account.address);
    expect(token0Log!.args.tokenId).to.equal(0n);
    expect(token0Log!.args.memo).to.equal("batch memo1");

    expect(token1Log).to.not.be.undefined;
    expect(token1Log!.args.from).to.equalAddress(classOwner.account.address);
    expect(token1Log!.args.to).to.equalAddress(randomSigner2.account.address);
    expect(token1Log!.args.tokenId).to.equal(1n);
    expect(token1Log!.args.memo).to.equal("batch memo2");
  });

  it("should not able to send token owned by other", async function () {
    const { nftClassContract, classOwner, randomSigner, likerLand } =
      await loadFixture(initBookNFTTokenBatch);
    await expect(
      nftClassContract.write.batchTransferWithMemo(
        [
          classOwner.account.address,
          [randomSigner.account.address],
          [2],
          ["batch memo1"],
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.rejectedWith(/ERC721InsufficientApproval/);
    expect(await nftClassContract.read.ownerOf([2])).to.equalAddress(
      likerLand.account.address,
    );
  });

  it("should fails all if one fails", async function () {
    const { nftClassContract, classOwner, randomSigner, likerLand } =
      await loadFixture(initBookNFTTokenBatch);
    await expect(
      nftClassContract.write.batchTransferWithMemo(
        [
          classOwner.account.address,
          [
            randomSigner.account.address,
            randomSigner.account.address,
            randomSigner.account.address,
          ],
          [0, 1, 2],
          ["batch memo1", "batch memo2", "batch memo3"],
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.rejectedWith(/ERC721InsufficientApproval/);
    expect(await nftClassContract.read.ownerOf([0])).to.equalAddress(
      classOwner.account.address,
    );
    expect(await nftClassContract.read.ownerOf([1])).to.equalAddress(
      classOwner.account.address,
    );
    expect(await nftClassContract.read.ownerOf([2])).to.equalAddress(
      likerLand.account.address,
    );
  });
});

describe("BookNFTToken Burnable", () => {
  async function initBookNFTTokenBurnable() {
    const {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      randomSigner2,
      publicClient,
    } = await loadFixture(deployProtocol);

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      const unwatch = likeProtocol.watchEvent.NewBookNFT({
        onLogs: (logs) => {
          const id = logs[0].args.bookNFT;
          if (id) {
            unwatch();
            resolve(id);
          }
        },
      });
    });

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    await likeProtocol.write.newBookNFT([
      {
        creator: classOwner.account.address,
        updaters: [classOwner.account.address, likerLand.account.address],
        minters: [classOwner.account.address, likerLand.account.address],
        config: bookConfig,
      },
    ]);

    const nftClassId = await NewClassEvent;
    const nftClassContract = await viem.getContractAt("BookNFT", nftClassId);

    await nftClassContract.write.mint(
      [
        classOwner.account.address,
        ["_mint1"],
        [
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
        ],
      ],
      {
        account: classOwner.account,
      },
    );

    return {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      randomSigner2,
      publicClient,
      nftClassId,
      nftClassContract,
    };
  }

  it("owner should be able to burn NFT", async function () {
    const { nftClassContract, classOwner } = await loadFixture(
      initBookNFTTokenBurnable,
    );
    expect(await nftClassContract.read.ownerOf([0])).to.equalAddress(
      classOwner.account.address,
    );
    await expect(
      nftClassContract.write.burn([0], {
        account: classOwner.account,
      }),
    ).to.be.not.rejected;
    await expect(nftClassContract.read.ownerOf([0])).to.be.rejectedWith(
      /ERC721NonexistentToken/,
    );
  });

  it("should not able to burn NFT owned by other", async function () {
    const { nftClassContract, classOwner, randomSigner } = await loadFixture(
      initBookNFTTokenBurnable,
    );
    await expect(
      nftClassContract.write.burn([0], {
        account: randomSigner.account,
      }),
    ).to.be.rejectedWith(/ERC721InsufficientApproval/);

    expect(await nftClassContract.read.ownerOf([0])).to.equalAddress(
      classOwner.account.address,
    );
  });

  it("should count the total supply correctly after burn", async function () {
    const { nftClassContract, classOwner } = await loadFixture(
      initBookNFTTokenBurnable,
    );
    expect(await nftClassContract.read.totalSupply()).to.equal(1n);
    await nftClassContract.write.burn([0], {
      account: classOwner.account,
    });
    expect(await nftClassContract.read.totalSupply()).to.equal(0n);
  });
});
