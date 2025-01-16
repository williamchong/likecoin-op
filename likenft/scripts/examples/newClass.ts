import { Listener } from "ethers";
import { ethers } from "hardhat";

async function newClass() {
  const signer = await ethers.provider.getSigner();

  const LikeNFT = await ethers.getContractFactory("LikeNFT", {
    signer,
  });

  const likeNFT = LikeNFT.attach(process.env.PROXY_ADDRESS!);

  const handleNewClass: Listener = (id, parameters, event) => {
    event.removeListener();
    console.log("newClassEventPayload", id, parameters);
  };
  await likeNFT.on("NewClass", handleNewClass);
  const tx = await likeNFT.newClass({
    creator: signer.address,
    parent: {
      type_: 1,
      iscn_id_prefix:
        "iscn://likecoin-chain/FyZ13m_hgwzUC6UoaS3vFdYvdG6QXfajU3vcatw7X1c/1",
    },
    input: {
      name: "My Book",
      symbol: "KOOB",
      metadata: JSON.stringify({
        name: "My Book 202412201604",
        symbol: "KOOB202412201604",
        description: "My description 202412201604",
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
  console.log(await tx.wait());
}

newClass().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
