// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

import {MsgMintNFT} from "../types/msgs/MsgMintNFT.sol";
import {MsgMintNFTs} from "../types/msgs/MsgMintNFTs.sol";
import {MsgNewClass} from "../types/msgs/MsgNewClass.sol";
import {MsgUpdateClass} from "../types/msgs/MsgUpdateClass.sol";

import {LikeNFTClass} from "./LikeNFTClass.sol";

error ErrNftClassNotFound();

contract LikeProtocol is
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable,
    PausableUpgradeable
{
    struct LikeNFTStorage {
        mapping(address class_id => LikeNFTClass) classIdClassMapping;
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

    event NewClass(address classId, MsgNewClass params);

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

    function newClass(MsgNewClass memory msgNewClass) public whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        LikeNFTClass class = new LikeNFTClass(msgNewClass);
        address id = address(class);
        $.classIdClassMapping[id] = class;
        emit NewClass(id, msgNewClass);
    }

    function updateClass(
        MsgUpdateClass memory msgUpdateClass
    ) public whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        LikeNFTClass class = $.classIdClassMapping[msgUpdateClass.class_id];
        if (address(class) == address(0)) {
            revert ErrNftClassNotFound();
        }
        class.update(msgUpdateClass.input);
    }

    function mintNFT(MsgMintNFT calldata msgMintNFT) external whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        LikeNFTClass class = $.classIdClassMapping[msgMintNFT.class_id];
        if (address(class) == address(0)) {
            revert ErrNftClassNotFound();
        }
        string[] memory metadata_list = new string[](1);
        metadata_list[0] = msgMintNFT.input.metadata;
        class.mint(msgMintNFT.to, metadata_list);
    }

    function mintNFTs(MsgMintNFTs calldata msgMintNFTs) external whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        LikeNFTClass class = $.classIdClassMapping[msgMintNFTs.class_id];
        if (address(class) == address(0)) {
            revert ErrNftClassNotFound();
        }
        string[] memory metadata_list = new string[](msgMintNFTs.inputs.length);
        for (uint i = 0; i < msgMintNFTs.inputs.length; i++) {
            metadata_list[i] = msgMintNFTs.inputs[i].metadata;
        }
        class.mint(msgMintNFTs.to, metadata_list);
    }

    function _authorizeUpgrade(
        address _newImplementation // solhint-disable-next-line no-empty-blocks
    ) internal override onlyOwner {}
}
