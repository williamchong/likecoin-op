import { Contract, ContractFactory, EventLog } from "ethers";
import { ethers } from "hardhat";

export interface TransactionMemo {
  transaction: string;
  from: string;
  to: string;
  tokenId: number;
  memo: string;
}

type AttachedContract = ReturnType<ContractFactory<any[], Contract>["attach"]>;

async function _queryTransactionMemos() {
  const classId = "0x73228898f17e2057D898475953c1018704155e25";
  const tokenId = 0;

  const signer = await ethers.provider.getSigner();

  const LikeNFTClass = await ethers.getContractAt(
    "LikeNFTClass",
    classId,
  );
  const likeNFTClass = LikeNFTClass.connect(signer);
  console.log(await queryTransactionMemos(likeNFTClass, tokenId));
}

export async function queryTransactionMemos(
  class_: AttachedContract,
  tokenId: number,
): Promise<TransactionMemo[]> {
  const filters = class_.filters.TransferWithMemo(null, null, tokenId);
  const logs = await class_.queryFilter(filters);
  const transactionMemos: TransactionMemo[] = [];
  for (const log of logs) {
    const { transactionHash } = log;
    if (log instanceof EventLog) {
      const [from, to, tokenId, memo] = log.args;
      transactionMemos.push({
        transaction: transactionHash,
        from,
        to,
        tokenId,
        memo,
      });
    }
  }

  return transactionMemos;
}



_queryTransactionMemos().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
