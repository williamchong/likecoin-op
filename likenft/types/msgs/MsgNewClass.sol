// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ClassInput} from "../ClassInput.sol";
import {ClassParentInput} from "../ClassParentInput.sol";

struct MsgNewClass {
    address creator;
    ClassParentInput parent;
    ClassInput input;
}
