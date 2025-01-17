// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {NFTInput} from "../NFTInput.sol";

struct MsgMintNFTs {
    address creator;
    address class_id;
    NFTInput[] inputs;
}
