// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {NFTInput} from "../NFTInput.sol";

/**
 * message MsgMintNFT {
 *   string creator = 1;
 *   string class_id = 2;
 *   string id = 3;
 *   NFTInput input = 4 [(gogoproto.nullable) = true];
 * }
 */
struct MsgMintNFT {
    address creator;
    address class_id;
    // string id;
    NFTInput input;
}
