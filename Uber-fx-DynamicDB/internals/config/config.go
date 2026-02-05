package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_DRIVER string
	DB_HOST   string
	DB_USER   string
	DB_PASS   string
	DB_NAME   string
	DB_PORT   string
	APP_PORT  string

	MY_DRIVER  string
	MY_HOST    string
	MY_USER    string
	MY_PASS    string
	MY_DB_NAME string
	MY_PORT    string

	JWT_SECRET string
}

func LoadEnv() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Failed to load env")
	}

	cfg := &Config{
		DB_DRIVER: os.Getenv("DB_DRIVER"),
		DB_HOST:   os.Getenv("DB_HOST"),
		DB_USER:   os.Getenv("DB_USER"),
		DB_PASS:   os.Getenv("DB_PASS"),
		DB_NAME:   os.Getenv("DB_NAME"),
		DB_PORT:   os.Getenv("DB_PORT"),
		APP_PORT:  os.Getenv("APP_PORT"),

		// MySQL
		MY_DRIVER:  os.Getenv("MY_DRIVER"),
		MY_HOST:    os.Getenv("MY_HOST"),
		MY_USER:    os.Getenv("MY_USER"),
		MY_PASS:    os.Getenv("MY_PASS"),
		MY_DB_NAME: os.Getenv("MY_DB_NAME"),
		MY_PORT:    os.Getenv("MY_PORT"),

		//secret key
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}
	log.Println("config loaded :", cfg)
	return cfg
}
