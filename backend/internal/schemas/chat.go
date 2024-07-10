package schemas

import "time"

type ChatToCreate struct {
	Name string `json:"name"`
}

type ChatUserToCreate struct {
	ChatID uint `json:"chat_id"`
	FromID uint `json:"from_id"`
}

type ChatMessageToCreate struct {
	ChatID uint `json:"chat_id"`
	FromID uint `json:"from_id"`
}

type ChatUser struct {
	ID        uint      `json:"id"`
	ChatID    uint      `json:"chat_id"`
	FromID    uint      `json:"from_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatMessage struct {
	ID        uint      `json:"id"`
	ChatID    uint      `json:"chat_id"`
	MessageID uint      `json:"message_id"`
	CreatedAt time.Time `json:"created_at"`
}

// Chat is a JSON chat
type Chat struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Chats struct {
	Chats []Chats `json:"chats"`
}
