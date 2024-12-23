// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

/**
 * message ClassParent {
 *   ClassParentType type = 1;
 *   string iscn_id_prefix = 2 [(gogoproto.nullable) = true];
 *   uint64 iscn_version_at_mint = 3 [(gogoproto.nullable) = true];
 *   string account = 4 [(gogoproto.nullable) = true];
 * }
 */
struct ClassParent {
    ClassParentType type_;
    string iscn_id_prefix;
    uint64 iscn_version_at_mint;
    address account;
}

/**
 * enum ClassParentType {
 *   UNKNOWN = 0;
 *   ISCN = 1;
 *   ACCOUNT = 2;
 * }
 */
enum ClassParentType {
    UNKNOWN,
    ISCN,
    ACCOUNT
}
