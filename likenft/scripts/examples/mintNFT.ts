import { ethers } from "hardhat";

async function mintNFT() {
  const classId = "0x1D146390C1D4E03C74b87D896b254a5468EDF804";
  const signer = await ethers.provider.getSigner();

  const LikeNFTClass = await ethers.getContractAt("BookNFT", classId);
  const likeNFTClass = LikeNFTClass.connect(signer);

  const tx = await likeNFTClass.mint(
    signer.address,
    ["mint via script"],
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
  );
  console.log(await tx.wait());
}

mintNFT().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
