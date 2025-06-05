package model

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type ContractLevelMetadataOpenSea struct {
	Name          string   `json:"name,omitempty" mapstructure:"name"`
	Symbol        string   `json:"symbol,omitempty"  mapstructure:"symbol"`
	Description   string   `json:"description,omitempty"  mapstructure:"description"`
	Image         string   `json:"image,omitempty"  mapstructure:"image"`
	BannerImage   string   `json:"banner_image,omitempty"  mapstructure:"banner_image"`
	FeaturedImage string   `json:"featured_image,omitempty"  mapstructure:"featured_image"`
	ExternalLink  string   `json:"external_link,omitempty"  mapstructure:"external_link"`
	Collaborators []string `json:"collaborators,omitempty"  mapstructure:"collaborators"`
}

type AdditionalProps map[string]any

type contractLevelMetadataMapStructure struct {
	ContractLevelMetadataOpenSea `mapstructure:",squash"`
	AdditionalProps              AdditionalProps `mapstructure:",remain"`
}

type ContractLevelMetadata struct {
	ContractLevelMetadataOpenSea
	AdditionalProps
}

func (m *ContractLevelMetadata) UnmarshalJSON(data []byte) error {
	_m := make(map[string]any)
	err := json.Unmarshal(data, &_m)
	if err != nil {
		return err
	}

	c := new(contractLevelMetadataMapStructure)
	err = mapstructure.Decode(_m, c)
	if err != nil {
		return err
	}

	m.ContractLevelMetadataOpenSea = c.ContractLevelMetadataOpenSea
	m.AdditionalProps = c.AdditionalProps
	return nil
}

func (m *ContractLevelMetadata) MarshalJSON() ([]byte, error) {
	openSeaBytes, err := json.Marshal(m.ContractLevelMetadataOpenSea)
	if err != nil {
		return nil, err
	}

	res := make(map[string]any)
	err = json.Unmarshal(openSeaBytes, &res)
	if err != nil {
		return nil, err
	}

	// Unmarshalling null will make the target dict empty
	// refs: https://go.dev/play/p/nxgfstsnRiL
	if m.AdditionalProps != nil {
		additionalPropsBytes, err := json.Marshal(m.AdditionalProps)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(additionalPropsBytes, &res)
		if err != nil {
			return nil, err
		}
	}

	return json.Marshal(res)
}

type ERC721MetadataOpenSea struct {
	Image           string                    `json:"image,omitempty" mapstructure:"image"`
	ImageData       string                    `json:"image_data,omitempty" mapstructure:"image_data"`
	ExternalUrl     string                    `json:"external_url,omitempty" mapstructure:"external_url"`
	Description     string                    `json:"description,omitempty" mapstructure:"description"`
	Name            string                    `json:"name,omitempty" mapstructure:"name"`
	Attributes      []ERC721MetadataAttribute `json:"attributes,omitempty" mapstructure:"attributes"`
	BackgroundColor string                    `json:"background_color,omitempty" mapstructure:"background_color"`
	AnimationUrl    string                    `json:"animation_url,omitempty" mapstructure:"animation_url"`
	YoutubeUrl      string                    `json:"youtube_url,omitempty" mapstructure:"youtube_url"`
}

type erc721MetadataMapStructure struct {
	ERC721MetadataOpenSea `mapstructure:",squash"`
	AdditionalProps       AdditionalProps `mapstructure:",remain"`
}

type ERC721Metadata struct {
	ERC721MetadataOpenSea
	AdditionalProps
}

func (m *ERC721Metadata) UnmarshalJSON(data []byte) error {
	_m := make(map[string]any)
	err := json.Unmarshal(data, &_m)
	if err != nil {
		return err
	}

	c := new(erc721MetadataMapStructure)
	err = mapstructure.Decode(_m, c)
	if err != nil {
		return err
	}

	m.ERC721MetadataOpenSea = c.ERC721MetadataOpenSea
	m.AdditionalProps = c.AdditionalProps
	return nil
}

func (m *ERC721Metadata) MarshalJSON() ([]byte, error) {
	openSeaBytes, err := json.Marshal(m.ERC721MetadataOpenSea)
	if err != nil {
		return nil, err
	}

	res := make(map[string]any)
	err = json.Unmarshal(openSeaBytes, &res)
	if err != nil {
		return nil, err
	}

	// Unmarshalling null will make the target dict empty
	// refs: https://go.dev/play/p/nxgfstsnRiL
	if m.AdditionalProps != nil {
		additionalPropsBytes, err := json.Marshal(m.AdditionalProps)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(additionalPropsBytes, &res)
		if err != nil {
			return nil, err
		}
	}

	return json.Marshal(res)
}

type ERC721MetadataAttributeDisplayType string

var (
	ERC721MetadataAttributeDisplayTypeNumber          ERC721MetadataAttributeDisplayType = "number"
	ERC721MetadataAttributeDisplayTypeBoostNumber     ERC721MetadataAttributeDisplayType = "boost_number"
	ERC721MetadataAttributeDisplayTypeBoostPercentage ERC721MetadataAttributeDisplayType = "boost_percentage"
)

type ERC721MetadataAttribute struct {
	DisplayType *ERC721MetadataAttributeDisplayType `json:"display_type,omitempty" mapstructure:"display_type"`
	TraitType   string                              `json:"trait_type,omitempty" mapstructure:"trait_type"`
	Value       interface{}                         `json:"value,omitempty" mapstructure:"value"`
}
