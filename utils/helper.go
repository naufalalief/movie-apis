package utils

import "os"

// GetEnv adalah fungsi untuk mengambil nilai dari environment variable
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
