// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

struct ClassParent {
    ClassParentType type_;
    string iscn_id_prefix;
    uint64 iscn_version_at_mint;
    address account;
}

enum ClassParentType {
    UNKNOWN,
    ISCN,
    ACCOUNT
}
