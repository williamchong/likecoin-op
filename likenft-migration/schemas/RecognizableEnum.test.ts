import { RecognizableEnum } from './RecognizableEnum';

describe('RecognizableEnum', () => {
  it('should parse', () => {
    expect(RecognizableEnum(['a', 'b']).nullable().parse(null)).toEqual(null);
    expect(RecognizableEnum(['a', 'b']).parse('a')).toEqual({
      isEnum: true,
      value: 'a',
    });
    expect(RecognizableEnum(['a', 'b']).parse('b')).toEqual({
      isEnum: true,
      value: 'b',
    });
    expect(RecognizableEnum(['a', 'b']).parse('c')).toEqual({
      isEnum: false,
      value: 'c',
    });
  });
});
