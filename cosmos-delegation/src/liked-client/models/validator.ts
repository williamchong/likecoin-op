import { z } from "zod";

export const ValidatorSchema = z.object({
  operator_address: z.string(),
  tokens: z.coerce.bigint(),
  description: z.object({
    moniker: z.string(),
    identity: z.string(),
    website: z.string(),
    security_contact: z.string(),
    details: z.string(),
  }),
});

export type Validator = z.infer<typeof ValidatorSchema>;
