import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const BookNFTModule = buildModule("BookNFTModule", (m) => {
  const bookNFTImpl = m.contract("BookNFT", [], { id: "BookNFTImpl" });

  return { bookNFTImpl };
});

export default BookNFTModule;
