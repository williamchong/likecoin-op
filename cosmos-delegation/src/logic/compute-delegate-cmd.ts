import Decimal from "decimal.js";
import { LikedClient } from "../liked-client";
import { Validator } from "../liked-client/models/validator";

interface ValidatorStakeToAverage {
  operator_address: string;
  tokens: bigint;
  average_tokens_with_delegated_tokens: bigint;
  tokens_to_average: bigint;
}

interface SortedValidatorInfo {
  __type: "sorted";
  sortedValidators: Validator[];
}

function sortValidators(validators: Validator[]): SortedValidatorInfo {
  function compareValidatorsByAmountAsc(a: Validator, b: Validator): number {
    if (a.tokens === b.tokens) {
      return 0;
    }
    if (b.tokens > a.tokens) {
      return -1;
    }
    return 1;
  }

  return {
    __type: "sorted",
    sortedValidators: validators.sort(compareValidatorsByAmountAsc),
  };
}

const predicateValidatorAboveThreshold =
  (thresholdAmount: bigint) => (validator: ValidatorStakeToAverage) => {
    return validator.tokens_to_average >= thresholdAmount;
  };

function getValidatorsAboveThreshold(
  validators: ValidatorStakeToAverage[],
  thresholdAmount: bigint,
): ValidatorStakeToAverage[] {
  return validators.filter(predicateValidatorAboveThreshold(thresholdAmount));
}

function* pickValidators(
  { sortedValidators }: SortedValidatorInfo,
  requestedCount: number,
): Generator<Validator[]> {
  function* _pickValidators(
    pickedValidators: Validator[],
    validators: Validator[],
    count: number,
  ): Generator<Validator[]> {
    if (pickedValidators.length === requestedCount) {
      yield pickedValidators;
      return;
    }

    if (validators.length === 0 || count === 0) {
      return;
    }

    yield* _pickValidators(
      [...pickedValidators, validators[0]],
      validators.slice(1),
      count - 1,
    );
    yield* _pickValidators(pickedValidators, validators.slice(1), count);
  }

  yield* _pickValidators([], sortedValidators, requestedCount);
}

function* generateValidatorCombinations(
  sortedValidators: SortedValidatorInfo,
): Generator<Validator[]> {
  const count = sortedValidators.sortedValidators.length;
  for (let i = count; i > 0; i--) {
    yield* pickValidators(sortedValidators, i);
  }
}

function computeAverageAmountWithDelegatedAmount(
  validatorCombination: Validator[],
  delegatedAmount: bigint,
): ValidatorStakeToAverage[] {
  const averageAmountWithDelegatedAmount = new Decimal(
    validatorCombination.reduce(
      (acc, validator) => acc + validator.tokens,
      delegatedAmount,
    ),
  )
    .div(validatorCombination.length)
    .floor();
  return validatorCombination.map((validator) => ({
    ...validator,
    average_tokens_with_delegated_tokens: BigInt(
      averageAmountWithDelegatedAmount.toString(),
    ),
    tokens_to_average: BigInt(
      averageAmountWithDelegatedAmount
        .minus(new Decimal(validator.tokens))
        .toString(),
    ),
  }));
}

function* generateAverageAmountWithDelegatedAmount(
  validatorCombinations: Generator<Validator[]>,
  delegatedAmount: bigint,
): Generator<ValidatorStakeToAverage[]> {
  for (const validatorCombination of validatorCombinations) {
    yield computeAverageAmountWithDelegatedAmount(
      validatorCombination,
      delegatedAmount,
    );
  }
}

interface CombinationResult {
  combination: ValidatorStakeToAverage[];
  ok: boolean;
  hasNegativeStakeAmountToAverage: boolean;
  hasAboveThresholdStakeAmountToAverage: boolean;
}

function logCombinationResult(
  logger: Console,
  result: CombinationResult,
  thresholdAmount: bigint,
) {
  logger.error("Combination:");
  for (const validator of result.combination) {
    const percentThreshold = new Decimal(
      validator.average_tokens_with_delegated_tokens,
    )
      .div(thresholdAmount)
      .mul(100);
    logger.error(
      `${validator.operator_address}: ${validator.tokens.toLocaleString()} (Diff: ${validator.tokens_to_average.toLocaleString()}) -> Average: ${validator.average_tokens_with_delegated_tokens.toLocaleString()} (${percentThreshold.toSignificantDigits(2)}% of Threshold)`,
    );
  }
  logger.error("--------------------------------");
  if (result.hasNegativeStakeAmountToAverage) {
    logger.error("Error: Combination has negative stake amount to average");
  }
  if (result.hasAboveThresholdStakeAmountToAverage) {
    logger.error(
      "Error: Combination has above threshold stake amount to average",
    );
  }
  if (result.ok) {
    logger.error("Combination ok");
  } else {
    logger.error("Combination not ok");
  }
  logger.error("================================================");
}

function determineCombination(
  combination: ValidatorStakeToAverage[],
  thresholdAmount: bigint,
): CombinationResult {
  const hasNegativeStakeAmountToAverage = combination.some(
    (validator) => validator.tokens_to_average < 0n,
  );
  const hasAboveThresholdStakeAmountToAverage =
    getValidatorsAboveThreshold(combination, thresholdAmount).length > 0;
  return {
    combination,
    ok:
      !hasNegativeStakeAmountToAverage &&
      !hasAboveThresholdStakeAmountToAverage,
    hasNegativeStakeAmountToAverage,
    hasAboveThresholdStakeAmountToAverage,
  };
}

function getOKCombination(
  combinations: Generator<ValidatorStakeToAverage[]>,
  thresholdAmount: bigint,
): ValidatorStakeToAverage[] | null {
  for (const combination of combinations) {
    const c = determineCombination(combination, thresholdAmount);
    logCombinationResult(console, c, thresholdAmount);
    if (c.ok) {
      return c.combination;
    }
  }
  return null;
}

const getStakeCommandGenerator = (
  likedClient: LikedClient,
  delegatorAddress: string,
  denom: string,
): ((combination: ValidatorStakeToAverage[]) => Generator<string>) => {
  return function* (combination: ValidatorStakeToAverage[]): Generator<string> {
    for (const validator of combination) {
      yield likedClient.getTxStakingDelegateCmd(
        validator.operator_address,
        delegatorAddress,
        validator.tokens_to_average,
        denom,
      );
    }
  };
};

async function computeDelegateCmd(
  likedClient: LikedClient,
  delegatorAddress: string,
  selectedValidatorAddresses: string[],
  denom: string,
  thresholdPercentage: Decimal,
): Promise<string[]> {
  const delegatedAmount =
    (await likedClient.queryBankBalances(delegatorAddress, denom))?.amount ??
    0n;
  console.error("Delegated Amount:", delegatedAmount.toLocaleString());

  const bankTotal = (await likedClient.queryBankTotal(denom))?.amount ?? 0n;
  console.error("Bank Total:", bankTotal.toLocaleString());
  const thresholdAmount = BigInt(
    Decimal(bankTotal).mul(thresholdPercentage).div(100).floor().toString(),
  );
  console.error(
    "Threshold Amount:",
    thresholdAmount.toLocaleString(),
    `(${thresholdPercentage}% of Bank Total)`,
  );

  const validators = await likedClient.queryStakingValidators(
    selectedValidatorAddresses,
  );
  const sortedValidators = sortValidators(validators);
  const combinations = generateValidatorCombinations(sortedValidators);
  const okCombination = getOKCombination(
    generateAverageAmountWithDelegatedAmount(combinations, delegatedAmount),
    thresholdAmount,
  );
  if (!okCombination) {
    return [];
  }

  const generateStakeCommands = getStakeCommandGenerator(
    likedClient,
    delegatorAddress,
    denom,
  );
  const commands: string[] = [];
  for (const cmd of generateStakeCommands(okCombination)) {
    commands.push(cmd);
  }
  return commands;
}

export default computeDelegateCmd;
