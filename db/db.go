package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/my-golang-project/config"
)

var DB *sql.DB

// InitDB menginisialisasi koneksi ke PostgreSQL
func InitDB() {
	var err error
	DB, err = sql.Open("postgres", config.DB_URI)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}
	log.Println("Database connected successfully")
}
