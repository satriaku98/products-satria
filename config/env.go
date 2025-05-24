package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadEnv membaca file .env dan menyimpan ke variabel environment
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}
}

// GetEnv mengambil nilai dari environment variable dengan default value jika kosong
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// GetEnvInt mengambil nilai integer dari environment variable dengan default value jika kosong
func GetEnvInt(key string, defaultValue int) int {
	valueStr := GetEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
