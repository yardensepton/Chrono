package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"my-go-server/repositories"
	"my-go-server/routers"
	"my-go-server/services"
	"my-go-server/users"
)

func main() {
	// Load environment variables from .env if present (not fatal)
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, falling back to environment variables")
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL is not set")
		return
	}

	// Connect DB (GORM)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
		return
	}

	err = db.AutoMigrate(&users.User{})  // Add other models here if needed
	if err != nil {
		log.Fatalf("failed to migrate tables: %v", err)
		return
	}

	// Initialize repositories that need the DB
	userRepo := repositories.NewPostgresUserRepository(db)
	usersService := services.NewUserService(userRepo)

	// Setup and run router
	router := routeres.SetupRouter(usersService)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
		return
	}
}
