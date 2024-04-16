package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBUser        string
	DBPassword    string
	DBName        string
	DBHost        string
	DBPort        string
	SSLMode       string
	WebServerPort string
}

func LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return Config{
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		SSLMode:       os.Getenv("SSL_MODE"),
		WebServerPort: os.Getenv("WEB_SERVER_PORT"),
	}
}
