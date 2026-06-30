import { task } from "hardhat/config";
import fs from "fs";
import {
  keccak256,
  encodeAbiParameters,
  parseAbiParameters,
  getAddress,
  toHex,
  type Address,
  type Hex,
} from "viem";

// ERC-7201 base slot of veLikeRewardStorage (see veLikeRewardNoLock.sol).
// Field order (StakingCondition is inline = 4 slots):
//   0 vault | 1 likecoin | 2 rewardPool | 3 totalStaked | 4 lastRewardTime
//   5..8 currentStakingCondition | 9 stakerInfos (mapping) | 10 drawer | 11 autoSyncEnabled
const STORAGE_BASE =
  0xe9672d2c676bb94d428d6ce523668c779079df8febe4142a9972a2a2313d2c00n;
const TOTAL_STAKED_SLOT = STORAGE_BASE + 3n;
const STAKER_INFOS_SLOT = STORAGE_BASE + 9n;

const VELIKE_DEFAULT = "0xE55C2b91E688BE70e5BbcEdE3792d723b4766e2B";
const REWARD_DEFAULT = "0x5806B7fb388C2D3D894fE40Ba4598938b01496BA"; // veLikeRewardNoLock proxy
const VAULT_DEPLOY_BLOCK = 37568967n; // veLikeV0Module#ERC1967Proxy

const slotToHex = (slot: bigint): Hex => toHex(slot, { size: 32 });

// Storage location of stakerInfos[account].stakedAmount (first member of the struct).
const stakedAmountSlot = (account: Address): Hex =>
  keccak256(
    encodeAbiParameters(parseAbiParameters("uint256, uint256"), [
      BigInt(account),
      STAKER_INFOS_SLOT,
    ]),
  );

const readU256 = (raw: Hex | undefined): bigint =>
  raw && raw !== "0x" ? BigInt(raw) : 0n;

// Reconstruct net veLIKE balance per account from the event log file,
// counting only events at or before `upToBlock`.
function reconstructBalances(
  file: string,
  upToBlock: bigint,
): { balances: Map<Address, bigint>; maxBlock: bigint } {
  const balances = new Map<Address, bigint>();
  let maxBlock = 0n;
  for (const raw of fs.readFileSync(file, "utf8").split("\n")) {
    const line = raw.trim();
    if (!line || line.startsWith("#")) continue;
    const [blkStr, acctStr, event, amtStr] = line.split(/\s+/);
    const block = BigInt(blkStr);
    if (block > maxBlock) maxBlock = block;
    if (block > upToBlock) continue;
    const account = getAddress(acctStr);
    const amount = BigInt(amtStr);
    const prev = balances.get(account) ?? 0n;
    if (event === "mint" || event === "in")
      balances.set(account, prev + amount);
    else if (event === "burn" || event === "out")
      balances.set(account, prev - amount);
    else throw new Error(`Unknown event "${event}" in ${file}: ${line}`);
  }
  return { balances, maxBlock };
}

// Highest block in [lo, hi] whose timestamp <= tsCutoff (or lo-1 if none).
async function blockAtOrBefore(
  client: any,
  tsCutoff: bigint,
  lo: bigint,
  hi: bigint,
): Promise<bigint> {
  const loBlk = await client.getBlock({ blockNumber: lo });
  if (loBlk.timestamp > tsCutoff) return lo - 1n;
  const hiBlk = await client.getBlock({ blockNumber: hi });
  if (hiBlk.timestamp <= tsCutoff) return hi;
  let loB = lo;
  let hiB = hi;
  let ans = lo;
  while (loB <= hiB) {
    const mid = (loB + hiB) / 2n;
    const b = await client.getBlock({ blockNumber: mid });
    if (b.timestamp <= tsCutoff) {
      ans = mid;
      loB = mid + 1n;
    } else {
      hiB = mid - 1n;
    }
  }
  return ans;
}

task(
  "buildSyncStakers",
  "Analyze the indexed transfer log + on-chain staker state and emit syncStakers command(s)",
)
  .addOptionalParam("velike", "veLike vault proxy address", VELIKE_DEFAULT)
  .addOptionalParam(
    "reward",
    "veLikeRewardNoLock proxy address",
    REWARD_DEFAULT,
  )
  .addOptionalParam("file", "Path to the indexed transfer log file", "")
  .addParam(
    "starttime",
    "Reward period startTime (unix seconds). Eligibility cutoff: only accounts holding veLIKE at/before this time are synced (retroactive index 0 is only fair for them). e.g. 1780614000",
  )
  .addParam(
    "endtime",
    "Reward period endTime (unix seconds), for period validation",
  )
  .addOptionalParam(
    "block",
    "Snapshot block (defaults to the file's last indexed block); on-chain reads are pinned to it",
    "",
  )
  .addOptionalParam("chunk", "Max addresses per syncStakers call", "100")
  .addOptionalParam(
    "out",
    "Output JSON path (default velike-events/unsynced-{block}.json)",
    "",
  )
  .setAction(async (args, hre) => {
    const publicClient = await hre.viem.getPublicClient();
    const chainId = await publicClient.getChainId();
    const velike = getAddress(args.velike);
    const reward = getAddress(args.reward);
    const startTime = BigInt(args.starttime);
    const endTime = BigInt(args.endtime);
    const file = args.file || `velike-events/${chainId}-velike-transfers.txt`;
    if (!fs.existsSync(file)) {
      throw new Error(
        `Log file not found: ${file}. Run \`indexVeLikeTransfers\` first.`,
      );
    }

    // Resolve snapshot block: explicit, else the file's last indexed block.
    const preScan = reconstructBalances(file, 2n ** 255n);
    const snapshotBlock =
      args.block === "" ? preScan.maxBlock : BigInt(args.block);
    const out = args.out || `velike-events/unsynced-${snapshotBlock}.json`;

    console.log(`Vault (veLike):   ${velike}`);
    console.log(`Reward (NoLock):  ${reward}`);
    console.log(`Log file:         ${file}`);
    console.log(`Snapshot @ block: ${snapshotBlock}`);
    console.log(`Period:           start=${startTime} end=${endTime}`);

    const rewardC = await hre.viem.getContractAt("veLikeRewardNoLock", reward);
    const vaultC = await hre.viem.getContractAt(
      "contracts/veLikeV2.sol:veLike",
      velike,
    );

    // --- Validate the passed period against the on-chain condition. ---
    const cond = (await publicClient.readContract({
      address: reward,
      abi: rewardC.abi,
      functionName: "getCurrentCondition",
      blockNumber: snapshotBlock,
    })) as { startTime: bigint; endTime: bigint };
    if (cond.startTime !== startTime || cond.endTime !== endTime) {
      throw new Error(
        `Period mismatch: on-chain condition is start=${cond.startTime} end=${cond.endTime}, ` +
          `but you passed start=${startTime} end=${endTime}. Check --starttime/--endtime.`,
      );
    }
    console.log("Period matches on-chain condition. ✅");

    // --- Storage-layout self-check: stored totalStaked must equal getConfig's. ---
    const cfg = (await publicClient.readContract({
      address: reward,
      abi: rewardC.abi,
      functionName: "getConfig",
      blockNumber: snapshotBlock,
    })) as readonly [Address, Address, bigint, bigint, bigint];
    const totalStakedView = cfg[3];
    const totalStakedStorage = readU256(
      await publicClient.getStorageAt({
        address: reward,
        slot: slotToHex(TOTAL_STAKED_SLOT),
        blockNumber: snapshotBlock,
      }),
    );
    if (totalStakedStorage !== totalStakedView) {
      throw new Error(
        `Storage layout check FAILED: slot=${totalStakedStorage} != getConfig=${totalStakedView}. Aborting.`,
      );
    }
    console.log(`Storage layout verified. totalStaked = ${totalStakedView}`);

    // --- Map the startTime cutoff to a block, and reconstruct balances twice:
    //     current (at snapshotBlock) and "as of period start" (at cutoffBlock). ---
    const cutoffBlock = await blockAtOrBefore(
      publicClient,
      startTime,
      VAULT_DEPLOY_BLOCK,
      snapshotBlock,
    );
    console.log(`Cutoff block (<= startTime): ${cutoffBlock}\n`);

    const { balances: nowBalances } = reconstructBalances(file, snapshotBlock);
    const { balances: startBalances } = reconstructBalances(file, cutoffBlock);

    const holders = [...nowBalances.entries()]
      .filter(([, bal]) => bal > 0n)
      .map(([account, bal]) => ({ account, fileBalance: bal }));
    console.log(`Holders now (file, balance > 0): ${holders.length}\n`);

    // --- For each holder read on-chain stakedAmount + balanceOf at snapshot block. ---
    type Row = {
      account: Address;
      fileBalance: bigint;
      chainBalance: bigint;
      staked: bigint;
      heldAtStart: boolean;
    };
    const rows: Row[] = [];
    const BATCH = 20;
    for (let i = 0; i < holders.length; i += BATCH) {
      const batch = await Promise.all(
        holders.slice(i, i + BATCH).map(async (h): Promise<Row> => {
          const [chainBalance, stakedRaw] = await Promise.all([
            publicClient.readContract({
              address: velike,
              abi: vaultC.abi,
              functionName: "balanceOf",
              args: [h.account],
              blockNumber: snapshotBlock,
            }) as Promise<bigint>,
            publicClient.getStorageAt({
              address: reward,
              slot: stakedAmountSlot(h.account),
              blockNumber: snapshotBlock,
            }),
          ]);
          return {
            account: h.account,
            fileBalance: h.fileBalance,
            chainBalance,
            staked: readU256(stakedRaw),
            heldAtStart: (startBalances.get(h.account) ?? 0n) > 0n,
          };
        }),
      );
      rows.push(...batch);
      process.stdout.write(
        `\r  read ${Math.min(i + BATCH, holders.length)}/${holders.length}`,
      );
    }
    process.stdout.write("\n\n");

    // --- Validate file reconstruction against chain balances. ---
    const mismatches = rows.filter((r) => r.fileBalance !== r.chainBalance);
    if (mismatches.length > 0) {
      console.log(
        `⚠️  ${mismatches.length} account(s) where file balance != on-chain balance (re-run indexVeLikeTransfers):`,
      );
      for (const m of mismatches.slice(0, 10))
        console.log(
          `    ${m.account}  file=${m.fileBalance}  chain=${m.chainBalance}`,
        );
      if (mismatches.length > 10)
        console.log(`    ... and ${mismatches.length - 10} more`);
      console.log("");
    } else {
      console.log("File balances match on-chain balances. ✅\n");
    }

    // --- Classify. ---
    const synced = rows.filter((r) => r.staked > 0n);
    const unsyncedAll = rows.filter(
      (r) => r.staked === 0n && r.chainBalance > 0n,
    );
    // Eligible: unsynced AND was already a holder at the period start (index 0 is fair).
    const eligible = unsyncedAll.filter((r) => r.heldAtStart);
    // Excluded: unsynced but joined AFTER the period start — syncing them with
    // index 0 would retroactively over-credit. The natural deposit path already
    // set their index, so they must NOT be force-synced here.
    const lateJoiners = unsyncedAll.filter((r) => !r.heldAtStart);

    console.log(`Already staked on reward (synced): ${synced.length}`);
    console.log(`Unsynced total:                    ${unsyncedAll.length}`);
    console.log(`  eligible (held at start):        ${eligible.length}`);
    console.log(`  EXCLUDED (joined after start):   ${lateJoiners.length}\n`);

    if (lateJoiners.length > 0) {
      console.log(
        "Excluded late joiners (NOT synced — would be over-credited):",
      );
      for (const r of lateJoiners)
        console.log(`    ${r.account}  balance=${r.chainBalance}`);
      console.log("");
    }

    if (eligible.length === 0) {
      console.log("No eligible unsynced holders. Nothing to do.");
    } else {
      eligible.sort((a, b) => (b.chainBalance > a.chainBalance ? 1 : -1));
      const addrs = eligible.map((r) => r.account);
      const chunkSize = parseInt(args.chunk, 10);
      console.log(
        "--- syncStakers command(s) (run while the period is ACTIVE) ---\n",
      );
      for (let i = 0; i < addrs.length; i += chunkSize) {
        const arg = `[${addrs.slice(i, i + chunkSize).join(",")}]`;
        console.log(
          `cast send ${reward} "syncStakers(address[])" "${arg}" \\\n` +
            `  --account likecoin-deployer.eth --rpc-url $RPC\n`,
        );
      }
    }

    fs.writeFileSync(
      out,
      JSON.stringify(
        {
          velike,
          reward,
          snapshotBlock: snapshotBlock.toString(),
          startTime: startTime.toString(),
          endTime: endTime.toString(),
          cutoffBlock: cutoffBlock.toString(),
          totalStaked: totalStakedView.toString(),
          holders: holders.length,
          synced: synced.length,
          eligible: eligible.map((r) => ({
            account: r.account,
            balance: r.chainBalance.toString(),
          })),
          excludedLateJoiners: lateJoiners.map((r) => ({
            account: r.account,
            balance: r.chainBalance.toString(),
          })),
          syncStakersChunks: (() => {
            const addrs = eligible.map((r) => r.account);
            const chunkSize = parseInt(args.chunk, 10);
            return Array.from(
              { length: Math.ceil(addrs.length / chunkSize) },
              (_, i) => addrs.slice(i * chunkSize, (i + 1) * chunkSize),
            );
          })(),
        },
        null,
        2,
      ),
    );
    console.log(`Wrote result to ${out}`);
  });
