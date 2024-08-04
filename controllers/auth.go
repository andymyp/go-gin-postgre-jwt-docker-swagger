package controllers

import (
	"net/http"

	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/config"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/helpers"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/models"
	"github.com/gin-gonic/gin"
)

// Register 			godoc
// @Summary      	Register user
// @Tags         	Auth
// @Accept       	json
// @Produce      	json
// @Param        	request body models.UserRegister true "Payload [Raw]"
// @Success      	200 "ok"
// @Router       	/auth/register [post]
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

// Login 					godoc
// @Summary      	Login user
// @Tags         	Auth
// @Accept       	json
// @Produce      	json
// @Param        	request body models.UserLogin true "Payload [Raw]"
// @Success      	200 "ok"
// @Router       	/auth/login [post]
func Login(c *gin.Context) {
	var input models.UserLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "message": err.Error()})
		return
	}

	if err := helpers.ValidateStruct(input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "message": err.Error()})
		return
	}

	var user models.User

	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 0, "message": "Email wrong!"})
		return
	}

	if !helpers.CheckPassword(user.Password, input.Password) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 0, "message": "Password wrong!"})
		return
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": 0, "message": "Failed to generate token!"})
		return
	}

	user.Token = token

	if err := config.DB.Save(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": 0, "message": err.Error()})
		return
	}

	response := models.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: user.Token,
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": response})
}

// TestAuth 			godoc
// @Security 			Bearer
// @Summary      	Test auth
// @Tags         	Auth
// @Accept       	json
// @Produce      	json
// @Success      	200 "ok"
// @Router       	/auth/test [get]
func TestAuth(c *gin.Context) {
	user, _ := c.Get("user")
	claims, _ := c.Get("claims")
	c.JSON(http.StatusOK, gin.H{"status": 1, "user": user, "claims": claims})
}
