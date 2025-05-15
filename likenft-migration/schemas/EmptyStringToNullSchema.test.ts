import { EmptyStringToNullSchema } from './EmptyStringToNullSchema';

describe('EmptyStringToNullSchema', () => {
  it('should parse', () => {
    expect(EmptyStringToNullSchema.nullable().parse(null)).toEqual(null);
    expect(EmptyStringToNullSchema.parse('')).toEqual(null);
    expect(EmptyStringToNullSchema.parse('abc')).toEqual('abc');
  });
});
