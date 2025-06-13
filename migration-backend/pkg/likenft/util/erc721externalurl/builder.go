package erc721externalurl

type ERC721ExternalURLBuilder interface {
	BuildSerial(classId string, tokenId uint64) string
	BuildArbitrary(classId string) string
}
