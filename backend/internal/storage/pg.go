package storage

import (
	"fmt"
	"os"
	"worki/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	var err error
	dsn := fmt.Sprintf("postgresql://%s:%s@postgres:5432/%s?sslmode=disable", dbUser, dbPassword, dbName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto migrate the User model
	DB.AutoMigrate(&models.User{})
}
