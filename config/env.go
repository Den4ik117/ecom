package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PublicHost       string
	Port             string
	DBUser           string
	DBPassword       string
	DBAddress        string
	DBName           string
	RabbitMQHost     string
	RabbitMQPort     string
	RabbitMQUser     string
	RabbitMQPassword string
}

var Envs = initConfig()

func initConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env file: %s", err)
	}

	return &Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://127.0.0.1"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBAddress: fmt.Sprintf(
			"%s:%s",
			getEnv("DB_HOST", "127.0.0.1"),
			getEnv("DB_PORT", "3306"),
		),
		DBName:           getEnv("DB_NAME", "app"),
		RabbitMQHost:     getEnv("RABBITMQ_HOST", "127.0.0.1"),
		RabbitMQPort:     getEnv("RABBITMQ_PORT", "5672"),
		RabbitMQUser:     getEnv("RABBITMQ_USER", "guest"),
		RabbitMQPassword: getEnv("RABBITMQ_PASS", "guest"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
