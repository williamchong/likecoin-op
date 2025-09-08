import { execSync } from "node:child_process";
import { ChainCoin, ChainCoinSchema } from "./models/chain-coin";
import { Validator, ValidatorSchema } from "./models/validator";
import { z } from "zod";

export interface LikedClient {
  queryStakingValidators(
    filterValidatorAddresses: string[],
  ): Promise<Validator[]>;
  queryStakingValidator(validatorAddress: string): Promise<Validator | null>;

  queryBankTotal(denom: string): Promise<ChainCoin | null>;

  queryBankBalances(address: string, denom: string): Promise<ChainCoin | null>;

  getTxStakingDelegateCmd(
    validatorAddress: string,
    delegatorAddress: string,
    amount: bigint,
    denom: string,
  ): string;
}

class CmdLikedClient implements LikedClient {
  constructor(
    private readonly cli: string,
    private readonly cliFlags: string,
  ) {}

  async queryStakingValidators(
    filterValidatorAddresses: string[] = [],
  ): Promise<Validator[]> {
    if (filterValidatorAddresses.length === 0) {
      // --page-key doesn't work
      // return await this.queryStakingValidatorsAll();
      throw new Error("Query all staking validators is not supported");
    }
    return await this.queryStakingValidatorsByValidatorAddresses(
      filterValidatorAddresses,
    );
  }

  /**
   * --page-key doesn't work
   *
   * private async queryStakingValidatorsAll(): Promise<Validator[]> {
   *   let { validators, nextKey } = await this.queryStakingValidatorsByPageKey();
   *   while (nextKey) {
   *     const { validators: newValidators, nextKey: newNextKey } =
   *       await this.queryStakingValidatorsByPageKey(nextKey);
   *     validators.push(...newValidators);
   *     nextKey = newNextKey;
   *   }
   *   return validators;
   * }
   *
   * private getQueryStakingValidatorsCmd(
   *   pageKey?: string,
   *   limit: number = 10,
   * ): string {
   *   return `${this.cli} query staking validators ${pageKey ? `--page-key ${pageKey}` : ""} --limit ${limit} ${this.cliFlags}`;
   * }
   *
   * private async queryStakingValidatorsByPageKey(
   *   pageKey?: string,
   * ): Promise<{ validators: Validator[]; nextKey: string | null }> {
   *   const responseSchema = z.object({
   *     validators: z.array(ValidatorSchema),
   *     pagination: z.object({
   *       next_key: z.string().nullable(),
   *     }),
   *   });
   *   try {
   *     const cmd = this.getQueryStakingValidatorsCmd(pageKey);
   *     const result = execSync(cmd);
   *     const resultJson = responseSchema.parse(JSON.parse(result.toString()));
   *     return {
   *       validators: resultJson.validators,
   *       nextKey: resultJson.pagination.next_key,
   *     };
   *   } catch (e) {
   *     throw new Error(e.message);
   *   }
   * }
   **/

  private getQueryStakingValidatorCmd(validatorAddress: string): string {
    return `${this.cli} query staking validator ${validatorAddress} -o json ${this.cliFlags}`;
  }

  async queryStakingValidator(
    validatorAddress: string,
  ): Promise<Validator | null> {
    try {
      const cmd = this.getQueryStakingValidatorCmd(validatorAddress);
      const result = execSync(cmd);
      const resultJson = ValidatorSchema.parse(JSON.parse(result.toString()));
      return resultJson;
    } catch (e) {
      throw new Error(e.message);
    }
  }

  private async queryStakingValidatorsByValidatorAddresses(
    validatorAddresses: string[],
  ): Promise<Validator[]> {
    const maybeValidators = await Promise.all(
      validatorAddresses.map(
        async (validatorAddress) =>
          await this.queryStakingValidator(validatorAddress).catch((err) => {
            console.error(
              `Failed to query validator ${validatorAddress}: ${err.message}`,
            );
            return null;
          }),
      ),
    );
    return maybeValidators.filter(
      (validator) => validator !== null,
    ) as Validator[];
  }

  private getQueryBankTotalCmd(): string {
    return `${this.cli} query bank total -o json ${this.cliFlags}`;
  }

  async queryBankTotal(denom: string): Promise<ChainCoin | null> {
    const responseSchema = z.object({
      supply: z.array(ChainCoinSchema),
    });
    try {
      const cmd = this.getQueryBankTotalCmd();
      const result = execSync(cmd);
      const resultJson = responseSchema.parse(JSON.parse(result.toString()));
      return resultJson.supply.find((supply) => supply.denom === denom) ?? null;
    } catch (e) {
      throw new Error(e.message);
    }
  }

  private getQueryBankBalancesCmd(address: string): string {
    return `${this.cli} query bank balances ${address} -o json ${this.cliFlags}`;
  }

  async queryBankBalances(
    address: string,
    denom: string,
  ): Promise<ChainCoin | null> {
    const responseSchema = z.object({
      balances: z.array(ChainCoinSchema),
    });
    try {
      const cmd = this.getQueryBankBalancesCmd(address);
      const result = execSync(cmd);
      const resultJson = responseSchema.parse(JSON.parse(result.toString()));
      return (
        resultJson.balances.find((balance) => balance.denom === denom) ?? null
      );
    } catch (e) {
      throw new Error(e.message);
    }
  }

  getTxStakingDelegateCmd(
    validatorAddress: string,
    delegatorAddress: string,
    amount: bigint,
    denom: string,
  ): string {
    return `${this.cli} tx staking delegate ${validatorAddress} ${amount}${denom} --from ${delegatorAddress} -o json ${this.cliFlags}`;
  }
}

export default CmdLikedClient;
