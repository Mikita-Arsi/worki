package chat

type AddUserToChatReq struct {
	ChatID uint `json:"chat_id" validate:"required"`
	UserID uint `json:"user_id" validate:"required"`
}

type AddUserToChatRes struct {
}

/*
func AddUserToChat(ctx context.Context, req AddUserToChatReq, httpReq *http.Request) (AddUserToChatRes, *wrapper.ErrorResult) {
	chat := &models.DBChat{}
	dbRequest := storage.GetDB().Model(chat).
		Where("id = ?", req.ChatID).
		First(chat)
	if dbRequest.Error != nil {
		errRes := wrapper.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: dbRequest.Error,
		}
		return AddUserToChatRes{}, &errRes
	}

	// Add the user to the chat
	chat.UsersID = append(chat.UsersID, req.UserID)
	updateRequest := storage.GetDB().Save(chat)
	if updateRequest.Error != nil {
		errRes := wrapper.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: updateRequest.Error,
		}
		return AddUserToChatRes{}, &errRes
	}

	return AddUserToChatRes{}, nil
}
*/
