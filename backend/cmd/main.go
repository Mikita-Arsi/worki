package main

import (
	"worki/internal/app"
	"worki/internal/config"
)

// @title Worki messenger API
// @version 0.1
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:1323
// @BasePath /v2
func main() {
	cfg := config.Config{}
	app.RunApp(&cfg)
}
