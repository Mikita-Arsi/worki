package message

type DeleteMessageReq struct {
	ID uint `json:"id" validate:"uuid"`
}

type DeleteMessageRes struct {
}

/*
func DeleteMessage(ctx context.Context, req DeleteMessageReq, httpReq *http.Request) (DeleteMessageRes, *wrapper.ErrorResult) {
	dbRequest := storage.GetDB().Delete(&models.DBMessage{}, req.ID)
	if dbRequest.Error != nil {
		errRes := wrapper.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: dbRequest.Error,
		}
		return DeleteMessageRes{}, &errRes
	}
	return DeleteMessageRes{}, nil
}
*/
