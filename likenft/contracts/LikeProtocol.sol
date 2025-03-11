// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

import {MsgMintNFT} from "../types/msgs/MsgMintNFT.sol";
import {MsgMintNFTs} from "../types/msgs/MsgMintNFTs.sol";
import {MsgMintNFTsFromTokenId} from "../types/msgs/MsgMintNFTsFromTokenId.sol";
import {MsgNewBookNFT} from "../types/msgs/MsgNewBookNFT.sol";
import {MsgUpdateBookNFT} from "../types/msgs/MsgUpdateBookNFT.sol";
import {BookConfig} from "../types/BookConfig.sol";

import {BookNFT} from "./BookNFT.sol";

error ErrNftClassNotFound();

contract LikeProtocol is
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable,
    PausableUpgradeable
{
    struct LikeNFTStorage {
        mapping(address classId => BookNFT) classIdClassMapping;
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

    event NewBookNFT(address bookNFT, BookConfig config);

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

    function isBookNFT(address classId) public view returns (bool) {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        return address($.classIdClassMapping[classId]) != address(0);
    }

    function newBookNFT(
        MsgNewBookNFT memory msgNewBookNFT
    ) public whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        BookNFT class = new BookNFT(msgNewBookNFT);
        address id = address(class);
        $.classIdClassMapping[id] = class;
        emit NewBookNFT(id, msgNewBookNFT.config);
    }

    /**
     * newBookNFTs
     *
     * Proxy call to create multiple BookNFT at once
     *
     * @param msgNewBookNFTs the BookNFTs to create
     */
    function newBookNFTs(
        MsgNewBookNFT[] calldata msgNewBookNFTs
    ) public whenNotPaused {
        for (uint i = 0; i < msgNewBookNFTs.length; i++) {
            new BookNFT(msgNewBookNFTs[i]);
        }
    }

    function updateBookNFT(
        MsgUpdateBookNFT memory msgUpdateBookNFT
    ) public whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        BookNFT class = $.classIdClassMapping[msgUpdateBookNFT.classId];
        if (address(class) == address(0)) {
            revert ErrNftClassNotFound();
        }
        class.update(msgUpdateBookNFT.config);
    }

    function mintNFT(MsgMintNFT calldata msgMintNFT) external whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        BookNFT class = $.classIdClassMapping[msgMintNFT.classId];
        if (address(class) == address(0)) {
            revert ErrNftClassNotFound();
        }
        string[] memory metadataList = new string[](1);
        metadataList[0] = msgMintNFT.input.metadata;
        class.mint(msgMintNFT.to, metadataList);
    }

    function mintNFTs(MsgMintNFTs calldata msgMintNFTs) external whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        BookNFT class = $.classIdClassMapping[msgMintNFTs.classId];
        if (address(class) == address(0)) {
            revert ErrNftClassNotFound();
        }
        string[] memory metadataList = new string[](msgMintNFTs.inputs.length);
        for (uint i = 0; i < msgMintNFTs.inputs.length; i++) {
            metadataList[i] = msgMintNFTs.inputs[i].metadata;
        }
        class.mint(msgMintNFTs.to, metadataList);
    }

    function safeMintNFTsWithTokenId(
        MsgMintNFTsFromTokenId calldata msgMintNFTsFromTokenId
    ) external whenNotPaused {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        BookNFT class = $.classIdClassMapping[msgMintNFTsFromTokenId.classId];
        if (address(class) == address(0)) {
            revert ErrNftClassNotFound();
        }
        string[] memory metadataList = new string[](
            msgMintNFTsFromTokenId.inputs.length
        );
        for (uint i = 0; i < msgMintNFTsFromTokenId.inputs.length; i++) {
            metadataList[i] = msgMintNFTsFromTokenId.inputs[i].metadata;
        }
        class.safeMintWithTokenId(
            msgMintNFTsFromTokenId.fromTokenId,
            msgMintNFTsFromTokenId.to,
            metadataList
        );
    }

    function _authorizeUpgrade(
        address _newImplementation // solhint-disable-next-line no-empty-blocks
    ) internal override onlyOwner {}
}
