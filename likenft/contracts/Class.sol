// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ERC721A} from "erc721a/contracts/ERC721A.sol";

import {ClassStorage} from "../types/Class.sol";
import {ClassInput} from "../types/ClassInput.sol";
import {MsgNewClass} from "../types/msgs/MsgNewClass.sol";
import {NFTData} from "../types/NFTData.sol";

error ErrNftNoSupply();
error ErrCannotUpdateClassWithMintedTokens();

contract Class is ERC721A, Ownable {
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

    event TransferWithMemo(
        address indexed from,
        address indexed to,
        uint256 indexed tokenId,
        string memo
    );

    constructor(
        MsgNewClass memory msgNewClass
    )
        ERC721A(msgNewClass.input.name, msgNewClass.input.symbol)
        Ownable(msg.sender)
    {
        ClassStorage storage $ = _getClassStorage();
        $.name = msgNewClass.input.name;
        $.symbol = msgNewClass.input.symbol;
        $.data.metadata = msgNewClass.input.metadata;
        $.data.config = msgNewClass.input.config;
    }

    function update(ClassInput memory classInput) public onlyOwner {
        ClassStorage storage $ = _getClassStorage();
        $.name = classInput.name;
        $.symbol = classInput.symbol;
        $.data.metadata = classInput.metadata;
        emit ContractURIUpdated();
    }

    function mint(
        address to,
        string[] calldata metadata_list
    ) external onlyOwner {
        ClassStorage storage $ = _getClassStorage();

        uint256 nextTokenId = _nextTokenId();
        uint64 maxSupply = $.data.config.max_supply;
        uint quantity = metadata_list.length;

        if (maxSupply != 0 && totalSupply() + quantity > maxSupply) {
            revert ErrNftNoSupply();
        }

        _safeMint(to, quantity);

        for (uint i = 0; i < quantity; i++) {
            uint256 _tokenId = nextTokenId + i;
            nftDataMap[_tokenId].metadata = metadata_list[i];
        }
    }

    function name() public view override returns (string memory) {
        ClassStorage storage $ = _getClassStorage();
        return $.name;
    }

    function symbol() public view override returns (string memory) {
        ClassStorage storage $ = _getClassStorage();
        return $.symbol;
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

    function transferWithMemo(
        address from,
        address to,
        uint256 _tokenId,
        string calldata memo
    ) external payable {
        transferFrom(from, to, _tokenId);

        emit TransferWithMemo(from, to, _tokenId, memo);
    }
}
