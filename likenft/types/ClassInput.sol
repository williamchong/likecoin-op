// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

struct ClassInput {
    string name;
    string symbol;
    string metadata;
    ClassConfig config;
}

struct ClassConfig {
    uint64 max_supply;
}