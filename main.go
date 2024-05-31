package main

import (
	"log"
	"user-management-api/config"
	"user-management-api/models"
	"user-management-api/routes"

	"github.com/joho/godotenv"
)

// Application initialization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
