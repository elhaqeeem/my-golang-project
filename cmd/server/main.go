package main

import (
	"log"
	"net/http"

	"github.com/elhaqeeem/my-golang-project/config"
	"github.com/elhaqeeem/my-golang-project/db"
	"github.com/elhaqeeem/my-golang-project/router"
	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
)

func main() {
	// Load konfigurasi
	config.LoadConfig()

	// Inisialisasi database
	db.InitDB()

	// Inisialisasi Gin
	r := gin.Default()

	// Route HTTP/3 utama
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello HTTP/3!")
	})

	// Setup routes
	router.SetupRoutes(r)

	// Sertifikat dan kunci privat (ganti dengan path file Anda)
	certFile := "secret.CERT_PEM"
	keyFile := "secret.KEY.PEM"

	// Konfigurasi server HTTP/3
	h3Server := &http3.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Jalankan server dengan HTTP/3 (QUIC)
	log.Println("Starting server on https://localhost:8080 with HTTP/3 support...")
	err := h3Server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatal("Error starting HTTP/3 server:", err)
	}
}
