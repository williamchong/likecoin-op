import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const BookNFTModule = buildModule("BookNFTModule", (m) => {
  const bookNFTImpl = m.contract("BookNFT", [], { id: "BookNFTV1Impl" });

  return { bookNFTImpl };
});

export default BookNFTModule;
