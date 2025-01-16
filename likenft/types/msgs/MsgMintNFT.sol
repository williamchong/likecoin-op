// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {NFTInput} from "../NFTInput.sol";

struct MsgMintNFT {
    address creator;
    string class_id;
    NFTInput input;
}
