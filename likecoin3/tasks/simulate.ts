import fs from "fs";
import { Log, Block } from "viem";
import { task } from "hardhat/config";
import yaml from "js-yaml";
import {
  SimulationSchema,
  State,
} from "../simulate/likecollective/models/schema";
import { simulate } from "../simulate/likecollective/simulate";

type logMarshaling = {
  address: `0x${string}`;
  blockHash: `0x${string}` | null;
  blockNumber: `0x${string}` | null;
  data: `0x${string}`;
  logIndex: `0x${string}` | null;
  transactionHash: `0x${string}` | null;
  transactionIndex: `0x${string}` | null;
  removed: boolean;
  topics: `0x${string}`[];
};

function marshalLog(log: Log): logMarshaling {
  return {
    address: log.address,
    blockHash: log.blockHash,
    blockNumber: log.blockNumber ? `0x${log.blockNumber.toString(16)}` : null,
    data: log.data,
    logIndex: log.logIndex ? `0x${log.logIndex.toString(16)}` : null,
    transactionHash: log.transactionHash,
    transactionIndex: log.transactionIndex
      ? `0x${log.transactionIndex.toString(16)}`
      : null,
    removed: log.removed,
    topics: log.topics,
  };
}

type headerMarshaling = {
  parentHash: `0x${string}`;
  sha3Uncles: `0x${string}`;
  miner: `0x${string}`;
  stateRoot: `0x${string}`;
  transactionsRoot: `0x${string}`;
  receiptsRoot: `0x${string}`;
  logsBloom: `0x${string}` | null;
  difficulty: `0x${string}`;
  number: `0x${string}` | null;
  gasLimit: `0x${string}`;
  gasUsed: `0x${string}`;
  timestamp: `0x${string}`;
  extraData: `0x${string}`;
  mixHash: `0x${string}`;
  nonce: `0x${string}` | null;
  baseFeePerGas: `0x${string}` | null;
  withdrawalsRoot: `0x${string}` | null;
  blobGasUsed: `0x${string}` | null;
  excessBlobGas: `0x${string}` | null;
  parentBeaconBlockRoot: `0x${string}` | null;
};

function marshalHeader(block: Block): headerMarshaling {
  return {
    parentHash: block.parentHash,
    sha3Uncles: block.sha3Uncles,
    miner: block.miner,
    stateRoot: block.stateRoot,
    transactionsRoot: block.transactionsRoot,
    receiptsRoot: block.receiptsRoot,
    logsBloom: block.logsBloom,
    difficulty: `0x${block.difficulty.toString(16)}`,
    number: block.number ? `0x${block.number.toString(16)}` : null,
    gasLimit: `0x${block.gasLimit.toString(16)}`,
    gasUsed: `0x${block.gasUsed.toString(16)}`,
    timestamp: `0x${block.timestamp.toString(16)}`,
    extraData: block.extraData,
    mixHash: block.mixHash,
    nonce: block.nonce,
    baseFeePerGas: block.baseFeePerGas
      ? `0x${block.baseFeePerGas.toString(16)}`
      : null,
    withdrawalsRoot: block.withdrawalsRoot ?? null,
    blobGasUsed: block.blobGasUsed
      ? `0x${block.blobGasUsed.toString(16)}`
      : null,
    excessBlobGas: block.excessBlobGas
      ? `0x${block.excessBlobGas.toString(16)}`
      : null,
    parentBeaconBlockRoot: block.parentBeaconBlockRoot ?? null,
  };
}

type stateMarshaling = {
  bookPendingRewards: Record<`0x${string}`, string | null>;
  bookStakedAmounts: Record<`0x${string}`, string | null>;
  userBalance: Record<`0x${string}`, string | null>;
  userPendingRewards: Record<
    `0x${string}`,
    Record<`0x${string}`, string | null>
  >;
  userStakedAmounts: Record<
    `0x${string}`,
    Record<`0x${string}`, string | null>
  >;
};

function marshalState(state: State): stateMarshaling {
  return {
    bookPendingRewards: Object.fromEntries(
      Object.entries(state.bookPendingRewards).map(([key, value]) => [
        key,
        value != null ? value.toString() : null,
      ]),
    ),
    bookStakedAmounts: Object.fromEntries(
      Object.entries(state.bookStakedAmounts).map(([key, value]) => [
        key,
        value != null ? value.toString() : null,
      ]),
    ),
    userBalance: Object.fromEntries(
      Object.entries(state.userBalance).map(([key, value]) => [
        key,
        value != null ? value.toString() : null,
      ]),
    ),
    userPendingRewards: Object.fromEntries(
      Object.entries(state.userPendingRewards).map(([key, value]) => [
        key,
        Object.fromEntries(
          Object.entries(value ?? {}).map(([key, value]) => [
            key,
            value != null ? value.toString() : null,
          ]),
        ),
      ]),
    ),
    userStakedAmounts: Object.fromEntries(
      Object.entries(state.userStakedAmounts).map(([key, value]) => [
        key,
        Object.fromEntries(
          Object.entries(value ?? {}).map(([key, value]) => [
            key,
            value != null ? value.toString() : null,
          ]),
        ),
      ]),
    ),
  };
}

function marshalSimulationResult({
  logs,
  state,
}: {
  logs: { log: Log; block: Block }[];
  state: State;
}) {
  return {
    logs: logs.map(({ log, block }) => ({
      log: marshalLog(log),
      block: marshalHeader(block),
    })),
    state: marshalState(state),
  };
}

task("simulate:likecollective", "Simulate the likecollective contract")
  .addOptionalParam("outputfile", "The output file to write the results to")
  .addPositionalParam("simulationfile", "The simulation file to run")
  .setAction(async ({ simulationfile, outputfile }, { ignition, viem }) => {
    const publicClient = await viem.getPublicClient();

    const simulationData = fs.readFileSync(simulationfile, "utf8");
    const simulation = SimulationSchema.parse(yaml.load(simulationData));
    const simulationResult = await simulate(ignition, publicClient, simulation);
    const marshaledSimulationResult = marshalSimulationResult(simulationResult);

    if (outputfile) {
      fs.writeFileSync(
        outputfile,
        JSON.stringify(marshaledSimulationResult, null, 2),
      );
    } else {
      console.log(JSON.stringify(marshaledSimulationResult));
    }
  });
