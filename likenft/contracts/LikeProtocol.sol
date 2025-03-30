// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

import {MsgNewBookNFT} from "../types/msgs/MsgNewBookNFT.sol";
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
    // keccak256(abi.encode(uint256(keccak256("likeprotocol.storage")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant DATA_STORAGE =
        0xe3ffde652b1592025b57f85d2c64876717f9cdf4e44b57422a295c18d0719a00;
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
            newBookNFT(msgNewBookNFTs[i]);
        }
    }

    function _authorizeUpgrade(
        address _newImplementation // solhint-disable-next-line no-empty-blocks
    ) internal override onlyOwner {}
}
