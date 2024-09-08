package utils

import (
	"fmt"
	"os"
)

func GetEnvWithDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func MustLookupEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic(fmt.Errorf("couldn't lookup env variable: %s", key))
}

func LookupEnv(key string) (string, error) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("env variable %q is not set", key)
	}
	return value, nil
}
