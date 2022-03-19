package helper

import "os"

// MustGetEnv get environment value
func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return ""
	}
	return value
}
