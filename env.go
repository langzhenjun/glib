package utils

import (
	"os"
	"strings"
)

func GetEnvOrDefault(key string, def string) string {
	env := os.Getenv(key)
	if "" == env {
		return def
	}
	return env
}

func IsDevelopmentMode() bool {
	env := GetEnvOrDefault("ENVIRONMENT", "DEVELOPMENT")
	env = strings.ToUpper(env)
	return "DEVELOPMENT" == env
}

func IsProductionMode() bool {
	env := GetEnvOrDefault("ENVIRONMENT", "")
	env = strings.ToUpper(env)
	return "PRODUCTION" == env
}
