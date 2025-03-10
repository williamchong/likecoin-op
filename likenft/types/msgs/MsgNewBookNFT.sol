// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {BookConfig} from "../BookConfig.sol";

struct MsgNewBookNFT {
    address creator;
    address[] updaters;
    address[] minters;
    BookConfig config;
}
