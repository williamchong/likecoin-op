import { ethers } from "hardhat";

async function mintNFT() {
  const signer = await ethers.provider.getSigner();

  const LikeNFT = await ethers.getContractFactory("LikeNFT", {
    signer,
  });

  const likeNFT = LikeNFT.attach(process.env.ERC721_PROXY_ADDRESS!);

  const tx = await likeNFT.mintNFT({
    creator: signer.address,
    class_id: "0x9c746861f90e8908975035c6593c27e6cef644d4",
    input: {
      metadata: JSON.stringify({
        image:
          "ipfs://bafybeie2zfhndeavdc7ebbitr3krbxlxtjsdwmzbdkkkmiscg4j5xf3rwi",
        image_data: "",
        external_url: "https://bit.ly/moneyverse-pdf",
        description: "#0017",
        name: "#0017",
        attributes: [
          {
            trait_type: "background",
            value: "Fade Out",
          },
          {
            trait_type: "publish_info_layout",
            value: "Bottom",
          },
          {
            trait_type: "coins_layout",
            value: "Random",
          },
          {
            trait_type: "coins_color",
            value: "Sliver",
          },
          {
            display_type: "number",
            trait_type: "coin_1_rotate_x",
            value: 0.0045,
          },
          {
            display_type: "number",
            trait_type: "coin_1_rotate_z",
            value: 0.001,
          },
          {
            display_type: "number",
            trait_type: "coin_2_rotate_x",
            value: 0.0052,
          },
          {
            display_type: "number",
            trait_type: "coin_2_rotate_z",
            value: 0.0057,
          },
          {
            display_type: "number",
            trait_type: "coin_1_position_x",
            value: 135,
          },
          {
            display_type: "number",
            trait_type: "coin_1_position_y",
            value: 21,
          },
          {
            display_type: "number",
            trait_type: "coin_2_position_x",
            value: -160,
          },
          {
            display_type: "number",
            trait_type: "coin_2_position_y",
            value: 158,
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
