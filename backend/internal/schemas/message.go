package schemas

import "time"

type MessageToCreate struct {
	ChatID int    `json:"chat_id" validate:"required"`
	FromID int    `json:"from_id" validate:"required"`
	Text   string `json:"text" validate:"required"`
}

// Message is a JSON message
type Message struct {
	ID        uint      `json:"id"`
	ChatID    int       `json:"chat_id" validate:"required"`
	FromID    int       `json:"from_id" validate:"required"`
	Text      string    `json:"text" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
