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

	e.Use(middleware.Recover())
	usersGroup := e.Group("/users")

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	usersGroup.POST("/", api.CreateUser)
	usersGroup.GET("/", api.GetUsers)
	usersGroup.GET("/:id", api.GetUserByID)
	usersGroup.GET("/:username", api.GetUserByUsername)
	usersGroup.PUT("/", api.UpdateUser)
	usersGroup.DELETE("/:id", api.DeleteUserByID)
	usersGroup.DELETE("/:username", api.DeleteUserByUsername)
	usersGroup.Use(logger.LogRequest)

	e.Logger.Fatal(e.Start(":1323"))
}
