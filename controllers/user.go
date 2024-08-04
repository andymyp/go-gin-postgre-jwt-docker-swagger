package controllers

import (
	"net/http"

	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/config"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserProfile(c *gin.Context) {
	var user models.User

	id := c.Param("id")

	if err := config.DB.Where("id=?", id).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": 0, "message": "User not found!"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
			return
		}
	}

	response := models.UserProfile{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": response})
}
