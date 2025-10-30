// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {ERC4626Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC4626Upgradeable.sol";
import {ERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC4626} from "@openzeppelin/contracts/interfaces/IERC4626.sol";
import {Likecoin} from "./Likecoin.sol";

interface IRewardContract {
    function getPendingReward(address account) external view returns (uint256);
    function claimReward(address account, bool restake) external returns (uint256);
    function deposit(address account, uint256 rewardAmount) external;
    function withdraw(address account, uint256 rewardAmount) external;
}

/// @custom:security-contact rickmak@oursky.com
contract veLike is
    ERC4626Upgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable
{
    struct veLikeStorage {
        address rewardContract;
    }

    // keccak256(abi.encode(uint256(keccak256("veLike.storage")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CLASS_DATA_STORAGE =
        0xb9e14b2a89d227541697d62a06ecbf5ccc9ad849800745b40b2826662a177600;

    function _getveLikeData() private pure returns (veLikeStorage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := CLASS_DATA_STORAGE
        }
    }

    // Errors
    error ErrNoRewardToClaim();
    error ErrNonTransferable();

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

    // Start of veLike specific functions

    /**
     * setRewardContract function
     *
     * Set the reward contract for the veLike.
     *
     * @param rewardContract - the reward contract to set
     */
    function setRewardContract(address rewardContract) public onlyOwner {
        veLikeStorage storage $ = _getveLikeData();
        $.rewardContract = rewardContract;
    }

    /**
     * getCurrentCondition function
     *
     * Get the current staking condition, it can be inactive. i.e. not started or already ended.
     *
     * @return currentCondition - the current staking condition
     */
    function getCurrentRewardContract() public view returns (IRewardContract) {
        veLikeStorage storage $ = _getveLikeData();
        return IRewardContract($.rewardContract);
    }

    /**
     * getPendingReward function
     *
     * Get the pending reward for the account. Calculated to the query block height.
     * In subsequent claim, the reward might be more as block height is updated.
     *
     * @param account - the account to get the pending reward for
     * @return pendingReward - the pending reward for the account
     */
    function getPendingReward(address account) public view returns (uint256) {
        IRewardContract rewardContract = getCurrentRewardContract();
        if (rewardContract == IRewardContract(address(0))) {
            return 0;
        }
        return rewardContract.getPendingReward(account);
    }

    /**
     * claimReward function
     *
     * Claim the reward for the account.
     *
     * @param account - the account to claim the reward for
     * @return reward - the reward for the account
     */
    function claimReward(
        address account
    ) public whenNotPaused nonReentrant returns (uint256) {
        IRewardContract rewardContract = getCurrentRewardContract();
        if (rewardContract == IRewardContract(address(0))) {
            revert ErrNoRewardToClaim();
        }
        uint256 reward = rewardContract.claimReward(account, false);
        return reward;
    }

    /**
     * restakeReward function
     *
     * Restake the reward for the account.
     *
     * @param account - the account to restake the reward
     * @return reward - the amount of asset restaked
     */
    function restakeReward(
        address account
    ) public whenNotPaused nonReentrant returns (uint256) {
        IRewardContract rewardContract = getCurrentRewardContract();
        if (rewardContract == IRewardContract(address(0))) {
            revert ErrNoRewardToClaim();
        }
        uint256 reward = rewardContract.claimReward(account, true);
        _mint(account, reward);
        return reward;
    }

    // End of veLike specific functions
    // Start of ERC20 Overrides

    /**
     * transfer function
     *
     * veLIKE is non-transferable voting escrow token, so it should not be transferred.
     * Override ERC20 transfer function to revert.
     *
     * @return bool - true if the transfer is successful
     */
    function transfer(
        address,
        uint256
    ) public virtual override(ERC20Upgradeable, IERC20) returns (bool) {
        revert ErrNonTransferable();
    }

    /**
     * transferFrom function
     *
     * veLIKE is non-transferable voting escrow token, so it should not be transferred.
     * Override ERC20 transferFrom function to revert.
     *
     * @return bool - true if the transfer is successful
     */
    function transferFrom(
        address,
        address,
        uint256
    ) public virtual override(ERC20Upgradeable, IERC20) returns (bool) {
        revert ErrNonTransferable();
    }
    // End of ERC20 Overrides

    // Start of ERC4626 Overrides
    /**
     * totalAssets function
     *
     * veLike to Like should be one to one mapping, so the total supply is equal to the total assets.
     * Note: Vault share is not veLike.
     */
    function totalAssets() public view override returns (uint256) {
        return totalSupply();
    }
    /**
     * _deposit function
     *
     * Override ERC4626 _deposit function to update staker info on vault share. mint
     *
     * @param caller - the caller of the deposit
     * @param receiver - the receiver of the vault share
     * @param assets - the amount of asset to deposit
     * @param shares - the amount of shares to mint
     */
    function _deposit(
        address caller,
        address receiver,
        uint256 assets,
        uint256 shares
    ) internal virtual override whenNotPaused {
        // Copying from ERC4626 _deposit function for clarity
        SafeERC20.safeTransferFrom(
            IERC20(asset()),
            caller,
            address(this),
            assets
        );
        _mint(receiver, shares);

        // Vault specific logic
        IRewardContract rewardContract = getCurrentRewardContract();
        if (rewardContract != IRewardContract(address(0))) {
            rewardContract.deposit(receiver, assets);
        }

        // Copying from ERC4626 _deposit function Event for clarity
        emit Deposit(caller, receiver, assets, shares);
    }

    /**
     * _withdraw function
     *
     * Override ERC4626 _withdraw function to update staker info on vault share. burn
     *
     * @param caller - the caller of the withdraw
     * @param receiver - the receiver of the vault share
     * @param assets - the amount of asset to withdraw
     * @param shares - the amount of shares to burn
     */
    function _withdraw(
        address caller,
        address receiver,
        address owner,
        uint256 assets,
        uint256 shares
    ) internal virtual override whenNotPaused {
        // Copying from ERC4626 _withdraw function for clarity
        // Same as calling super._withdraw(caller, receiver, assets, shares);
        if (caller != owner) {
            _spendAllowance(owner, caller, shares);
        }

        // Vault specific logic
        IRewardContract rewardContract = getCurrentRewardContract();
        if (rewardContract != IRewardContract(address(0))) {
            rewardContract.withdraw(owner, assets);
        }

        // Copying from ERC4626 _withdraw function Event for clarity
        _burn(owner, shares);
        SafeERC20.safeTransfer(IERC20(asset()), receiver, assets);
        emit Withdraw(caller, receiver, owner, assets, shares);
    }
    // End of ERC4626 Overrides

    // Start of Admin functions

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }
    // End of Admin functions
}
