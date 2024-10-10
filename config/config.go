package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Server struct {
		Address string
	}
	Database struct {
		DSN string
	}
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	var config Config

	config.Server.Address = os.Getenv("SERVER_ADDRESS")
	config.Database.DSN = os.Getenv("DATABASE_DSN")
	if config.Server.Address == "" {
		config.Server.Address = ":8081" // Значение по умолчанию
	}

	return &config, nil
}

func InitDB(dsn string) *gorm.DB {
	fmt.Println("Using DSN:", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Connecting...")
	return db
}
