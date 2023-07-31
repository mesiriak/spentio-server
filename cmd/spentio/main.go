package main

import (
	"spentio-server/configs"
	"spentio-server/pkg/api"
)

func main() {
	config := configs.CreateConfig()

	app := api.NewApp(&config)

	app.RunApp()
}
