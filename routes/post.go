package routes

import (
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/controllers"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/middlewares"
	"github.com/gin-gonic/gin"
)

func PostRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.Use(middlewares.AuthMiddleware())
	v1.POST("/post", controllers.CreatePost)
	v1.GET("/posts", controllers.GetPosts)
	v1.GET("/post/:id", controllers.GetPost)
	v1.PUT("/post/:id", controllers.UpdatePost)
	v1.DELETE("/post/:id", controllers.DeletePost)
}
