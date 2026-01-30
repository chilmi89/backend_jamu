package main

import (
	"backend_jamu/database/seeders"
	"backend_jamu/internal/database"
	"backend_jamu/middleware"
	"backend_jamu/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// @title           Jamu Nusantara API
// @version         1.0
// @description     Simple REST API for Jamu Nusantara Backend.
// @host            localhost:8080
// @BasePath        /api

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Type "Bearer" followed by a space and then your token.

func main() {
	// Load environment variables
	_ = godotenv.Load()

	// Hubungkan ke Database
	database.Connect()
	defer database.DB.Close()

	// Jalankan Seeder
	seeders.SeedAdmin(database.DB)

	// Inisialisasi Router
	mux := routes.SetupRoutes()

	// Global Middleware (CORS)
	handler := middleware.CORS(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("==================================================")
	fmt.Println("      JAMU NUSANTARA - BACKEND API v2.0")
	fmt.Println("==================================================")
	fmt.Printf(" [✓] DB Connected  : Bun ORM (MySQL)\n")
	fmt.Printf(" [✓] Version       : Professional Structure\n")
	fmt.Printf(" [✓] Server Port   : %s\n", port)
	fmt.Println("--------------------------------------------------")
	fmt.Println(" [!] Server is running at http://localhost:" + port)
	fmt.Println(" [!] Press CTRL+C to stop")
	fmt.Println("==================================================")

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
