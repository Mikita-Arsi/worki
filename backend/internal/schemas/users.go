package schemas

import "time"

type UserToCreate struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email"  validate:"email"`
	Password  string `json:"password" validate:"required"`
}

type UserToUpdate struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email"  validate:"email"`
}

type User struct {
	ID        uint      `json:"id"`
	Firstname string    `json:"firstname" validate:"required"`
	Lastname  string    `json:"lastname" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email"  validate:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Users struct {
	Users []User `json:"users"`
}

type UserWithPassword struct {
	User
	Password string `json:"password" validate:"required"`
}
