// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC4626} from "@openzeppelin/contracts/interfaces/IERC4626.sol";

/// @custom:security-contact rickmak@oursky.com
contract veLikeRewardNoLock is
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
        uint256 rewardClaimed; // Not use for calculation, only for tracking.
    }

    struct veLikeRewardStorage {
        address vault;
        address likecoin;
        uint256 rewardPool; // Tracking the likecoin pool authorized for reward distribution.
        uint256 totalStaked;
        uint256 lastRewardTime;
        StakingCondition currentStakingCondition;
        mapping(address account => StakerInfo stakerInfo) stakerInfos;
        address drawer;
        bool autoSyncEnabled; // Set by initTotalStaked() to enable lazy staker sync.
    }

    uint256 public constant ACC_REWARD_PRECISION = 1e18; // Precision scalar for reward index

    // keccak256(abi.encode(uint256(keccak256("veLikeReward.storage")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant CLASS_DATA_STORAGE =
        0xe9672d2c676bb94d428d6ce523668c779079df8febe4142a9972a2a2313d2c00;

    function _getveLikeRewardData()
        private
        pure
        returns (veLikeRewardStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := CLASS_DATA_STORAGE
        }
    }

    // Errors
    error ErrNoRewardToClaim();
    error ErrConflictCondition();
    error ErrUnauthorized();
    error ErrNotActive();
    error ErrAlreadySynced();
    error ErrMismatchSync();

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address initialOwner) public initializer {
        __Pausable_init();
        __ReentrancyGuard_init();
        __Ownable_init(initialOwner);
        __UUPSUpgradeable_init();
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}

    modifier onlyVault() {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        if (_msgSender() != $.vault) {
            revert ErrUnauthorized();
        }
        _;
    }

    // Start of veLikeRewardNoLock specific functions

    function setVault(address vault) public onlyOwner {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        $.vault = vault;
    }
    function setLikecoin(address likecoin) public onlyOwner {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        $.likecoin = likecoin;
    }
    function getConfig()
        public
        view
        returns (address, address, uint256, uint256, uint256)
    {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        return (
            $.vault,
            $.likecoin,
            $.rewardPool,
            $.totalStaked,
            $.lastRewardTime
        );
    }

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
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        return $.currentStakingCondition;
    }

    function getClaimedReward(address account) public view returns (uint256) {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        StakerInfo memory stakerInfo = $.stakerInfos[account];
        return stakerInfo.rewardClaimed;
    }

    /**
     * getPendingReward function
     *
     * Get the pending reward for the account. Calculated to the query block height.
     * In subsequent claim, the reward might be more as block height is updated.
     *
     * For un-synced stakers (pre-rotation stakers who have vault balance but
     * stakedAmount == 0), the vault balance is used as the effective stake.
     *
     * @param account - the account to get the pending reward for
     * @return pendingReward - the pending reward for the account
     */
    function getPendingReward(address account) public view returns (uint256) {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        uint256 calculatedReward = _pendingReward(account);
        uint256 stakedAmount = _effectiveStakedAmount(account);
        if (
            stakedAmount == 0 ||
            $.totalStaked == 0 ||
            $.currentStakingCondition.endTime <=
                $.currentStakingCondition.startTime
        ) {
            return calculatedReward;
        }
        uint256 targetTime = block.timestamp;
        if (targetTime > $.currentStakingCondition.endTime) {
            targetTime = $.currentStakingCondition.endTime;
        }
        uint256 timePassed = targetTime - $.lastRewardTime;
        uint256 newReward = timePassed *
            _rewardPerTimeWithPrecision($.currentStakingCondition);
        uint256 nonCalculatedReward = (newReward * stakedAmount) /
            ($.totalStaked * ACC_REWARD_PRECISION);
        return calculatedReward + nonCalculatedReward;
    }

    /**
     * _pendingReward function
     *
     * Internal function to calculate the pending reward for the account.
     * Uses _effectiveStakedAmount to handle un-synced pre-rotation stakers.
     *
     */
    function _pendingReward(address account) internal view returns (uint256) {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        StakerInfo memory stakerInfo = $.stakerInfos[account];
        uint256 stakedAmount = _effectiveStakedAmount(account);
        return
            (stakedAmount *
                ($.currentStakingCondition.rewardIndex -
                    stakerInfo.rewardIndex)) / ACC_REWARD_PRECISION;
    }

    /**
     * _effectiveStakedAmount function
     *
     * Returns the effective staked amount for reward calculation.
     * For synced users, returns stakerInfo.stakedAmount.
     * For un-synced pre-rotation stakers (stakedAmount == 0 but vault balance > 0),
     * returns the vault balance so they earn retroactive rewards.
     * This fallback only applies when autoSyncEnabled is true (set by initTotalStaked).
     */
    function _effectiveStakedAmount(
        address account
    ) internal view returns (uint256) {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        uint256 stakedAmount = $.stakerInfos[account].stakedAmount;
        if (stakedAmount == 0 && $.autoSyncEnabled) {
            return IERC4626($.vault).balanceOf(account);
        }
        return stakedAmount;
    }

    /**
     * _syncStaker function
     *
     * Lazy-sync a pre-rotation staker into this reward contract.
     * Only operates when autoSyncEnabled is true (set by initTotalStaked).
     * If stakerInfo.stakedAmount == 0 but the user has a vault balance,
     * sets stakedAmount to match the vault balance. The user's rewardIndex
     * stays at 0, so they earn retroactive rewards from the period start
     * (since addReward resets rewardIndex to 0).
     *
     * totalStaked is NOT adjusted because it was pre-initialized via
     * initTotalStaked() to include all vault holders.
     */
    function _syncStaker(address account) internal {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        if (!$.autoSyncEnabled) {
            return;
        }
        StakerInfo storage stakerInfo = $.stakerInfos[account];
        if (stakerInfo.stakedAmount != 0) {
            return;
        }
        uint256 vaultBalance = IERC4626($.vault).balanceOf(account);
        if (vaultBalance == 0) {
            return;
        }
        stakerInfo.stakedAmount = vaultBalance;
    }

    function _isActive() internal view returns (bool) {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        if (
            block.timestamp < $.currentStakingCondition.startTime ||
            block.timestamp > $.currentStakingCondition.endTime
        ) {
            return false;
        }
        return true;
    }

    /**
     * _updateVault function
     *
     * Update the vault reward index and reward debt.
     *
     */
    function _updateVault() internal {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        StakingCondition storage currentCondition = $.currentStakingCondition;
        uint256 targetTime = block.timestamp;
        if (targetTime < currentCondition.startTime) {
            targetTime = currentCondition.startTime;
        }
        if (targetTime > currentCondition.endTime) {
            targetTime = currentCondition.endTime;
        }
        if (targetTime == $.lastRewardTime) {
            return;
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

    // End of veLikeRewardNoLock specific functions

    // Start of Vault functions

    function deposit(
        address account,
        uint256 stakedAmount
    ) public whenNotPaused onlyVault {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        _syncStaker(account);
        _updateVault();
        _claimReward(account, false);
        $.stakerInfos[account].stakedAmount += stakedAmount;
        $.totalStaked += stakedAmount;
    }

    function withdraw(
        address account,
        uint256 amount
    ) public whenNotPaused onlyVault {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        _syncStaker(account);
        _updateVault();
        _claimReward(account, false);
        $.totalStaked -= amount;
        $.stakerInfos[account].stakedAmount -= amount;
    }

    /**
     * claimReward function
     *
     * Claim the reward for the account, only caller by vault.
     *
     * @param account - the account to claim the reward for
     * @param restake - true if the reward should be restaked, false if the reward should be claimed
     * @return reward - the reward for the account
     */
    function claimReward(
        address account,
        bool restake
    ) public whenNotPaused onlyVault returns (uint256) {
        _syncStaker(account);
        uint256 currentPendingReward = getPendingReward(account);
        if (currentPendingReward == 0) {
            revert ErrNoRewardToClaim();
        }
        return _claimReward(account, restake);
    }

    /**
     * _claimReward function
     *
     * Claim the reward for the account.
     *
     * @param account - the account to claim the reward for
     * @param restake - true if the reward should be restaked, false if the reward should be claimed
     * @return reward - the reward for the account
     */
    function _claimReward(
        address account,
        bool restake
    ) public onlyVault returns (uint256) {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        StakerInfo storage stakerInfo = $.stakerInfos[account];

        _updateVault();
        uint256 rewardClaimed = _pendingReward(account);
        stakerInfo.rewardClaimed += rewardClaimed;
        stakerInfo.rewardIndex = $.currentStakingCondition.rewardIndex;
        $.rewardPool -= rewardClaimed;
        if (rewardClaimed == 0) {
            return 0;
        }
        if (restake) {
            stakerInfo.stakedAmount += rewardClaimed;
            $.totalStaked += rewardClaimed;
            // Relay on the Vault to _mint the veLIKE.
        } else {
            SafeERC20.safeTransferFrom(
                IERC20($.likecoin),
                $.drawer,
                account,
                rewardClaimed
            );
        }
        return rewardClaimed;
    }
    // End of Vault functions

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
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        return $.lastRewardTime;
    }

    function getRewardPool() public view returns (uint256) {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        return $.rewardPool;
    }

    /**
     * initTotalStaked function
     *
     * Initialize totalStaked from the vault's totalSupply and enable
     * auto-sync for pre-rotation stakers. Called once during deployment
     * setup (after setVault) to ensure the reward accumulator uses the
     * correct denominator that includes all existing vault holders.
     */
    function initTotalStaked() external onlyOwner {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        require(!$.autoSyncEnabled, "Already initialized");
        $.totalStaked = IERC4626($.vault).totalSupply();
        $.autoSyncEnabled = true;
    }

    /**
     * syncStakers function
     *
     * Admin function to eagerly sync pre-rotation stakers into this reward
     * contract. Must be called during the active reward period (between
     * startTime and endTime). For each account, sets stakedAmount to the
     * current vault balance. The staker's rewardIndex stays at 0 so they
     * earn retroactive rewards from the period start.
     *
     * totalStaked is NOT adjusted because it was pre-initialized via
     * initTotalStaked() to include all vault holders.
     *
     * Reverts with ErrAlreadySynced if the account is already synced and
     * the stakedAmount matches the vault balance. Reverts with
     * ErrMismatchSync if the account is already synced but the
     * stakedAmount differs from the vault balance.
     *
     * @param accounts - the accounts to sync
     */
    function syncStakers(address[] calldata accounts) external onlyOwner {
        if (!_isActive()) {
            revert ErrNotActive();
        }
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            uint256 vaultBalance = IERC4626($.vault).balanceOf(account);
            StakerInfo storage stakerInfo = $.stakerInfos[account];
            if (stakerInfo.stakedAmount != 0) {
                if (stakerInfo.stakedAmount == vaultBalance) {
                    revert ErrAlreadySynced();
                } else {
                    revert ErrMismatchSync();
                }
            }
            stakerInfo.stakedAmount = vaultBalance;
        }
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
        address drawer,
        uint256 rewardAmount,
        uint256 startTime,
        uint256 endTime
    ) external onlyOwner {
        if (_isActive()) {
            revert ErrConflictCondition();
        }
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        if (startTime <= $.lastRewardTime) {
            revert ErrConflictCondition();
        }
        if (endTime < startTime) {
            revert ErrConflictCondition();
        }
        if (endTime < block.timestamp) {
            revert ErrConflictCondition();
        }
        $.lastRewardTime = startTime;
        $.drawer = drawer;
        // perform last update if needed
        $.rewardPool += rewardAmount;
        $.currentStakingCondition = StakingCondition({
            startTime: startTime,
            endTime: endTime,
            rewardAmount: rewardAmount,
            rewardIndex: 0
        });
    }

    // End of Admin functions
}
