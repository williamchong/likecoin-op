import { task } from "hardhat/config";
import fs from "fs";
import path from "path";
import { parseAbiItem, getAddress, type Address } from "viem";
import { VELIKE_DEFAULT, vaultDeployBlock } from "./velikeConstants";

// veLIKE is a non-transferable ERC4626; acquisition is always a mint
// (Transfer from the zero address) and exit is a burn (Transfer to zero).
// So the full Transfer history is enough to reconstruct every holder balance.
const TRANSFER_EVENT = parseAbiItem(
  "event Transfer(address indexed from, address indexed to, uint256 value)",
);

const ZERO: Address = "0x0000000000000000000000000000000000000000";

const DEFAULT_FILE = (chainId: number) =>
  `velike-events/${chainId}-velike-transfers.txt`;

type Line = { block: bigint; logIndex: number; text: string };

// Find the highest block already recorded in the file (resume point).
function lastIndexedBlock(file: string): bigint | null {
  if (!fs.existsSync(file)) return null;
  let max: bigint | null = null;
  for (const raw of fs.readFileSync(file, "utf8").split("\n")) {
    const line = raw.trim();
    if (!line || line.startsWith("#")) continue;
    const blk = BigInt(line.split(/\s+/)[0]);
    if (max === null || blk > max) max = blk;
  }
  return max;
}

task(
  "indexVeLikeTransfers",
  "Append new veLIKE Transfer events to a git-tracked log file (resumes from the last indexed block)",
)
  .addOptionalParam("velike", "veLike vault proxy address", VELIKE_DEFAULT)
  .addOptionalParam("block", "Tip block to index up to (or 'latest')", "latest")
  .addOptionalParam("file", "Path to the append-only log file", "")
  .setAction(async (args, hre) => {
    const publicClient = await hre.viem.getPublicClient();
    const chainId = await publicClient.getChainId();
    const velike = getAddress(args.velike);
    const file = args.file || DEFAULT_FILE(chainId);
    const tip =
      args.block === "latest"
        ? await publicClient.getBlockNumber()
        : BigInt(args.block);

    const last = lastIndexedBlock(file);
    const fromBlock = last === null ? vaultDeployBlock(chainId) : last + 1n;

    console.log(`Vault (veLike): ${velike}`);
    console.log(`Log file:       ${file}`);
    console.log(
      `Resume from:    ${fromBlock}${last === null ? " (fresh / deploy block)" : ` (last indexed ${last})`}`,
    );
    console.log(`Index up to:    ${tip}`);

    if (fromBlock > tip) {
      console.log("Already up to date. Nothing to index.");
      return;
    }

    // Adaptive log scan: split the range on provider errors (too many results).
    const collected: Line[] = [];
    const scan = async (from: bigint, to: bigint): Promise<void> => {
      try {
        const logs = await publicClient.getLogs({
          address: velike,
          event: TRANSFER_EVENT,
          fromBlock: from,
          toBlock: to,
        });
        for (const log of logs) {
          const f = log.args.from as Address | undefined;
          const t = log.args.to as Address | undefined;
          const amount = (log.args.value as bigint).toString();
          const block = log.blockNumber!;
          const logIndex = log.logIndex!;
          // mint: from == 0 -> credit `to`; burn: to == 0 -> debit `from`.
          // (Defensive `in`/`out` lines in case transferability is ever enabled.)
          if (f === ZERO && t && t !== ZERO) {
            collected.push({
              block,
              logIndex,
              text: `${block} ${getAddress(t)} mint ${amount}`,
            });
          } else if (t === ZERO && f && f !== ZERO) {
            collected.push({
              block,
              logIndex,
              text: `${block} ${getAddress(f)} burn ${amount}`,
            });
          } else {
            if (t && t !== ZERO)
              collected.push({
                block,
                logIndex,
                text: `${block} ${getAddress(t)} in ${amount}`,
              });
            if (f && f !== ZERO)
              collected.push({
                block,
                logIndex,
                text: `${block} ${getAddress(f)} out ${amount}`,
              });
          }
        }
      } catch (e) {
        if (to > from) {
          const mid = from + (to - from) / 2n;
          await scan(from, mid);
          await scan(mid + 1n, to);
        } else {
          throw e;
        }
      }
    };
    await scan(fromBlock, tip);

    // Stable chronological order (block, then logIndex).
    collected.sort((a, b) =>
      a.block === b.block
        ? a.logIndex - b.logIndex
        : a.block < b.block
          ? -1
          : 1,
    );

    fs.mkdirSync(path.dirname(path.resolve(file)), { recursive: true });
    if (collected.length > 0) {
      fs.appendFileSync(file, collected.map((l) => l.text).join("\n") + "\n");
    }
    console.log(
      `\nAppended ${collected.length} event line(s) for blocks ${fromBlock}..${tip}.`,
    );
    console.log(
      `Commit ${file} so the next rotation resumes from block ${tip + 1n}.`,
    );
  });
