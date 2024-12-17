package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
)

func main() {
	// Inisialisasi Gin
	r := gin.Default()

	// Route HTTP/3 utama
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello HTTP/3!")
	})

	// Mengambil sertifikat dan kunci privat dari environment variables
	certPEM := os.Getenv("CERT_PEM")
	keyPEM := os.Getenv("KEY_PEM")

	if certPEM == "" || keyPEM == "" {
		log.Fatal("Sertifikat atau kunci privat tidak ditemukan dalam environment variables")
	}

	// Memuat pasangan sertifikat dan kunci
	cert, err := tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))
	if err != nil {
		log.Fatalf("Gagal memuat pasangan sertifikat dan kunci: %v", err)
	}

	// Konfigurasi server HTTP/3
	h3Server := &http3.Server{
		Addr: ":8080",
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
		Handler: r,
	}

	// Jalankan server dengan HTTP/3 (QUIC)
	log.Println("Starting server on https://localhost:8080 with HTTP/3 support...")
	err = h3Server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting HTTP/3 server:", err)
	}
}
