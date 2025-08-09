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

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing..")
	}
func ConnectDB() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn) , &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&models.User{}, &models.Note{},)

	DB = db
	fmt.Println("Database connected successfully")
}
