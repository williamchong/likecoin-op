import { ethers } from "hardhat";

async function updateClass() {
  const signer = await ethers.provider.getSigner();

  const LikeNFT = await ethers.getContractFactory("LikeNFT", {
    signer,
  });

  const likeNFT = LikeNFT.attach(process.env.PROXY_ADDRESS!);

  const tx = await likeNFT.updateClass({
    creator: signer.address,
    class_id: "0x14CE6632272552E676b53FE6202edA8F1Be4992c",
    input: {
      name: "My Book",
      symbol: "KOOB",
      description: "Description",
      uri: "",
      uri_hash: "",
      metadata: JSON.stringify({
        name: "My Book 202412201605 Updated",
        symbol: "KOOB202412201605 Updated",
        description: "My description 202412201604 Updated",
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

updateClass().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
