package models

import (
	"time"
)

type MessageToCreate struct {
	ChatID int    `json:"chat_id" validate:"required"`
	FromID int    `json:"from_id" validate:"required"`
	Text   string `json:"text" validate:"required"`
}

// ToDB converts MessageToCreate to DBMessage
func (message *MessageToCreate) ToDB(createdAt time.Time) *DBMessage {
	return &DBMessage{
		ChatID:    message.ChatID,
		FromID:    message.FromID,
		Text:      message.Text,
		CreatedAt: createdAt,
	}
}

// Message is a JSON message
type Message struct {
	ID        uint      `json:"id"`
	ChatID    int       `json:"chat_id" validate:"required"`
	FromID    int       `json:"from_id" validate:"required"`
	Text      string    `json:"text" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// DBMessage is a Postgres message
type DBMessage struct {
	tableName struct{}  `pg:"messages" gorm:"primaryKey"`
	ID        uint      `pg:"id,notnull,pk"`
	ChatID    int       `pg:"chat_id,notnull"`
	FromID    int       `pg:"from_id,notnull"`
	Text      string    `pg:"text,notnull"`
	CreatedAt time.Time `pg:"created_at,notnull"`
}

// TableName overrides default table name for gorm
func (DBMessage) TableName() string {
	return "messages"
}

// ToWeb converts DBMessage to Message
func (dbMessage *DBMessage) ToWeb() *Message {
	return &Message{
		ID:        dbMessage.ID,
		ChatID:    dbMessage.ChatID,
		FromID:    dbMessage.FromID,
		Text:      dbMessage.Text,
		CreatedAt: dbMessage.CreatedAt,
	}
}
