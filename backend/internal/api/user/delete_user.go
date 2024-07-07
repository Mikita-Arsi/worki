package user

import (
	"context"
	"net/http"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/AlhimicMan/goswag/wrapper"
)

type DeleteUserReq struct {
	ID uint `json:"id" validate:"uuid"`
}

type DeleteUserRes struct {
}

func DeleteUser(ctx context.Context, req DeleteUserReq, httpReq *http.Request) (DeleteUserRes, *wrapper.ErrorResult) {
	dbRequest := storage.DB.Delete(&models.DBUser{}, req.ID)
	if dbRequest.Error != nil {
		errRes := wrapper.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: dbRequest.Error,
		}
		return DeleteUserRes{}, &errRes
	}
	return DeleteUserRes{}, nil
}
