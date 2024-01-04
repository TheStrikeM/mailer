package main

import (
	"mailer/internal/app"
	"mailer/internal/config"
	configManager "mailer/pkg/config"
)

func main() {
	cfg := configManager.MustLoad[config.Config]()
	app.Run(cfg)
}
