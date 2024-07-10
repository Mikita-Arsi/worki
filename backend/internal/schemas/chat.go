package schemas

import "time"

type ChatToCreate struct {
	Name string `json:"name"`
}

type ChatUserToCreate struct {
	ChatID uint `json:"chat_id"`
	FromID uint `json:"from_id"`
}

type ChatUser struct {
	ID        uint      `json:"id"`
	ChatID    uint      `json:"chat_id"`
	FromID    uint      `json:"from_id"`
	CreatedAt time.Time `json:"created_at"`
}

// Chat is a JSON chat
type Chat struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// Chat is a JSON chat
type ChatToUpdate struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Chats struct {
	Chats []Chat `json:"chats"`
}
