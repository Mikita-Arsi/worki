package models

import (
	"time"
	"worki/internal/schemas"

	"gorm.io/gorm"
)

type MessageToCreate struct {
	ChatID int    `json:"chat_id" validate:"required"`
	FromID int    `json:"from_id" validate:"required"`
	Text   string `json:"text" validate:"required"`
}

// DBMessage is a Postgres message
type MessageDB struct {
	gorm.Model
	ID        int       `pg:"id,notnull,pk"`
	ChatID    int       `pg:"chat_id,notnull"`
	FromID    int       `pg:"from_id,notnull"`
	Text      string    `pg:"text,notnull"`
	CreatedAt time.Time `pg:"created_at,notnull"`
}

// ToWeb converts DBMessage to Message
func (dbMessage *MessageDB) ToWeb() *schemas.Message {
	return &schemas.Message{
		ID:        dbMessage.ID,
		ChatID:    dbMessage.ChatID,
		FromID:    dbMessage.FromID,
		Text:      dbMessage.Text,
		CreatedAt: dbMessage.CreatedAt,
	}
}
