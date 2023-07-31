package rest

import (
	"fmt"
	"log"
	"net/http"
	"spentio-server/configs"
	"spentio-server/pkg/api/rest/handlers"
)

func NewHttpApp(config *configs.Config) *http.ServeMux {
	mux := http.NewServeMux()

	if err := registerHttpHandlers(mux); err != nil {
		log.Fatalf("Failed to register default handlers: %s", err)
	}

	return mux
}

func registerHttpHandlers(mux *http.ServeMux) error {
	/* Register all handlers here. */

	if err := registerDefaultHandlers(mux); err != nil {
		log.Fatalf("Failed to register default handlers: %s", err)
	}

	return nil
}

func registerDefaultHandlers(mux *http.ServeMux) error {
	prefix := "/"

	mux.HandleFunc(fmt.Sprintf("%s", prefix), handlers.RootHandler)

	return nil
}
