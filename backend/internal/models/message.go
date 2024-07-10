package models

import (
	"time"
	"worki/internal/schemas"

	"gorm.io/gorm"
)

type DBMessage struct {
	gorm.Model
	ChatID    int       `pg:"chat_id,notnull" json:"chat_id" validate:"required"`
	FromID    int       `pg:"from_id,notnull" json:"from_id" validate:"required"`
	Text      string    `pg:"text,notnull" json:"text" validate:"required"`
	CreatedAt time.Time `pg:"created_at,notnull"`
}

// TableName overrides default table name for gorm
func (DBMessage) TableName() string {
	return "messages"
}

// ToWeb converts DBMessage to Message
func (dbMessage *DBMessage) ToWeb() *schemas.Message {
	return &schemas.Message{
		ID:        dbMessage.ID,
		ChatID:    dbMessage.ChatID,
		FromID:    dbMessage.FromID,
		Text:      dbMessage.Text,
		CreatedAt: dbMessage.CreatedAt,
	}
}
