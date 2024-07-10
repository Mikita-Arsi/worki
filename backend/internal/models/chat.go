package models

import (
	"time"
	"worki/internal/schemas"

	"gorm.io/gorm"
)

// DBChat is a Postgres chat
type DBChat struct {
	gorm.Model
}

// ToWeb converts DBChat to Chat
func (dbChat *DBChat) ToWeb() *schemas.Chat {
	return &schemas.Chat{
		ID:         dbChat.ID,
		CreatedAt:  dbChat.CreatedAt,
		Name: 		dbChat.Name,
	}
}

// ToWeb converts ChatUserDB to Chat
func (ChatUserDB *ChatUserDB) ToWeb() *schemas.ChatUser {
	return &schemas.Chat{
		ID:         ChatUserDB.ID,
		CreatedAt:  ChatUserDB.CreatedAt,
		ChatID:     ChatUserDB.ChatID,
		FromID:     ChatUserDB.FromID,
	}
}

// ToWeb converts ChatMessageDB to Chat
func (ChatMessageDB *ChatMessageDB) ToWeb() *schemas.ChatMessage {
	return &schemas.Chat{
		ID:         ChatMessageDB.ID,
		CreatedAt:  ChatMessageDB.CreatedAt,
		ChatID:     ChatMessageDB.ChatID,
		MessageID:  ChatMessageDB.FromID,
	}
}
