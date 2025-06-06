package cosmosnftidclassifier

type cosmosNFTIDMatcher interface {
	ExtractSerialID(cosmosNFTID string) (uint64, bool)
}
