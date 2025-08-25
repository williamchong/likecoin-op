import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import {
  BaseContractMethod,
  Contract,
  ContractDeployTransaction,
  ContractFactory,
  ContractMethodArgs,
  TransactionRequest,
} from "ethers";
// @ts-ignore-next-line
import { copyOverrides } from "ethers/contract";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import { IGasEstimationInterceptor } from "./interceptors";

export async function deployWithGasEstimation<
  A extends any[] = any[],
  I = Contract,
>(
  hre: HardhatRuntimeEnvironment,
  contractFactory: ContractFactory<A, I>,
  deployArgs: ContractMethodArgs<A>,
  options?: {
    interceptor?: IGasEstimationInterceptor;
    deployer?: HardhatEthersSigner;
  },
): ReturnType<typeof contractFactory.deploy> {
  const { ethers } = hre;

  let deployer = options?.deployer;
  if (deployer == null) {
    const [signer] = await ethers.getSigners();
    deployer = signer;
  }

  contractFactory = contractFactory.connect(deployer);

  let overrides: Omit<ContractDeployTransaction, "data"> = {};

  const fragment = contractFactory.interface.deploy;

  if (fragment.inputs.length + 1 === deployArgs.length) {
    overrides = await copyOverrides(deployArgs.pop());
  }

  const feeData = await ethers.provider.getFeeData();
  const { gasPrice, maxFeePerGas, maxPriorityFeePerGas } = feeData;

  const deployTx = await contractFactory.getDeployTransaction(
    ...deployArgs,
    overrides,
  );

  const gasLimit = await deployer.estimateGas(deployTx);

  const gasEstimation = {
    gasPrice,
    maxFeePerGas,
    maxPriorityFeePerGas,
    gasLimit,
  };

  options?.interceptor?.onGasEstimate?.(gasEstimation);

  const adjustedGasEstimation =
    options?.interceptor?.transformGasEstimation?.(gasEstimation) ??
    gasEstimation;

  const overridesToBeCalled = {
    ...overrides,
    ...adjustedGasEstimation,
  };

  options?.interceptor?.onCallArgs?.(...deployArgs, overridesToBeCalled);

  const contract = await contractFactory.deploy(...deployArgs, {
    ...overridesToBeCalled,
  });

  return contract;
}

export async function callWithGasEstimation<A extends any[] = any[]>(
  hre: HardhatRuntimeEnvironment,
  method: BaseContractMethod,
  args: ContractMethodArgs<A>,
  options?: {
    interceptor?: IGasEstimationInterceptor;
    deployer?: HardhatEthersSigner;
  },
) {
  const { ethers } = hre;

  let deployer = options?.deployer;
  if (deployer == null) {
    const [signer] = await ethers.getSigners();
    deployer = signer;
  }

  let overrides: Omit<TransactionRequest, "data"> = {};

  const fragment = method.fragment;

  if (fragment.inputs.length + 1 === args.length) {
    overrides = await copyOverrides(args.pop());
  }

  const feeData = await ethers.provider.getFeeData();
  const { gasPrice, maxFeePerGas, maxPriorityFeePerGas } = feeData;

  const gasLimit = await method.estimateGas(...args);

  const gasEstimation = {
    gasPrice,
    maxFeePerGas,
    maxPriorityFeePerGas,
    gasLimit,
  };

  options?.interceptor?.onGasEstimate?.(gasEstimation);

  const adjustedGasEstimation =
    options?.interceptor?.transformGasEstimation?.(gasEstimation) ??
    gasEstimation;

  const overridesToBeCalled = {
    ...overrides,
    ...adjustedGasEstimation,
  };

  options?.interceptor?.onCallArgs?.(...args, overridesToBeCalled);

  return method(...args, overridesToBeCalled);
}
