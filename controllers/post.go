package controllers

import (
	"net/http"

	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/config"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/helpers"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func GetPosts(c *gin.Context) {
	var post []models.Post

	if err := config.DB.Debug().Preload("User").Find(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
		return
	}

	var response []models.PostResponse

	for _, post := range post {
		response = append(response, models.PostResponse{
			ID:        post.ID,
			User:      models.UserResponse{ID: post.User.ID, Name: post.User.Name, Email: post.User.Email},
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": response})
}

func GetPost(c *gin.Context) {
	var post models.Post

	id := c.Param("id")

	if err := config.DB.Debug().Preload("User").Where("id=?", id).First(&post).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": 0, "message": "Post not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}
	}

	response := models.PostResponse{
		ID:        post.ID,
		User:      models.UserResponse{ID: post.User.ID, Name: post.User.Name, Email: post.User.Email},
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": response})
}
