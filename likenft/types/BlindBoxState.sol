// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/**
 * message BlindBoxState {
 *   uint64 content_count = 1;
 *   bool to_be_revealed = 2;
 * }
 */
struct BlindBoxState {
    uint64 content_count;
    bool to_be_revealed;
}
