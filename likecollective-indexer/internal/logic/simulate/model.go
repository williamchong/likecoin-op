package simulate

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/holiman/uint256"
)

type SimulationLogWithHeader struct {
	Log    types.Log     `json:"log"`
	Header *types.Header `json:"block"`
}

type Simulation struct {
	Logs  []SimulationLogWithHeader `json:"logs"`
	State State                     `json:"state"`
}

type State struct {
	BookPendingRewards map[common.Address]*uint256.Int                    `json:"bookPendingRewards"`
	BookStakedAmounts  map[common.Address]*uint256.Int                    `json:"bookStakedAmounts"`
	UserPendingRewards map[common.Address]map[common.Address]*uint256.Int `json:"userPendingRewards"`
	UserStakedAmounts  map[common.Address]map[common.Address]*uint256.Int `json:"userStakedAmounts"`
}

var (
	ErrBookPendingRewardsMismatch = errors.New("book pending rewards mismatch")
	ErrBookStakedAmountsMismatch  = errors.New("book staked amounts mismatch")
	ErrUserPendingRewardsMismatch = errors.New("user pending rewards mismatch")
	ErrUserStakedAmountsMismatch  = errors.New("user staked amounts mismatch")
)

func (a *State) Verify(b *State) error {
	// bookPendingRewards

	for bookNFTAddress, bookPendingRewardAmount := range a.BookPendingRewards {
		if bookPendingRewardAmount.Cmp(uint256.NewInt(0)) == 0 {
			continue
		}
		if _, ok := b.BookPendingRewards[bookNFTAddress]; !ok {
			return fmt.Errorf(
				"%w: %s: not found",
				ErrBookPendingRewardsMismatch,
				bookNFTAddress.Hex(),
			)
		}
		if bookPendingRewardAmount.Cmp(b.BookPendingRewards[bookNFTAddress]) != 0 {
			return fmt.Errorf(
				"%w: %s: %s != %s",
				ErrBookPendingRewardsMismatch,
				bookNFTAddress.Hex(),
				bookPendingRewardAmount.Hex(),
				b.BookPendingRewards[bookNFTAddress].Hex(),
			)
		}
	}

	for bookNFTAddress, bookPendingRewardAmount := range b.BookPendingRewards {
		if bookPendingRewardAmount.Cmp(uint256.NewInt(0)) == 0 {
			continue
		}
		if _, ok := a.BookPendingRewards[bookNFTAddress]; !ok {
			return fmt.Errorf(
				"%w: %s: not found",
				ErrBookPendingRewardsMismatch,
				bookNFTAddress.Hex(),
			)
		}
		if bookPendingRewardAmount.Cmp(a.BookPendingRewards[bookNFTAddress]) != 0 {
			return fmt.Errorf(
				"%w: %s: %s != %s",
				ErrBookPendingRewardsMismatch,
				bookNFTAddress.Hex(),
				bookPendingRewardAmount.Hex(),
				a.BookPendingRewards[bookNFTAddress].Hex(),
			)
		}
	}

	// bookStakedAmounts
	for bookNFTAddress, bookStakedAmount := range a.BookStakedAmounts {
		if bookStakedAmount.Cmp(uint256.NewInt(0)) == 0 {
			continue
		}
		if _, ok := b.BookStakedAmounts[bookNFTAddress]; !ok {
			return fmt.Errorf(
				"%w: %s: not found",
				ErrBookStakedAmountsMismatch,
				bookNFTAddress.Hex(),
			)
		}
		if bookStakedAmount.Cmp(b.BookStakedAmounts[bookNFTAddress]) != 0 {
			return fmt.Errorf(
				"%w: %s: %s != %s",
				ErrBookStakedAmountsMismatch,
				bookNFTAddress.Hex(),
				bookStakedAmount.Hex(),
				b.BookStakedAmounts[bookNFTAddress].Hex(),
			)
		}
	}

	for bookNFTAddress, bookStakedAmount := range b.BookStakedAmounts {
		if bookStakedAmount.Cmp(uint256.NewInt(0)) == 0 {
			continue
		}
		if _, ok := a.BookStakedAmounts[bookNFTAddress]; !ok {
			return fmt.Errorf(
				"%w: %s: not found",
				ErrBookStakedAmountsMismatch,
				bookNFTAddress.Hex(),
			)
		}
		if bookStakedAmount.Cmp(a.BookStakedAmounts[bookNFTAddress]) != 0 {
			return fmt.Errorf(
				"%w: %s: %s != %s",
				ErrBookStakedAmountsMismatch,
				bookNFTAddress.Hex(),
				bookStakedAmount.Hex(),
				a.BookStakedAmounts[bookNFTAddress].Hex(),
			)
		}
	}

	// userPendingRewards
	for userEVMAddress, userPendingRewards := range a.UserPendingRewards {
		if _, ok := b.UserPendingRewards[userEVMAddress]; !ok {
			return fmt.Errorf(
				"%w: %s: not found",
				ErrUserPendingRewardsMismatch,
				userEVMAddress.Hex(),
			)
		}
		for bookNFTAddress, userPendingRewardAmount := range userPendingRewards {
			if userPendingRewardAmount.Cmp(uint256.NewInt(0)) == 0 {
				continue
			}
			if _, ok := b.UserPendingRewards[userEVMAddress][bookNFTAddress]; !ok {
				return fmt.Errorf(
					"%w: %s: %s not found",
					ErrUserPendingRewardsMismatch,
					userEVMAddress.Hex(),
					bookNFTAddress.Hex(),
				)
			}
			if userPendingRewardAmount.Cmp(b.UserPendingRewards[userEVMAddress][bookNFTAddress]) != 0 {
				return fmt.Errorf(
					"%w: %s: %s != %s",
					ErrUserPendingRewardsMismatch,
					userEVMAddress.Hex(),
					userPendingRewardAmount.Hex(),
					b.UserPendingRewards[userEVMAddress][bookNFTAddress].Hex(),
				)
			}
		}
	}

	for userEVMAddress, userPendingRewards := range b.UserPendingRewards {
		if _, ok := a.UserPendingRewards[userEVMAddress]; !ok {
			return fmt.Errorf(
				"%w: %s: not found",
				ErrUserPendingRewardsMismatch,
				userEVMAddress.Hex(),
			)
		}
		for bookNFTAddress, userPendingRewardAmount := range userPendingRewards {
			if userPendingRewardAmount.Cmp(uint256.NewInt(0)) == 0 {
				continue
			}
			if _, ok := a.UserPendingRewards[userEVMAddress][bookNFTAddress]; !ok {
				return fmt.Errorf(
					"%w: %s: %s not found",
					ErrUserPendingRewardsMismatch,
					userEVMAddress.Hex(),
					bookNFTAddress.Hex(),
				)
			}
			if userPendingRewardAmount.Cmp(a.UserPendingRewards[userEVMAddress][bookNFTAddress]) != 0 {
				return fmt.Errorf(
					"%w: %s: %s != %s",
					ErrUserPendingRewardsMismatch,
					userEVMAddress.Hex(),
					userPendingRewardAmount.Hex(),
					a.UserPendingRewards[userEVMAddress][bookNFTAddress].Hex(),
				)
			}
		}
	}

	// userStakedAmounts
	for userEVMAddress, userStakedAmounts := range a.UserStakedAmounts {
		if _, ok := b.UserStakedAmounts[userEVMAddress]; !ok {
			return fmt.Errorf(
				"%w: %s: not found",
				ErrUserStakedAmountsMismatch,
				userEVMAddress.Hex(),
			)
		}
		for bookNFTAddress, userStakedAmount := range userStakedAmounts {
			if userStakedAmount.Cmp(uint256.NewInt(0)) == 0 {
				continue
			}
			if _, ok := b.UserStakedAmounts[userEVMAddress][bookNFTAddress]; !ok {
				return fmt.Errorf(
					"%w: %s: %s not found",
					ErrUserStakedAmountsMismatch,
					userEVMAddress.Hex(),
					bookNFTAddress.Hex(),
				)
			}
			if userStakedAmount.Cmp(b.UserStakedAmounts[userEVMAddress][bookNFTAddress]) != 0 {
				return fmt.Errorf(
					"%w: %s: %s != %s",
					ErrUserStakedAmountsMismatch,
					userEVMAddress.Hex(),
					userStakedAmount.Hex(),
					b.UserStakedAmounts[userEVMAddress][bookNFTAddress].Hex(),
				)
			}
		}
	}

	for userEVMAddress, userStakedAmounts := range b.UserStakedAmounts {
		if _, ok := a.UserStakedAmounts[userEVMAddress]; !ok {
			return fmt.Errorf(
				"%w: %s: not found",
				ErrUserStakedAmountsMismatch,
				userEVMAddress.Hex(),
			)
		}
		for bookNFTAddress, userStakedAmount := range userStakedAmounts {
			if userStakedAmount.Cmp(uint256.NewInt(0)) == 0 {
				continue
			}
			if _, ok := a.UserStakedAmounts[userEVMAddress][bookNFTAddress]; !ok {
				return fmt.Errorf(
					"%w: %s: %s not found",
					ErrUserStakedAmountsMismatch,
					userEVMAddress.Hex(),
					bookNFTAddress.Hex(),
				)
			}
			if userStakedAmount.Cmp(a.UserStakedAmounts[userEVMAddress][bookNFTAddress]) != 0 {
				return fmt.Errorf(
					"%w: %s: %s != %s",
					ErrUserStakedAmountsMismatch,
					userEVMAddress.Hex(),
					userStakedAmount.Hex(),
					a.UserStakedAmounts[userEVMAddress][bookNFTAddress].Hex(),
				)
			}
		}
	}

	return nil
}
