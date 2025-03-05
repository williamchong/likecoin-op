import { ethers } from "hardhat";

async function mintNFT() {
  const classId = "0xAf0f5e9a70349947961691DB04f642c480200315";
  const signer = await ethers.provider.getSigner();

  const LikeProtocol = await ethers.getContractAt(
    "LikeProtocol",
    process.env.ERC721_PROXY_ADDRESS!,
  );

  const likeProtocol = LikeProtocol.connect(signer);

  const tx = await likeProtocol.mintNFT({
    to: signer.address,
    classId: classId,
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
