package routes

import (
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/controllers"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.Use(middlewares.AuthMiddleware())
	v1.GET("/user/profile", controllers.UserProfile)
	v1.GET("/user/posts", controllers.GetMyPosts)
}
