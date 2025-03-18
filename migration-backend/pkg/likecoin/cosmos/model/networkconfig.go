package model

import "encoding/json"

type CoinLookup struct {
	ViewDenom                   string `json:"viewDenom"`
	ChainDenom                  string `json:"chainDenom"`
	ChainToViewConversionFactor string `json:"chainToViewConversionFactor"`
}

type NetworkConfig struct {
	CoinLookup []CoinLookup `json:"coinLookup"`
}

func LoadNetworkConfig(jsonData []byte) (*NetworkConfig, error) {
	var data NetworkConfig
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
