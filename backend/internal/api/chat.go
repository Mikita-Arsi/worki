package api

import (
	"net/http"
	"worki/internal/models"
	"worki/internal/schemas"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)


// @Summary Create chat
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

// @Summary Add user to chat
// @Accept json
// @Param request body schemas.ChatUserToCreate true "Add User to Chat data"
// @Success 204 {object} nil
// @Failure 400 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/ [post]
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
// @Accept json
// @Param id path int true "Chat ID"
// @Success 200 {object} schemas.GetChatMessagesRes
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/msg/{id} [get]
func GetChatMessages(c echo.Context) error {
	chatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, schemas.HTTPError{Message: "Invalid chat ID"})
	}

	chat := &models.MessagesDB{}
	dbRequest := storage.GetDB().Where("id = ?", chatID).Find(chat)
	if dbRequest.Error != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPError{Message: dbRequest.Error.Error()})
	}
	return c.JSON(http.StatusOK, chat)
}
/*
// @Summary Get chat users
// @Accept json
// @Param id path int true "Chat ID"
// @Success 200 {object} schemas.GetChatUsersRes
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/usr/{id}/users [get]
func GetChatUsers(c echo.Context) error {
	chatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, schemas.HTTPError{Message: "Invalid chat ID"})
	}

	chat := &models.DBChat{}
	dbRequest := storage.GetDB().Where("id = ?", chatID).First(chat)
	if dbRequest.Error != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPError{Message: dbRequest.Error.Error()})
	}

	return c.JSON(http.StatusOK, schemas.GetChatUsersRes{UsersID: chat.UsersID})
}

// @Summary Remove user from chat
// @Accept json
// @Param request body schemas.RemoveUserFromChatReq true "Remove User from Chat data"
// @Success 204 {object} nil
// @Failure 400 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/removeUser [post]
func RemoveUserFromChat(c echo.Context) error {
	req := new(schemas.GetChatUsersRes)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	chat := &models.DBChat{}
	dbRequest := storage.GetDB().Where("id = ?", req.ChatID).First(chat)
	if dbRequest.Error != nil {
		return c.JSON(http.StatusBadRequest, schemas.HTTPError{Message: dbRequest.Error.Error()})
	}

	var updatedUsersID []int
	for _, id := range chat.UsersID {
		if id != req.UserID {
			updatedUsersID = append(updatedUsersID, id)
		}
	}
	chat.UsersID = updatedUsersID
	updateRequest := storage.GetDB().Save(chat)
	if updateRequest.Error != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPError{Message: updateRequest.Error.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// @Summary Delete chat by ID
// @Accept json
// @Param id path int true "Chat ID"
// @Success 204 {object} nil
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/{id} [delete]
func DeleteChat(c echo.Context) error {


	dbRequest := storage.GetDB().Delete(&models.DBChat{}, chatID)
	if dbRequest.Error != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPError{Message: dbRequest.Error.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
*/
