package rest

import (
	"net/http"

	"github.com/Octy-ai/octy-shopify/pkg/config"
)

type Adapter struct {
	httpClient *http.Client
	config     *config.Config
}

func NewAdapter(config *config.Config) (*Adapter, error) {
	httpClient := &http.Client{}
	return &Adapter{httpClient: httpClient, config: config}, nil
}
