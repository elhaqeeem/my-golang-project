package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
)

func main() {
	// Load konfigurasi
	// config.LoadConfig() // Pastikan Anda memuat konfigurasi jika diperlukan

	// Inisialisasi Gin
	r := gin.Default()

	// Route HTTP/3 utama
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello HTTP/3!")
	})

	// Setup routes
	// router.SetupRoutes(r) // Pastikan Anda mengatur routes jika diperlukan

	// Mengambil sertifikat dan kunci privat dari environment variables
	certPEM := os.Getenv("secret.CERT_PEM")
	keyPEM := os.Getenv("secret.KEY_PEM")

	if certPEM == "" || keyPEM == "" {
		log.Fatal("Sertifikat atau kunci privat tidak ditemukan dalam environment variables")
	}

	// Konfigurasi server HTTP/3
	h3Server := &http3.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Jalankan server dengan HTTP/3 (QUIC)
	log.Println("Starting server on https://localhost:8080 with HTTP/3 support...")
	err := h3Server.ListenAndServeTLS(certPEM, keyPEM)
	if err != nil {
		log.Fatal("Error starting HTTP/3 server:", err)
	}
}
