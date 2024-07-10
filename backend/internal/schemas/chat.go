package schemas

type AddUserToChatReq struct {
	ChatID int `json:"chat_id" validate:"required"`
	UserID int `json:"user_id" validate:"required"`
}

type GetChatUsersRes struct {
	AddUserToChatReq
}

type MessageToCreate struct {
	ChatID int    `json:"chat_id" validate:"required"`
	FromID int    `json:"from_id" validate:"required"`
	Text   string `json:"text" validate:"required"`
}
