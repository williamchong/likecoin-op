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

async function _queryTransactionMemos() {
  const classId = "0x7f2a8B018075A412bE100EFF15b0F3E4c6DE96B4";
  const tokenId = 0;

  const signer = await ethers.provider.getSigner();

  const Class = await ethers.getContractFactory("Class", {
    signer,
  });

  const class_ = Class.attach(classId);

  console.log(await queryTransactionMemos(class_, tokenId));
}

_queryTransactionMemos().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
