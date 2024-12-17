package main

import (
	"log"
	"net/http"

	"github.com/elhaqeeem/my-golang-project/config"
	"github.com/elhaqeeem/my-golang-project/db"
	"github.com/elhaqeeem/my-golang-project/router"
	"github.com/gin-gonic/gin"
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
	err := r.RunTLS(":8080", "otp", "keyotp") // Port dan path sertifikat
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
