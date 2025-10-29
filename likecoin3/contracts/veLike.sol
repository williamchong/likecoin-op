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

/// @custom:security-contact rickmak@oursky.com
contract veLike is
    ERC4626Upgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable
{
    struct StakingCondition {
        uint256 startTime;
        uint256 endTime;
        uint256 rewardAmount;
        uint256 rewardIndex;
    }

    struct StakerInfo {
        uint256 stakedAmount;
        uint256 rewardIndex;
        uint256 lockedUntil;
        uint256 rewardClaimed; // Not use for calculation, only for tracking.
    }

    struct veLikeStorage {
        uint256 rewardPool; // Tracking the likecoin hold by vault for reward distribution.
        uint256 totalStaked;
        uint256 lastRewardTime;
        uint256 currentConditionIndex;
        mapping(address account => StakerInfo stakerInfo) stakerInfos;
        mapping(uint256 index => StakingCondition condition) conditions;
    }

    uint256 public constant ACC_REWARD_PRECISION = 1e18; // Precision scalar for reward index

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
    error ErrConflictCondition();
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
     * getCurrentCondition function
     *
     * Get the current staking condition, it can be inactive. i.e. not started or already ended.
     *
     * @return currentCondition - the current staking condition
     */
    function getCurrentCondition()
        public
        view
        returns (StakingCondition memory)
    {
        StakingCondition memory currentCondition;
        bool _isActive;
        (currentCondition, _isActive) = _activeCondition();
        return currentCondition;
    }

    function getClaimedReward(address account) public view returns (uint256) {
        veLikeStorage storage $ = _getveLikeData();
        StakerInfo memory stakerInfo = $.stakerInfos[account];
        return stakerInfo.rewardClaimed;
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
        veLikeStorage storage $ = _getveLikeData();
        StakingCondition memory currentCondition;
        bool _isActive;
        (currentCondition, _isActive) = _activeCondition();
        uint256 calculatedReward = _pendingReward(account);
        if (!_isActive) {
            return calculatedReward;
        }
        uint256 veLikeAmount = balanceOf(account);
        if (veLikeAmount == 0) {
            return 0;
        }
        uint256 timePassed = block.timestamp - $.lastRewardTime;
        uint256 newReward = timePassed *
            _rewardPerTimeWithPrecision(currentCondition);
        uint256 nonCalculatedReward = (newReward * veLikeAmount) /
            ($.totalStaked * ACC_REWARD_PRECISION);
        return calculatedReward + nonCalculatedReward;
    }

    /**
     * _pendingReward function
     *
     * Internal function to calculate the pending reward for the account.
     *
     */
    function _pendingReward(address account) internal view returns (uint256) {
        veLikeStorage storage $ = _getveLikeData();
        StakerInfo memory stakerInfo = $.stakerInfos[account];
        return
            (stakerInfo.stakedAmount *
                ($.conditions[$.currentConditionIndex].rewardIndex -
                    stakerInfo.rewardIndex)) / ACC_REWARD_PRECISION;
    }

    /**
     * _claimReward function
     *
     * Internal function to claim the reward for the account.
     *
     * @param account - the account to claim the reward for
     * @param restake - true if the reward should be restaked, false if the reward should be claimed
     */
    function _claimReward(address account, bool restake) internal returns (uint256) {
        veLikeStorage storage $ = _getveLikeData();
        StakingCondition storage condition = $.conditions[
            $.currentConditionIndex
        ];
        StakerInfo storage stakerInfo = $.stakerInfos[account];
        uint256 rewardClaimed = _pendingReward(account);
        stakerInfo.rewardClaimed += rewardClaimed;
        stakerInfo.rewardIndex = condition.rewardIndex;
        $.rewardPool -= rewardClaimed;
        if (restake) {
            stakerInfo.stakedAmount += rewardClaimed;
            $.totalStaked += rewardClaimed;
            _mint(account, rewardClaimed);
        } else {
            SafeERC20.safeTransfer(IERC20(asset()), account, rewardClaimed);
        }
        return rewardClaimed;    }

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
        uint256 pendingReward = getPendingReward(account);
        if (pendingReward == 0) {
            revert ErrNoRewardToClaim();
        }

        _updateVault();
        uint256 rewardClaimed = _claimReward(account, false);
        return rewardClaimed;
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
    ) public nonReentrant returns (uint256) {
        uint256 pendingReward = getPendingReward(account);
        if (pendingReward == 0) {
            revert ErrNoRewardToClaim();
        }

        _updateVault();
        uint256 rewardClaimed = _claimReward(account, true);
        return rewardClaimed;
    }

    function _activeCondition()
        internal
        view
        returns (StakingCondition storage, bool)
    {
        veLikeStorage storage $ = _getveLikeData();
        StakingCondition storage currentCondition = $.conditions[
            $.currentConditionIndex
        ];
        if (
            block.timestamp < currentCondition.startTime ||
            block.timestamp > currentCondition.endTime
        ) {
            return (currentCondition, false);
        }
        return (currentCondition, true);
    }
    /**
     * _updateVault function
     *
     * Update the vault reward index and reward debt.
     *
     */
    function _updateVault() internal {
        veLikeStorage storage $ = _getveLikeData();
        StakingCondition storage currentCondition;
        bool _isActive;
        (currentCondition, _isActive) = _activeCondition();
        uint256 targetTime = block.timestamp;
        if (!_isActive) {
            // Perform last update if needed
            if (
                $.lastRewardTime > currentCondition.startTime &&
                $.lastRewardTime < currentCondition.endTime
            ) {
                targetTime = currentCondition.endTime;
            } else {
                return;
            }
        }
        if ($.totalStaked > 0) {
            uint256 timePassed = targetTime - $.lastRewardTime;
            uint256 reward = timePassed *
                _rewardPerTimeWithPrecision(currentCondition);
            currentCondition.rewardIndex += reward / $.totalStaked;
            $.lastRewardTime = targetTime;
        }
    }

    function _rewardPerTimeWithPrecision(
        StakingCondition memory condition
    ) internal pure returns (uint256) {
        return
            (ACC_REWARD_PRECISION * condition.rewardAmount) /
            (condition.endTime - condition.startTime);
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
        veLikeStorage storage $ = _getveLikeData();
        // Copying from ERC4626 _deposit function for clarity
        SafeERC20.safeTransferFrom(
            IERC20(asset()),
            caller,
            address(this),
            assets
        );
        _mint(receiver, shares);

        // Vault specific logic
        _updateVault();
        // Note, we must claim the reward, othereise the denominator will be wrong on next claim.
        _claimReward(receiver, false);
        $.totalStaked += assets;
        $.stakerInfos[receiver].stakedAmount += assets;

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

        _burn(owner, shares);
        SafeERC20.safeTransfer(IERC20(asset()), receiver, assets);

        // Vault specific logic
        _updateVault();

        // Copying from ERC4626 _withdraw function Event for clarity
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

    /**
     * getLastRewardTime function
     *
     * Get the last reward time.
     *
     * @return lastRewardTime - the last reward time
     */
    function getLastRewardTime() public view returns (uint256) {
        veLikeStorage storage $ = _getveLikeData();
        return $.lastRewardTime;
    }

    function getRewardPool() public view returns (uint256) {
        veLikeStorage storage $ = _getveLikeData();
        return $.rewardPool;
    }

    /**
     * addReward function
     *
     * Admin function for authorized address too deposit asset as reward. This
     * staking vault rewards is linearly over time. reward calculation is update in the current block timestamp.
     *
     * @param rewardAmount - the amount of reward to deposit, asset ERC20(likecoin)
     * @param endTime - the end time of the staking condition
     */
    function addReward(
        uint256 rewardAmount,
        uint256 startTime,
        uint256 endTime
    ) external onlyOwner {
        veLikeStorage storage $ = _getveLikeData();
        if (startTime <= $.lastRewardTime) {
            revert ErrConflictCondition();
        }
        if (endTime < startTime) {
            revert ErrConflictCondition();
        }
        if (endTime < block.timestamp) {
            revert ErrConflictCondition();
        }
        if ($.currentConditionIndex == 0) {
            $.lastRewardTime = startTime;
        }
        SafeERC20.safeTransferFrom(
            IERC20(asset()),
            _msgSender(),
            address(this),
            rewardAmount
        );
        // perform last update if needed
        _updateVault();
        uint256 lastRewardIndex = $
            .conditions[$.currentConditionIndex]
            .rewardIndex;
        $.rewardPool += rewardAmount;
        $.currentConditionIndex++;
        $.conditions[$.currentConditionIndex] = StakingCondition({
            startTime: startTime,
            endTime: endTime,
            rewardAmount: rewardAmount,
            rewardIndex: lastRewardIndex
        });
    }

    // End of Admin functions
}
