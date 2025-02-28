package evm_test

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"testing"

	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"
	goyaml "gopkg.in/yaml.v2"
)

func TestContractLevelMetadataFromCosmosClass(t *testing.T) {
	Convey("ContractLevelMetadataFromCosmosClass", t, func() {
		f, err := os.Open("testdata/contract_level_metadata_from_cosmos_class.tests.yaml")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		type TestCase struct {
			Name                  string `json:"name"`
			CosmosClassResponse   string `json:"cosmosclassresponse"`
			ContractLevelMetadata string `json:"contractlevelmetadata"`
		}

		decoder := goyaml.NewDecoder(f)

		for {
			var testCase TestCase
			err := decoder.Decode(&testCase)
			if errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				panic(err)
			}

			Convey(testCase.Name, func() {
				var cosmosClass struct {
					Class *cosmosmodel.Class `json:"class"`
				}
				err := json.Unmarshal([]byte(testCase.CosmosClassResponse), &cosmosClass)
				if err != nil {
					panic(err)
				}

				contractLevelMetadata := evm.ContractLevelMetadataFromCosmosClass(cosmosClass.Class)
				contractLevelMetadataStr, err := json.Marshal(contractLevelMetadata)
				if err != nil {
					panic(err)
				}
				require.JSONEq(t, testCase.ContractLevelMetadata, string(contractLevelMetadataStr))
			})
		}
	})
}

func TestContractLevelMetadataFromCosmosClassListItem(t *testing.T) {
	Convey("ContractLevelMetadataFromCosmosClassListItem", t, func() {
		f, err := os.Open("testdata/contract_level_metadata_from_cosmos_class_list_item.tests.yaml")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		type TestCase struct {
			Name                  string `json:"name"`
			CosmosClass           string `json:"cosmosclass"`
			ContractLevelMetadata string `json:"contractlevelmetadata"`
		}

		decoder := goyaml.NewDecoder(f)

		for {
			var testCase TestCase
			err := decoder.Decode(&testCase)
			if errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				panic(err)
			}

			Convey(testCase.Name, func() {
				var cosmosClass cosmosmodel.ClassListItem
				err := json.Unmarshal([]byte(testCase.CosmosClass), &cosmosClass)
				if err != nil {
					panic(err)
				}

				contractLevelMetadata := evm.ContractLevelMetadataFromCosmosClassListItem(&cosmosClass)
				contractLevelMetadataStr, err := json.Marshal(contractLevelMetadata)
				if err != nil {
					panic(err)
				}
				require.JSONEq(t, testCase.ContractLevelMetadata, string(contractLevelMetadataStr))
			})
		}
	})
}

func TestERC721MetadataFromCosmosNFT(t *testing.T) {
	Convey("ERC721MetadataFromCosmosNFT", t, func() {
		f, err := os.Open("testdata/erc721_metadata_from_cosmos_nft.tests.yaml")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		type TestCase struct {
			Name           string `json:"name"`
			CosmosNFT      string `json:"cosmosnft"`
			ERC721Metadata string `json:"erc721metadata"`
		}

		decoder := goyaml.NewDecoder(f)

		for {
			var testCase TestCase
			err := decoder.Decode(&testCase)
			if errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				panic(err)
			}

			Convey(testCase.Name, func() {
				var cosmosNFT cosmosmodel.NFT
				err := json.Unmarshal([]byte(testCase.CosmosNFT), &cosmosNFT)
				if err != nil {
					panic(err)
				}

				contractLevelMetadata := evm.ERC721MetadataFromCosmosNFT(&cosmosNFT)
				contractLevelMetadataStr, err := json.Marshal(contractLevelMetadata)
				if err != nil {
					panic(err)
				}
				require.JSONEq(t, testCase.ERC721Metadata, string(contractLevelMetadataStr))
			})
		}
	})
}
