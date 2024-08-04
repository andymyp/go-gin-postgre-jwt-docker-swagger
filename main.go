package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/config"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error: .env file not found!")
	}

	config.ConnectDatabase()

	router := gin.Default()

	//! Idle
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API server is running"})
	})

	//! All Routes
	routes.AuthRoute(router)
	routes.UserRoute(router)
	routes.PostRoute(router)

	APP_PORT := os.Getenv("APP_PORT")
	APP_PORT = fmt.Sprintf(":%s", APP_PORT)

	router.Run(APP_PORT)
}
