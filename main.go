package main

import (
	"log"
	"net/http"
	"os"
	"sistem-presensi/api"
	"sistem-presensi/db"
	"sistem-presensi/models"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
	"github.com/rs/cors"
)

// func init() {
// 	// Load .env file
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func main() {

	// Connect to Database
	dbInstance := db.NewDB()
	dbCredential := models.Credential{
		Host:         os.Getenv("DB_HOST"),
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		Port:         getEnvAsInt("DB_PORT", 5432),
		Schema:       os.Getenv("DB_SCHEMA"),
	}

	conn, err := dbInstance.Connect(&dbCredential)
	if err != nil {
		panic(err)
	}

	// Auto-Migrate Models
	conn.AutoMigrate(
		&models.User{}, &models.Mahasiswa{}, &models.Dosen{},
		&models.MataKuliah{}, &models.Pertemuan{}, &models.JadwalKuliah{},
		&models.Presensi{}, &models.PertemuanRekap{}, &models.RekapPresensi{},
		&models.MahasiswaRekap{}, &models.Session{},
	)

	// Initialize Gin Router
	router := gin.Default()

	// Run API Server
	router = api.RunServer(router, conn)

	// Cors handler
	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow your React app's origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))

	// Start the server
	// err = router.Run(":8080")
	// if err != nil {
	// 	panic(err)
	// }
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
