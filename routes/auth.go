package routes

import (
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.POST("/auth/register", controllers.Register)
	v1.POST("/auth/login", controllers.Login)
}
