package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MySQLUser         string
	MySQLPassword     string
	MySQLHost         string
	MySQLPort         string
	MySQLDatabase     string
	RedisHost         string
	RedisPort         string
	ServerPort        string
	URLSalt           string
	MaxLenthToShorten int
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	maxLenthToShorten, err := strconv.Atoi(os.Getenv("MAX_LENTH_TO_SHORTEN"))
	if err != nil {
		log.Fatalf("Error converting MAX_LENTH_TO_SHORTEN to int: %v", err)
	}

	return &Config{
		MySQLUser:         os.Getenv("MYSQL_USER"),
		MySQLPassword:     os.Getenv("MYSQL_PASSWORD"),
		MySQLHost:         os.Getenv("MYSQL_HOST"),
		MySQLPort:         os.Getenv("MYSQL_PORT"),
		MySQLDatabase:     os.Getenv("MYSQL_DATABASE"),
		RedisHost:         os.Getenv("REDIS_HOST"),
		RedisPort:         os.Getenv("REDIS_PORT"),
		ServerPort:        os.Getenv("SERVER_PORT"),
		URLSalt:           os.Getenv("URL_SALT"),
		MaxLenthToShorten: maxLenthToShorten,
	}
}
