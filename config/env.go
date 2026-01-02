package config

import (
	"os"
	"strconv"
)

type Config struct {
	PublicHost string
	Port       string

	DBName string

	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() *Config {
	return &Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBName:                 getEnv("DB_NAME", "db.db"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "necesserily_set_jwt_secret_in_env"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return v
	}

	return fallback
}
