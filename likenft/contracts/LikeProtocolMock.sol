// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {LikeProtocol} from "../contracts/LikeProtocol.sol";

contract LikeProtocolMock is LikeProtocol {
    function version() public pure returns (uint256) {
        return 2;
    }

    function protocolDataStorage() external pure returns (bytes32) {
        return
            keccak256(
                abi.encode(uint256(keccak256("likeprotocol.storage")) - 1)
            ) & ~bytes32(uint256(0xff));
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
