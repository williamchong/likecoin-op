import { loadFixture } from "@nomicfoundation/hardhat-toolbox-viem/network-helpers";
import { expect } from "chai";
import { EventLog, BaseContract } from "ethers";
import { ethers, viem } from "hardhat";

import "./setup";
import { BookConfigLoader, BookTokenConfigLoader } from "./BookConfigLoader";
import { deployProtocol } from "./factory";

describe("BookNFTClass", () => {
  async function initMint() {
    const {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      publicClient,
    } = await loadFixture(deployProtocol);
    const protocolOwner = deployer;

    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      const unwatch = likeProtocol.watchEvent.NewBookNFT({
        onLogs: (logs) => {
          const id = logs[0].args.bookNFT;
          unwatch();
          resolve(id);
        },
      });
    });
    await likeProtocol.write.newBookNFT([
      {
        creator: classOwner.account.address,
        updaters: [classOwner.account.address, likerLand.account.address],
        minters: [classOwner.account.address, likerLand.account.address],
        config: bookConfig,
      },
    ]);

    const nftClassId = await NewClassEvent;
    const book0NFT = await viem.getContractAt("BookNFT", nftClassId);
    expect(await book0NFT.read.owner()).to.equalAddress(
      classOwner.account.address,
    );
    return {
      likeProtocol,
      bookNFTImpl,
      deployer,
      protocolOwner,
      classOwner,
      likerLand,
      randomSigner,
      publicClient,
      book0NFT,
    };
  }

  it("should have the correct STORAGE_SLOT", async function () {
    const bookNFTMock = await viem.deployContract("BookNFTMock");
    expect(await bookNFTMock.read.bookNFTStorage()).to.equal(
      "0x8303e9d27d04c843c8d4a08966b1e1be0214fc0b3375d79db0a8252068c41f00",
    );
  });

  it("should support EIP 721 interface", async function () {
    const { book0NFT } = await loadFixture(initMint);
    // https://eips.ethereum.org/EIPS/eip-721#specification
    expect(await book0NFT.read.supportsInterface(["0x80ac58cd"])).to.equal(
      true,
    );
  });

  it("should support EIP 2981 interface", async function () {
    const { book0NFT } = await loadFixture(initMint);
    // https://eips.ethereum.org/EIPS/eip-2981#specification
    expect(await book0NFT.read.supportsInterface(["0x2a55205a"])).to.equal(
      true,
    );
  });

  it("should support EIP 4906 interface", async function () {
    const { book0NFT } = await loadFixture(initMint);
    // https://eips.ethereum.org/EIPS/eip-4906#specification
    expect(await book0NFT.read.supportsInterface(["0x49064906"])).to.equal(
      true,
    );
  });

  it("should not able to re-initialize", async function () {
    const { book0NFT, randomSigner } = await loadFixture(initMint);
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    const owner = await book0NFT.read.owner();
    expect(owner).is.not.equalAddress(randomSigner.account.address);

    await expect(
      book0NFT.write.initialize([bookConfig.name, bookConfig.symbol], {
        account: randomSigner.account,
      }),
    ).to.be.rejectedWith("InvalidInitialization()");
  });

  it("should able to initConfig by class owner", async function () {
    const { book0NFT, classOwner , likerLand} = await loadFixture(initMint);
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    await expect(
      book0NFT.write.initConfig([
        classOwner.account.address,
        [classOwner.account.address, likerLand.account.address],
        [classOwner.account.address, likerLand.account.address],
        bookConfig,
      ], {
        account: classOwner.account,
      }),
    ).to.be.not.rejected;
  });

  it("should not able to initConfig after mint", async function () {
    const { book0NFT, classOwner , likerLand} = await loadFixture(initMint);
    await book0NFT.write.mint([
      classOwner.account.address,
      ["_mint"],
      [JSON.stringify({
        name: "My Book",
        symbol: "KOOB",
      })],
    ], {
      account: classOwner.account,
    });
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    await expect(
      book0NFT.write.initConfig([
        classOwner.account.address,
        [classOwner.account.address, likerLand.account.address],
        [classOwner.account.address, likerLand.account.address],
        bookConfig,
      ], {
        account: classOwner.account,
      }),
    ).to.be.rejectedWith("InvalidInitialization()");
  });

  it("should not able to initConfig by random signer", async function () {
    const { book0NFT, randomSigner, classOwner, likerLand } = await loadFixture(initMint);
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    await expect(
      book0NFT.write.initConfig([
        randomSigner.account.address,
        [randomSigner.account.address, likerLand.account.address],
        [randomSigner.account.address, likerLand.account.address],
        bookConfig,
      ], {
        account: randomSigner.account,
      }),
    ).to.be.rejected;
  });

  it("should have the right roles assigned", async function () {
    const { book0NFT, protocolOwner, classOwner, likerLand, randomSigner } =
      await loadFixture(initMint);
    const MINTER_ROLE = await book0NFT.read.MINTER_ROLE();
    const UPDATER_ROLE = await book0NFT.read.UPDATER_ROLE();
    expect(
      await book0NFT.read.hasRole([MINTER_ROLE, protocolOwner.account.address]),
    ).to.equal(false);
    expect(
      await book0NFT.read.hasRole([MINTER_ROLE, classOwner.account.address]),
    ).to.equal(true);
    expect(
      await book0NFT.read.hasRole([MINTER_ROLE, likerLand.account.address]),
    ).to.equal(true);
    expect(
      await book0NFT.read.hasRole([MINTER_ROLE, randomSigner.account.address]),
    ).to.equal(false);
    expect(
      await book0NFT.read.hasRole([
        UPDATER_ROLE,
        protocolOwner.account.address,
      ]),
    ).to.equal(false);
    expect(
      await book0NFT.read.hasRole([UPDATER_ROLE, classOwner.account.address]),
    ).to.equal(true);
    expect(
      await book0NFT.read.hasRole([UPDATER_ROLE, likerLand.account.address]),
    ).to.equal(true);
    expect(
      await book0NFT.read.hasRole([UPDATER_ROLE, randomSigner.account.address]),
    ).to.equal(false);
  });

  it("should return the right current index", async function () {
    const { book0NFT, classOwner } = await loadFixture(initMint);
    expect(await book0NFT.read.getCurrentIndex()).to.equal(0n);

    await expect(
      book0NFT.write.mint(
        [
          classOwner.account.address,
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
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.not.rejected;

    expect(await book0NFT.read.getCurrentIndex()).to.equal(1n);
  });

  it("should allow class owner to update class and mint NFT", async function () {
    const { book0NFT, classOwner } = await loadFixture(initMint);
    expect(await book0NFT.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    await expect(
      book0NFT.write.update(
        [
          {
            name: "My Book updated",
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
            max_supply: 20,
          },
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.not.rejected;
    await expect(await book0NFT.read.totalSupply()).to.equal(0n);

    await expect(
      book0NFT.write.mint(
        [
          classOwner.account.address,
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
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.not.rejected;
    await expect(await book0NFT.read.totalSupply()).to.equal(1n);
    await expect(
      book0NFT.write.update(
        [
          {
            name: "My Book updated",
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
            max_supply: 20,
          },
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.not.rejected;
    await expect(await book0NFT.read.totalSupply()).to.equal(1n);
  });

  it("should reject update class with decreasing max supply", async function () {
    const { book0NFT, classOwner } = await loadFixture(initMint);
    await expect(await book0NFT.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    await expect(
      book0NFT.write.update(
        [
          {
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
          },
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.rejectedWith("ErrSupplyDecrease");

    expect(await book0NFT.read.symbol()).to.equal("KOOB");
  });

  it("should not allow random address update class", async function () {
    const { book0NFT, classOwner, randomSigner } = await loadFixture(initMint);

    await expect(
      book0NFT.write.update(
        [
          {
            name: "Hi Jack",
            symbol: "HIJACK",
            metadata: JSON.stringify({}),
            config: {
              max_supply: 0,
            },
          },
        ],
        {
          account: randomSigner.account,
        },
      ),
    ).to.be.rejected;
    expect(await book0NFT.read.owner()).to.equalAddress(
      classOwner.account.address,
    );
    expect(await book0NFT.read.symbol()).to.equal("KOOB");
  });

  it("should allow class owner to mintNFTs in batch", async function () {
    const { book0NFT, classOwner } = await loadFixture(initMint);
    expect(await book0NFT.read.totalSupply()).to.equal(0n);
    await expect(
      book0NFT.write.batchMint(
        [
          [classOwner.account.address, classOwner.account.address],
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
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.not.rejected;
    expect(await book0NFT.read.totalSupply()).to.equal(2n);
    expect(await book0NFT.read.tokenURI([0n])).to.equal(
      "data:application/json;base64,eyJpbWFnZSI6ImlwZnM6Ly9RbVVFVjQxSGJpN3FreGVZU1ZVdG9FNXhrZlJGbnFTZDYyZmE1djhOYXlhNVlzIiwiaW1hZ2VfZGF0YSI6IiIsImV4dGVybmFsX3VybCI6Imh0dHBzOi8vd3d3Lmdvb2dsZS5jb20iLCJkZXNjcmlwdGlvbiI6IjIwMjQxMjE5MTcyOSAjMDAwMSBEZXNjcmlwdGlvbiIsIm5hbWUiOiIyMDI0MTIxOTE3MjkgIzAwMDEiLCJhdHRyaWJ1dGVzIjpbeyJ0cmFpdF90eXBlIjoiSVNDTiBJRCIsInZhbHVlIjoiaXNjbjovL2xpa2Vjb2luLWNoYWluL0Z5WjEzbV9oZ3d6VUM2VW9hUzN2RmRZdmRHNlFYZmFqVTN2Y2F0dzdYMWMvMSJ9XSwiYmFja2dyb3VuZF9jb2xvciI6IiIsImFuaW1hdGlvbl91cmwiOiIiLCJ5b3V0dWJlX3VybCI6IiJ9",
    );
  });

  it("should check token id when safe mint with token id", async function () {
    const { book0NFT, classOwner } = await loadFixture(initMint);
    expect(await book0NFT.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    const mintNFT = async () => {
      await book0NFT.write.mint(
        [
          classOwner.account.address,
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
        ],
        {
          account: classOwner.account,
        },
      );
    };

    const safeMintWithTokenId = async (fromTokenId: number) => {
      await book0NFT.write.safeMintWithTokenId(
        [
          fromTokenId,
          [classOwner.account.address, classOwner.account.address],
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
        ],
        {
          account: classOwner.account,
        },
      );
    };

    await expect(mintNFT()).to.be.not.rejected;
    await expect(await book0NFT.read.totalSupply()).to.equal(1n);
    await expect(safeMintWithTokenId([0])).to.be.rejectedWith(
      "ErrTokenIdMintFails(1)",
    );
    await expect(safeMintWithTokenId(1)).to.be.not.rejected;
    await expect(await book0NFT.read.totalSupply()).to.equal(3n);
    await expect(safeMintWithTokenId([1])).to.be.rejectedWith(
      "ErrTokenIdMintFails(3)",
    );
  });
});

describe("BookNFT permission control", () => {
  async function initPermissionControl() {
    const {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      publicClient,
    } = await loadFixture(deployProtocol);

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      const unwatch = likeProtocol.watchEvent.NewBookNFT({
        onLogs: (logs) => {
          const id = logs[0].args.bookNFT;
          unwatch();
          resolve(id);
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
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    return {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      publicClient,
      nftClassId,
      nftClassContract,
    };
  }

  it("should not allow random address to grant/revoke role", async function () {
    const { nftClassContract, classOwner, randomSigner } = await loadFixture(
      initPermissionControl,
    );
    const MINTER_ROLE = await nftClassContract.read.MINTER_ROLE();
    const UPDATER_ROLE = await nftClassContract.read.UPDATER_ROLE();

    await expect(
      nftClassContract.write.ownerGrantRole(
        [MINTER_ROLE, classOwner.account.address],
        {
          account: randomSigner.account,
        },
      ),
    ).to.be.rejected;
    await expect(
      nftClassContract.write.ownerRevokeRole(
        [UPDATER_ROLE, classOwner.account.address],
        {
          account: randomSigner.account,
        },
      ),
    ).to.be.rejected;

    expect(
      await nftClassContract.read.hasRole([
        MINTER_ROLE,
        classOwner.account.address,
      ]),
    ).to.equal(true);
    expect(
      await nftClassContract.read.hasRole([
        UPDATER_ROLE,
        classOwner.account.address,
      ]),
    ).to.equal(true);
  });

  it("should allow class owner with minter role to mint NFT", async function () {
    const { nftClassContract, classOwner, likerLand } = await loadFixture(
      initPermissionControl,
    );

    const tokenConfig0 = BookTokenConfigLoader.load(
      "./test/fixtures/TokenConfig0.json",
    );
    const tokenConfig1 = BookTokenConfigLoader.load(
      "./test/fixtures/TokenConfig1.json",
    );
    const MINTER_ROLE = await nftClassContract.read.MINTER_ROLE();
    expect(
      await nftClassContract.read.hasRole([
        MINTER_ROLE,
        classOwner.account.address,
      ]),
    ).to.equal(true);
    await expect(
      nftClassContract.write.mint(
        [
          likerLand.account.address,
          ["_mint1", "_mint2"],
          [tokenConfig0, tokenConfig1],
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.not.rejected;
    expect(await nftClassContract.read.totalSupply()).to.equal(2n);
    expect(
      await nftClassContract.read.balanceOf([likerLand.account.address]),
    ).to.equal(2n);
    expect(
      await nftClassContract.read.balanceOf([classOwner.account.address]),
    ).to.equal(0n);
  });

  it("should not allow random address to mint NFT", async function () {
    const { nftClassContract, randomSigner, likerLand } = await loadFixture(
      initPermissionControl,
    );

    const tokenConfig0 = BookTokenConfigLoader.load(
      "./test/fixtures/TokenConfig0.json",
    );
    const tokenConfig1 = BookTokenConfigLoader.load(
      "./test/fixtures/TokenConfig1.json",
    );
    await expect(
      nftClassContract.write.mint(
        [
          likerLand.account.address,
          ["_mint1", "_mint2"],
          [tokenConfig0, tokenConfig1],
        ],
        {
          account: randomSigner.account,
        },
      ),
    ).to.be.rejectedWith("ErrUnauthorized()");

    await expect(await nftClassContract.read.totalSupply()).to.equal(0n);
    await expect(
      await nftClassContract.read.balanceOf([likerLand.account.address]),
    ).to.equal(0n);
  });

  it("should not allow new owner without minter role to mint NFT", async function () {
    const { nftClassContract, classOwner, randomSigner, likerLand } =
      await loadFixture(initPermissionControl);
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    await expect(
      nftClassContract.write.transferOwnership([randomSigner.account.address], {
        account: classOwner.account,
      }),
    ).to.not.be.rejected;
    expect(await nftClassContract.read.owner()).to.equalAddress(
      randomSigner.account.address,
    );

    const tokenConfig0 = BookTokenConfigLoader.load(
      "./test/fixtures/TokenConfig0.json",
    );
    const tokenConfig1 = BookTokenConfigLoader.load(
      "./test/fixtures/TokenConfig1.json",
    );
    await expect(
      nftClassContract.write.mint(
        [
          randomSigner.account.address,
          ["_mint1", "_mint2"],
          [tokenConfig0, tokenConfig1],
        ],
        {
          account: randomSigner.account,
        },
      ),
    ).to.be.rejectedWith("ErrUnauthorized()");
    await expect(await nftClassContract.read.totalSupply()).to.equal(0n);
    await expect(
      await nftClassContract.read.balanceOf([likerLand.account.address]),
    ).to.equal(0n);
  });

  it("should allow new owner to add minter role", async function () {
    const { nftClassContract, classOwner, randomSigner, likerLand } =
      await loadFixture(initPermissionControl);
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    await expect(
      nftClassContract.write.transferOwnership([randomSigner.account.address], {
        account: classOwner.account,
      }),
    ).to.not.be.rejected;
    expect(await nftClassContract.read.owner()).to.equalAddress(
      randomSigner.account.address,
    );

    const MINTER_ROLE = await nftClassContract.read.MINTER_ROLE();
    await expect(
      nftClassContract.write.ownerGrantRole(
        [MINTER_ROLE, randomSigner.account.address],
        {
          account: randomSigner.account,
        },
      ),
    ).to.not.be.rejected;
    await expect(
      nftClassContract.write.ownerGrantRole(
        [MINTER_ROLE, likerLand.account.address],
        {
          account: randomSigner.account,
        },
      ),
    ).to.not.be.rejected;
    expect(
      await nftClassContract.read.hasRole([
        MINTER_ROLE,
        randomSigner.account.address,
      ]),
    ).to.equal(true);
    expect(
      await nftClassContract.read.hasRole([
        MINTER_ROLE,
        likerLand.account.address,
      ]),
    ).to.equal(true);
  });
});

describe("BookNFT ownership transfer", () => {
  async function initOwnershipTransfer() {
    const {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      publicClient,
    } = await loadFixture(deployProtocol);

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      const unwatch = likeProtocol.watchEvent.NewBookNFT({
        onLogs: (logs) => {
          const id = logs[0].args.bookNFT;
          unwatch();
          resolve(id);
        },
      });
    });

    const bookConfig = {
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
      max_supply: 10n,
    };

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
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    return {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      publicClient,
      nftClassId,
      nftClassContract,
    };
  }

  it("should allow class owner to transfer ownership", async function () {
    const { nftClassContract, classOwner, randomSigner } = await loadFixture(
      initOwnershipTransfer,
    );
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    await expect(
      nftClassContract.write.transferOwnership([randomSigner.account.address], {
        account: classOwner.account,
      }),
    ).to.not.be.rejected;
    expect(await nftClassContract.read.owner()).to.equalAddress(
      randomSigner.account.address,
    );
  });

  it("should not allow the original owner to transfer ownership again", async function () {
    const { nftClassContract, classOwner, randomSigner } = await loadFixture(
      initOwnershipTransfer,
    );
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    await expect(
      nftClassContract.write.transferOwnership([randomSigner.account.address], {
        account: classOwner.account,
      }),
    ).to.not.be.rejected;
    expect(await nftClassContract.read.owner()).to.equalAddress(
      randomSigner.account.address,
    );

    await expect(
      nftClassContract.write.transferOwnership([classOwner.account.address], {
        account: classOwner.account,
      }),
    ).to.be.rejected;
    expect(await nftClassContract.read.owner()).to.equalAddress(
      randomSigner.account.address,
    );
  });

  it("should not allow random signer to transfer ownership", async function () {
    const { nftClassContract, classOwner, randomSigner } = await loadFixture(
      initOwnershipTransfer,
    );
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    await expect(
      nftClassContract.write.transferOwnership([randomSigner.account.address], {
        account: randomSigner.account,
      }),
    ).to.be.rejected;
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );
  });

  it("should not modify minter permission when transfer ownership", async function () {
    const { nftClassContract, classOwner, randomSigner, likerLand } =
      await loadFixture(initOwnershipTransfer);
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    await expect(
      nftClassContract.write.transferOwnership([randomSigner.account.address], {
        account: classOwner.account,
      }),
    ).to.not.be.rejected;
    expect(await nftClassContract.read.owner()).to.equalAddress(
      randomSigner.account.address,
    );

    const tokenMetadata = [
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
    ];

    await expect(
      nftClassContract.write.mint(
        [classOwner.account.address, ["_mint1", "_mint2"], tokenMetadata],
        {
          account: likerLand.account,
        },
      ),
    ).to.be.not.rejected;
    expect(await nftClassContract.read.totalSupply()).to.equal(2n);
    expect(
      await nftClassContract.read.balanceOf([classOwner.account.address]),
    ).to.equal(2n);
  });

  it("should not allow next owner to mint NFT without minter permission", async function () {
    const { nftClassContract, classOwner, randomSigner } = await loadFixture(
      initOwnershipTransfer,
    );
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    await expect(
      nftClassContract.write.transferOwnership([randomSigner.account.address], {
        account: classOwner.account,
      }),
    ).to.not.be.rejected;
    expect(await nftClassContract.read.owner()).to.equalAddress(
      randomSigner.account.address,
    );

    const tokenMetadata = [
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
    ];

    await expect(
      nftClassContract.write.mint(
        [randomSigner.account.address, ["_mint1", "_mint2"], tokenMetadata],
        {
          account: randomSigner.account,
        },
      ),
    ).to.be.rejected;
    expect(await nftClassContract.read.totalSupply()).to.equal(0n);
    expect(
      await nftClassContract.read.balanceOf([randomSigner.account.address]),
    ).to.equal(0n);
  });

  it("should allow original owner to renounce its minter/updater role", async function () {
    const { nftClassContract, classOwner, randomSigner } = await loadFixture(
      initOwnershipTransfer,
    );
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    await expect(
      nftClassContract.write.transferOwnership([randomSigner.account.address], {
        account: classOwner.account,
      }),
    ).to.not.be.rejected;
    expect(await nftClassContract.read.owner()).to.equalAddress(
      randomSigner.account.address,
    );

    const MINTER_ROLE = await nftClassContract.read.MINTER_ROLE();
    const UPDATER_ROLE = await nftClassContract.read.UPDATER_ROLE();

    await expect(
      nftClassContract.write.renounceRole(
        [MINTER_ROLE, classOwner.account.address],
        {
          account: classOwner.account,
        },
      ),
    ).to.not.be.rejected;
    expect(
      await nftClassContract.read.hasRole([
        MINTER_ROLE,
        classOwner.account.address,
      ]),
    ).to.equal(false);

    await expect(
      nftClassContract.write.renounceRole(
        [UPDATER_ROLE, classOwner.account.address],
        {
          account: classOwner.account,
        },
      ),
    ).to.not.be.rejected;
    expect(
      await nftClassContract.read.hasRole([
        MINTER_ROLE,
        classOwner.account.address,
      ]),
    ).to.equal(false);
  });
});

describe("BookNFT config validation", () => {
  async function initConfigValidation() {
    const {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      publicClient,
    } = await loadFixture(deployProtocol);

    const NewClassEvent = new Promise<{ id: string }>((resolve, reject) => {
      const unwatch = likeProtocol.watchEvent.NewBookNFT({
        onLogs: (logs) => {
          const id = logs[0].args.bookNFT;
          unwatch();
          resolve(id);
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
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    return {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      publicClient,
      nftClassId,
      nftClassContract,
    };
  }

  it("should retuen correct config", async function () {
    const { nftClassContract } = await loadFixture(initConfigValidation);
    const config = await nftClassContract.read.getBookConfig();
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    expect(config.max_supply).to.equal(BigInt(bookConfig.max_supply));
    expect(config.metadata).to.equal(bookConfig.metadata);
    expect(config.name).to.equal(bookConfig.name);
    expect(config.symbol).to.equal(bookConfig.symbol);
  });

  it("should reject empty name in constructor", async function () {
    const { likeProtocol, classOwner } =
      await loadFixture(initConfigValidation);

    await expect(
      likeProtocol.write.newBookNFT([
        {
          creator: classOwner.account.address,
          updaters: [classOwner.account.address],
          minters: [classOwner.account.address],
          config: {
            name: "",
            symbol: "KOOB",
            metadata: JSON.stringify({
              name: "Test Collection",
              description: "Test Description",
            }),
            max_supply: 10n,
          },
        },
      ]),
    ).to.be.rejectedWith("ErrEmptyName()");
  });

  it("should reject decreasing max supply in update", async function () {
    const { nftClassContract, classOwner } =
      await loadFixture(initConfigValidation);

    await expect(
      nftClassContract.write.update(
        [
          {
            name: "Valid Name",
            symbol: "KOOB",
            metadata: JSON.stringify({
              name: "Test Collection",
              description: "Test Description",
            }),
            max_supply: 5n,
          },
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.rejectedWith("ErrSupplyDecrease");
  });

  it("should reject empty symbol in update", async function () {
    const { nftClassContract, classOwner } =
      await loadFixture(initConfigValidation);

    await expect(
      nftClassContract.write.update(
        [
          {
            name: "Valid Name",
            symbol: "",
            metadata: JSON.stringify({
              name: "Test Collection",
              description: "Test Description",
            }),
            max_supply: 10n,
          },
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.rejectedWith("ErrInvalidSymbol()");
  });

  it("should not allow change of symbol in update", async function () {
    const { nftClassContract, classOwner } =
      await loadFixture(initConfigValidation);
    const originalSymbol = await nftClassContract.read.symbol();

    await expect(
      nftClassContract.write.update(
        [
          {
            name: "Valid Name",
            symbol: "NEWSYMBOL", // Different from original symbol
            metadata: JSON.stringify({
              name: "Test Collection",
              description: "Test Description",
            }),
            max_supply: 10n,
          },
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.rejectedWith("ErrInvalidSymbol()");

    // Verify symbol remains unchanged
    expect(await nftClassContract.read.symbol()).to.equal(originalSymbol);
  });

  it("should not allow zero max supply in update", async function () {
    const { nftClassContract, classOwner } =
      await loadFixture(initConfigValidation);

    await expect(
      nftClassContract.write.update(
        [
          {
            name: "Valid Name",
            symbol: "KOOB",
            metadata: JSON.stringify({
              name: "Test Collection",
              description: "Test Description",
            }),
            max_supply: 0n,
          },
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.rejectedWith("ErrMaxSupplyZero()");
  });

  it("should not verify metadata JSON", async function () {
    const { nftClassContract, classOwner } =
      await loadFixture(initConfigValidation);

    await expect(
      nftClassContract.write.update(
        [
          {
            name: "Valid Name",
            symbol: "KOOB",
            metadata: "invalid json",
            max_supply: 10n,
          },
        ],
        {
          account: classOwner.account,
        },
      ),
    ).to.be.not.rejected;
  });
});

describe("BookNFT royalty", () => {
  async function initRoyalty() {
    const {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
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

    await likeProtocol.write.setRoyaltyReceiver([deployer.account.address]);
    await likeProtocol.write.newBookNFTWithRoyalty([
      {
        creator: classOwner.account.address,
        updaters: [classOwner.account.address, likerLand.account.address],
        minters: [classOwner.account.address, likerLand.account.address],
        config: bookConfig,
      },
      1000n,
    ]);

    const nftClassId = await NewClassEvent;
    const nftClassContract = await viem.getContractAt("BookNFT", nftClassId);
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );
    expect(await nftClassContract.read.getProtocolBeacon()).to.equalAddress(
      likeProtocol.address,
    );
    const [receiver, royaltyAmount] = await nftClassContract.read.royaltyInfo([
      0n,
      100n,
    ]);
    expect(receiver).to.equalAddress(deployer.account.address);
    expect(royaltyAmount).to.equal(10n);

    return {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      publicClient,
      nftClassId,
      nftClassContract,
    };
  }

  it("should not allow owner to set royalty fraction", async function () {
    const { nftClassContract, classOwner } = await loadFixture(initRoyalty);
    await expect(
      nftClassContract.write.setRoyaltyFraction([1000n], {
        account: classOwner.account,
      }),
    ).to.be.rejectedWith("ErrUnauthorized()");
  });

  it("should not allow updater to set royalty fraction", async function () {
    const { nftClassContract, likerLand } = await loadFixture(initRoyalty);
    await expect(
      nftClassContract.write.setRoyaltyFraction([500n], {
        account: likerLand.account,
      }),
    ).to.be.rejectedWith("ErrUnauthorized()");
  });

  it("should not allow non-owner to set royalty fraction", async function () {
    const { nftClassContract, randomSigner } = await loadFixture(initRoyalty);
    await expect(
      nftClassContract.write.setRoyaltyFraction([1000n], {
        account: randomSigner.account,
      }),
    ).to.be.rejectedWith("ErrUnauthorized()");
  });

  it("should not allow non-protocol beacon to set royalty fraction", async function () {
    const { nftClassContract, deployer } = await loadFixture(initRoyalty);

    await expect(
      nftClassContract.write.setRoyaltyFraction([1000n], {
        account: deployer.account,
      }),
    ).to.be.rejectedWith("ErrUnauthorized()");
  });
});

describe("BookNFT version", () => {
  async function initVersion() {
    const {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
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
    expect(await nftClassContract.read.owner()).to.equalAddress(
      classOwner.account.address,
    );

    // Deploy V2 but not upgrade
    const v2NFTClassContract = await viem.deployContract("BookNFTMock");

    return {
      likeProtocol,
      bookNFTImpl,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
      publicClient,
      nftClassId,
      nftClassContract,
      v2NFTClassContract,
    };
  }

  it("should have the correct version on protocol replace implementation", async function () {
    const { likeProtocol, v2NFTClassContract, nftClassId } =
      await loadFixture(initVersion);
    await likeProtocol.write.upgradeTo([v2NFTClassContract.address]);
    const version = await v2NFTClassContract.read.version();
    expect(version).to.equal(2n);
    const beaconProxy = await viem.getContractAt("BookNFTMock", nftClassId);
    expect(await beaconProxy.read.version()).to.equal(2n);
  });

  it("should not able to initialize v2 NFT class", async function () {
    const {
      likeProtocol,
      v2NFTClassContract,
      nftClassId,
      classOwner,
      likerLand,
    } = await loadFixture(initVersion);
    await likeProtocol.write.upgradeTo([v2NFTClassContract.address]);

    const beaconProxy = await viem.getContractAt("BookNFTMock", nftClassId);
    const bookConfig = BookConfigLoader.load(
      "./test/fixtures/BookConfig0.json",
    );
    await expect(
      beaconProxy.write.initialize([bookConfig.name, bookConfig.symbol], {
        account: classOwner.account,
      }),
    ).to.be.rejectedWith("InvalidInitialization()");
  });

  it("should preserve owner on implementation upgrade", async function () {
    const { nftClassContract, likeProtocol, v2NFTClassContract, nftClassId } =
      await loadFixture(initVersion);
    const owner = await nftClassContract.read.owner();

    await likeProtocol.write.upgradeTo([v2NFTClassContract.address]);

    const beaconProxy = await viem.getContractAt("BookNFTMock", nftClassId);
    expect(await beaconProxy.read.owner()).to.equalAddress(owner);
  });

  it("should preserve minters on implementation upgrade", async function () {
    const {
      nftClassContract,
      likeProtocol,
      v2NFTClassContract,
      nftClassId,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
    } = await loadFixture(initVersion);
    const MINTER_ROLE = await nftClassContract.read.MINTER_ROLE();
    expect(
      await nftClassContract.read.hasRole([
        MINTER_ROLE,
        deployer.account.address,
      ]),
    ).to.equal(false);
    expect(
      await nftClassContract.read.hasRole([
        MINTER_ROLE,
        classOwner.account.address,
      ]),
    ).to.equal(true);
    expect(
      await nftClassContract.read.hasRole([
        MINTER_ROLE,
        likerLand.account.address,
      ]),
    ).to.equal(true);
    expect(
      await nftClassContract.read.hasRole([
        MINTER_ROLE,
        randomSigner.account.address,
      ]),
    ).to.equal(false);

    await likeProtocol.write.upgradeTo([v2NFTClassContract.address]);

    const beaconProxy = await viem.getContractAt("BookNFTMock", nftClassId);
    expect(
      await beaconProxy.read.hasRole([MINTER_ROLE, deployer.account.address]),
    ).to.equal(false);
    expect(
      await beaconProxy.read.hasRole([MINTER_ROLE, classOwner.account.address]),
    ).to.equal(true);
    expect(
      await beaconProxy.read.hasRole([MINTER_ROLE, likerLand.account.address]),
    ).to.equal(true);
    expect(
      await beaconProxy.read.hasRole([
        MINTER_ROLE,
        randomSigner.account.address,
      ]),
    ).to.equal(false);
  });

  it("should preserve updaters on implementation upgrade", async function () {
    const {
      nftClassContract,
      likeProtocol,
      v2NFTClassContract,
      nftClassId,
      deployer,
      classOwner,
      likerLand,
      randomSigner,
    } = await loadFixture(initVersion);
    const UPDATER_ROLE = await nftClassContract.read.UPDATER_ROLE();
    expect(
      await nftClassContract.read.hasRole([
        UPDATER_ROLE,
        deployer.account.address,
      ]),
    ).to.equal(false);
    expect(
      await nftClassContract.read.hasRole([
        UPDATER_ROLE,
        classOwner.account.address,
      ]),
    ).to.equal(true);
    expect(
      await nftClassContract.read.hasRole([
        UPDATER_ROLE,
        likerLand.account.address,
      ]),
    ).to.equal(true);
    expect(
      await nftClassContract.read.hasRole([
        UPDATER_ROLE,
        randomSigner.account.address,
      ]),
    ).to.equal(false);

    await likeProtocol.write.upgradeTo([v2NFTClassContract.address]);

    const beaconProxy = await viem.getContractAt("BookNFTMock", nftClassId);
    expect(
      await beaconProxy.read.hasRole([UPDATER_ROLE, deployer.account.address]),
    ).to.equal(false);
    expect(
      await beaconProxy.read.hasRole([
        UPDATER_ROLE,
        classOwner.account.address,
      ]),
    ).to.equal(true);
    expect(
      await beaconProxy.read.hasRole([UPDATER_ROLE, likerLand.account.address]),
    ).to.equal(true);
    expect(
      await beaconProxy.read.hasRole([
        UPDATER_ROLE,
        randomSigner.account.address,
      ]),
    ).to.equal(false);
  });

  it("should preserve name, symbol and max supply on implementation upgrade", async function () {
    const { nftClassContract, likeProtocol, v2NFTClassContract, nftClassId } =
      await loadFixture(initVersion);
    const originalName = await nftClassContract.read.name();
    const originalSymbol = await nftClassContract.read.symbol();
    const originalMaxSupply = await nftClassContract.read.maxSupply();

    await likeProtocol.write.upgradeTo([v2NFTClassContract.address]);

    const beaconProxy = await viem.getContractAt("BookNFTMock", nftClassId);
    expect(await beaconProxy.read.name()).to.equal(originalName);
    expect(await beaconProxy.read.symbol()).to.equal(originalSymbol);
    expect(await beaconProxy.read.maxSupply()).to.equal(originalMaxSupply);
  });

  it("should preserve metadata on implementation upgrade", async function () {
    const { nftClassContract, likeProtocol, v2NFTClassContract, nftClassId } =
      await loadFixture(initVersion);
    const originalMetadata = await nftClassContract.read.contractURI();

    await likeProtocol.write.upgradeTo([v2NFTClassContract.address]);

    const beaconProxy = await viem.getContractAt("BookNFTMock", nftClassId);
    expect(await beaconProxy.read.contractURI()).to.equal(originalMetadata);
  });
});
