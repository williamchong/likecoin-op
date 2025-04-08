package model

import (
	likenftmodel "github.com/likecoin/like-migration-backend/pkg/likenft/model"
)

type ContractLevelMetadataAttributes struct {
	TraitType string `json:"trait_type,omitempty"`
	Value     any    `json:"value,omitempty"`
}

type ContractLevelMetadataLikeCoin struct {
	ISCNIdPrefix string `json:"iscnIdPrefix,omitempty"`
	ClassId      string `json:"classId,omitempty"`
}

type ContractLevelMetadataISCN struct {
	// Additional metadata from cosmos
	Message                      string `json:"message,omitempty"`
	NftMetaCollectionId          string `json:"nft_meta_collection_id,omitempty"`
	NftMetaCollectionName        string `json:"nft_meta_collection_name,omitempty"`
	NftMetaCollectionDescrption  string `json:"nft_meta_collection_descrption,omitempty"`
	NftMetaCollectionDescription string `json:"nft_meta_collection_description,omitempty"`

	// Content metadata
	Context             string                            `json:"@context,omitempty"`
	Type                string                            `json:"@type,omitempty"`
	Author              *likenftmodel.Author              `json:"author,omitempty"`
	InLanguage          string                            `json:"inLanguage,omitempty"`
	ISBN                string                            `json:"isbn,omitempty"`
	Keywords            string                            `json:"keywords,omitempty"`
	Publisher           string                            `json:"publisher,omitempty"`
	SameAs              []string                          `json:"sameAs,omitempty"`
	UsageInfo           string                            `json:"usageInfo,omitempty"`
	Version             any                               `json:"version,omitempty"`
	ThumbnailUrl        string                            `json:"thumbnailUrl,omitempty"`
	Url                 string                            `json:"url,omitempty"`
	RecordNotes         string                            `json:"recordNotes,omitempty"`
	DateCreated         string                            `json:"dateCreated,omitempty"`
	DatePublished       string                            `json:"datePublished,omitempty"`
	ExifInfo            []any                             `json:"exifInfo,omitempty"`
	ContentFingerprints []string                          `json:"contentFingerprints,omitempty"`
	Attributes          []ContractLevelMetadataAttributes `json:"attributes"`
	LikeCoin            *ContractLevelMetadataLikeCoin    `json:"likecoin,omitempty"`
}

// https://eips.ethereum.org/EIPS/eip-7572
type ContractLevelMetadata struct {
	Name          string   `json:"name,omitempty"`
	Symbol        string   `json:"symbol,omitempty"`
	Description   string   `json:"description,omitempty"`
	Image         string   `json:"image,omitempty"`
	BannerImage   string   `json:"banner_image,omitempty"`
	FeaturedImage string   `json:"featured_image,omitempty"`
	ExternalLink  string   `json:"external_link,omitempty"`
	Collaborators []string `json:"collaborators,omitempty"`

	ContractLevelMetadataISCN
}
