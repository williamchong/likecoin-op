import { RawCreateParams, z } from 'zod';

function IsEnum<U extends string, T extends Readonly<[U, ...U[]]>>(
  values: T,
  params?: RawCreateParams
) {
  return z.enum(values, params).transform((value) => ({
    isEnum: true as const,
    value,
  }));
}

export const UnrecognizableEnum = z.string().transform((value) => ({
  isEnum: false as const,
  value,
}));

export function RecognizableEnum<
  U extends string,
  T extends Readonly<[U, ...U[]]>
>(values: T, params?: RawCreateParams) {
  return z.union([IsEnum(values, params), UnrecognizableEnum]);
}
