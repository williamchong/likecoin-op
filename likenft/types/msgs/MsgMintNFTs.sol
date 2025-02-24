// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {NFTData} from "../NFTData.sol";

struct MsgMintNFTs {
    address classId;
    address to;
    NFTData[] inputs;
}
