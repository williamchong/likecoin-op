// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;
// Neet to import this for HH Ignition
// solhint-disable-next-line no-unused-import
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Likecoin} from "../contracts/Likecoin.sol";

contract LikecoinMock is Likecoin {
    function version() public pure returns (uint256) {
        return 2;
    }
}
