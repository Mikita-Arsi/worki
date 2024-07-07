package user

import (
	"net/http"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)

type GetUserReq struct {
	ID uint `json:"id" validate:"uuid"`
}

func GetUser(e echo.Context, req GetUserReq, httpReq *http.Request) (models.User, error) {
	user := &models.DBUser{}
	dbRequest := storage.DB.Model(user).
		Where("id = ?", req.ID).
		First(user)
	if dbRequest.Error != nil {
		return models.User{}, dbRequest.Error
	}
	return *user.ToWeb(), nil
}
