// Package main My-Go-Server
// @title Chrono API
// @version 1.0
// @description API server for therapist-patient management
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by your JWT token, e.g. "Bearer <token>"


package main

import (
	"context"
	"log"
	"os"

	// "github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	// "github.com/nicksnyder/go-i18n/v2/i18n"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	// "golang.org/x/text/language"

	"my-go-server/repositories"
	"my-go-server/routers"
	"my-go-server/services"
)

func main() {
	// Load environment variables from .env if present (not fatal)
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, falling back to environment variables")
	}

	// // Determine locales path - try relative first, then server/
	// localesPath := "locales/"
	// if _, err := os.Stat("locales/en.toml"); os.IsNotExist(err) {
	// 	localesPath = "server/locales/"
	// }

	// // Initialize i18n bundle
	// bundle := i18n.NewBundle(language.English)
	// bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// bundle.MustLoadMessageFile(localesPath + "en.toml")
	// bundle.MustLoadMessageFile(localesPath + "he.toml")

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI is not set")
		return
	}

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
		return
	}

	// Ping MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("failed to ping MongoDB: %v", err)
		return
	}

	Database := client.Database("chrono")
	log.Printf("Connected to MongoDB database: %s", Database.Name())

	// Initialize MongoDB repositories
	userRepo := repositories.NewMongoUserRepository(client, "chrono")
	clinicRepo := repositories.NewMongoClinicRepository(client, "chrono")
	refreshTokenRepo := repositories.NewMongoRefreshTokenRepository(client, "chrono")

	// Initialize services
	userService := services.NewUserService(userRepo)
	clinicService := services.NewClinicService(clinicRepo)
	authService := services.NewAuthService(userRepo, refreshTokenRepo)

	// Setup and run router
	router := routers.SetupRouter(userService, authService,clinicService)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
		return
	}
}
