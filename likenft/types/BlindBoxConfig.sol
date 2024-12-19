// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {MintPeriod} from "./MintPeriod.sol";

/**
 * message BlindBoxConfig {
 *   repeated MintPeriod mint_periods = 1 [(gogoproto.nullable) = false];
 *   google.protobuf.Timestamp reveal_time = 2 [(gogoproto.stdtime) = true, (gogoproto.nullable) = false];
 * }
 */
struct BlindBoxConfig {
    MintPeriod[] mint_periods;
    uint64 reveal_time;
}
