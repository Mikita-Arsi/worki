package user

import (
	"net/http"
	"time"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateUser(ctx echo.Context, req models.UserToCreate, httpReq *http.Request) (models.User, error) {
	userId := uuid.New()
	createdAt := time.Now()
	dbUser := req.ToDB(userId, createdAt)
	dbRequest := storage.DB.Create(dbUser)
	if dbRequest.Error != nil {
		return models.User{}, dbRequest.Error
	}
	return *dbUser.ToWeb(), nil
}
