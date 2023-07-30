package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type EmailConfig struct {
	Address  string
	Password string
}

func InitDbConfig() (*DbConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("error loading .env file")
		return nil, err
	}
	config := &DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return config, nil
}

func InitEmailConfig() (*EmailConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("error loading .env file")
		return nil, err
	}
	config := &EmailConfig{
		Address:  os.Getenv("MAIL_ADDR"),
		Password: os.Getenv("MAIL_PASS"),
	}
	return config, nil
}
