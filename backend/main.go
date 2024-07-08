package main

import (
	"worki/internal/app"
	"worki/internal/config"
)

// @title Worki messenger API
// @version 0.1
// @description This is API for Worki Messanger.

// @contact.name Worki Support
// @contact.email support@worki.io

// @host localhost:1323
func main() {
	cfg := config.Config{}
	cfg.Load()
	app.RunApp(&cfg)
}
