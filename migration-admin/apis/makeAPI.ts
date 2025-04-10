import type { NuxtAxiosInstance } from "@nuxtjs/axios";
import { AxiosRequestConfig, Method } from "axios";
import { z, ZodTypeDef } from "zod";

export function makeAPI<Resp, Req = void>(def: {
  url: string;
  method: "GET" | "get";
  requestSchema?: z.ZodSchema<URLSearchParams, ZodTypeDef, Req>;
  responseSchema: z.ZodSchema<Resp>;
}): (axios: NuxtAxiosInstance) => (req: Req) => Promise<Resp>;
export function makeAPI<Resp, Req = void>(def: {
  url: string;
  method: "DELETE" | "delete";
  requestSchema?: z.ZodSchema<URLSearchParams, ZodTypeDef, Req>;
  responseSchema: z.ZodSchema<Resp>;
}): (axios: NuxtAxiosInstance) => (req: Req) => Promise<Resp>;
export function makeAPI<Resp, Req = void>(def: {
  url: string;
  method: Method;
  requestSchema?: z.ZodSchema<Req>;
  responseSchema: z.ZodSchema<Resp>;
}): (axios: NuxtAxiosInstance) => (req: Req) => Promise<Resp> {
  return (axios: NuxtAxiosInstance) => async (req: Req) => {
    const request = def.requestSchema ? def.requestSchema.parse(req) : req;
    const requestParams: Pick<
      Partial<AxiosRequestConfig>,
      "data" | "params"
    > = def.method === "GET" || def.method === "get"
      ? {
          params: request,
        }
      : {
          data: request,
        };
    const resp = await axios.$request({
      method: def.method,
      url: def.url,
      ...requestParams,
    });
    const response = def.responseSchema.parse(resp);
    return response;
  };
}
