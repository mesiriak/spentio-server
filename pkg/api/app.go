package api

import (
	"log"
	"net/http"
	"spentio-server/configs"
	"spentio-server/pkg/api/rest"
)

type App struct {
	Http      *http.ServeMux
	AppConfig *configs.Config
}

func NewApp(config *configs.Config) App {
	httpApp := rest.NewHttpApp(config)

	return App{
		Http:      httpApp,
		AppConfig: config,
	}
}

func (a *App) RunApp() {
	// Now runs only HTTP application
	// We can use contexts to run multiple servers at same moment

	err := http.ListenAndServe(a.AppConfig.AppAddress, a.Http)

	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
