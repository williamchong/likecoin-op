// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {LikeProtocol} from "../contracts/LikeProtocol.sol";

contract LikeProtocolMock is LikeProtocol {
    function version() public pure returns (uint256) {
        return 2;
    }
}
