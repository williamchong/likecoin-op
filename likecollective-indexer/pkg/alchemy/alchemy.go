package alchemy

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type AlchemyClient interface {
	CreateWebhook(request *CreateWebhookRequest) (*CreateWebhookResponseData, error)
}

type alchemyClient struct {
	config *AlchemyConfig
}

func NewAlchemyClient(
	config *AlchemyConfig,
) AlchemyClient {
	return &alchemyClient{
		config,
	}
}
