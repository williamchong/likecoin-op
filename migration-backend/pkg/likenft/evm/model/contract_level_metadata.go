package model

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
}
