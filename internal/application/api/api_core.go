package api

import (
	"github.com/Octy-ai/octy-shopify/internal/ports"
	c "github.com/Octy-ai/octy-shopify/pkg/config"
)

// Application implements the APIPort interface
type Application struct {
	rest    ports.RestPort
	db      ports.DatabasePort
	content Content
	config  c.Config
}

// NewApplication creates a new Application instances
func NewApplication(rest ports.RestPort, db ports.DatabasePort, content Content, config c.Config) *Application {
	return &Application{
		rest:    rest,
		db:      db,
		content: content,
		config:  config,
	}
}
