package config

import (
	"fmt"
	"os"
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
	

	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	JWTSecret = os.Getenv("JWT_SECRET")

	return nil

}

func GetDBDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", DBHost, DBPort, DBUser, DBPassword, DBName)
}
