package rest

import (
	"net/http"
	"spentio-server/configs"
	"spentio-server/pkg/api/rest/handlers"
)

func NewHttpApp(config *configs.Config) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)

	return mux
}
