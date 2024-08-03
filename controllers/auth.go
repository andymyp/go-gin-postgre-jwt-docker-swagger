package controllers

import (
	"net/http"

	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/config"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/helpers"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "message": err.Error()})
		return
	}

	if err := helpers.ValidateStruct(user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "message": err.Error()})
		return
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
		return
	}

	user.Password = hashedPassword

	if err := config.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Register success."})
}
