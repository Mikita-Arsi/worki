package config

import "os"

type Config struct {
	UserDB     string `env:"POSTGRES_USER"`
	PasswordDB string `env:"POSTGRES_PASSWORD"`
	NameDB     string `env:"POSTGRES_DB"`
}

func (cfg *Config) Load() {
	cfg.NameDB = os.Getenv("POSTGRES_DB")
	cfg.PasswordDB = os.Getenv("POSTGRES_PASSWORD")
	cfg.UserDB = os.Getenv("POSTGRES_USER")
}
