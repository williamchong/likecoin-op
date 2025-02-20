import { ethers } from "hardhat";

async function mintNFTInClass() {
  const classId = "0x84ce8AaB5aceCaE283083761498440539a5DD8dE";
  const signer = await ethers.provider.getSigner();

  const LikeNFTClass = await ethers.getContractAt("LikeNFTClass", classId);
  const likeNFTClass = LikeNFTClass.connect(signer);

  const tx = await likeNFTClass.mint(signer.address, [
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
  ]);
  console.log(await tx.wait());
}

mintNFTInClass().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
