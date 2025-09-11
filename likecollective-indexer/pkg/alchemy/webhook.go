package alchemy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type CreateWebhookRequest struct {
	Network      string `json:"network"`
	WebhookType  string `json:"webhook_type"`
	WebhookUrl   string `json:"webhook_url"`
	GraphqlQuery string `json:"graphql_query"`
}

type CreateWebhookResponseData struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Network            string   `json:"network"`
	Networks           []string `json:"networks"`
	WebhookType        string   `json:"webhook_type"`
	WebhookUrl         string   `json:"webhook_url"`
	IsActive           bool     `json:"is_active"`
	TimeCreated        int64    `json:"time_created"`
	SigningKey         string   `json:"signing_key"`
	Version            string   `json:"version"`
	DeactivationReason string   `json:"deactivation_reason"`
}

type CreateWebhookResponse struct {
	Data *CreateWebhookResponseData `json:"data"`
	*ErrorResponse
}

var ErrWebhookCreationFailed = errors.New("webhook creation failed")

func (c *alchemyClient) CreateWebhook(requestData *CreateWebhookRequest) (*CreateWebhookResponseData, error) {
	url := fmt.Sprintf("%s/api/create-webhook", c.config.AlchemyBaseUrl)
	body, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Alchemy-Token", c.config.AlchemyAuthToken)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	responseData := new(CreateWebhookResponse)
	if err := decoder.Decode(responseData); err != nil {
		return nil, err
	}
	if responseData.ErrorResponse != nil {
		return nil, errors.Join(
			ErrWebhookCreationFailed,
			errors.New(responseData.ErrorResponse.Message),
		)
	}
	return responseData.Data, nil
}
