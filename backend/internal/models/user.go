package models

import (
	"worki/internal/schemas"

	"gorm.io/gorm"
)

type UserDB struct {
	gorm.Model
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Username  string `pg:"unique" json:"username" validate:"required"`
	Email     string `pg:"unique" json:"email"  validate:"email"`
	Password  string `json:"password" validate:"required"`
}

func (dbUser *UserDB) ToWeb() *schemas.User {
	return &schemas.User{
		ID:        dbUser.ID,
		Firstname: dbUser.Firstname,
		Lastname:  dbUser.Lastname,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		CreatedAt: dbUser.CreatedAt,
	}
}
