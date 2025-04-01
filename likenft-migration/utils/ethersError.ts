import { ErrorCode, EthersError } from 'ethers';

export function isEthersError<T extends ErrorCode = ErrorCode>(
  e: any
): e is EthersError<T> {
  return 'code' in e && 'shortMessage' in e;
}
