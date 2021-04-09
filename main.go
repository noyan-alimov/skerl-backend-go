package main

import (
	"github.com/noyan-alimov/skerl-backend/app"
	"github.com/noyan-alimov/skerl-backend/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8000")
}
