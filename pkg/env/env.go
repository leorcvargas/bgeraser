// Package env exposes environment variable values through
// the GetEnvOrDie function
package env

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
)

func GetEnvOrDie(key string) string {
	value := os.Getenv(key)

	if value == "" {
		err := fmt.Errorf("missing environment variable %s", key)
		log.Fatal(err)
	}

	return value
}
