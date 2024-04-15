package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var cfg *Config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBname     string
	Port       string
	Secret     string
	TokenExp   string
}

func newConfig() *Config {
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBname:     os.Getenv("DB_DATABASE"),
		Port:       os.Getenv("PORT"),
		Secret:     os.Getenv("SECRET"),
		TokenExp:   os.Getenv("TOKEN_EXP"),
	}
}

func Get() *Config {
	if cfg == nil {
		cfg = newConfig()
	}

	return cfg
}
