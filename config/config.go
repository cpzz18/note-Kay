package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"note-kay/models"
)

var DB *gorm.DB

// Load environment variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing with system environment variables...")
	}
}

// Connect to the database
func ConnectDB() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto migrate
	if err := db.AutoMigrate(&models.User{}, &models.Note{}, models.Folder{}, models.Tag{}, models.NoteTag{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
}
