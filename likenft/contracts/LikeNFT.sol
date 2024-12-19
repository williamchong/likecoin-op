// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

import {MsgMintNFT} from "../types/msgs/MsgMintNFT.sol";
import {MsgNewClass} from "../types/msgs/MsgNewClass.sol";
import {MsgUpdateClass} from "../types/msgs/MsgUpdateClass.sol";

import {Class} from "./Class.sol";

error ErrNftClassNotFound();
error ErrNftClassAlreadyExists();

contract LikeNFT is
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable,
    PausableUpgradeable
{
    struct LikeNFTStorage {
        address minter;
        mapping(address creator => mapping(string class_id => Class)) creatorClassIdClassMapping;
        Class[] classes;
    }
    // keccak256(abi.encode(uint256(keccak256("likenft.storage")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant DATA_STORAGE =
        0xf59cae2d8704429a88f4a038c4cff8d2643dc6b4647d519013fb42e0b4344200;
    function _getLikeNFTStorage()
        private
        pure
        returns (LikeNFTStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := DATA_STORAGE
        }
    }

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    function initialize(address initialOwner) public initializer {
        __UUPSUpgradeable_init();
        __Ownable_init(initialOwner);
        __Pausable_init();
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function newClass(
        MsgNewClass memory msgNewClass,
        string memory id
    ) public onlyOwner whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        if (
            address($.creatorClassIdClassMapping[msgNewClass.creator][id]) !=
            address(0)
        ) {
            revert ErrNftClassAlreadyExists();
        }
        Class class = new Class(msgNewClass, id);
        $.classes.push(class);
        $.creatorClassIdClassMapping[msgNewClass.creator][id] = class;
    }

    function updateClass(
        MsgUpdateClass memory msgUpdateClass
    ) public onlyOwner whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        Class class = $.creatorClassIdClassMapping[msgUpdateClass.creator][
            msgUpdateClass.class_id
        ];
        if (address(class) == address(0)) {
            revert ErrNftClassNotFound();
        }
        class.update(msgUpdateClass.input);
    }

    function mintNFT(
        MsgMintNFT memory msgMintNFT
    ) public onlyOwner whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        Class class = $.creatorClassIdClassMapping[msgMintNFT.creator][
            msgMintNFT.class_id
        ];
        if (address(class) == address(0)) {
            revert ErrNftClassNotFound();
        }
        class.mint(msgMintNFT.creator, msgMintNFT.input.metadata);
    }

    function _authorizeUpgrade(
        address _newImplementation
    )
        internal
        override
        onlyOwner // solhint-disable-next-line no-empty-blocks
    {}
}
