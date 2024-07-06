package user

import (
	"net/http"
	"time"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)

func CreateUser(ctx echo.Context, req models.UserToCreate, httpReq *http.Request) (models.User, error) {
	createdAt := time.Now()
	dbUser := req.ToDB(createdAt)
	dbRequest := storage.DB.Create(dbUser)
	if dbRequest.Error != nil {
		return models.User{}, dbRequest.Error
	}
	return *dbUser.ToWeb(), nil
}
