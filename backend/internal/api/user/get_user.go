package user

import (
	"context"
	"net/http"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/AlhimicMan/goswag/wrapper"
	"github.com/google/uuid"
)

type GetUserReq struct {
	ID uuid.UUID `json:"id" validate:"uuid"`
}

func GetUser(ctx context.Context, req GetUserReq, httpReq *http.Request) (models.User, error) {
	user := &models.DBUser{}
	dbRequest := storage.DB.Model(user).
		Where("id = ?", req.ID).
		First(user)
	if dbRequest.Error != nil {
		errRes := wrapper.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: dbRequest.Error,
		}
		return models.User{}, &errRes
	}
	return *user.ToWeb(), nil
}
