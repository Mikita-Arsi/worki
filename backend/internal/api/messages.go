package api

import (
	"net/http"
	"strconv"
	"worki/internal/models"
	"worki/internal/schemas"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)

// @Summary Create message
// @Accept json
// @Param request body schemas.MessageToCreate true "Message data"
// @Success 201 {object} schemas.Message
// @Failure 400 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /messages/ [post]
func CreateMessage(c echo.Context) error {
	req := new(models.MessageDB)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}
	tx := storage.GetDB().Create(&req)
	if tx.Error != nil {
		return c.JSON(500, schemas.HTTPError{
			Message: tx.Error,
		})
	}

	return c.JSON(http.StatusCreated, req.ToWeb())
}

// @Summary Get messages
// @Accept json
// @Success 200 {object} schemas.Messages
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /messages/ [get]
func GetMessages(c echo.Context) error {
	var messages []models.MessageDB
	storage.GetDB().Find(&messages)
	if len(messages) == 0 {
		return c.JSON(404, echo.ErrNotFound)
	}
	return c.JSON(200, messages)
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

	dbRequest := storage.GetDB().Delete(&models.MessageDB{}, id)
	if dbRequest.Error != nil {
		return c.JSON(http.StatusInternalServerError, schemas.HTTPError{Message: dbRequest.Error.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
