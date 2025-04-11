import type { NuxtAxiosInstance } from '@nuxtjs/axios';
import { Method } from 'axios';
import { z, ZodTypeDef } from 'zod';

export function makeAPI<Resp, Req = void>(def: {
  url: string;
  method: Method;
  requestSchema?: z.ZodSchema<any, ZodTypeDef, Req>;
  responseSchema: z.ZodSchema<Resp, ZodTypeDef, any>;
}): (axios: NuxtAxiosInstance) => (req: Req) => Promise<Resp> {
  return (axios: NuxtAxiosInstance) => async (req: Req) => {
    const data = def.requestSchema?.parse(req);
    try {
      const resp = await axios.$request({
        method: def.method,
        url: def.url,
        data,
      });
      const response = def.responseSchema.parse(resp);
      return response;
    } catch (err) {
      alert('Unexpected error happened. Please try again later. \n\n' + err);
      throw err;
    }
  };
}
