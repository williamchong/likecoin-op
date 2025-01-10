import z from "zod";

import { useConfig } from "../hooks/useConfig";
import { useMakeAPI } from "./api";

const MigrationRecordSchema = z.object({
  cosmos_tx_hash: z.string(),
  eth_tx_hash: z.string(),
  cosmos_address: z.string(),
  eth_address: z.string(),
});

const MigrationRecordStatusSchema = z.enum([
  "unknown",
  "pending",
  "success",
  "failed",
]);

const ResponseSchema = z.object({
  migration_record: MigrationRecordSchema,
  status: MigrationRecordStatusSchema,
});

export const queryKey = (cosmosTxHash: string) => [
  "migration_record",
  cosmosTxHash,
];

export default function useGetMigrationRecord(cosmosTxHash: string) {
  const config = useConfig();
  return useMakeAPI({
    method: "GET",
    url: new URL(
      `${config.apiBaseUri}/migration_record/${cosmosTxHash}`,
    ).toString(),
    responseSchema: ResponseSchema,
  });
}
