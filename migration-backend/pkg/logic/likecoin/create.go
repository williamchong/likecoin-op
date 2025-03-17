package likecoin

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/model"
)

var ErrLikeCoinMigrationInProgress = fmt.Errorf("err like coin migration in progress")
var ErrEvmSignatureNotVerified = fmt.Errorf("err evm signature not verified")

func CreateIfAllEnded(
	db *sql.DB,
	mintingEthAddress string,
	userEthAddress string,
	evmSignature string,
	evmSignatureMessage string,
	userCosmosAddress string,
	burningCosmosAddress string,
	amount string,
) (*model.LikeCoinMigration, error) {
	m, err := appdb.QueryNonEndedLikeCoinMigration(db, userCosmosAddress)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// continue
		} else {
			return nil, err
		}
	}

	if m != nil {
		return nil, ErrLikeCoinMigrationInProgress
	}

	recoveredAddr, err := ethereum.RecoverAddress(evmSignature, []byte(evmSignatureMessage))

	if err != nil {
		return nil, err
	}

	if !strings.EqualFold(recoveredAddr.Hex(), userEthAddress) {
		return nil, ErrEvmSignatureNotVerified
	}

	m, err = appdb.InsertLikeCoinMigration(db, &model.LikeCoinMigration{
		MintingEthAddress:    mintingEthAddress,
		UserEthAddress:       userEthAddress,
		EvmSignature:         evmSignature,
		EvmSignatureMessage:  evmSignatureMessage,
		UserCosmosAddress:    userCosmosAddress,
		BurningCosmosAddress: burningCosmosAddress,
		Amount:               amount,
		Status:               model.LikeCoinMigrationStatusPendingCosmosTxHash,
	})

	if err != nil {
		return nil, err
	}

	return m, nil
}
