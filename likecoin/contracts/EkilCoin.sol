pragma solidity ^0.8.0;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract EkilCoin is ERC20 {
    string private _gitHash;

    constructor(
        uint256 initialSupply,
        string memory gitHash
    ) ERC20("EkilCoin", "EKIL") {
        _mint(msg.sender, initialSupply);

        _gitHash = gitHash;
    }

    function getGitHash() external view returns (string memory) {
        return _gitHash;
    }
}
