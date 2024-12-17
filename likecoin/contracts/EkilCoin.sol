pragma solidity ^0.8.0;

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {ERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

contract EkilCoin is
    Initializable,
    ERC20Upgradeable,
    PausableUpgradeable,
    OwnableUpgradeable,
    AccessControlUpgradeable,
    UUPSUpgradeable
{
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    address private _grantedMinter;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address initialOwner,
        address initialMinter
    ) public initializer {
        __ERC20_init("EkilCoin", "EKIL");
        __Pausable_init();
        __AccessControl_init();
        __Ownable_init(initialOwner);
        __UUPSUpgradeable_init();

        _setMinter(initialMinter);
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function mint(
        address to,
        uint256 amount
    ) public onlyRole(MINTER_ROLE) whenNotPaused {
        _mint(to, amount);
    }

    function _authorizeUpgrade(address) internal view override onlyOwner {}

    function _setMinter(address minter) internal returns (address) {
        require(minter != address(0), "INVALID_NULL_ADDRESS");
        _revokeRole(MINTER_ROLE, _grantedMinter);
        _grantRole(MINTER_ROLE, minter);
        _grantedMinter = minter;
        return _grantedMinter;
    }

    function setMinter(address minter) public onlyOwner returns (address) {
        return _setMinter(minter);
    }

    function getMinter() public view returns (address) {
        return _grantedMinter;
    }
}
