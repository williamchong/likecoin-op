package cosmosnftidclassifier

import "slices"

type SerialNFTIDs []uint64

func (s SerialNFTIDs) Max() uint64 {
	return slices.Max(s)
}
