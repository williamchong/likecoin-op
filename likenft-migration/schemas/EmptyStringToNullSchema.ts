import { z } from 'zod';

export const EmptyStringToNullSchema = z
  .string()
  .transform((s) => {
    if (s === '') {
      return null;
    }
    return s;
  })
  .pipe(z.string().min(1).nullable());
