export type Params = Record<string, any>;

export class NonNullParameter {
  params: Params;

  constructor(params: Params) {
    this.params = params;
  }

  get(): Params {
    return Object.entries(this.params)
      .filter(([_, value]) => value !== null)
      .filter(([_, value]) => value !== undefined)
      .reduce<Params>((acc, [key, value]) => {
        acc[key] = value;
        return acc;
      }, {});
  }
}
