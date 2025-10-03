package newnftclass

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/prepareactions"
)

type Input struct {
	prepareactions.PrepareNewClassActionOutput
}

type bookConfig struct {
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	Metadata  string `json:"metadata"`
	MaxSupply uint64 `json:"max_supply"`
}

type msgNewBookNFT struct {
	Creator  string     `json:"creator"`
	Updaters []string   `json:"updaters"`
	Minters  []string   `json:"minters"`
	Config   bookConfig `json:"config"`
}

type Output struct {
	Input
	msgNewBookNFT
}

type PrepareNewNFTClassParam interface {
	Prepare(
		ctx context.Context,
		logger *slog.Logger,
		input *Input,
	) (*Output, error)
}

type prepareNewNFTClassParam struct {
	httpClient *http.Client
}

func NewPrepareNewNFTClassParam(
	httpClient *http.Client,
) PrepareNewNFTClassParam {
	return &prepareNewNFTClassParam{
		httpClient,
	}
}

func (p *prepareNewNFTClassParam) Prepare(
	ctx context.Context,
	logger *slog.Logger,
	input *Input,
) (*Output, error) {
	mylogger := logger.WithGroup("PrepareNewNFTClassParam").With("opAddress", input.OpAddress)
	mylogger.Info("PrepareNewNFTClassParam")

	symbol := input.Metadata.Symbol
	name := input.Metadata.Name

	metadataBytes, err := json.Marshal(input.Metadata)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %v", err)
	}
	metadataStr := string(metadataBytes)
	maxSupply := input.MaxSupply

	return &Output{
		Input: *input,
		msgNewBookNFT: msgNewBookNFT{
			Creator:  input.InitialOwner,
			Updaters: input.InitialUpdaters,
			Minters:  input.InitialMinters,
			Config: bookConfig{
				Name:      name,
				Symbol:    symbol,
				Metadata:  metadataStr,
				MaxSupply: maxSupply,
			},
		},
	}, nil
}
