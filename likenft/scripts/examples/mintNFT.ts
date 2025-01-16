import { ethers } from "hardhat";

async function mintNFT() {
  const signer = await ethers.provider.getSigner();

  const LikeNFT = await ethers.getContractFactory("LikeNFT", {
    signer,
  });

  const likeNFT = LikeNFT.attach(process.env.PROXY_ADDRESS!);

  const tx = await likeNFT.mintNFT({
    creator: signer.address,
    class_id: "0x14CE6632272552E676b53FE6202edA8F1Be4992c",
    input: {
      uri: "",
      uri_hash: "",
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
  });
  console.log(await tx.wait());
}

mintNFT().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
