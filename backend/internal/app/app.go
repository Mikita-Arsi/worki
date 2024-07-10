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
	/*chatsGroup := e.Group("/chats")
	messagesGroup := e.Group("/messages")*/

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	usersGroup.POST("/", api.CreateUser)
	usersGroup.GET("/", api.GetUsers)
	usersGroup.GET("/:id", api.GetUserByID)
	usersGroup.GET("/username/:username", api.GetUserByUsername)
	usersGroup.PUT("/", api.UpdateUser)
	usersGroup.DELETE("/:id", api.DeleteUserByID)
	usersGroup.DELETE("/username/:username", api.DeleteUserByUsername)
	usersGroup.Use(logger.LogRequest)

	/*chatsGroup.POST("/", api.CreateChat)
	chatsGroup.GET("/", api.GetChats)
	chatsGroup.GET("/messages/:id", api.GetChatMessages)
	chatsGroup.GET("/users/:id", api.GetChatUsers)
	chatsGroup.POST("/addUser", api.AddUserToChat)
	chatsGroup.POST("/removeUser", api.RemoveUserFromChat)
	chatsGroup.DELETE("/:id", api.DeleteChat)
	chatsGroup.Use(logger.LogRequest)

	messagesGroup.POST("/:id", api.CreateMessage)
	messagesGroup.GET("/", api.GetMessages)
	messagesGroup.DELETE("/:id", api.DeleteMessageByID)
	messagesGroup.DELETE("/:chatID/:messageID", api.DeleteMessageByChatIDAndMessageID)
	messagesGroup.Use(logger.LogRequest)*/

	e.Logger.Fatal(e.Start(":1323"))
}
