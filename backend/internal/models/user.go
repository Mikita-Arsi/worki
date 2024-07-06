package models

import (
	"time"

	"github.com/google/uuid"
)

type UserToCreate struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email"  validate:"email"`
	Password  string `json:"password" validate:"required"`
}

// ToDB converts UserToCreate to DBUser
func (user *UserToCreate) ToDB(id uuid.UUID, createdAt time.Time) *DBUser {
	return &DBUser{
		ID:        id,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: createdAt,
	}
}

// User is a JSON user
type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname" validate:"required"`
	Lastname  string    `json:"lastname" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email"  validate:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserWithEmail struct {
	User
}

type UserWithPassword struct {
	User
	Password string `json:"password" validate:"required"`
}

// ToDB converts User to DBUser
func (user *UserWithPassword) ToDB() *DBUser {
	return &DBUser{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// DBUser is a Postgres user
type DBUser struct {
	tableName struct{}  `pg:"users" gorm:"primaryKey"`
	ID        uuid.UUID `pg:"id,notnull,pk"`
	Firstname string    `pg:"firstname,notnull"`
	Lastname  string    `pg:"lastname,notnull"`
	Password  string    `pg:"password,notnull"`
	Email     string    `pg:"password,notnull"`
	Username  string    `pg:"username,notnull"`
	CreatedAt time.Time `pg:"created_at,notnull"`
}

// TableName overrides default table name for gorm
func (DBUser) TableName() string {
	return "users"
}

// ToWeb converts DBUser to User
func (dbUser *DBUser) ToWeb() *User {
	return &User{
		ID:        dbUser.ID,
		Firstname: dbUser.Firstname,
		Lastname:  dbUser.Lastname,
		Username:  dbUser.Username,
		CreatedAt: dbUser.CreatedAt,
	}
}
