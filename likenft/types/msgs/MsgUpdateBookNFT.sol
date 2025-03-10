// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {BookConfig} from "../BookConfig.sol";

struct MsgUpdateBookNFT {
    address classId;
    BookConfig config;
}
