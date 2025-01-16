// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";

import {ClassStorage} from "../types/Class.sol";
import {ClassInput} from "../types/ClassInput.sol";
import {MsgNewClass} from "../types/msgs/MsgNewClass.sol";
import {NFTData} from "../types/NFTData.sol";

error ErrNftNoSupply();
error ErrCannotUpdateClassWithMintedTokens();

contract Class is ERC721, Ownable {
    // keccak256(abi.encode(uint256(keccak256("likenft.storage.class")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CLASS_DATA_STORAGE =
        0x99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f200;
    function _getClassStorage() private pure returns (ClassStorage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := CLASS_DATA_STORAGE
        }
    }

    mapping(uint256 => NFTData) private nftDataMap;

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    uint256 public tokenId;

    address private _grantedMinter;

    event ContractURIUpdated();

    constructor(
        MsgNewClass memory msgNewClass
    )
        ERC721(msgNewClass.input.name, msgNewClass.input.symbol)
        Ownable(msg.sender)
    {
        ClassStorage storage $ = _getClassStorage();
        $.name = msgNewClass.input.name;
        $.symbol = msgNewClass.input.symbol;
        $.description = msgNewClass.input.description;
        $.uri = msgNewClass.input.uri;
        $.uri_hash = msgNewClass.input.uri_hash;
        $.data.metadata = msgNewClass.input.metadata;
        $.data.parent.type_ = msgNewClass.parent.type_;
        $.data.parent.iscn_id_prefix = msgNewClass.parent.iscn_id_prefix;
        $.data.config = msgNewClass.input.config;
    }

    function update(ClassInput memory classInput) public onlyOwner {
        if (tokenId > 0) {
            revert ErrCannotUpdateClassWithMintedTokens();
        }
        ClassStorage storage $ = _getClassStorage();
        $.name = classInput.name;
        $.symbol = classInput.symbol;
        $.description = classInput.description;
        $.uri = classInput.uri;
        $.uri_hash = classInput.uri_hash;
        $.data.metadata = classInput.metadata;
        emit ContractURIUpdated();
    }

    function mint(address to, string memory metadata) public onlyOwner {
        ClassStorage storage $ = _getClassStorage();
        uint64 maxSupply = $.data.config.max_supply;
        if (maxSupply != 0 && tokenId >= maxSupply) {
            revert ErrNftNoSupply();
        }

        _safeMint(to, tokenId);
        nftDataMap[tokenId].class_parent = $.data.parent;
        nftDataMap[tokenId].metadata = metadata;
        nftDataMap[tokenId].to_be_revealed = false;
        tokenId++;
    }

    function contractURI() public view returns (string memory) {
        return
            string.concat(
                "data:application/json;utf8,",
                _getClassStorage().data.metadata
            );
    }

    function tokenURI(
        uint256 _tokenId
    ) public view virtual override returns (string memory) {
        return
            string.concat(
                "data:application/json;utf8,",
                nftDataMap[_tokenId].metadata
            );
    }
}
