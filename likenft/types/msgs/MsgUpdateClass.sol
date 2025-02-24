// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ClassInput} from "../ClassInput.sol";

struct MsgUpdateClass {
    address class_id;
    ClassInput input;
}
