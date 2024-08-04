package controllers

import (
	"net/http"

	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/config"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserProfile 		godoc
// @Security 			Bearer
// @Summary      	User profile
// @Tags         	User
// @Accept       	json
// @Produce      	json
// @Success      	200 "ok"
// @Router       	/user/profile [get]
func UserProfile(c *gin.Context) {
	getuser, _ := c.Get("user")
	actualUser, _ := getuser.(models.UserResponse)

	var user models.User

	if err := config.DB.Where("id=?", actualUser.ID).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": 0, "message": "User not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}
	}

	response := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": response})
}

// GetMyPosts 		godoc
// @Security 			Bearer
// @Summary      	Get my posts
// @Tags         	User
// @Accept       	json
// @Produce      	json
// @Success      	200 "ok"
// @Router       	/user/posts [get]
func GetMyPosts(c *gin.Context) {
	user, _ := c.Get("user")
	actualUser, _ := user.(models.UserResponse)

	var post []models.Post

	if err := config.DB.Debug().Preload("User").Where("user_id=?", actualUser.ID).Find(&post).Error; err != nil {
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
