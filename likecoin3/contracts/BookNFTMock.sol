// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {BookNFT} from "../contracts/BookNFT.sol";

contract BookNFTMock is BookNFT {
    function version() public pure returns (uint256) {
        return 2;
    }

    function bookNFTStorage() external pure returns (bytes32) {
        return
            keccak256(
                abi.encode(
                    uint256(keccak256("likeprotocol.booknft.storage")) - 1
                )
            ) & ~bytes32(uint256(0xff));
    }
}
