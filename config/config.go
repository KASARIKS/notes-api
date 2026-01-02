package config

import "os"

type Config struct {
	PublicHost string
	Port       string
}

var Envs = initConfig()

func initConfig() *Config {
	return &Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
