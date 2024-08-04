package config

import (
	"fmt"
	"os"
	"time"

	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	HOST := os.Getenv("POSTGRES_HOST")
	PORT := os.Getenv("POSTGRES_PORT")
	USERNAME := os.Getenv("POSTGRES_USERNAME")
	PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DATABASE := os.Getenv("POSTGRES_DATABASE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", USERNAME, PASSWORD, HOST, PORT, DATABASE)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for sqlDB.Ping() != nil {
		if start.After(start.Add(30 * time.Second)) {
			fmt.Println("Failed to connect Database after 10 seconds")
			break
		}
	}

	fmt.Println("Connected to Database: ", sqlDB.Ping() == nil)

	db.AutoMigrate(&models.User{}, &models.Post{})

	DB = db
}
