package controllers

import (
	"net/http"

	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/config"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/helpers"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var input struct {
		Title   string `json:"title" validate:"required"`
		Content string `json:"content" validate:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "message": err.Error()})
		return
	}

	if err := helpers.ValidateStruct(input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "message": err.Error()})
		return
	}

	var post models.Post

	user, _ := c.Get("user")
	actualUser, _ := user.(models.UserResponse)
	post.UserID = actualUser.ID
	post.Title = input.Title
	post.Content = input.Content

	if err := config.DB.Create(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Post created."})
}
