package system

import (
	"os"
)

func mustHaveEnv(key string) string {
	env := os.Getenv("ENV")
	value := os.Getenv(key)
	if value == "" && env != "test" {
		panic("Missing environment variable: " + key)
	}
	return value
}

var (
    LOG_LEVEL            = mustHaveEnv("LOG_LEVEL")
	HTTP_PORT            = mustHaveEnv("HTTP_PORT")
)
