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
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {Base64} from "@openzeppelin/contracts/utils/Base64.sol";

/// @title LikeStakePosition
/// @notice ERC721 receipt representing a user's stake position in a specific book (BookNFT)
/// @dev Upgradeable, mint/burn/update restricted to LikeCollective (manager)
contract LikeStakePosition is
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    ERC721Upgradeable,
    ERC721EnumerableUpgradeable,
    ERC721URIStorageUpgradeable
{
    using Strings for uint256;

    struct Position {
        address bookNFT;
        uint256 stakedAmount;
        uint256 rewardIndex;
        address initialStaker;
    }

    // keccak256(abi.encode(uint256(keccak256("likestakeposition.storage")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CONTRACT_STORAGE_SLOT =
        0x2c3a4a1c92b0f847cbe6b33689f93d825b12d6a2f2a7bdb9c9a6a1cf7e6bb200;

    struct ContractData {
        address manager;
        uint256 nextTokenId;
        string baseURI;
        mapping(uint256 => Position) positions;
    }

    function _getStorage() private pure returns (ContractData storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := CONTRACT_STORAGE_SLOT
        }
    }

    // Errors
    error ErrNotManager();
    error ErrZeroAddress();
    error ErrInvalidOwner();

    // Events
    event ManagerUpdated(address indexed manager);
    event BaseURIUpdated(string baseURI);
    event PositionMinted(uint256 indexed tokenId, address indexed to, address indexed bookNFT, uint256 amount, uint256 rewardIndex);
    event PositionUpdated(uint256 indexed tokenId, uint256 amount, uint256 rewardIndex);
    event PositionBurned(uint256 indexed tokenId);

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

    // Ownership / Admin
    function setManager(address manager_) external onlyOwner {
        if (manager_ == address(0)) revert ErrZeroAddress();
        ContractData storage $ = _getStorage();
        $.manager = manager_;
        emit ManagerUpdated(manager_);
    }

    function setBaseURI(string calldata baseURI_) external onlyOwner {
        ContractData storage $ = _getStorage();
        $.baseURI = baseURI_;
        emit BaseURIUpdated(baseURI_);
    }

    function pause() external onlyOwner { _pause(); }
    function unpause() external onlyOwner { _unpause(); }

    // Manager-only modifiers
    modifier onlyManager() {
        if (_msgSender() != _getStorage().manager) revert ErrNotManager();
        _;
    }

    // Mint/Burn/Update: callable only by LikeCollective (manager)
    function mintPosition(
        address to,
        address bookNFT,
        uint256 stakedAmount,
        uint256 rewardIndex
    ) external whenNotPaused nonReentrant onlyManager returns (uint256 tokenId) {
        if (to == address(0) || bookNFT == address(0)) revert ErrZeroAddress();
        ContractData storage $ = _getStorage();
        tokenId = $.nextTokenId;
        $.nextTokenId++;
        _safeMint(to, tokenId);
        $.positions[tokenId] = Position({
            bookNFT: bookNFT,
            stakedAmount: stakedAmount,
            rewardIndex: rewardIndex,
            initialStaker: to
        });
        emit PositionMinted(tokenId, to, bookNFT, stakedAmount, rewardIndex);
    }

    function burnPosition(uint256 tokenId) external whenNotPaused nonReentrant onlyManager {
        _burn(tokenId);
        delete _getStorage().positions[tokenId];
        emit PositionBurned(tokenId);
    }

    function updatePosition(
        uint256 tokenId,
        uint256 newStakedAmount,
        uint256 newRewardIndex
    ) external whenNotPaused nonReentrant onlyManager {
        ContractData storage $ = _getStorage();
        Position storage p = $.positions[tokenId];
        // rely on ERC721 _ownerOf revert for non-existent token
        if (ownerOf(tokenId) == address(0)) revert ErrInvalidOwner();
        p.stakedAmount = newStakedAmount;
        p.rewardIndex = newRewardIndex;
        emit PositionUpdated(tokenId, newStakedAmount, newRewardIndex);
        emit MetadataUpdate(tokenId);
    }
    
    function updatePositionRewardIndex(
        uint256 tokenId,
        uint256 newRewardIndex
    ) external whenNotPaused nonReentrant onlyManager {
        ContractData storage $ = _getStorage();
        Position storage p = $.positions[tokenId];
        p.rewardIndex = newRewardIndex;
        emit PositionUpdated(tokenId, p.stakedAmount, newRewardIndex);
        emit MetadataUpdate(tokenId);
    }

    // Views
    function getNextTokenId() external view returns (uint256) {
        return _getStorage().nextTokenId;
    }

    function getPosition(uint256 tokenId) external view returns (Position memory) {
        return _getStorage().positions[tokenId];
    }

    function manager() external view returns (address) {
        return _getStorage().manager;
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    // ERC721 metadata base URI
    function _baseURI() internal view override returns (string memory) {
        return _getStorage().baseURI;
    }

    // View functions
    function getUserPositions(address user) external view returns (uint256[] memory positions) {
        uint256 balance = balanceOf(user);
        for (uint256 i = 0; i < balance; i++) {
            positions[i] = tokenOfOwnerByIndex(user, i);
        }
        return positions;
    }

    function getUserPositionByBookNFT(address user, address bookNFT) external view returns (uint256[] memory) {
        uint256 balance = balanceOf(user);
        uint256[] memory positions = new uint256[](balance);
        uint256 index = 0;
        for (uint256 i = 0; i < balance; i++) {
            uint256 tokenId = tokenOfOwnerByIndex(user, i);
            if (_getStorage().positions[tokenId].bookNFT == bookNFT) {
                positions[index] = tokenId;
                index++;
            }
        }
        return positions;
    }

    function tokenURI(uint256 tokenId) public view override(ERC721Upgradeable, ERC721URIStorageUpgradeable) returns (string memory) {
        Position memory position = _getStorage().positions[tokenId];
        string memory baseURI = _baseURI();
        string memory tokenURL = Base64.encode(abi.encodePacked(
            position.bookNFT,
            position.stakedAmount,
            position.rewardIndex,
            position.initialStaker
        ));
        return string.concat(baseURI, tokenURL);
    }

    // The following functions are overrides required by Solidity.

        function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable)
        returns (address)
    {
        return super._update(to, tokenId, auth);
    }

    function _increaseBalance(address account, uint128 value)
        internal
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable)
    {
        super._increaseBalance(account, value);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable, ERC721URIStorageUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}


