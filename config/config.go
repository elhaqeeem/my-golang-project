package config

import (
	"log"
	"os"
)

// Database URL
var DB_URI string

// LoadConfig mengatur variabel lingkungan
func LoadConfig() {
	DB_URI = os.Getenv("DATABASE_URL")
	if DB_URI == "" {
		DB_URI = "postgres://user:password@localhost:5432/dbname?sslmode=disable"
	}

	log.Println("Database URI:", DB_URI)
}
