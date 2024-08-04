package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/config"
	_ "github.com/andymyp/go-gin-postgre-jwt-docker-swagger/docs"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//! For Generate Swagger Docs
// @title           							Go Gin API
// @version         							1.0
// @description     							Golang API with Gin, Postgre, JWT, Docker, and Swagger
// @contact.name   								API Support
// @contact.email  								andymyp1997@gmail.com
// @schemes 											http https
// @host      										localhost:3000
// @BasePath  										/api/v1
// @securityDefinitions.apiKey  	Bearer
// @in 														header
// @name 													Authorization
// @description										Enter the token with the `Bearer prefix`, e.g. 'Bearer abcde12345'

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

	//! API Docs
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//! All Routes
	routes.AuthRoute(router)
	routes.UserRoute(router)
	routes.PostRoute(router)

	APP_PORT := os.Getenv("APP_PORT")
	APP_PORT = fmt.Sprintf(":%s", APP_PORT)

	router.Run(APP_PORT)
}
