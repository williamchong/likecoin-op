// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {NFTData} from "../NFTData.sol";

struct MsgMintNFTsFromTokenId {
    address classId;
    address to;
    uint256 fromTokenId;
    NFTData[] inputs;
}
