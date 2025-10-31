import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import LikeCollectiveV0Module from "./LikeCollectiveV0";
import LikeStakePositionModule from "./LikeStakePosition";
import LikecoinModule from "./Likecoin";

/**
 * 
 * DOTENV_CONFIG_PATH=.env \
   npx hardhat ignition deploy \
    ignition/modules/LikeCollective.ts \
    --verify --strategy create2 \
    --parameters ignition/parameters.json --network baseSepolia
 */
const LikeCollectiveModule = buildModule("LikeCollectiveModule", (m) => {
  const { likeCollectiveV0 } = m.useModule(LikeCollectiveV0Module);
  const { likeStakePosition } = m.useModule(LikeStakePositionModule);
  const { likecoin } = m.useModule(LikecoinModule);

  const likeCollectiveImpl = m.contract("LikeCollective", [], {
    id: "LikeCollectiveImpl",
  });

  m.call(likeCollectiveV0, "upgradeToAndCall", [likeCollectiveImpl, "0x"]);

  const likeCollective = m.contractAt("LikeCollective", likeCollectiveV0);

  m.call(likeCollective, "setLikeStakePosition", [likeStakePosition]);
  m.call(likeCollective, "setLikecoin", [likecoin]);
  m.call(likeStakePosition, "setManager", [likeCollective]);

  return {
    likeCollectiveImpl,
    likeCollective,
    likeStakePosition,
    likecoin,
  };
});

export default LikeCollectiveModule;
