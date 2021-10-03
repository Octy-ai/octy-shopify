package main

import (
	"log"

	"github.com/Octy-ai/octy-shopify/internal/adapters/framework/primary/rest"
	dbSec "github.com/Octy-ai/octy-shopify/internal/adapters/framework/secondary/database"
	restSec "github.com/Octy-ai/octy-shopify/internal/adapters/framework/secondary/rest"
	"github.com/Octy-ai/octy-shopify/internal/application/api"
	"github.com/Octy-ai/octy-shopify/internal/application/domain/content"
	"github.com/Octy-ai/octy-shopify/pkg/config"
)

func main() {

	conf, err := config.NewConfig("./pkg/config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load configurations: %v", err)
	}

	// init right side (secondary) adapters
	restDrivenAdapter, err := restSec.NewAdapter(conf)
	if err != nil {
		log.Fatalf("failed to initialize rest driven adapter: %v", err)
	}

	dbDrivenAdapter, err := dbSec.NewAdapter(conf)
	if err != nil {
		log.Fatalf("failed to initialize database driven adapter: %v", err)
	}
	dbDrivenAdapter.Connect()

	// init core domain logic services
	content := content.New()

	applicationAPI := api.NewApplication(restDrivenAdapter, dbDrivenAdapter, content, *conf)

	// init left side (primary) adapters
	restAdapter := rest.NewAdapter(applicationAPI, conf)
	// start rest server
	restAdapter.Run()
}
