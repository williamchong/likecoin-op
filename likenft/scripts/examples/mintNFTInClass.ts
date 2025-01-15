import { ethers } from "hardhat";

async function mintNFTInClass() {
  const signer = await ethers.provider.getSigner();

  const Class_ = await ethers.getContractFactory("Class", {
    signer,
  });

  // Extract and update the class id from newClass's NewClass event
  const class_ = Class_.attach("0x14CE6632272552E676b53FE6202edA8F1Be4992c");

  const tx = await class_.mint(
    signer.address,
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
  );
  console.log(await tx.wait());
}

mintNFTInClass().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
