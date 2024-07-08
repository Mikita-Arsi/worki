package app

import (
	_ "worki/docs"
	"worki/internal/api"
	"worki/internal/config"
	"worki/internal/logger"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RunApp(cfg *config.Config) {
	storage.InitDB(cfg)
	e := echo.New()

	e.Use(logger.LogRequest)
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/users", api.CreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}
