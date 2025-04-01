package model

type ContractLevelMetadata struct {
	Name          string   `json:"name,omitempty"`
	Symbol        string   `json:"symbol,omitempty"`
	Description   string   `json:"description,omitempty"`
	Image         string   `json:"image,omitempty"`
	BannerImage   string   `json:"banner_image,omitempty"`
	FeaturedImage string   `json:"featured_image,omitempty"`
	ExternalLink  string   `json:"external_link,omitempty"`
	Collaborators []string `json:"collaborators,omitempty"`
}

type ERC721Metadata struct {
	Image           string                    `json:"image,omitempty"`
	ImageData       string                    `json:"image_data,omitempty"`
	ExternalUrl     string                    `json:"external_url,omitempty"`
	Description     string                    `json:"description,omitempty"`
	Name            string                    `json:"name,omitempty"`
	Attributes      []ERC721MetadataAttribute `json:"attributes,omitempty"`
	BackgroundColor string                    `json:"background_color,omitempty"`
	AnimationUrl    string                    `json:"animation_url,omitempty"`
	YoutubeUrl      string                    `json:"youtube_url,omitempty"`
}

type ERC721MetadataAttributeDisplayType string

var (
	ERC721MetadataAttributeDisplayTypeNumber          ERC721MetadataAttributeDisplayType = "number"
	ERC721MetadataAttributeDisplayTypeBoostNumber     ERC721MetadataAttributeDisplayType = "boost_number"
	ERC721MetadataAttributeDisplayTypeBoostPercentage ERC721MetadataAttributeDisplayType = "boost_percentage"
)

type ERC721MetadataAttribute struct {
	DisplayType *ERC721MetadataAttributeDisplayType `json:"display_type,omitempty"`
	TraitType   string                              `json:"trait_type,omitempty"`
	Value       interface{}                         `json:"value,omitempty"`
}
