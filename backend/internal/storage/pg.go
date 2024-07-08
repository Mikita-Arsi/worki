package storage

import (
	"fmt"
	"worki/internal/config"
	"worki/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDB(cfg *config.Config) {
	test := "dihiq"
	fmt.Printf("%v %v %v %v", test, cfg.UserDB,
		cfg.PasswordDB,
		cfg.NameDB)
	var err error
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@postgres:5432/%s?sslmode=disable",
		cfg.UserDB,
		cfg.PasswordDB,
		cfg.NameDB,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.UserDB{})
}

func GetDB() *gorm.DB {
	return db
}
