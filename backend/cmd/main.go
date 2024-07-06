package main

import (
	"worki/internal/routes"

	"github.com/AlhimicMan/goswag/wrapper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router := wrapper.NewRouter(e)
	group := router.Group("/users", "Users")
	routes.RegisterRoutes(group)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	_, err := router.GenerateSwagger()
	if err != nil {
		e.Logger.Fatalf("cannot generate swagger: %w", err)
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
