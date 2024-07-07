package storage

import (
	"fmt"
	"worki/internal/config"
	"worki/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB(cfg *config.Config) {
	fmt.Print(cfg.NameDB, cfg.PasswordDB, cfg.UserDB)
	var err error
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@postgres:5432/%s?sslmode=disable",
		cfg.UserDB,
		cfg.PasswordDB,
		cfg.NameDB,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	DB.AutoMigrate(&models.User{})
}
