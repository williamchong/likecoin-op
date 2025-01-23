// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ClassConfig} from "./ClassConfig.sol";

struct ClassStorage {
    string name;
    string symbol;
    ClassDataStorage data;
}

struct ClassDataStorage {
    string metadata;
    ClassConfig config;
}
