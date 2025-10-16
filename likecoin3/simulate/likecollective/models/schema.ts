import { z } from "zod";

const AddressSchema = z
  .string()
  .regex(/^0x[0-9a-fA-F]{40}$/)
  .transform((val) => val as `0x${string}`);
export type Address = z.infer<typeof AddressSchema>;

const BigIntStringSchema = z
  .string()
  .regex(/^\d+$/)
  .transform((val) => BigInt(val));
export type BigIntString = z.infer<typeof BigIntStringSchema>;

const NewStakePositionCallParamsSchema = z.object({
  type: z.literal("newStakePosition"),
  bookNFT: AddressSchema,
  account: AddressSchema,
  amount: BigIntStringSchema,
});
export type NewStakePositionCallParams = z.infer<
  typeof NewStakePositionCallParamsSchema
>;

const IncreaseStakeToPositionCallParamsSchema = z.object({
  type: z.literal("increaseStakeToPosition"),
  tokenID: BigIntStringSchema,
  account: AddressSchema,
  amount: BigIntStringSchema,
});
export type IncreaseStakeToPositionCallParams = z.infer<
  typeof IncreaseStakeToPositionCallParamsSchema
>;

const DecreaseStakePositionCallParamsSchema = z.object({
  type: z.literal("decreaseStakePosition"),
  tokenID: BigIntStringSchema,
  account: AddressSchema,
  amount: BigIntStringSchema,
});
export type DecreaseStakePositionCallParams = z.infer<
  typeof DecreaseStakePositionCallParamsSchema
>;

const RemoveStakePositionCallParamsSchema = z.object({
  type: z.literal("removeStakePosition"),
  tokenID: BigIntStringSchema,
  account: AddressSchema,
});
export type RemoveStakePositionCallParams = z.infer<
  typeof RemoveStakePositionCallParamsSchema
>;

const ClaimRewardsCallParamsSchema = z.object({
  type: z.literal("claimRewards"),
  tokenID: BigIntStringSchema,
  account: AddressSchema,
});
export type ClaimRewardsCallParams = z.infer<
  typeof ClaimRewardsCallParamsSchema
>;

const ClaimAllRewardsCallParamsSchema = z.object({
  type: z.literal("claimAllRewards"),
  account: AddressSchema,
});
export type ClaimAllRewardsCallParams = z.infer<
  typeof ClaimAllRewardsCallParamsSchema
>;

const RestakeRewardPositionCallParamsSchema = z.object({
  type: z.literal("restakeRewardPosition"),
  tokenID: BigIntStringSchema,
  account: AddressSchema,
});
export type RestakeRewardPositionCallParams = z.infer<
  typeof RestakeRewardPositionCallParamsSchema
>;

const DepositRewardCallParamsSchema = z.object({
  type: z.literal("depositReward"),
  bookNFT: AddressSchema,
  amount: BigIntStringSchema,
  account: AddressSchema,
});
export type DepositRewardCallParams = z.infer<
  typeof DepositRewardCallParamsSchema
>;

const TransferStakePositionCallParamsSchema = z.object({
  type: z.literal("transferStakePosition"),
  from: AddressSchema,
  to: AddressSchema,
  tokenID: BigIntStringSchema,
});
export type TransferStakePositionCallParams = z.infer<
  typeof TransferStakePositionCallParamsSchema
>;

const CallParamsSchema = z.discriminatedUnion("type", [
  NewStakePositionCallParamsSchema,
  IncreaseStakeToPositionCallParamsSchema,
  DecreaseStakePositionCallParamsSchema,
  RemoveStakePositionCallParamsSchema,
  ClaimRewardsCallParamsSchema,
  ClaimAllRewardsCallParamsSchema,
  RestakeRewardPositionCallParamsSchema,
  DepositRewardCallParamsSchema,
  TransferStakePositionCallParamsSchema,
]);
export type CallParams = z.infer<typeof CallParamsSchema>;

const StateSchema = z.object({
  bookPendingRewards: z.record(AddressSchema, BigIntStringSchema),
  bookStakedAmounts: z.record(AddressSchema, BigIntStringSchema),
  userBalance: z.record(AddressSchema, BigIntStringSchema),
  userPendingRewards: z.record(
    AddressSchema,
    z.record(AddressSchema, BigIntStringSchema),
  ),
  userStakedAmounts: z.record(
    AddressSchema,
    z.record(AddressSchema, BigIntStringSchema),
  ),
});
export type State = z.infer<typeof StateSchema>;

const SetupSchema = z.object({
  deployer: AddressSchema,
  accounts: z.array(
    z.object({
      address: AddressSchema,
      likecoin: BigIntStringSchema,
    }),
  ),
});
export type Setup = z.infer<typeof SetupSchema>;

export const SimulationSchema = z.object({
  name: z.string(),
  setup: SetupSchema,
  steps: z.array(
    z.object({
      name: z.string(),
      calls: z.array(CallParamsSchema),
      expectedLogs: z.array(z.record(z.string(), z.any())),
      expectedState: StateSchema,
    }),
  ),
});
export type Simulation = z.infer<typeof SimulationSchema>;
