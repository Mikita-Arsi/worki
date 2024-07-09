package api

import (
	"net/http"
	"time"
	"worki/internal/models"
	"worki/internal/schemas"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)

// @Summary Create message
// @Accept json
// @Param chatID path int true "Chat ID"
// @Param request body schemas.MessageToCreate true "Message data"
// @Success 201 {object} schemas.Message
// @Failure 400 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/{chatID}/messages [post]
func CreateMessage(c echo.Context) error {
	chatID, err := strconv.Atoi(c.Param("chatID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, schemas.HTTPError{Message: "Invalid chat ID"})
	}

	req := new(schemas.MessageToCreate)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	date := time.Now()
	dbMessage := req.ToDB(date)
	dbMessage.ChatID = uint(chatID)
	dbRequest := storage.GetDB().Create(dbMessage)
	if dbRequest.Error != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPError{Message: dbRequest.Error.Error()})
	}

	return c.JSON(http.StatusCreated, dbMessage.ToWeb())
}

// @Summary Get messages
// @Accept json
// @Success 200 {object} []schemas.Message
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /messages [get]
func GetMessages(c echo.Context) error {
	var messages []models.DBMessage
	dbRequest := storage.GetDB().Find(&messages)
	if dbRequest.Error != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPError{Message: dbRequest.Error.Error()})
	}
	if len(messages) == 0 {
		return c.JSON(http.StatusNotFound, echo.ErrNotFound)
	}

	var result []schemas.Message
	for _, dbMessage := range messages {
		result = append(result, *dbMessage.ToWeb())
	}
	return c.JSON(http.StatusOK, result)
}

// @Summary Delete message by ID
// @Accept json
// @Param id path int true "Message ID"
// @Success 204 {object} schemas.Message
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /messages/{id} [delete]
func DeleteMessageByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, schemas.HTTPError{Message: "Invalid message ID"})
	}

	dbRequest := storage.GetDB().Delete(&models.DBMessage{}, id)
	if dbRequest.Error != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPError{Message: dbRequest.Error.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// @Summary Delete message by chat ID and message ID
// @Accept json
// @Param chatID path int true "Chat ID"
// @Param messageID path int true "Message ID"
// @Success 204 {object} schemas.Message
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /chats/{chatID}/messages/{messageID} [delete]
func DeleteMessageByChatIDAndMessageID(c echo.Context) error {
	chatID, err := strconv.Atoi(c.Param("chatID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, schemas.HTTPError{Message: "Invalid chat ID"})
	}

	messageID, err := strconv.Atoi(c.Param("messageID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, schemas.HTTPError{Message: "Invalid message ID"})
	}

	dbRequest := storage.GetDB().Where("chat_id = ?", chatID).Delete(&models.DBMessage{}, messageID)
	if dbRequest.Error != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPError{Message: dbRequest.Error.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
