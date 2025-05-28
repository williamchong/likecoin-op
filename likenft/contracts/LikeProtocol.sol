// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {IBeacon} from "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";
import {BeaconProxy} from "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";

import {MsgNewBookNFT} from "../types/msgs/MsgNewBookNFT.sol";
import {BookConfig} from "../types/BookConfig.sol";

import {BookNFT} from "./BookNFT.sol";

error ErrNftClassNotFound();
interface BookNFTInterface {
    function initialize(MsgNewBookNFT memory msgNewBookNFT) external;
}

contract LikeProtocol is
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable,
    PausableUpgradeable,
    IBeacon
{
    struct LikeNFTStorage {
        mapping(address classId => bool isBookNFT) classIdMapping;
        address bookNFTImplementation;
        address royaltyReceiver;
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
    event BookNFTImplementationUpgraded(address newImplementation);
    error BookNFTInvalidImplementation(address implementation);

    function initialize(
        address initialOwner,
        address bookNFTImplementation
    ) public initializer {
        __UUPSUpgradeable_init();
        __Ownable_init(initialOwner);
        __Pausable_init();
        LikeNFTStorage storage $ = _getLikeNFTStorage();

        if (bookNFTImplementation.code.length == 0) {
            revert BookNFTInvalidImplementation(bookNFTImplementation);
        }
        $.bookNFTImplementation = bookNFTImplementation;
        $.royaltyReceiver = initialOwner;
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function isBookNFT(address classId) public view returns (bool) {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        return $.classIdMapping[classId];
    }

    function newBookNFT(
        MsgNewBookNFT memory msgNewBookNFT
    ) public whenNotPaused returns (address bookAddress) {
        LikeNFTStorage storage $ = _getLikeNFTStorage();
        bytes memory initData = abi.encodeWithSelector(
            BookNFTInterface.initialize.selector,
            msgNewBookNFT
        );
        BeaconProxy proxy = new BeaconProxy(address(this), initData);
        bookAddress = address(proxy);
        $.classIdMapping[bookAddress] = true;
        emit NewBookNFT(bookAddress, msgNewBookNFT.config);
    }

    /**
     * newBookNFTWithRoyalty
     *
     * Proxy call to create a BookNFT with a royalty fraction
     *
     * @param msgNewBookNFT - the message to create the BookNFT
     * @param royaltyFraction - the royalty fraction to set
     */
    function newBookNFTWithRoyalty(
        MsgNewBookNFT memory msgNewBookNFT,
        uint96 royaltyFraction
    ) public whenNotPaused returns (address bookAddress) {
        bookAddress = newBookNFT(msgNewBookNFT);
        BookNFT(bookAddress).setRoyaltyFraction(royaltyFraction);
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

    // Beacon implementation
    /**
     * @notice Get the implementation address of the BookNFT
     * @return The implementation address of the BookNFT
     */
    function implementation() external view override returns (address) {
        return _getLikeNFTStorage().bookNFTImplementation;
    }

    /**
     * @notice Upgrade the implementation address of the BookNFT
     * @param newImplementation The new implementation address of the BookNFT
     */
    function upgradeTo(address newImplementation) external onlyOwner {
        if (newImplementation.code.length == 0) {
            revert BookNFTInvalidImplementation(newImplementation);
        }
        _getLikeNFTStorage().bookNFTImplementation = newImplementation;
        emit BookNFTImplementationUpgraded(newImplementation);
    }
    // End of Beacon implementation

    // Royalty
    function getRoyaltyReceiver() external view returns (address) {
        return _getLikeNFTStorage().royaltyReceiver;
    }

    function setRoyaltyReceiver(address royaltyReceiver) external onlyOwner {
        _getLikeNFTStorage().royaltyReceiver = royaltyReceiver;
    }
    // End of Royalty
}
