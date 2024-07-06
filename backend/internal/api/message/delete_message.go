package message

import (
	"context"
	"net/http"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/AlhimicMan/goswag/wrapper"
)

type DeleteMessageReq struct {
	ID uint `json:"id" validate:"uuid"`
}

type DeleteMessageRes struct {
}

func DeleteMessage(ctx context.Context, req DeleteMessageReq, httpReq *http.Request) (DeleteMessageRes, *wrapper.ErrorResult) {
	dbRequest := storage.DB.Delete(&models.DBMessage{}, req.ID)
	if dbRequest.Error != nil {
		errRes := wrapper.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: dbRequest.Error,
		}
		return DeleteMessageRes{}, &errRes
	}
	return DeleteMessageRes{}, nil
}
