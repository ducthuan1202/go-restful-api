package helpers

import "os"

// GetenvWithDefaultValue is function
func GetenvWithDefaultValue(key string, defaultValue string) string {
	if value := os.Getenv(key); value == "" {
		return defaultValue
	} else {
		return value
	}
}

// GetEnv is function
func Getenv(key string) string {
	return os.Getenv(key)
}
