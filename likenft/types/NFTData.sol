// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ClassParent} from "./ClassParent.sol";

/**
 * message NFTData {
 *   bytes metadata = 1 [
 *     (gogoproto.nullable) = false,
 *     (gogoproto.customtype) = "JsonInput"
 *   ];
 *   ClassParent class_parent = 2 [(gogoproto.nullable) = false];
 *   bool to_be_revealed = 3;
 * }
 */
struct NFTData {
    string metadata;
    ClassParent class_parent;
}
