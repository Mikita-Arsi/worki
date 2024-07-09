package message

import (
	"net/http"
	"time"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)

func CreateMessage(ctx echo.Context, req models.MessageToCreate, chatID uint, httpReq *http.Request) (models.Message, error) {
	date := time.Now()
	dbMessage := req.ToDB(date)
	dbMessage.ChatID = chatID
	dbRequest := storage.GetDB().Create(dbMessage)
	if dbRequest.Error != nil {
		return models.Message{}, dbRequest.Error
	}
	return *dbMessage.ToWeb(), nil
}
