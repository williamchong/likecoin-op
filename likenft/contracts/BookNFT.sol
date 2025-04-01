// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ERC721, ERC721Enumerable} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

import {BookConfig} from "../types/BookConfig.sol";
import {MsgNewBookNFT} from "../types/msgs/MsgNewBookNFT.sol";

error ErrUnauthorized();
error ErrEmptyName();
error ErrEmptySymbol();
error ErrInvalidMetadata();
error ErrMaxSupplyZero();
error ErrNftNoSupply();
error ErrTokenIdMintFails(uint256 nextTokenId);

contract BookNFT is ERC721Enumerable, Ownable, AccessControl {
    struct BookNFTStorage {
        string name;
        string symbol;
        string metadata;
        uint64 max_supply;
        uint256 _currentIndex;
        mapping(uint256 => string) tokenURIMap;
    }

    // keccak256(abi.encode(uint256(keccak256("likeprotocol.booknft.storage")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CLASS_DATA_STORAGE =
        0x8303e9d27d04c843c8d4a08966b1e1be0214fc0b3375d79db0a8252068c41f00;
    function _getClassStorage()
        private
        pure
        returns (BookNFTStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := CLASS_DATA_STORAGE
        }
    }

    // Constants
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant UPDATER_ROLE = keccak256("UPDATER_ROLE");
    // End Constants

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
        if (!hasRole(MINTER_ROLE, _msgSender())) {
            revert ErrUnauthorized();
        }
        _;
    }

    modifier onlyUpdater() {
        if (!hasRole(UPDATER_ROLE, _msgSender())) {
            revert ErrUnauthorized();
        }
        _;
    }

    constructor(
        MsgNewBookNFT memory msgNewBookNFT
    )
        ERC721(msgNewBookNFT.config.name, msgNewBookNFT.config.symbol)
        Ownable(msgNewBookNFT.creator)
    {
        _validateBookConfig(msgNewBookNFT.config);

        BookNFTStorage storage $ = _getClassStorage();
        $.name = msgNewBookNFT.config.name;
        $.symbol = msgNewBookNFT.config.symbol;
        $.max_supply = msgNewBookNFT.config.max_supply;
        $.metadata = msgNewBookNFT.config.metadata;

        $._currentIndex = 0;

        for (uint i = 0; i < msgNewBookNFT.minters.length; i++) {
            _grantRole(MINTER_ROLE, msgNewBookNFT.minters[i]);
        }
        for (uint i = 0; i < msgNewBookNFT.updaters.length; i++) {
            _grantRole(UPDATER_ROLE, msgNewBookNFT.updaters[i]);
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

    function _validateBookConfig(BookConfig memory config) internal pure {
        if (bytes(config.name).length == 0) {
            revert ErrEmptyName();
        }
        if (bytes(config.symbol).length == 0) {
            revert ErrEmptySymbol();
        }
        if (config.max_supply == 0) {
            revert ErrMaxSupplyZero();
        }
    }

    function update(BookConfig calldata config) public onlyUpdater {
        _validateBookConfig(config);
        BookNFTStorage storage $ = _getClassStorage();
        require(config.max_supply >= $.max_supply, "ErrSupplyDecrease");

        $.name = config.name;
        $.symbol = config.symbol;
        $.max_supply = config.max_supply;
        $.metadata = config.metadata;
        emit ContractURIUpdated();
    }

    function mint(
        address to,
        string[] calldata metadataList
    ) external onlyMinter {
        _ensureEnoughSupply(metadataList.length);
        for (uint i = 0; i < metadataList.length; i++) {
            _mintWithEvent(_msgSender(), to, metadataList[i]);
        }
    }

    /**
     * batchMint
     *
     * batch mint with metadata list
     *
     * @param tos - owner address to hold the new minted token
     * @param metadataList - list of metadata to supply, the length of the list should be the same as the length of the tos. Metadata will fill the corresponding position of the tos.
     */
    function batchMint(
        address[] calldata tos,
        string[] calldata metadataList
    ) external onlyMinter {
        _ensureEnoughSupply(metadataList.length);
        for (uint i = 0; i < tos.length; i++) {
            _mintWithEvent(_msgSender(), tos[i], metadataList[i]);
        }
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
        _ensureEnoughSupply(metadataList.length);
        for (uint i = 0; i < metadataList.length; i++) {
            _mintWithEvent(_msgSender(), to, metadataList[i]);
        }
    }

    /**
     * _ensureEnoughtSupply
     *
     * ensure the supply is enough
     *
     * @param quantity - the quantity of the tokens to mint
     */
    function _ensureEnoughSupply(uint quantity) internal view {
        BookNFTStorage storage $ = _getClassStorage();
        if (totalSupply() + quantity > $.max_supply) {
            revert ErrNftNoSupply();
        }
    }

    /**
     * _mintWithEvent
     *
     * mint a new token with metadata, caller should ensure the supply is enough.
     *
     * @param from - the address that is transferring the token
     * @param to - owner address to hold the new minted token
     * @param metadata - metadata to supply
     */
    function _mintWithEvent(
        address from,
        address to,
        string calldata metadata
    ) internal {
        BookNFTStorage storage $ = _getClassStorage();
        $.tokenURIMap[$._currentIndex] = metadata;
        _safeMint(to, $._currentIndex);
        emit TransferWithMemo(from, to, $._currentIndex, "_mint");
        $._currentIndex++;
    }

    function transferWithMemo(
        address from,
        address to,
        uint256 _tokenId,
        string calldata memo
    ) external payable {
        safeTransferFrom(from, to, _tokenId);

        emit TransferWithMemo(from, to, _tokenId, memo);
    }

    /**
     * batchTransferWithMemo
     *
     * batch transfer with memo from one address to multiple addresses, it
     * assume the parameters array length are the same.
     * The tokens in `tokenIds` will be transferred to the addresses in the same
     * position in `tos`
     *
     * @param from - the start token ids,
     * @param tos - owner address to hold the new minted token
     * @param tokenIds - list of metadata to supply
     * @param memos - list of memo to supply
     */
    function batchTransferWithMemo(
        address from,
        address[] calldata tos,
        uint256[] calldata tokenIds,
        string[] calldata memos
    ) external payable {
        for (uint i = 0; i < tokenIds.length; i++) {
            safeTransferFrom(from, tos[i], tokenIds[i]);
            emit TransferWithMemo(from, tos[i], tokenIds[i], memos[i]);
        }
    }

    // Start Querying functions
    /**
     * getBookConfig
     *
     * getting the book config, owner can modify the book config field and use
     * it in update function
     *
     * @return the book config
     */
    function getBookConfig() public view returns (BookConfig memory) {
        BookNFTStorage storage $ = _getClassStorage();
        return
            BookConfig({
                name: $.name,
                symbol: $.symbol,
                metadata: $.metadata,
                max_supply: $.max_supply
            });
    }

    /**
     * getCurrentIndex
     *
     * getting the current index of the book nft, this is the index of the next token to be minted
     *
     * @return the current index
     */
    function getCurrentIndex() public view returns (uint256) {
        BookNFTStorage storage $ = _getClassStorage();
        return $._currentIndex;
    }

    function name() public view override returns (string memory) {
        BookNFTStorage storage $ = _getClassStorage();
        return $.name;
    }

    function symbol() public view override returns (string memory) {
        BookNFTStorage storage $ = _getClassStorage();
        return $.symbol;
    }

    function contractURI() public view returns (string memory) {
        BookNFTStorage storage $ = _getClassStorage();
        return string.concat("data:application/json;utf8,", $.metadata);
    }

    function tokenURI(
        uint256 _tokenId
    ) public view virtual override returns (string memory) {
        BookNFTStorage storage $ = _getClassStorage();
        return
            string.concat(
                "data:application/json;utf8,",
                $.tokenURIMap[_tokenId]
            );
    }
    // End Querying functions
}
