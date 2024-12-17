import { expect } from "chai";
import { ethers, upgrades } from "hardhat";

describe("EkilCoin", () => {
  before(async function () {
    this.EkilCoin = await ethers.getContractFactory("EkilCoin");
    const [ownerSigner, signer1] = await ethers.getSigners();

    this.ownerSigner = ownerSigner;
    this.signer1 = signer1;
  });

  beforeEach(async function () {
    const ekilCoin = await upgrades.deployProxy(
      this.EkilCoin,
      [this.ownerSigner.address, this.ownerSigner.address],
      {
        initializer: "initialize",
      },
    );
    const deployment = await ekilCoin.waitForDeployment();
    this.contractAddress = await deployment.getAddress();
  });

  it("should only allow admin to update permissions", async function () {
    const EkilCoinOwnerSigner = await ethers.getContractFactory("EkilCoin", {
      signer: this.ownerSigner,
    });
    const ekilCoinOwnerSigner = EkilCoinOwnerSigner.attach(
      this.contractAddress,
    );

    const EkilCoinSigner1 = await ethers.getContractFactory("EkilCoin", {
      signer: this.signer1,
    });
    const ekilCoinSigner1 = EkilCoinSigner1.attach(this.contractAddress);

    // Admin should be able to update permissions
    await expect(ekilCoinOwnerSigner.setMinter(this.ownerSigner)).to.be.not
      .rejected;

    // Other signers should not be able to update permissions
    await expect(ekilCoinSigner1.setMinter(this.signer1)).to.be.rejected;
  });

  it("should allow owner to pauser", async function () {
    const EkilCoinOwnerSigner = await ethers.getContractFactory("EkilCoin", {
      signer: this.ownerSigner,
    });
    const ekilCoinOwnerSigner = EkilCoinOwnerSigner.attach(
      this.contractAddress,
    );
    const EkilCoinSigner1 = await ethers.getContractFactory("EkilCoin", {
      signer: this.signer1,
    });
    const ekilCoinSigner1 = EkilCoinSigner1.attach(this.contractAddress);

    // Signer 1 should not be not able to pause
    await expect(ekilCoinOwnerSigner.pause()).to.be.not.rejected;
    await expect(ekilCoinSigner1.pause()).to.be.rejected;
  });

  it("should allow minter to mint", async function () {
    const EkilCoinOwnerSigner = await ethers.getContractFactory("EkilCoin", {
      signer: this.ownerSigner,
    });
    const ekilCoinOwnerSigner = EkilCoinOwnerSigner.attach(
      this.contractAddress,
    );
    const EkilCoinSigner1 = await ethers.getContractFactory("EkilCoin", {
      signer: this.signer1,
    });
    const ekilCoinSigner1 = EkilCoinSigner1.attach(this.contractAddress);

    // Signer 1 should not be not able to mint
    await expect(ekilCoinOwnerSigner.mint(this.signer1.address, 1)).to.be.not
      .rejected;
    await expect(ekilCoinSigner1.mint(this.signer1.address, 1)).to.be.rejected;

    // Update minter to signer 1. Should be able to mint
    await ekilCoinOwnerSigner.setMinter(this.signer1.address);
    await expect(ekilCoinSigner1.mint(this.signer1.address, 1)).to.be.not
      .rejected;
  });

  it("should pause minting", async function () {
    const EkilCoinOwnerSigner = await ethers.getContractFactory("EkilCoin", {
      signer: this.ownerSigner,
    });
    const ekilCoinOwnerSigner = EkilCoinOwnerSigner.attach(
      this.contractAddress,
    );

    // Should be able to mint
    await expect(ekilCoinOwnerSigner.mint(this.signer1.address, 1)).to.be.not
      .rejected;

    // Pause contract
    await ekilCoinOwnerSigner.pause();

    // Should be not able to mint
    await expect(ekilCoinOwnerSigner.mint(this.signer1.address, 1)).to.be
      .rejected;

    // Unpause contract
    await ekilCoinOwnerSigner.unpause();

    // Should be able to mint again
    await expect(ekilCoinOwnerSigner.mint(this.signer1.address, 1)).to.be.not
      .rejected;
  });

  it("allow owner to upgrade", async function () {
    const EkilCoinOwnerSigner = await ethers.getContractFactory("EkilCoin", {
      signer: this.ownerSigner,
    });
    const EkilCoinSigner1 = await ethers.getContractFactory("EkilCoin", {
      signer: this.signer1,
    });

    // Signer 1 should not be not able to upgrade
    await expect(
      upgrades.upgradeProxy(this.contractAddress, EkilCoinOwnerSigner),
    ).to.be.not.rejected;
    await expect(upgrades.upgradeProxy(this.contractAddress, EkilCoinSigner1))
      .to.be.rejected;
  });
});
