package erc721externalurl

type ERC721ExternalURLBuilder interface {
	Build(classId string, tokenId uint64) string
}
