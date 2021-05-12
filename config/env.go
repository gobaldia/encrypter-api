package config

import (
	"os"
)

func getStringFromEnv(key string, prefix string) string {
	if v, ok := os.LookupEnv(prefix + key); ok {
		return v
	}
	return ""
}

type GlobalConfig struct {
	SecretKey string
}

func GetConfig() GlobalConfig {
	config := GlobalConfig{
		SecretKey: getStringFromEnv("SECRET_KEY", ""),
	}

	return config
}
