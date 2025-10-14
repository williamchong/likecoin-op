package evmeventprocessor

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"slices"

	"likenft-indexer/ent"
	"likenft-indexer/internal/database"
	"likenft-indexer/internal/evm"
	"likenft-indexer/internal/evm/book_nft"
	"likenft-indexer/internal/evm/util/logconverter"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	minterRole  = crypto.Keccak256([]byte("MINTER_ROLE"))
	updaterRole = crypto.Keccak256([]byte("UPDATER_ROLE"))
)

type roleChangedProcessor struct {
	httpClient         *http.Client
	evmClient          *evm.EvmClient
	nftClassRepository database.NFTClassRepository
}

func MakeRoleChangedProcessor(
	httpClient *http.Client,
	evmClient *evm.EvmClient,
	nftClassRepository database.NFTClassRepository,
) *roleChangedProcessor {
	return &roleChangedProcessor{
		httpClient:         httpClient,
		evmClient:          evmClient,
		nftClassRepository: nftClassRepository,
	}
}

func (e *roleChangedProcessor) Process(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	switch evmEvent.Topic0 {
	case "RoleGranted":
		return e.handleRoleGranted(ctx, logger, evmEvent)
	case "RoleRevoked":
		return e.handleRoleRevoked(ctx, logger, evmEvent)
	}
	return errors.Join(UnknownEvent, fmt.Errorf("%s", evmEvent.Topic0))
}

func (e *roleChangedProcessor) handleRoleGranted(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	bookNFTLogConverter := logconverter.NewLogConverter(e.evmClient.BookNFTABI)
	bookNFTLog := bookNFTLogConverter.ConvertEvmEventToLog(evmEvent)

	roleGrantedEvent := new(book_nft.BookNftRoleGranted)
	if err := bookNFTLogConverter.UnpackLog(bookNFTLog, roleGrantedEvent); err == nil {
		if bytes.Equal(roleGrantedEvent.Role[:], minterRole) {
			nftClass, err := e.nftClassRepository.QueryNFTClassByAddress(ctx, evmEvent.Address)
			if err != nil {
				logger.Error("e.nftClassRepository.QueryNFTClassByAddress", "err", err)
				return err
			}
			return e.handleBookNFTMinterRoleGranted(ctx, nftClass, roleGrantedEvent.Account)
		}
		if bytes.Equal(roleGrantedEvent.Role[:], updaterRole) {
			nftClass, err := e.nftClassRepository.QueryNFTClassByAddress(ctx, evmEvent.Address)
			if err != nil {
				logger.Error("e.nftClassRepository.QueryNFTClassByAddress", "err", err)
				return err
			}
			return e.handleBookNFTUpdaterRoleGranted(ctx, nftClass, roleGrantedEvent.Account)
		}
	}

	return errors.Join(UnknownEvent, fmt.Errorf("no candidate to unpack log"))
}

func (e *roleChangedProcessor) handleRoleRevoked(
	ctx context.Context,
	logger *slog.Logger,

	evmEvent *ent.EVMEvent,
) error {
	bookNFTLogConverter := logconverter.NewLogConverter(e.evmClient.BookNFTABI)
	bookNFTLog := bookNFTLogConverter.ConvertEvmEventToLog(evmEvent)

	roleRevokedEvent := new(book_nft.BookNftRoleRevoked)
	if err := bookNFTLogConverter.UnpackLog(bookNFTLog, roleRevokedEvent); err == nil {
		if bytes.Equal(roleRevokedEvent.Role[:], minterRole) {
			nftClass, err := e.nftClassRepository.QueryNFTClassByAddress(ctx, evmEvent.Address)
			if err != nil {
				logger.Error("e.nftClassRepository.QueryNFTClassByAddress", "err", err)
				return err
			}
			return e.handleBookNFTMinterRoleRevoked(ctx, nftClass, roleRevokedEvent.Account)
		}
		if bytes.Equal(roleRevokedEvent.Role[:], updaterRole) {
			nftClass, err := e.nftClassRepository.QueryNFTClassByAddress(ctx, evmEvent.Address)
			if err != nil {
				logger.Error("e.nftClassRepository.QueryNFTClassByAddress", "err", err)
				return err
			}
			return e.handleBookNFTUpdaterRoleRevoked(ctx, nftClass, roleRevokedEvent.Account)
		}
	}

	return errors.Join(UnknownEvent, fmt.Errorf("no candidate to unpack log"))
}

func (e *roleChangedProcessor) handleBookNFTMinterRoleGranted(
	ctx context.Context,

	nftClass *ent.NFTClass,
	grantedAccount common.Address,
) error {
	minterAddresses := nftClass.MinterAddresses

	if minterAddresses == nil {
		minterAddresses = []string{}
	}

	if slices.Contains(minterAddresses, grantedAccount.Hex()) {
		return nil
	}

	minterAddresses = append(minterAddresses, grantedAccount.Hex())
	return e.nftClassRepository.UpdateMinterAddresses(ctx, nftClass.Address, minterAddresses)
}

func (e *roleChangedProcessor) handleBookNFTUpdaterRoleGranted(
	ctx context.Context,

	nftClass *ent.NFTClass,
	grantedAccount common.Address,
) error {
	updaterAddresses := nftClass.UpdaterAddresses

	if updaterAddresses == nil {
		updaterAddresses = []string{}
	}

	if slices.Contains(updaterAddresses, grantedAccount.Hex()) {
		return nil
	}

	updaterAddresses = append(updaterAddresses, grantedAccount.Hex())
	return e.nftClassRepository.UpdateUpdaterAddresses(ctx, nftClass.Address, updaterAddresses)
}

func (e *roleChangedProcessor) handleBookNFTMinterRoleRevoked(
	ctx context.Context,

	nftClass *ent.NFTClass,
	revokedAccount common.Address,
) error {
	minterAddresses := nftClass.MinterAddresses
	if minterAddresses == nil {
		minterAddresses = []string{}
	}

	if !slices.Contains(minterAddresses, revokedAccount.Hex()) {
		return nil
	}

	minterAddresses = slices.Delete(
		minterAddresses,
		slices.Index(minterAddresses, revokedAccount.Hex()),
		1,
	)

	return e.nftClassRepository.UpdateMinterAddresses(ctx, nftClass.Address, minterAddresses)
}

func (e *roleChangedProcessor) handleBookNFTUpdaterRoleRevoked(
	ctx context.Context,

	nftClass *ent.NFTClass,
	revokedAccount common.Address,
) error {
	updaterAddresses := nftClass.UpdaterAddresses
	if updaterAddresses == nil {
		updaterAddresses = []string{}
	}

	if !slices.Contains(updaterAddresses, revokedAccount.Hex()) {
		return nil
	}

	updaterAddresses = slices.Delete(
		updaterAddresses,
		slices.Index(updaterAddresses, revokedAccount.Hex()),
		1,
	)

	return e.nftClassRepository.UpdateUpdaterAddresses(ctx, nftClass.Address, updaterAddresses)
}

func init() {
	registerEventProcessor(
		"RoleGranted",
		func(inj *eventProcessorDeps) eventProcessor {
			return MakeRoleChangedProcessor(
				inj.httpClient,
				inj.evmClient,
				inj.nftClassRepository,
			)
		},
	)
	registerEventProcessor(
		"RoleRevoked",
		func(inj *eventProcessorDeps) eventProcessor {
			return MakeRoleChangedProcessor(
				inj.httpClient,
				inj.evmClient,
				inj.nftClassRepository,
			)
		},
	)
}
