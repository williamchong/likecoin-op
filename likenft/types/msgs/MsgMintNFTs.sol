// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {NFTData} from "../NFTData.sol";

struct MsgMintNFTs {
    address class_id;
    address to;
    NFTData[] inputs;
}
