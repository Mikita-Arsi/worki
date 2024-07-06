package user

import (
	"context"
	"worki/internal/models"
	"worki/internal/storage"
)

func UpdateUser(ctx context.Context, req models.UserWithPassword) (models.User, error) {

	dbModel := req.ToDB()
	dbRequest := storage.DB.Save(dbModel)
	if dbRequest.Error != nil {
		return models.User{}, dbRequest.Error
	}

	return *req.ToDB().ToWeb(), nil
}
