package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser       string
	PostgresPassword   string
	PostgresHost       string
	PostgresPort       string
	PostgresAddress    string
	LogLevel           string
	ServerAddress      string
	SSLMode            string
	RedisAddress       string
	RedisPassword      string
	RedisDB            int
	RedisHost          string
	AccessTokenSecret  string
	RefreshTokenSecret string
	LogFilePath        string
	AccessTTL          int
	RefreshTTL         int
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load("../.env"); err != nil {
		return nil, err
	}

	redisDB := os.Getenv("REDIS_DB")

	redisDBint, err := strconv.Atoi(redisDB)
	if err != nil {
		return nil, err
	}

	accessTTL := os.Getenv("ACCESS_TTL")

	accessTTLint, err := strconv.Atoi(accessTTL)
	if err != nil {
		return nil, err
	}

	refreshTTL := os.Getenv("REFRESH_TTL")

	refreshTTLint, err := strconv.Atoi(refreshTTL)
	if err != nil {
		return nil, err
	}

	return &Config{
		PostgresUser:       os.Getenv("POSTGRES_USER"),
		PostgresPassword:   os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost:       os.Getenv("POSTGRES_HOST"),
		PostgresPort:       os.Getenv("POSTGRES_PORT"),
		PostgresAddress:    os.Getenv("POSTGRES_ADDRESS"),
		SSLMode:            os.Getenv("SSL_MODE"),
		ServerAddress:      os.Getenv("SERVER_ADDRESS"),
		LogLevel:           os.Getenv("LOG_LEVEL"),
		LogFilePath:        os.Getenv("LOG_FILE_PATH"),
		RedisAddress:       os.Getenv("REDIS_ADDRESS"),
		RedisPassword:      os.Getenv("REDIS_PASSWORD"),
		RedisDB:            redisDBint,
		RedisHost:          os.Getenv("REDIS_HOST"),
		AccessTTL:          accessTTLint,
		RefreshTTL:         refreshTTLint,
		AccessTokenSecret:  os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret: os.Getenv("REFRESH_TOKEN_SECRET"),
	}, nil
}
