// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import {ERC721URIStorageUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import {ERC721EnumerableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";

/// @custom:security-contact rickmak@oursky.com
contract LikeStakePositionV0 is
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    ERC721Upgradeable,
    ERC721EnumerableUpgradeable,
    ERC721URIStorageUpgradeable
{
    // keccak256(abi.encode(uint256(keccak256("likestakeposition.storage")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CONTRACT_STORAGE_SLOT =
        0x2c3a4a1c92b0f847cbe6b33689f93d825b12d6a2f2a7bdb9c9a6a1cf7e6bb200;

    struct ContractData {
        address manager;
        uint256 nextTokenId;
        string baseURI;
    }

    function _getStorage() private pure returns (ContractData storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := CONTRACT_STORAGE_SLOT
        }
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address initialOwner) public initializer {
        __ERC721_init("LikeStakePosition", "LIKESP");
        __ERC721Enumerable_init();
        __ERC721URIStorage_init();
        __UUPSUpgradeable_init();
        __Ownable_init(initialOwner);
        __Pausable_init();
        __ReentrancyGuard_init();

        _getStorage().nextTokenId = 1;
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}

    // The following functions are overrides required by Solidity.

    function _update(
        address to,
        uint256 tokenId,
        address auth
    )
        internal
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable)
        returns (address)
    {
        return super._update(to, tokenId, auth);
    }

    function _increaseBalance(
        address account,
        uint128 value
    ) internal override(ERC721Upgradeable, ERC721EnumerableUpgradeable) {
        super._increaseBalance(account, value);
    }

    function tokenURI(
        uint256 tokenId
    )
        public
        view
        override(ERC721Upgradeable, ERC721URIStorageUpgradeable)
        returns (string memory)
    {
        return "";
    }

    function supportsInterface(
        bytes4 interfaceId
    )
        public
        view
        override(
            ERC721Upgradeable,
            ERC721EnumerableUpgradeable,
            ERC721URIStorageUpgradeable
        )
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}
