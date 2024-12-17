package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/my-golang-project/config"
	"github.com/my-golang-project/db"
	"github.com/my-golang-project/router"
)

func main() {
	// Load konfigurasi
	config.LoadConfig()

	// Inisialisasi database
	db.InitDB()

	// Inisialisasi Gin
	r := gin.Default()

	// Setup HTTP/3 jika diperlukan
	r.Any("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello HTTP/3!")
	})

	// Setup routes
	router.SetupRoutes(r)

	// Jalankan server dengan HTTP/3
	err := r.RunTLS(":8080", "cert.pem", "key.pem") // Port dan path sertifikat
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
