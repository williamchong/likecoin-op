// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ClassConfig} from "./ClassConfig.sol";
import {ClassParent} from "./ClassParent.sol";

struct ClassStorage {
    string name;
    string symbol;
    string description;
    string uri;
    string uri_hash;
    ClassDataStorage data;
}

struct ClassDataStorage {
    string metadata;
    ClassParent parent;
    ClassConfig config;
}
