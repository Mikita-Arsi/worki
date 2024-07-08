package chat

import (
	"net/http"
	"time"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)

func CreateChat(ctx echo.Context, req models.ChatToCreate, httpReq *http.Request) (models.Chat, error) {
	createdAt := time.Now()
	dbChat := req.ToDB(createdAt)
	dbRequest := storage.GetDB().Create(dbChat)
	if dbRequest.Error != nil {
		return models.Chat{}, dbRequest.Error
	}
	return *dbChat.ToWeb(), nil
}
