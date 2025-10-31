import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";
import LikeCollectiveV0Module from "./LikeCollectiveV0";

/**
 * V1 LikeCollective deployment is designed to recurrently upgrade of V0
 * DOTENV_CONFIG_PATH=.env \
   npx hardhat ignition deploy \
    ignition/modules/LikeCollectiveV1.ts \
    --verify --strategy create2 \
    --parameters ignition/parameters.json --network baseSepolia

# Swipe previous deployment
npx hardhat ignition wipe chain-902 \
  "LikeCollectiveV1Module#LikeCollectiveV0Module~LikeCollectiveV0.upgradeToAndCall"
npx hardhat ignition wipe chain-902 \
  "LikeCollectiveV1Module#LikeCollectiveV1Impl"

 */
const LikeCollectiveV1Module = buildModule("LikeCollectiveV1Module", (m) => {
  const { likeCollectiveV0 } = m.useModule(LikeCollectiveV0Module);

  const likeCollectiveV1Impl = m.contract("LikeCollective", [], {
    id: "LikeCollectiveV1Impl",
  });

  m.call(likeCollectiveV0, "upgradeToAndCall", [likeCollectiveV1Impl, "0x"]);

  return {
    likeCollectiveV1Impl,
  };
});

export default LikeCollectiveV1Module;
