package chat

import (
	"net/http"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)

type GetChatMessagesReq struct {
	ID uint `json:"id" validate:"uuid"`
}

type GetChatMessagesRes struct {
	MessagesID []uint `json:"messages_id"`
}

func GetChatMessages(e echo.Context, req GetChatMessagesReq, httpReq *http.Request) (GetChatMessagesRes, error) {
	chat := &models.DBChat{}
	dbRequest := storage.DB.Model(chat).
		Where("id = ?", req.ID).
		First(chat)
	if dbRequest.Error != nil {
		return GetChatMessagesRes{}, dbRequest.Error
	}
	return GetChatMessagesRes{MessagesID: chat.MessagesID}, nil
}
