// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {ERC4626Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC4626Upgradeable.sol";
import {IERC4626} from "@openzeppelin/contracts/interfaces/IERC4626.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Likecoin} from "./Likecoin.sol";

/// @custom:security-contact rickmak@oursky.com
contract veLike is
    ERC4626Upgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable
{
    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address initialOwner, address like) public initializer {
        __ERC4626_init(IERC20(address(like)));
        __ERC20_init("vote-escrowed LikeCoin", "veLIKE");
        __Pausable_init();
        __ReentrancyGuard_init();
        __Ownable_init(initialOwner);
        __UUPSUpgradeable_init();
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}

    /* Start of IERC4626 — Tokenized Vaults interface (OpenZeppelin Contracts v5) */
    /**
     * @dev Returns the underlying asset of the vault.
     */
    function asset() public view override returns (address) {
        return address(asset());
    }

    /**
     * @dev Returns the total amount of the underlying asset managed by the vault.
     */
    function totalAssets() public view override returns (uint256) {
        return totalAssets();
    }

    /**
     * @dev Returns the current amount of shares that `assets` would mint.
     */
    function convertToShares(
        uint256 assets
    ) public view override returns (uint256) {
        return convertToShares(assets);
    }

    /**
     * @dev Returns the current amount of assets that `shares` would redeem.
     */
    function convertToAssets(
        uint256 shares
    ) public view override returns (uint256) {
        return convertToAssets(shares);
    }

    /**
     * @dev Returns the maximum amount of assets that `receiver` can deposit right now.
     */
    function maxDeposit(
        address receiver
    ) public view override returns (uint256) {
        return maxDeposit(receiver);
    }

    /**
     * @dev Allows an on-chain or off-chain user to simulate the effects of their deposit at the current block, given current on-chain conditions.
     */
    function previewDeposit(
        uint256 assets
    ) public view override returns (uint256) {
        return previewDeposit(assets);
    }

    /**
     * @dev Mints shares Vault shares to receiver by depositing exactly amount of underlying tokens.
     */
    function deposit(
        uint256 assets,
        address receiver
    ) public override returns (uint256 shares) {
        return deposit(assets, receiver);
    }

    /**
  @dev Returns the maximum amount of the Vault shares that can be minted for the receiver, through a mint call.


**/
    function maxMint(address receiver) public view override returns (uint256) {
        return maxMint(receiver);
    }

    /**
  @dev Allows an on-chain or off-chain user to simulate the effects of their mint at the current block, given current on-chain conditions.
**/
    function previewMint(
        uint256 shares
    ) public view override returns (uint256) {
        return previewMint(shares);
    }

    /**
  @dev Mints exactly shares Vault shares to receiver by depositing amount of underlying tokens.
**/
    function mint(
        uint256 shares,
        address receiver
    ) public override returns (uint256 assets) {
        return mint(shares, receiver);
    }

    /**
  @dev Returns the maximum amount of the underlying asset that can be withdrawn from the owner balance in the Vault, through a withdraw call.
**/
    function maxWithdraw(address owner) public view override returns (uint256) {
        return maxWithdraw(owner);
    }

    /**
  @dev Allows an on-chain or off-chain user to simulate the effects of their withdraw at the current block, given current on-chain conditions.
**/
    function previewWithdraw(
        uint256 assets
    ) public view override returns (uint256) {
        return previewWithdraw(assets);
    }

    /**
  @dev Burns shares from owner and sends exactly assets of underlying tokens to receiver.
**/
    function withdraw(
        uint256 assets,
        address receiver,
        address owner
    ) public override returns (uint256 shares) {
        return withdraw(assets, receiver, owner);
    }

    /**
  @dev Returns the maximum amount of Vault shares that can be redeemed from the owner balance in the Vault, through a redeem call.
**/
    function maxRedeem(address owner) public view override returns (uint256) {
        return maxRedeem(owner);
    }

    /**
  @dev Allows an on-chain or off-chain user to simulate the effects of their redemption at the current block, given current on-chain conditions.
**/
    function previewRedeem(
        uint256 shares
    ) public view override returns (uint256) {
        return previewRedeem(shares);
    }

    /**
  @dev Burns exactly shares from owner and sends assets of underlying tokens to receiver.


**/
    function redeem(
        uint256 shares,
        address receiver,
        address owner
    ) public override returns (uint256 assets) {
        return redeem(shares, receiver, owner);
    }
}
