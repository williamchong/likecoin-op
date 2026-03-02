import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const BookNFTV2Module = buildModule("BookNFTV2Module", (m) => {
  const bookNFTV2Impl = m.contract("BookNFT", [], { id: "BookNFTV2Impl" });

  return { bookNFTV2Impl };
});

export default BookNFTV2Module;
