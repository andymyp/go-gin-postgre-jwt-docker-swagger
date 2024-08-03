package routes

import (
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/controllers"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.POST("/auth/register", controllers.Register)
	v1.POST("/auth/login", controllers.Login)
	v1.GET("/auth/test", middlewares.AuthMiddleware(), controllers.TestAuth)
}
