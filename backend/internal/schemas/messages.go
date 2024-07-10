package schemas

import "time"

// Message is a JSON message
type Message struct {
	ID        int       `json:"id"`
	ChatID    int       `json:"chat_id" validate:"required"`
	FromID    int       `json:"from_id" validate:"required"`
	Text      string    `json:"text" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
