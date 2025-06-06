package cosmosnftidclassifier

type result struct {
	SerialNFTIDs    SerialNFTIDs
	ArbitraryNFTIDs []string
}

func (r *result) AllSerial() (SerialNFTIDs, bool) {
	if len(r.SerialNFTIDs) > 0 && len(r.ArbitraryNFTIDs) == 0 {
		return r.SerialNFTIDs, true
	}
	return nil, false
}

func (r *result) AllArbitrary() ([]string, bool) {
	if len(r.SerialNFTIDs) == 0 && len(r.ArbitraryNFTIDs) > 0 {
		return r.ArbitraryNFTIDs, true
	}
	return nil, false
}
