package config

import (
	"os"
)

type Config struct {
	PostgresConfig PgConfig
}

type PgConfig struct {
	User     string
	NameDb   string
	Password string
	Port     string
}

func GetConfig() (Config, error) {
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "user" // сделал так, потому что всё время вылезает ошибка, а как исправить пока не знаю, хочу в дальнейшим погрузиться в докер
	}
	nameDb := os.Getenv("POSTGRES_DB")
	if nameDb == "" {
		nameDb = "mydatabase"
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "password"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}

	return Config{PostgresConfig: PgConfig{
		User:     user,
		NameDb:   nameDb,
		Password: password,
		Port:     port}}, nil
}
