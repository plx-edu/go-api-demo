package db

import (
	"go-api-demo/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/driver/postgres"
)

var Instance *gorm.DB
var err error

func Connect() {
	// env := godotenv.Load()

	// if env  != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	
	Instance, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	// Instance, err := gorm.Open("postgres", os.Getenv("DB_URL"))
	
	if err != nil {
			log.Fatal("Error ", err)
	}
	log.Println("Connected to DB")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.Post{})

	log.Println("DB Migration Completed")
}