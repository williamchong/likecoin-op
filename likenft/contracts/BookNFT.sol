// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ERC721, ERC721Enumerable} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {ClassStorage} from "../types/Class.sol";
import {ClassInput} from "../types/ClassInput.sol";
import {MsgNewClass} from "../types/msgs/MsgNewClass.sol";
import {NFTData} from "../types/NFTData.sol";

error ErrUnauthorized();
error ErrNftNoSupply();
error ErrTokenIdMintFails(uint256 nextTokenId);

contract BookNFT is ERC721Enumerable, Ownable, AccessControl {
    // keccak256(abi.encode(uint256(keccak256("likenft.storage.class")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CLASS_DATA_STORAGE =
        0x99391ccf5d97dbb7711a73831d943712d1774ca037a259af20891dc6f0d9f200;
    function _getClassStorage() private pure returns (ClassStorage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := CLASS_DATA_STORAGE
        }
    }

    // Constants
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant UPDATER_ROLE = keccak256("UPDATER_ROLE");
    // End Constants

    // Storage
    uint256 private _currentIndex;
    mapping(uint256 => NFTData) private nftDataMap;
    // End Storage

    // Events
    event ContractURIUpdated();

    event TransferWithMemo(
        address indexed from,
        address indexed to,
        uint256 indexed tokenId,
        string memo
    );
    // End Events

    modifier onlyMinter() {
        // FIXME: tx.origin is prone to phishing attacks
        if (!hasRole(MINTER_ROLE, tx.origin)) {
            revert ErrUnauthorized();
        }
        _;
    }

    modifier onlyUpdater() {
        // FIXME: tx.origin is prone to phishing attacks
        if (!hasRole(UPDATER_ROLE, tx.origin)) {
            revert ErrUnauthorized();
        }
        _;
    }

    constructor(
        MsgNewClass memory msgNewClass
    )
        ERC721(msgNewClass.input.name, msgNewClass.input.symbol)
        Ownable(msgNewClass.creator)
    {
        ClassStorage storage $ = _getClassStorage();
        $.name = msgNewClass.input.name;
        $.symbol = msgNewClass.input.symbol;
        $.data.metadata = msgNewClass.input.metadata;
        $.data.config = msgNewClass.input.config;

        _currentIndex = 0;

        for (uint i = 0; i < msgNewClass.minters.length; i++) {
            _grantRole(MINTER_ROLE, msgNewClass.minters[i]);
        }
        for (uint i = 0; i < msgNewClass.updaters.length; i++) {
            _grantRole(UPDATER_ROLE, msgNewClass.updaters[i]);
        }
    }

    function supportsInterface(
        bytes4 interfaceId
    )
        public
        view
        virtual
        override(ERC721Enumerable, AccessControl)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    function update(ClassInput memory classInput) public onlyUpdater {
        ClassStorage storage $ = _getClassStorage();
        $.name = classInput.name;
        $.symbol = classInput.symbol;
        $.data.metadata = classInput.metadata;
        emit ContractURIUpdated();
    }

    function mint(
        address to,
        string[] calldata metadataList
    ) external onlyMinter {
        _mint(to, metadataList);
    }

    /**
     * safeMintWithTokenId
     *
     * a fast fails function call to ensure the transaction sender
     * is getting the desired tokenId(in stead of next Id) in the result.
     *
     * Expect caller to check and specify correct start token id
     *
     * @param fromTokenId - the start token id
     * @param to - owner address to hold the new minted token
     * @param metadataList - list of metadata to supply
     */
    function safeMintWithTokenId(
        uint256 fromTokenId,
        address to,
        string[] calldata metadataList
    ) external onlyMinter {
        if (totalSupply() != fromTokenId) {
            revert ErrTokenIdMintFails(totalSupply());
        }
        _mint(to, metadataList);
    }

    function _mint(address to, string[] memory metadataList) internal {
        ClassStorage storage $ = _getClassStorage();

        uint64 maxSupply = $.data.config.max_supply;
        uint quantity = metadataList.length;

        if (maxSupply != 0 && totalSupply() + quantity > maxSupply) {
            revert ErrNftNoSupply();
        }

        for (uint i = 0; i < quantity; i++) {
            nftDataMap[_currentIndex].metadata = metadataList[i];
            _safeMint(to, _currentIndex);
            _currentIndex++;
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
