import { ViemIgnitionHelper } from "@nomicfoundation/hardhat-ignition-viem/dist/src/viem-ignition-helper";
import { PublicClient } from "@nomicfoundation/hardhat-viem/src/types";
import { CallExecutor } from "./calls";
import { EventRetriever } from "./events";
import { Simulation } from "./models/schema";
import { deployCollective, setupFundLikecoin } from "./setup";
import { StateRetriever } from "./state";

export async function simulate(
  ignition: ViemIgnitionHelper,
  publicClient: PublicClient,
  simulate: Simulation,
) {
  const { likecoin, likeCollective, likeStakePosition } =
    await deployCollective(ignition, simulate.setup);

  for (const account of simulate.setup.accounts) {
    await setupFundLikecoin(
      likecoin,
      simulate.setup.deployer,
      account.address,
      account.likecoin,
    );
  }

  const callExecutor = new CallExecutor(
    publicClient,
    likecoin,
    likeCollective,
    likeStakePosition,
  );
  const eventRetriever = new EventRetriever(
    publicClient,
    likeCollective,
    likeStakePosition,
  );
  const stateRetriever = new StateRetriever(
    likeCollective,
    likeStakePosition,
    likecoin,
  );

  const startBlock = await publicClient.getBlockNumber();
  let endBlock = startBlock;

  for (const step of simulate.steps) {
    const { calls } = step;
    const txHash = await callExecutor.execute(calls[0], ...calls.slice(1));
    const txReceipt = await publicClient.waitForTransactionReceipt({
      hash: txHash,
    });
    endBlock = txReceipt.blockNumber;
  }

  const logs = await eventRetriever.getEvents(endBlock, startBlock + 1n);
  const state = await stateRetriever.retrieve();

  return {
    logs,
    state,
  };
}
