package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Octy-ai/octy-shopify/internal/ports"
	"github.com/Octy-ai/octy-shopify/pkg/config"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

type Adapter struct {
	api    ports.APIPort
	config *config.Config
}

func NewAdapter(api ports.APIPort, config *config.Config) *Adapter {
	return &Adapter{api: api, config: config}
}

//Middlewares

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"%s\t%s",
			r.Method,
			r.RequestURI,
		)
		next.ServeHTTP(w, r)
	})
}

func responseManagerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

//Rest server definition and implementation

func registerHandlers(resta Adapter, r *mux.Router) {

	// Customer handlers
	r.HandleFunc("/api/customers/createupdate", createCustomerController(resta)).Methods("POST")

	// Event handlers
	r.HandleFunc("/api/events/create", createEventController(resta)).Methods("POST")
	r.HandleFunc("/api/hooks/events/charge/create", createChargedEventSHWHController(resta)).Methods("POST")
	/*
		REQUIRED WEBHOOKS:
			{
				"webhook": {
					"topic": "orders/paid",
					"address": "https://{host | ngrok_tunnel}/api/hooks/events/charge/create",
					"format": "json"
				}
			}
	*/

	// Item handlers
	r.HandleFunc("/api/hooks/items/create", createItemController(resta)).Methods("POST")
	r.HandleFunc("/api/hooks/items/update", updateItemController(resta)).Methods("POST")
	/*
		REQUIRED WEBHOOKS:
			{
				"webhook": {
					"topic": "products/create",
					"address": "https://{host | ngrok_tunnel}/api/hooks/items/create",
					"format": "json"
				}
			}
			{
				"webhook": {
					"topic": "products/update",
					"address": "https://{host | ngrok_tunnel}/api/hooks/items/update",
					"format": "json"
				}
			}
	*/

	// Content handlers
	r.HandleFunc("/api/content", getContentController(resta)).Methods("POST")

	// Recommendations handlers
	r.HandleFunc("/api/recommendations", getRecController(resta)).Methods("POST")
}

func (resta Adapter) Run() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	router.Use(responseManagerMiddleware)
	registerHandlers(resta, router)
	if resta.config.App.NgrokTunnel != "" {
		log.Printf("Server running on host %v and listenting on port %v >> Localhost forwarded to ngrok tunnel : %v . ctrl + c to quit!",
			resta.config.App.Host, resta.config.App.Port, resta.config.App.NgrokTunnel)
	} else {
		log.Printf("Server running on host %v and listenting on port %v . ctrl + c to quit!",
			resta.config.App.Host, resta.config.App.Port)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*",
		},
	})
	if err := http.ListenAndServe(":"+strconv.Itoa(resta.config.App.Port), corsOpts.Handler(router)); err != nil {
		log.Fatalf("failed to serve rest server over port %v: %v", resta.config.App.Port, err)
	}
}
