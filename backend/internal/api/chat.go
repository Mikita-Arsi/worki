package api

import (
	"net/http"
	"strconv"
	"worki/internal/models"
	"worki/internal/schemas"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)

// @Summary Create chat
// @Tags Chats
// @Accept json
// @Param request body schemas.ChatToCreate true "Chat data"
// @Success 201 {object} schemas.Chat
// @Failure 400 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/ [post]
func CreateChat(c echo.Context) error {
	req := new(models.ChatDB)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}
	tx := storage.GetDB().Create(&req)
	if tx.Error != nil {
		return c.JSON(500, schemas.HTTPError{
			Message: tx.Error,
		})
	}
	return c.JSON(201, req.ToWeb())
}

// @Summary Get chats
// @Tags Chats
// @Accept json
// @Success 200 {object} schemas.Chats
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/ [get]
func GetChats(c echo.Context) error {
	var chats []models.ChatDB
	var chatsSchema schemas.Chats
	storage.GetDB().Find(&chats)
	if len(chats) == 0 {
		return c.JSON(404, echo.ErrNotFound)
	}
	for _, value := range chats {
		chatsSchema.Chats = append(chatsSchema.Chats, *value.ToWeb())
	}
	return c.JSON(200, chatsSchema.Chats)
}

// @Summary Get chat by ID
// @Tags Chats
// @Accept json
// @Param id path int true "Chat ID"
// @Success 200 {object} schemas.Chat
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/{id} [get]
func GetChat(c echo.Context) error {
	id := c.Param("id")
	var chat models.ChatDB
	tx := storage.GetDB().First(&chat, id)
	if tx.Error != nil {
		return c.JSON(404, echo.ErrNotFound)
	}
	return c.JSON(200, chat.ToWeb())
}

// @Summary Update chat name
// @Tags Chats
// @Accept json
// @Param request body schemas.ChatToUpdate true "Chat data"
// @Success 201 {object} schemas.User
// @Failure 400 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/ [put]
func UpdateChatName(c echo.Context) error {
	chatSchema := new(schemas.ChatToUpdate)
	if err := c.Bind(chatSchema); err != nil {
		return c.JSON(400, echo.ErrBadRequest)
	}
	chat := models.ChatDB{}
	storage.GetDB().First(&chat, chatSchema.ID)
	chat.Name = chatSchema.Name
	tx := storage.GetDB().Save(&chat)
	if tx.Error != nil {
		return c.JSON(500, echo.ErrInternalServerError)
	}
	return c.JSON(201, chat.ToWeb())
}

// @Summary Add user to chat
// @Tags Chats
// @Accept json
// @Param request body schemas.ChatUserToCreate true "Add User to Chat data"
// @Success 204 {object} nil
// @Failure 400 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/add [post]
func AddUserToChat(c echo.Context) error {
	req := new(models.ChatUserDB)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}
	anotherUser := models.ChatUserDB{}
	storage.GetDB().Find(&anotherUser, "fromid = ?", req.FromID)
	if anotherUser.FromID == req.FromID {
		return c.JSON(400, echo.ErrBadRequest)
	}
	tx := storage.GetDB().Create(&req)
	if tx.Error != nil {
		return c.JSON(500, schemas.HTTPError{
			Message: tx.Error,
		})
	}

	return c.NoContent(http.StatusNoContent)
}

// @Summary Get chat messages
// @Tags Chats
// @Accept json
// @Param id path int true "Chat ID"
// @Success 200 {object} schemas.Messages
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/{id}/messages [get]
func GetChatMessages(c echo.Context) error {
	chatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, schemas.HTTPError{Message: "Invalid chat ID"})
	}

	var messages []models.MessageDB
	req := schemas.Messages{}
	dbRequest := storage.GetDB().Find(&messages, "chatid = ?", chatID)
	if len(messages) == 0 {
		return c.JSON(404, echo.ErrNotFound)
	}
	if dbRequest.Error != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPError{Message: dbRequest.Error.Error()})
	}
	for _, value := range messages {
		req.Messages = append(req.Messages, *value.ToWeb())
	}

	return c.JSON(http.StatusOK, req)
}

// @Summary Get chat users
// @Tags Chats
// @Accept json
// @Param id path int true "Chat ID"
// @Success 200 {object} schemas.Users
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/{id}/users [get]
func GetChatUsers(c echo.Context) error {
	chatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, schemas.HTTPError{Message: "Invalid chat ID"})
	}

	chat := []models.ChatUserDB{}
	users := schemas.Users{}
	dbRequest := storage.GetDB().Where("id = ?", chatID).Find(&chat)
	if dbRequest.Error != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPError{Message: dbRequest.Error.Error()})
	}
	for _, value := range chat {
		user := &models.UserDB{}
		storage.GetDB().First(user, "userid = ?", value.FromID)
		users.Users = append(users.Users, *user.ToWeb())
	}

	return c.JSON(http.StatusOK, users)
}

// @Summary Remove user from chat
// @Tags Chats
// @Accept json
// @Param request body schemas.ChatUserToCreate true "Remove User from Chat data"
// @Success 204 {object} nil
// @Failure 400 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/user [delete]
func RemoveUserFromChat(c echo.Context) error {
	req := new(models.ChatUserDB)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}
	tx := storage.GetDB().Delete(&req)
	if tx.Error != nil {
		return c.JSON(404, echo.ErrNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}

// @Summary Delete chat by ID
// @Tags Chats
// @Accept json
// @Param id path int true "Chat ID"
// @Success 204 {object} nil
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/{id} [delete]
func DeleteChat(c echo.Context) error {
	chatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, schemas.HTTPError{Message: "Invalid chat ID"})
	}

	dbRequest := storage.GetDB().Delete(&models.ChatDB{}, chatID)
	if dbRequest.Error != nil {
		return c.JSON(404, echo.ErrNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}
