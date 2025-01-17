import { ethers } from "hardhat";

async function mintNFT() {
  const signer = await ethers.provider.getSigner();

  const LikeNFT = await ethers.getContractFactory("LikeNFT", {
    signer,
  });

  const likeNFT = LikeNFT.attach(process.env.PROXY_ADDRESS!);

  const tx = await likeNFT.mintNFTs({
    creator: signer.address,
    class_id: "0x210DA2E40318339B84F243980E49120d190C3713",
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
