import z from "zod";

import { useConfig } from "../hooks/useConfig";
import { useMakeAPI } from "./api";

const RequestSchema = z.object({
  cosmos_tx_hash: z.string(),
});

const ResponseSchema = z.object({
  eth_tx_hash: z.string(),
});

export default function useInitLikeCoinMigrationFromCosmos() {
  const config = useConfig();
  return useMakeAPI({
    method: "POST",
    url: `${config.apiBaseUri}/init_likecoin_migration_from_cosmos`,
    requestSchema: RequestSchema,
    responseSchema: ResponseSchema,
  });
}
