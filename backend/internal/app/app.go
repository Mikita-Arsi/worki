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
	chatsGroup := e.Group("/chats")
	messagesGroup := e.Group("/messages")

	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	usersGroup.POST("/", api.CreateUser)
	usersGroup.GET("/", api.GetUsers)
	usersGroup.GET("/:id", api.GetUserByID)
	usersGroup.GET("/usr/:username", api.GetUserByUsername)
	usersGroup.PUT("/:id", api.UpdateUser)
	usersGroup.DELETE("/:id", api.DeleteUserByID)
	usersGroup.DELETE("/usr/:username", api.DeleteUserByUsername)
	usersGroup.Use(logger.LogRequest)

	messagesGroup.POST("/", api.CreateMessage)
	messagesGroup.GET("/", api.GetMessages)
	messagesGroup.DELETE("/:id", api.DeleteMessageByID)
	messagesGroup.Use(logger.LogRequest)

	chatsGroup.POST("/", api.CreateChat)
	chatsGroup.GET("/:id", api.GetChat)
	chatsGroup.GET("/", api.GetChats)
	chatsGroup.PUT("/:id", api.UpdateChatName)
	chatsGroup.GET("/:id/messages", api.GetChatMessages)
	chatsGroup.GET("/:id/users", api.GetChatUsers)
	chatsGroup.DELETE("/", api.RemoveUserFromChat)
	chatsGroup.DELETE("/:id", api.DeleteChat)
	chatsGroup.POST("/add", api.AddUserToChat)
	chatsGroup.Use(logger.LogRequest)

	e.Logger.Fatal(e.Start(":1323"))
}
