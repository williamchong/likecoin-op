package alchemy

type AlchemyConfig struct {
	AlchemyBaseUrl   string `envconfig:"ALCHEMY_BASE_URL" default:""`
	AlchemyAuthToken string `envconfig:"ALCHEMY_AUTH_TOKEN" default:""`
}
