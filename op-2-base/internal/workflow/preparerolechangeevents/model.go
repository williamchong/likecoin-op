package preparerolechangeevents

import (
	"fmt"
	"slices"

	"github.com/ethereum/go-ethereum/crypto"
)

type EventType string

const (
	EventTypeRoleGranted EventType = "RoleGranted"
	EventTypeRoleRevoked EventType = "RoleRevoked"
)

func (e EventType) String() string {
	return string(e)
}

type Output struct {
	BookNFTId            string    `json:"booknft"`
	EventType            EventType `json:"event"`
	RoleBytesArrayString string    `json:"role_byte_array_string"`
	To                   string    `json:"to"`
	By                   string    `json:"by"`
	BlockNumber          uint64    `json:"block_number"`
	TransactionIndex     uint64    `json:"transaction_index"`
	LogIndex             uint64    `json:"log_index"`
}

type Role string

const (
	MinterRole  Role = "MINTER_ROLE"
	UpdaterRole Role = "UPDATER_ROLE"
)

func ComputeGrantedAddresses(role Role, events []Output) []string {
	roleBytes := crypto.Keccak256([]byte(role))
	roleBytesArrayString := fmt.Sprintf("%v", roleBytes)

	orderedEvents := orderEvents(events)

	addresses := make([]string, 0)
	for _, event := range orderedEvents {
		switch event.EventType {
		case EventTypeRoleGranted:
			if event.RoleBytesArrayString == roleBytesArrayString {
				addresses = append(addresses, event.To)
			}
		case EventTypeRoleRevoked:
			if event.RoleBytesArrayString == roleBytesArrayString {
				addresses = slices.Delete(addresses, slices.Index(addresses, event.To), 1)
			}
		}
	}

	return addresses
}

func orderEvents(events []Output) []Output {
	orderedEvents := slices.Clone(events)
	compareUint64 := func(a, b uint64) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	slices.SortFunc(orderedEvents, func(a, b Output) int {
		if a.BlockNumber != b.BlockNumber {
			return compareUint64(a.BlockNumber, b.BlockNumber)
		}
		if a.TransactionIndex != b.TransactionIndex {
			return compareUint64(a.TransactionIndex, b.TransactionIndex)
		}
		return compareUint64(a.LogIndex, b.LogIndex)
	})
	return orderedEvents
}
