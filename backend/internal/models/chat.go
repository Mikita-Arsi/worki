package models

import (
	"time"
)

type ChatToCreate struct {
	UsersID    []uint `json:"users_id" validate:"required"`
	MessagesID []uint `json:"messages_id" validate:"required"`
}

// ToDB converts ChatToCreate to DBChat
func (chat *ChatToCreate) ToDB(createdAt time.Time) *DBChat {
	return &DBChat{
		UsersID:    chat.UsersID,
		MessagesID: chat.MessagesID,
		CreatedAt:  createdAt,
	}
}

// Chat is a JSON chat
type Chat struct {
	ID         uint      `json:"id"`
	UsersID    []uint    `json:"users_id" validate:"required"`
	MessagesID []uint    `json:"messages_id" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
}

// ToDB converts Chat to DBChat
func (chat *Chat) ToDB() *DBChat {
	return &DBChat{
		ID:         chat.ID,
		UsersID:    chat.UsersID,
		MessagesID: chat.MessagesID,
		CreatedAt:  chat.CreatedAt,
	}
}

// DBChat is a Postgres chat
type DBChat struct {
	tableName  struct{}  `pg:"chats" gorm:"primaryKey"`
	ID         uint      `pg:"id,notnull,pk"`
	UsersID    []uint    `pg:"users_id,notnull"`
	MessagesID []uint    `pg:"messages_id,notnull"`
	CreatedAt  time.Time `pg:"created_at,notnull"`
}

// TableName overrides default table name for gorm
func (DBChat) TableName() string {
	return "chats"
}

// ToWeb converts DBChat to Chat
func (dbChat *DBChat) ToWeb() *Chat {
	return &Chat{
		ID:         dbChat.ID,
		UsersID:    dbChat.UsersID,
		MessagesID: dbChat.MessagesID,
		CreatedAt:  dbChat.CreatedAt,
	}
}
