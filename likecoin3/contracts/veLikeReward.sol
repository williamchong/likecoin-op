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
contract veLikeReward is
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
        uint256 rewardPool; // Tracking the likecoin hold by vault for reward distribution.
        uint256 totalStaked;
        uint256 lastRewardTime;
        uint256 currentConditionIndex;
        mapping(address account => StakerInfo stakerInfo) stakerInfos;
        mapping(uint256 index => StakingCondition condition) conditions;
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

    // Start of veLikeReward specific functions

    function setVault(address vault) public onlyOwner {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        $.vault = vault;
    }
    function setLikecoin(address likecoin) public onlyOwner {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        $.likecoin = likecoin;
    }
    function getConfig() public view returns (address, address, uint256, uint256, uint256) {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        return ($.vault, $.likecoin, $.rewardPool, $.totalStaked, $.lastRewardTime);
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
        StakingCondition memory currentCondition;
        bool _isActive;
        (currentCondition, _isActive) = _activeCondition();
        return currentCondition;
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
     * @param account - the account to get the pending reward for
     * @return pendingReward - the pending reward for the account
     */
    function getPendingReward(address account) public view returns (uint256) {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        StakingCondition memory currentCondition;
        bool _isActive;
        (currentCondition, _isActive) = _activeCondition();
        uint256 calculatedReward = _pendingReward(account);
        if (!_isActive) {
            return calculatedReward;
        }
        uint256 veLikeAmount = $.stakerInfos[account].stakedAmount;
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
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        StakerInfo memory stakerInfo = $.stakerInfos[account];
        return
            (stakerInfo.stakedAmount *
                ($.conditions[$.currentConditionIndex].rewardIndex -
                    stakerInfo.rewardIndex)) / ACC_REWARD_PRECISION;
    }

    function _activeCondition()
        internal
        view
        returns (StakingCondition storage, bool)
    {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
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
        veLikeRewardStorage storage $ = _getveLikeRewardData();
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

    // End of veLikeReward specific functions

    // Start of Vault functions

    function deposit(address account, uint256 rewardAmount) public onlyVault {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        _updateVault();
        // Note, we must claim the reward, othereise the denominator will be wrong on next claim.
        _claimReward(account, false);
        $.stakerInfos[account].stakedAmount += rewardAmount;
        $.totalStaked += rewardAmount;
    }

    function withdraw(address account, uint256 rewardAmount) public onlyVault {
        veLikeRewardStorage storage $ = _getveLikeRewardData();
        $.rewardPool -= rewardAmount;
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
    function claimReward(address account, bool restake) public onlyVault returns (uint256) {
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
        StakingCondition storage condition = $.conditions[
            $.currentConditionIndex
        ];
        StakerInfo storage stakerInfo = $.stakerInfos[account];

        _updateVault();
        uint256 rewardClaimed = _pendingReward(account);
        stakerInfo.rewardClaimed += rewardClaimed;
        stakerInfo.rewardIndex = condition.rewardIndex;
        $.rewardPool -= rewardClaimed;
        if (restake) {
            stakerInfo.stakedAmount += rewardClaimed;
            $.totalStaked += rewardClaimed;
            // Relay on the Vault to _mint the veLIKE.
        } else {
            SafeERC20.safeTransfer(IERC20($.likecoin), account, rewardClaimed);
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
        if ($.currentConditionIndex == 0) {
            $.lastRewardTime = startTime;
        }
        SafeERC20.safeTransferFrom(
            IERC20($.likecoin),
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
