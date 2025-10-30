// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

// solhint-disable-next-line no-unused-import
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {veLike} from "../contracts/veLike.sol";

contract veLikeMock is veLike {
    function version() public pure returns (uint256) {
        return 2;
    }

    function dataStorage() external pure returns (bytes32) {
        return
            keccak256(abi.encode(uint256(keccak256("veLike.storage")) - 1)) &
            ~bytes32(uint256(0xff));
    }
    function veLikeRewardDataStorage() external pure returns (bytes32) {
        return
            keccak256(abi.encode(uint256(keccak256("veLikeReward.storage")) - 1)) &
            ~bytes32(uint256(0xff));
    }
}
