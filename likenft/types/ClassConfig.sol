// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {BlindBoxConfig} from "./BlindBoxConfig.sol";

/**
 * message ClassConfig {
 *   bool burnable = 1;
 *   uint64 max_supply = 2;
 *   BlindBoxConfig blind_box_config = 3 [(gogoproto.nullable) = true];
 * }
 */
struct ClassConfig {
    bool burnable;
    uint64 max_supply;
    BlindBoxConfig blind_box_config;
}
