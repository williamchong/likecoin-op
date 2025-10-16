import { ViemIgnitionHelper } from "@nomicfoundation/hardhat-ignition-viem/dist/src/viem-ignition-helper";
import LikeCollectiveModule from "../../ignition/modules/LikeCollective";
import { Address, Setup } from "./models/schema";

export async function deployCollective(
  ignition: ViemIgnitionHelper,
  setup: Setup,
) {
  const { likeCollective, likecoin, likeStakePosition } = await ignition.deploy(
    LikeCollectiveModule,
    {
      parameters: {
        LikecoinModule: {
          initOwner: setup.deployer,
        },
        LikeCollectiveV0Module: {
          initOwner: setup.deployer,
        },
        LikeStakePositionV0Module: {
          initOwner: setup.deployer,
        },
      },
      defaultSender: setup.deployer,
      strategy: "create2",
    },
  );

  return {
    likeCollective,
    likecoin,
    likeStakePosition,
  };
}

export async function setupFundLikecoin(
  likecoin: LikeCoin,
  callerAddress: Address,
  fundedAddress: Address,
  amount: bigint,
): Promise<void> {
  await likecoin.write.mint([fundedAddress, amount], {
    account: callerAddress,
  });
}

export type Contracts = Awaited<ReturnType<typeof deployCollective>>;
export type LikeCollective = Contracts["likeCollective"];
export type LikeCoin = Contracts["likecoin"];
export type LikeStakePosition = Contracts["likeStakePosition"];
