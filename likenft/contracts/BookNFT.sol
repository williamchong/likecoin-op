// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import {ERC721EnumerableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";
import {ERC721BurnableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721BurnableUpgradeable.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {IERC2981} from "@openzeppelin/contracts/interfaces/IERC2981.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

import {BookConfig} from "../types/BookConfig.sol";
import {MsgNewBookNFT} from "../types/msgs/MsgNewBookNFT.sol";

error ErrUnauthorized();
error ErrEmptyName();
error ErrInvalidSymbol();
error ErrInvalidMetadata();
error ErrMaxSupplyZero();
error ErrNftNoSupply();
error ErrTokenIdMintFails(uint256 nextTokenId);

contract BookNFT is
    Initializable,
    ERC721EnumerableUpgradeable,
    ERC721BurnableUpgradeable,
    OwnableUpgradeable,
    AccessControlUpgradeable,
    IERC2981
{
    struct BookNFTStorage {
        string name;
        string symbol;
        string metadata;
        uint64 max_supply;
        uint256 _currentIndex;
        mapping(uint256 => string) tokenURIMap;
        uint96 royaltyFraction;
        address protocolBeacon;
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

    modifier onlyProtocol() {
        BookNFTStorage storage $ = _getClassStorage();
        if (_msgSender() != $.protocolBeacon) {
            revert ErrUnauthorized();
        }
        _;
    }

    function initialize(MsgNewBookNFT memory msgNewBookNFT) public initializer {
        __ERC721_init(msgNewBookNFT.config.name, msgNewBookNFT.config.symbol);
        __ERC721Enumerable_init();
        __ERC721Burnable_init();
        __Ownable_init(msgNewBookNFT.creator);
        __AccessControl_init();

        _validateBookConfig(msgNewBookNFT.config);

        BookNFTStorage storage $ = _getClassStorage();
        $.name = msgNewBookNFT.config.name;
        $.symbol = msgNewBookNFT.config.symbol;
        $.max_supply = msgNewBookNFT.config.max_supply;
        $.metadata = msgNewBookNFT.config.metadata;

        $._currentIndex = 0;
        $.protocolBeacon = _msgSender();

        for (uint i = 0; i < msgNewBookNFT.minters.length; i++) {
            _grantRole(MINTER_ROLE, msgNewBookNFT.minters[i]);
        }
        for (uint i = 0; i < msgNewBookNFT.updaters.length; i++) {
            _grantRole(UPDATER_ROLE, msgNewBookNFT.updaters[i]);
        }
    }

    // Start of inheritence resolve
    function supportsInterface(
        bytes4 interfaceId
    )
        public
        view
        virtual
        override(
            ERC721Upgradeable,
            ERC721EnumerableUpgradeable,
            AccessControlUpgradeable,
            IERC165
        )
        returns (bool)
    {
        return
            interfaceId == type(IERC2981).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    function _increaseBalance(
        address account,
        uint128 amount
    ) internal override(ERC721Upgradeable, ERC721EnumerableUpgradeable) {
        super._increaseBalance(account, amount);
    }

    function _update(
        address to,
        uint256 tokenId,
        address auth
    )
        internal
        override(ERC721Upgradeable, ERC721EnumerableUpgradeable)
        returns (address)
    {
        return super._update(to, tokenId, auth);
    }
    // End of inheritence resolve

    function _validateBookConfig(BookConfig memory config) internal pure {
        if (bytes(config.name).length == 0) {
            revert ErrEmptyName();
        }
        if (bytes(config.symbol).length == 0) {
            revert ErrInvalidSymbol();
        }
        if (config.max_supply == 0) {
            revert ErrMaxSupplyZero();
        }
    }

    function update(BookConfig calldata config) public onlyUpdater {
        _validateBookConfig(config);
        BookNFTStorage storage $ = _getClassStorage();
        require(config.max_supply >= $.max_supply, "ErrSupplyDecrease");
        require(
            keccak256(bytes(config.symbol)) == keccak256(bytes($.symbol)),
            ErrInvalidSymbol()
        );

        $.name = config.name;
        $.symbol = config.symbol;
        $.max_supply = config.max_supply;
        $.metadata = config.metadata;
        emit ContractURIUpdated();
    }

    /**
     * mint function
     *
     * mint a new token with metadata, caller should ensure the supply is enough.
     *
     * @param to - owner address to hold the new minted token
     * @param metadataList - list of metadata to supply
     */
    function mint(
        address to,
        string[] calldata memos,
        string[] calldata metadataList
    ) external onlyMinter {
        require(
            memos.length == metadataList.length,
            "ErrMemoMetadataLengthMismatch"
        );
        _ensureEnoughSupply(metadataList.length);
        for (uint i = 0; i < metadataList.length; i++) {
            _mintWithEvent(_msgSender(), to, memos[i], metadataList[i]);
        }
    }

    /**
     * batchMint
     *
     * batch mint with metadata list
     *
     * @param tos - owner address to hold the new minted token
     * @param memos - list of memo to supply
     * @param metadataList - list of metadata to supply, the length of the list should be the same as the length of the tos. Metadata will fill the corresponding position of the tos.
     */
    function batchMint(
        address[] calldata tos,
        string[] calldata memos,
        string[] calldata metadataList
    ) external onlyMinter {
        _ensureEnoughSupply(metadataList.length);
        for (uint i = 0; i < tos.length; i++) {
            _mintWithEvent(_msgSender(), tos[i], memos[i], metadataList[i]);
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
     * @param tos - owner address to hold the new minted token
     * @param memos - list of memo to supply
     * @param metadataList - list of metadata to supply
     */
    function safeMintWithTokenId(
        uint256 fromTokenId,
        address[] calldata tos,
        string[] calldata memos,
        string[] calldata metadataList
    ) external onlyMinter {
        if (totalSupply() != fromTokenId) {
            revert ErrTokenIdMintFails(totalSupply());
        }
        _ensureEnoughSupply(metadataList.length);
        for (uint i = 0; i < metadataList.length; i++) {
            _mintWithEvent(_msgSender(), tos[i], memos[i], metadataList[i]);
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
     * @param memo - memo to supply
     * @param metadata - metadata to supply
     */
    function _mintWithEvent(
        address from,
        address to,
        string calldata memo,
        string calldata metadata
    ) internal {
        BookNFTStorage storage $ = _getClassStorage();
        $.tokenURIMap[$._currentIndex] = metadata;
        _safeMint(to, $._currentIndex);
        emit TransferWithMemo(from, to, $._currentIndex, memo);
        $._currentIndex++;
    }

    function transferWithMemo(
        address from,
        address to,
        uint256 _tokenId,
        string calldata memo
    ) external {
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
    ) external {
        for (uint i = 0; i < tokenIds.length; i++) {
            safeTransferFrom(from, tos[i], tokenIds[i]);
            emit TransferWithMemo(from, tos[i], tokenIds[i], memos[i]);
        }
    }

    /**
     * setRoyaltyFraction
     *
     * set the royalty fraction for the book nft.
     * The feeDenominator is 10000.
     * Intended to only support BookNFTs based royalty, not per token based royalty.
     *
     * @param royaltyFraction - the royalty fraction to set
     */
    function setRoyaltyFraction(uint96 royaltyFraction) external onlyProtocol {
        BookNFTStorage storage $ = _getClassStorage();
        $.royaltyFraction = royaltyFraction;
    }

    /**
     * royaltyInfo
     *
     * getting the royalty info for a token sale.
     * In phase 1 of likeprotocol, all royalties will be sent to the BookNFT address as vested.
     * In later phase, the royalties withdrwal logic will be implemented.
     * The royalty is designed to be tied with the LikeProtocol contract.
     *
     * @param - To confronyt the token ID to get royalty info for
     * @param salePrice - the sale price of the token
     * @return receiver - the address that should receive the royalty payment
     * @return royaltyAmount - the amount of royalty to be paid
     */
    function royaltyInfo(
        uint256,
        uint256 salePrice
    ) external view override returns (address receiver, uint256 royaltyAmount) {
        BookNFTStorage storage $ = _getClassStorage();
        royaltyAmount = (salePrice * $.royaltyFraction) / 10000;
        receiver = address(this);
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

    function maxSupply() public view returns (uint64) {
        BookNFTStorage storage $ = _getClassStorage();
        return $.max_supply;
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

    function getProtocolBeacon() public view returns (address) {
        BookNFTStorage storage $ = _getClassStorage();
        return $.protocolBeacon;
    }
    // End Querying functions
}
