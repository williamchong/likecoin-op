import { ethers } from "hardhat";

async function mintNFT() {
  const classId = "0xAf0f5e9a70349947961691DB04f642c480200315";
  const signer = await ethers.provider.getSigner();

  const LikeProtocol = await ethers.getContractAt(
    "LikeProtocol",
    process.env.ERC721_PROXY_ADDRESS!,
  );

  const likeProtocol = LikeProtocol.connect(signer);

  const tx = await likeProtocol.mintNFTs({
    to: signer.address,
    classId: classId,
    inputs: [
      {
        metadata: JSON.stringify({
          image:
            "ipfs://bafybeie2zfhndeavdc7ebbitr3krbxlxtjsdwmzbdkkkmiscg4j5xf3rwi",
          image_data: "",
          external_url: "https://bit.ly/moneyverse-pdf",
          description: "#0001",
          name: "#0001",
          attributes: [],
          background_color: "",
          animation_url: "",
          youtube_url: "",
        }),
      },
      {
        metadata: JSON.stringify({
          image:
            "ipfs://bafybeie2zfhndeavdc7ebbitr3krbxlxtjsdwmzbdkkkmiscg4j5xf3rwi",
          image_data: "",
          external_url: "https://bit.ly/moneyverse-pdf",
          description: "#0002",
          name: "#0002",
          attributes: [],
          background_color: "",
          animation_url: "",
          youtube_url: "",
        }),
      },
      {
        metadata: JSON.stringify({
          image:
            "ipfs://bafybeie2zfhndeavdc7ebbitr3krbxlxtjsdwmzbdkkkmiscg4j5xf3rwi",
          image_data: "",
          external_url: "https://bit.ly/moneyverse-pdf",
          description: "#0003",
          name: "#0003",
          attributes: [],
          background_color: "",
          animation_url: "",
          youtube_url: "",
        }),
      },
      {
        metadata: JSON.stringify({
          image:
            "ipfs://bafybeie2zfhndeavdc7ebbitr3krbxlxtjsdwmzbdkkkmiscg4j5xf3rwi",
          image_data: "",
          external_url: "https://bit.ly/moneyverse-pdf",
          description: "#0004",
          name: "#0004",
          attributes: [],
          background_color: "",
          animation_url: "",
          youtube_url: "",
        }),
      },
      {
        metadata: JSON.stringify({
          image:
            "ipfs://bafybeie2zfhndeavdc7ebbitr3krbxlxtjsdwmzbdkkkmiscg4j5xf3rwi",
          image_data: "",
          external_url: "https://bit.ly/moneyverse-pdf",
          description: "#0005",
          name: "#0005",
          attributes: [],
          background_color: "",
          animation_url: "",
          youtube_url: "",
        }),
      },
    ],
  });
  console.log(await tx.wait());
}

mintNFT().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
