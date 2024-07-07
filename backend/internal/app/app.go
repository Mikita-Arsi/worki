package app

import (
	"worki/internal/config"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RunApp(cfg *config.Config) {
	// init DB
	// init services (db) в себя включают DB
	storage.InitDB(cfg)
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// init controllers в себя включают сервисы
