package schemas

import "time"


type ChatToCreate struct {
	Name 	   string	  `json:"name"`
}

type ChatUserToCreate struct {
	ChatID     uint      `json:"chat_id"`
	FromID     uint      `json:"from_id"`
}

type ChatMessageToCreate struct {
	ChatID     uint      `json:"chat_id"`
	FromID     uint      `json:"from_id"`
}

// Chat is a JSON chat
type Chat struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
}

type Chats struct {
	Chats []Chats `json:"chats"`
}