import { ContractMethodArgs } from "ethers";
import { GasEstimation } from "./types/GasEstimation";
import prompt from "prompt-sync";

export interface IGasEstimationInterceptor {
  onGasEstimate: (gasEstimation: GasEstimation) => void;
  onCallArgs: (...args: ContractMethodArgs<any[]>) => void;
  transformGasEstimation: (gasEstimation: GasEstimation) => GasEstimation;
}

export class BaseGasEstimationInterceptor implements IGasEstimationInterceptor {
  onGasEstimate(gasEstimation: GasEstimation) {}
  onCallArgs(...args: any[]) {}
  transformGasEstimation(gasEstimation: GasEstimation) {
    return gasEstimation;
  }
}

export class GasEstimationInterceptors implements IGasEstimationInterceptor {
  private interceptors: IGasEstimationInterceptor[] = [];

  constructor(...interceptors: IGasEstimationInterceptor[]) {
    this.interceptors = interceptors;
  }

  addInterceptor(interceptor: IGasEstimationInterceptor) {
    this.interceptors.push(interceptor);
  }

  onGasEstimate(gasEstimation: GasEstimation) {
    for (const interceptor of this.interceptors) {
      interceptor.onGasEstimate(gasEstimation);
    }
  }

  onCallArgs(...args: any[]) {
    for (const interceptor of this.interceptors) {
      interceptor.onCallArgs(...args);
    }
  }

  transformGasEstimation(gasEstimation: GasEstimation) {
    for (const interceptor of this.interceptors) {
      gasEstimation = interceptor.transformGasEstimation(gasEstimation);
    }
    return gasEstimation;
  }
}

export class GasEstimationLoggerInterceptor extends BaseGasEstimationInterceptor {
  private operation: string;

  constructor(operation: string) {
    super();
    this.operation = operation;
  }

  onGasEstimate(gasEstimation: GasEstimation) {
    console.log(
      `Operation '${this.operation}', Gas estimation:,`,
      gasEstimation,
    );
  }

  onCallArgs(...args: any[]) {
    console.log(`Operation '${this.operation}' to be call with args:`, args);
  }
}

export class GasEstimationConfirmationInterceptor extends BaseGasEstimationInterceptor {
  private p = prompt({
    sigint: true,
  });

  onCallArgs(...args: any[]) {
    this.p("Confirm call args and bumped gas estimation");
  }
}

export class GasEstimationAdjustmentInterceptor extends BaseGasEstimationInterceptor {
  private limitBump: number;
  private feeIncrement: number;

  constructor(limitBump: number, feeIncrement: number) {
    super();
    this.limitBump = limitBump;
    this.feeIncrement = feeIncrement;
  }

  transformGasEstimation(gasEstimation: GasEstimation) {
    return {
      ...gasEstimation,
      maxFeePerGas: gasEstimation.maxFeePerGas + BigInt(this.feeIncrement),
      maxPriorityFeePerGas:
        gasEstimation.maxPriorityFeePerGas + BigInt(this.feeIncrement),
      gasLimit: gasEstimation.gasLimit + BigInt(this.limitBump),
    };
  }
}

export class GasEstimationPickParamsInterceptor extends BaseGasEstimationInterceptor {
  transformGasEstimation(gasEstimation: GasEstimation) {
    return {
      gasPrice: null,
      maxFeePerGas: gasEstimation.maxFeePerGas,
      maxPriorityFeePerGas: gasEstimation.maxPriorityFeePerGas,
      gasLimit: gasEstimation.gasLimit,
    };
  }
}

function makeDefaultInterceptors(
  operation: string,
  limitBump: number,
  feeIncrement: number = 0,
  confirmation: boolean = true,
) {
  const interceptors = new GasEstimationInterceptors(
    new GasEstimationLoggerInterceptor(operation),
    new GasEstimationAdjustmentInterceptor(limitBump, feeIncrement),
    new GasEstimationPickParamsInterceptor(),
  );

  if (confirmation) {
    interceptors.addInterceptor(new GasEstimationConfirmationInterceptor());
  }

  return interceptors;
}

export default makeDefaultInterceptors;
