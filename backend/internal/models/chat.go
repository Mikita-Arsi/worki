package models

import (
	"worki/internal/schemas"

	"gorm.io/gorm"
)

// DBChat is a Postgres chat
type ChatDB struct {
	gorm.Model
	Name string `json:"name"`
}

type ChatUserDB struct {
	gorm.Model
	ChatID uint `json:"chat_id"`
	FromID uint `json:"from_id"`
}

type ChatUsersDB struct {
	Users []ChatUserDB `json:"users"`
}

// ToWeb converts DBChat to Chat
func (dbChat *ChatDB) ToWeb() *schemas.Chat {
	return &schemas.Chat{
		ID:        dbChat.ID,
		CreatedAt: dbChat.CreatedAt,
		Name:      dbChat.Name,
	}
}

// ToWeb converts ChatUserDB to Chat
func (ChatUserDB *ChatUserDB) ToWeb() *schemas.ChatUser {
	return &schemas.ChatUser{
		ID:        ChatUserDB.ID,
		CreatedAt: ChatUserDB.CreatedAt,
		ChatID:    ChatUserDB.ChatID,
		FromID:    ChatUserDB.FromID,
	}
}
