import { expect } from "chai";
import hardhat from "hardhat";

describe("EkilCoin", () => {
  before(async function () {
    this.EkilCoin = await hardhat.ethers.getContractFactory("EkilCoin");
  });

  beforeEach(async function () {
    this.ekilCoin = await this.EkilCoin.deploy(1000000, "GIT_HASH");
    await this.ekilCoin.waitForDeployment();
  });

  it("retrieve git hash", async function () {
    expect((await this.ekilCoin.getGitHash()).toString()).to.equal("GIT_HASH");
  });
});
