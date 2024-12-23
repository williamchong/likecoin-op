// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/**
 * message MintPeriod {
 *   google.protobuf.Timestamp start_time = 1 [(gogoproto.stdtime) = true, (gogoproto.nullable) = false];
 *   repeated string allowed_addresses = 2 ;
 *   uint64 mint_price = 3;
 * }
 */
struct MintPeriod {
    uint64 start_time;
    string[] allowed_addresses;
    uint64 mint_price;
}
