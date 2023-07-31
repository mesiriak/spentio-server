package api

import (
	"spentio-server/configs"
	"spentio-server/pkg/api/http"
)

type App struct {
	Http      any
	Grpc      any
	Websocket any
	AppConfig *configs.Config
}

func NewApp(config *configs.Config) App {
	httpApp := http.NewHttpApp(config)

	return App{
		Http:      &httpApp,
		Grpc:      nil,
		Websocket: nil,
		AppConfig: config,
	}
}

func (a *App) RunApp() {
	// Now runs only HTTP application
}
