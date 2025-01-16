// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ClassConfig} from "./ClassConfig.sol";
import {ClassParent} from "./ClassParent.sol";

/**
 * Class {
 *   string id          = 1;
 *   string name        = 2;
 *   string symbol      = 3;
 *   string description = 4;
 *   string uri         = 5;
 *   string uri_hash    = 6;
 *   google.protobuf.Any data = 7;
 * }
 */
struct ClassStorage {
    string id;
    string name;
    string symbol;
    string description;
    string uri;
    string uri_hash;
    ClassDataStorage data;
}

/**
 * message ClassData {
 *   bytes metadata = 1 [
 *     (gogoproto.nullable) = false,
 *     (gogoproto.customtype) = "JsonInput"
 *   ];
 *   ClassParent parent = 2 [(gogoproto.nullable) = false];
 *   ClassConfig config = 3 [(gogoproto.nullable) = false];
 *   BlindBoxState blind_box_state = 4 [(gogoproto.nullable) = false];
 * }
 */
struct ClassDataStorage {
    string metadata;
    ClassParent parent;
    ClassConfig config;
}
