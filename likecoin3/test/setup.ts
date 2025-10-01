import { isAddressEqual } from "viem";
import chai from "chai";

chai.Assertion.addMethod("equalAddress", function (expected: `0x${string}`) {
  const actual = this._obj as `0x${string}`;
  this.assert(
    isAddressEqual(actual, expected),
    `expected #{this} to equal address #{exp}`,
    `expected #{this} to not equal address #{exp}`,
    expected,
    actual,
  );
});
