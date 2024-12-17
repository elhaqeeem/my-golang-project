package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
)

// Cari port yang tersedia di sistem
func findAvailablePort() (string, error) {
	// Membuka listener di port acak
	listener, err := net.Listen("tcp", ":0") // Port 0 berarti sistem akan memilih port bebas
	if err != nil {
		return "", fmt.Errorf("gagal menemukan port yang tersedia: %v", err)
	}
	defer listener.Close()

	// Mengembalikan port yang digunakan listener
	return listener.Addr().String(), nil
}

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

	// Menemukan port yang tersedia
	port, err := findAvailablePort()
	if err != nil {
		log.Fatal(err)
	}

	// Konfigurasi server HTTP/3
	h3Server := &http3.Server{
		Addr:      port, // Gunakan port yang ditemukan
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
		Handler:   r,
	}

	// Jalankan server dengan HTTP/3 (QUIC)
	log.Printf("Starting server on %s with HTTP/3 support...\n", port)

	// Health check TCP
	// Memeriksa jika server dapat menerima koneksi TCP di port yang ditemukan
	go func() {
		for {
			conn, err := net.Dial("tcp", port)
			if err == nil {
				conn.Close()
				log.Printf("Health check sukses di %s\n", port)
				break
			}
		}
	}()

	// Jalankan server HTTP/3
	err = h3Server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting HTTP/3 server:", err)
	}
}
