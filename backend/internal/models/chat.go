package models

import (
	"worki/internal/schemas"

	"gorm.io/gorm"
)

// DBChat is a Postgres chat
type DBChat struct {
	gorm.Model
	Name string `json:"name"`
}

type ChatUserDB struct {
	gorm.Model
	ChatID uint `json:"chat_id"`
	FromID uint `json:"from_id"`
}

type ChatMessageDB struct {
	gorm.Model
	ChatID    uint `json:"chat_id"`
	MessageID uint `json:"message_id"`
}

type ChatMessagesDB struct {
	messages []ChatMessageDB `json:"messages"`
}

type ChatUsersDB struct {
	users []ChatUserDB `json:"users"`
}

// ToWeb converts DBChat to Chat
func (dbChat *DBChat) ToWeb() *schemas.Chat {
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

// ToWeb converts ChatMessageDB to Chat
func (ChatMessageDB *ChatMessageDB) ToWeb() *schemas.ChatMessage {
	return &schemas.ChatMessage{
		ID:        ChatMessageDB.ID,
		CreatedAt: ChatMessageDB.CreatedAt,
		ChatID:    ChatMessageDB.ChatID,
		MessageID: ChatMessageDB.MessageID,
	}
}
