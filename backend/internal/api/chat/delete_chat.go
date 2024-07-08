package chat

import (
	"context"
	"net/http"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/AlhimicMan/goswag/wrapper"
)

type DeleteChatReq struct {
	ID uint `json:"id" validate:"uuid"`
}

type DeleteChatRes struct {
}

func DeleteChat(ctx context.Context, req DeleteChatReq, httpReq *http.Request) (DeleteChatRes, *wrapper.ErrorResult) {
	dbRequest := storage.DB.Delete(&models.DBChat{}, req.ID)
	if dbRequest.Error != nil {
		errRes := wrapper.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: dbRequest.Error,
		}
		return DeleteChatRes{}, &errRes
	}
	return DeleteChatRes{}, nil
}
