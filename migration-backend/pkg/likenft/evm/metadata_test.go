package evm_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"testing"

	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/model"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"
	goyaml "gopkg.in/yaml.v2"
)

func TestContractLevelMetadataFromCosmosClass(t *testing.T) {
	Convey("ContractLevelMetadataFromCosmosClass", t, func() {
		rootDir := "testdata/contract_level_metadata_from_cosmos_class/"
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
		}
	})
}

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
					var iscn = model.ISCN{}
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

func TestContractLevelMetadataFromCosmosClassListItem(t *testing.T) {
	Convey("ContractLevelMetadataFromCosmosClassListItem", t, func() {
		rootDir := "testdata/contract_level_metadata_from_cosmos_class_list_item/"
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
		}
	})
}

func TestERC721MetadataFromCosmosNFT(t *testing.T) {
	Convey("ERC721MetadataFromCosmosNFT", t, func() {
		rootDir := "testdata/erc721_metadata_from_cosmos_nft/"
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
		}
	})
}
