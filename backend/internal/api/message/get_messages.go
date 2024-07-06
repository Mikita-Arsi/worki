package message

import (
	"context"
	"net/http"
	"worki/internal/models"
	"worki/internal/storage"

	"github.com/AlhimicMan/goswag/wrapper"
)

func GetMessages(ctx context.Context, httpReq *http.Request) ([]models.Message, error) {
	var messages []models.DBMessage
	dbRequest := storage.DB.Model(&messages).Find(&messages)
	if dbRequest.Error != nil {
		errRes := wrapper.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: dbRequest.Error,
		}
		return nil, &errRes
	}
	var result []models.Message
	for _, dbMessage := range messages {
		result = append(result, *dbMessage.ToWeb())
	}
	return result, nil
}
