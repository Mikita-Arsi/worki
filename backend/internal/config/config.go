package config

type Config struct {
	UserDB     string `env:"POSTGRES_USER"`
	PasswordDB string `env:"POSTGRES_PASSWORD"`
	NameDB     string `env:"POSTGRES_DB"`
}
