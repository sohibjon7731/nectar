package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
)

func LoadConfig() error {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		return err
	}
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	JWTSecret = os.Getenv("JWT_SECRET")

	return nil

}

func GetDBDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		DBUser, DBPassword, DBHost, DBPort, DBName)
}
