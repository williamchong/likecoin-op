package ent

import (
	"errors"
	"io"
	"os"
	"time"

	goyaml "gopkg.in/yaml.v2"
)

type BookNFT struct {
	EvmAddress      string    `json:"evmaddress"`
	StakedAmount    string    `json:"stakedamount"`
	LastStakedAt    time.Time `json:"laststakedat"`
	NumberOfStakers int       `json:"numberofstakers"`
}

type BookNFTClient struct {
}

func (c *BookNFTClient) Query() ([]*BookNFT, error) {
	f, err := os.Open("dbmockdata/book_nft.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := goyaml.NewDecoder(f)

	bookNFTs := make([]*BookNFT, 0)

	for {
		var bookNFT BookNFT
		if err := decoder.Decode(&bookNFT); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		bookNFTs = append(bookNFTs, &bookNFT)
	}

	return bookNFTs, nil
}
