import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import LikecoinModule from "./Likecoin";

const veLikeV0Module = buildModule("veLikeV0Module", (m) => {
  const initOwner = m.getParameter("initOwner");
  const { likecoin } = m.useModule(LikecoinModule);

  const veLikeV0Impl = m.contract("veLikeV0", [], {
    id: "veLikeV0Impl",
  });

  const initDataV0 = m.encodeFunctionCall(veLikeV0Impl, "initialize", [
    initOwner,
    likecoin,
  ]);

  const veLikeProxy = m.contract("ERC1967Proxy", [veLikeV0Impl, initDataV0]);

  const veLikeV0 = m.contractAt("veLikeV0", veLikeProxy);

  return {
    veLikeV0,
    veLikeV0Impl,
    veLikeProxy,
    likecoin,
  };
});

export default veLikeV0Module;
