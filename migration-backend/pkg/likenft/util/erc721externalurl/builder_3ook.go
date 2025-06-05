package erc721externalurl

import (
	"fmt"
	"net/url"
	"strconv"
)

type erc721ExternalURLBuilder3ook struct {
	baseUrl *url.URL
}

// The url is expected in format https://sepolia.3ook.com/store/${classId}/${tokenId}
func MakeErc721ExternalURLBuilder3ook(baseUrlStr string) (ERC721ExternalURLBuilder, error) {
	baseUrl, err := url.Parse(baseUrlStr)
	if err != nil {
		return nil, err
	}
	return &erc721ExternalURLBuilder3ook{
		baseUrl,
	}, nil
}

func (b *erc721ExternalURLBuilder3ook) Build(classId string, tokenId uint64) string {
	url := b.baseUrl.JoinPath(fmt.Sprintf("%s/%s", classId, strconv.FormatUint(tokenId, 10)))
	return url.String()
}
