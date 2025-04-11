package evm_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"
	goyaml "gopkg.in/yaml.v2"

	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
)

func TestContractLevelMetadataFromCosmosClassAndISCN(t *testing.T) {
	Convey("ContractLevelMetadataFromCosmosClassAndISCN", t, func() {
		rootDir := "testdata/contract_level_metadata_from_cosmos_class_and_iscn/"
		entries, err := os.ReadDir(rootDir)
		if err != nil {
			t.Fatal(err)
		}
		for _, e := range entries {
			fullPath := path.Join(rootDir, e.Name())
			f, err := os.Open(fullPath)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			type TestCase struct {
				Name                  string `json:"name"`
				CosmosClassResponse   string `json:"cosmosclassresponse"`
				ISCNDataResponse      string `json:"iscndataresponse"`
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

				Convey(fmt.Sprintf("%s/%s", fullPath, testCase.Name), func() {
					var cosmosClass struct {
						Class *cosmosmodel.Class `json:"class"`
					}
					var iscn = cosmosmodel.ISCN{}
					err := json.Unmarshal([]byte(testCase.CosmosClassResponse), &cosmosClass)
					if err != nil {
						panic(err)
					}
					err = json.Unmarshal([]byte(testCase.ISCNDataResponse), &iscn)
					if err != nil {
						panic(err)
					}

					contractLevelMetadata := evm.ContractLevelMetadataFromCosmosClassAndISCN(cosmosClass.Class, &iscn)
					contractLevelMetadataStr, err := json.Marshal(contractLevelMetadata)
					if err != nil {
						panic(err)
					}
					require.JSONEq(t, testCase.ContractLevelMetadata, string(contractLevelMetadataStr))
				})
			}
		}
	})
}

func TestERC721MetadataFromCosmosNFTAndClassAndISCNData(t *testing.T) {
	Convey("ERC721MetadataFromCosmosNFTAndClassAndISCNData", t, func() {
		rootDir := "testdata/erc721_metadata_from_cosmos_nft_and_class_and_iscn/"
		entries, err := os.ReadDir(rootDir)
		if err != nil {
			t.Fatal(err)
		}
		for _, e := range entries {
			fullPath := path.Join(rootDir, e.Name())
			f, err := os.Open(fullPath)
			if err != nil {
				panic(err)
			}
			defer f.Close()

			type TestCase struct {
				Name                string  `json:"name"`
				CosmosNFT           string  `json:"cosmosnft"`
				CosmosClassResponse string  `json:"cosmosclassresponse"`
				ISCNDataResponse    string  `json:"iscndataresponse"`
				MetadataOverride    *string `json:"metadataoverride"`
				ERC721Metadata      string  `json:"erc721metadata"`
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
					var cosmosClass struct {
						Class *cosmosmodel.Class `json:"class"`
					}
					var iscn = cosmosmodel.ISCN{}
					var metadataOverride *cosmosmodel.NFTMetadata
					err := json.Unmarshal([]byte(testCase.CosmosClassResponse), &cosmosClass)
					if err != nil {
						panic(err)
					}
					err = json.Unmarshal([]byte(testCase.ISCNDataResponse), &iscn)
					if err != nil {
						panic(err)
					}
					err = json.Unmarshal([]byte(testCase.CosmosNFT), &cosmosNFT)
					if err != nil {
						panic(err)
					}
					if testCase.MetadataOverride != nil {
						err = json.Unmarshal([]byte(*testCase.MetadataOverride), &metadataOverride)
						if err != nil {
							panic(err)
						}
					}

					contractLevelMetadata := evm.ERC721MetadataFromCosmosNFTAndClassAndISCNData(
						&cosmosNFT,
						cosmosClass.Class,
						&iscn,
						metadataOverride,
					)
					contractLevelMetadataStr, err := json.Marshal(contractLevelMetadata)
					if err != nil {
						panic(err)
					}
					require.JSONEq(t, testCase.ERC721Metadata, string(contractLevelMetadataStr))
				})
			}
		}
	})
}
