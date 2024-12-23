// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ClassConfig} from "./ClassConfig.sol";

struct ClassInput {
    string name;
    string symbol;
    string description;
    string uri;
    string uri_hash;
    string metadata;
    ClassConfig config;
}
