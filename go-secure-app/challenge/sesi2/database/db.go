package database

import (
	"fmt"
	"log"
	"os"
	"sesi2/models"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	DB_HOST     := os.Getenv("DB_HOST")
	DB_USER     := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME     := os.Getenv("DB_NAME")
	
	PORT     := os.Getenv("DB_PORT")
	DB_PORT, err := strconv.Atoi(PORT)
	if err != nil {
		log.Fatalf("Invalid PORT: %s", err)
	}

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Successfully connected to database.")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}