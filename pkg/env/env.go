// Package env exposes environment variable values through
// the GetEnvOrDie function
package env

import (
	"fmt"
	"os"
)

func GetEnvOrDie(key string) string {
	value := os.Getenv(key)

	if value == "" {
		err := fmt.Errorf("missing environment variable %s", key)
		panic(err)
	}

	return value
}
