package chat

import (
	"context"
	"net/http"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/AlhimicMan/goswag/wrapper"
)

type RemoveUserFromChatReq struct {
	ChatID uint `json:"chat_id" validate:"required"`
	UserID uint `json:"user_id" validate:"required"`
}

type RemoveUserFromChatRes struct {
}

func RemoveUserFromChat(ctx context.Context, req RemoveUserFromChatReq, httpReq *http.Request) (RemoveUserFromChatRes, *wrapper.ErrorResult) {
	chat := &models.DBChat{}
	dbRequest := storage.DB.Model(chat).
		Where("id = ?", req.ChatID).
		First(chat)
	if dbRequest.Error != nil {
		errRes := wrapper.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: dbRequest.Error,
		}
		return RemoveUserFromChatRes{}, &errRes
	}

	// Remove the user from the chat
	var updatedUsersID []uint
	for _, id := range chat.UsersID {
		if id != req.UserID {
			updatedUsersID = append(updatedUsersID, id)
		}
	}
	chat.UsersID = updatedUsersID
	updateRequest := storage.DB.Save(chat)
	if updateRequest.Error != nil {
		errRes := wrapper.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: updateRequest.Error,
		}
		return RemoveUserFromChatRes{}, &errRes
	}

	return RemoveUserFromChatRes{}, nil
}
