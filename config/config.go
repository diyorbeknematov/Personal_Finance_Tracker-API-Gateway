package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT              int      `yaml:"http_port"`
	BUDGETING_SERVICE_PORT int      `yaml:"budgeting_service_port"`
	USER_SERVICE_PORT      int      `yaml:"user_service_port"`
	DB_HOST                string   `yaml:"db_host"`
	DB_PORT                int      `yaml:"db_port"`
	DB_USER                string   `yaml:"db_user"`
	DB_PASSWORD            string   `yaml:"db_password"`
	DB_NAME                string   `yaml:"db_name"`
	KafkaBrokers           []string `yaml:"kafka_brokers"`
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	config := &Config{}

	config.HTTP_PORT = cast.ToInt(coalesce("HTTP_PORT", 8080))
	config.USER_SERVICE_PORT = cast.ToInt(coalesce("USER_SERVICE_PORT", 50050))
	config.BUDGETING_SERVICE_PORT = cast.ToInt(coalesce("BUDGETING_SERVICE_PORT", 50051))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "your_password"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "your_db"))

	config.KafkaBrokers = cast.ToStringSlice(coalesce("KAFKA_BROKERS", "localhost:9092"))

	return config
}

func coalesce(key string, defaults interface{}) interface{} {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaults
}
