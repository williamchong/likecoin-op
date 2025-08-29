import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const LikecoinModule = buildModule("LikecoinModule", (m) => {
  const initOwner = m.getParameter("initOwner");
  const likecoinImpl = m.contract("Likecoin", [], { id: "LikecoinImpl" });

  const initData = m.encodeFunctionCall(likecoinImpl, "initialize", [
    initOwner,
  ]);
  const likecoinProxy = m.contract("ERC1967Proxy", [likecoinImpl, initData]);

  const likecoin = m.contractAt("Likecoin", likecoinProxy);

  return { likecoin, likecoinImpl, likecoinProxy };
});

export default LikecoinModule;
