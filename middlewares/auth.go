package middlewares

import (
	"net/http"
	"os"

	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/config"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKeyEnv = os.Getenv("JWT_SECRET_KEY")
var jwtSecretKey = []byte(secretKeyEnv)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 0, "message": "Authorization header is missing!"})
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 0, "message": "Invalid or expired token!"})
			return
		}

		claims, ok := token.Claims.(*models.Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 0, "message": "Invalid token claims!"})
			return
		}

		var user models.User

		if err := config.DB.Where("id=? AND token=?", claims.ID, tokenString).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": 0, "message": "Token not found or expired"})
			return
		}

		response := models.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Token: user.Token,
		}

		c.Set("user", response)
		c.Set("claims", claims)
		c.Next()
	}
}
