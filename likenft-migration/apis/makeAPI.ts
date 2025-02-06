import type { NuxtAxiosInstance } from '@nuxtjs/axios';
import { Method } from 'axios';
import { z } from 'zod';

export function makeAPI<Resp, Req = void>(def: {
  url: string;
  method: Method;
  requestSchema?: z.ZodSchema<Req>;
  responseSchema: z.ZodSchema<Resp>;
}): (axios: NuxtAxiosInstance) => (req: Req) => Promise<Resp> {
  return (axios: NuxtAxiosInstance) => async (req: Req) => {
    const resp = await axios.$request({
      method: def.method,
      url: def.url,
      data: req,
    });
    const response = def.responseSchema.parse(resp);
    return response;
  };
}
