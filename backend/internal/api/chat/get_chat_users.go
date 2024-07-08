package chat

import (
	"net/http"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)

type GetChatUsersReq struct {
	ID uint `json:"id" validate:"uuid"`
}

type GetChatUsersRes struct {
	UsersID []uint `json:"users_id"`
}

func GetChatUsers(e echo.Context, req GetChatUsersReq, httpReq *http.Request) (GetChatUsersRes, error) {
	chat := &models.DBChat{}
	dbRequest := storage.GetDB().Model(chat).
		Where("id = ?", req.ID).
		First(chat)
	if dbRequest.Error != nil {
		return GetChatUsersRes{}, dbRequest.Error
	}
	return GetChatUsersRes{UsersID: chat.UsersID}, nil
}
