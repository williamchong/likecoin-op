import { useMemo } from "react";
import z from "zod";

import { APIError } from "../models/error";

export const ErrorResponseSchema = z.object({
  error_description: z.string(),
});

export type ErrorResponse = z.infer<typeof ErrorResponseSchema>;

export interface API<Request, Response> {
  method: RequestInit["method"];
  url: string;
  requestSchema?: z.ZodSchema<Request>;
  responseSchema: z.ZodSchema<Response>;
}

export default function makeAPI<Response, Request = void>(
  api: API<Request, Response>,
) {
  return async (request: Request) => {
    const resp = await fetch(api.url, {
      method: api.method,
      body: JSON.stringify(request),
      headers: {
        "Content-Type": "application/json",
      },
    });
    if (resp.ok) {
      const d = await api.responseSchema.parseAsync(await resp.json());
      return d;
    }
    const t = await resp.text();
    try {
      const j = JSON.parse(t);
      const errorResponse = ErrorResponseSchema.safeParse(j);
      if (errorResponse.data != null) {
        throw new APIError(errorResponse.data.error_description);
      }
    } catch {
      // empty
    }
    throw new Error(t);
  };
}

export function useMakeAPI<Response, Request = void>(
  api: API<Request, Response>,
) {
  return useMemo(() => makeAPI(api), [api]);
}
