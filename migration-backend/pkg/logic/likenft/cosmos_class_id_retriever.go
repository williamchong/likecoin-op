package likenft

import (
	"database/sql"

	appdb "github.com/likecoin/like-migration-backend/pkg/db"
)

type CosmosClassIDRetriever interface {
	GetByEvmClassId(evmClassId string) (string, error)
}

type dbCosmosClassIDRetriever struct {
	db *sql.DB
}

func MakeDBCosmosClassIDRetriever(db *sql.DB) CosmosClassIDRetriever {
	return &dbCosmosClassIDRetriever{
		db: db,
	}
}

func (r *dbCosmosClassIDRetriever) GetByEvmClassId(evmClassId string) (string, error) {
	newClassAction, err := appdb.QueryLikeNFTMigrationActionNewClass(r.db, appdb.QueryLikeNFTMigrationActionNewClassFilter{
		EvmClassId: &evmClassId,
	})
	if err != nil {
		return "", err
	}
	return newClassAction.CosmosClassId, nil
}
