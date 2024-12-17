package config

import (
	"log"
	"os"
)

// Database URL
var DB_URI string

// LoadConfig mengatur variabel lingkungan
func LoadConfig() {
	// Memuat variabel dari file .env
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Error loading .env file: %v", err)
	//}

	// Ambil DATABASE_URL dari environment
	DB_URI = os.Getenv("DATABASE_URL")
	if DB_URI == "" {
		log.Fatal("DATABASE_URL is not set in environment")
	}

	log.Println("Database URI:", DB_URI)
}
