package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var cfg *Config

type Config struct {
	DBUrl         string
	DBPassword    string
	DBName        string
	Port          string
	Secret        string
	TokenExp      string
	ModelPath     string
	ONNXPath      string
	MigrationPath string
}

func newConfig() *Config {
	return &Config{
		DBUrl:         os.Getenv("DB_URL"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_DATABASE"),
		Port:          os.Getenv("PORT"),
		Secret:        os.Getenv("SECRET"),
		TokenExp:      os.Getenv("TOKEN_EXP"),
		ModelPath:     os.Getenv("MODEL_PATH"),
		ONNXPath:      os.Getenv("ONNX_PATH"),
		MigrationPath: os.Getenv("MIGRATION_PATH"),
	}
}

func Get() *Config {
	if cfg == nil {
		cfg = newConfig()
	}

	return cfg
}
