package controllers

import (
	"net/http"

	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/config"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/helpers"
	"github.com/andymyp/go-gin-postgre-jwt-docker-swagger/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreatePost 		godoc
// @Security 			Bearer
// @Summary      	Create post
// @Tags         	User
// @Accept       	json
// @Produce      	json
// @Param        	request body models.InputPost true "Payload [Raw]"
// @Success      	200 "ok"
// @Router       	/post [post]
func CreatePost(c *gin.Context) {
	var input models.InputPost

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

// GetPosts 			godoc
// @Security 			Bearer
// @Summary      	Get all posts
// @Tags         	Post
// @Accept       	json
// @Produce      	json
// @Success      	200 "ok"
// @Router       	/posts [get]
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

// GetPost 				godoc
// @Security 			Bearer
// @Summary      	Get post
// @Tags         	Post
// @Accept       	json
// @Produce      	json
// @Param        	id path string true "Post ID"
// @Success      	200 "ok"
// @Router       	/post/{id} [get]
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

// UpdatePost 		godoc
// @Security 			Bearer
// @Summary      	Update post
// @Tags         	User
// @Accept       	json
// @Produce      	json
// @Param        	id path string true "Post ID"
// @Param        	request body models.InputPost true "Payload [Raw]"
// @Success      	200 "ok"
// @Router       	/post/{id} [put]
func UpdatePost(c *gin.Context) {
	var input models.InputPost

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "message": err.Error()})
		return
	}

	if err := helpers.ValidateStruct(input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "message": err.Error()})
		return
	}

	user, _ := c.Get("user")
	actualUser, _ := user.(models.UserResponse)

	var post models.Post

	id := c.Param("id")

	if config.DB.Model(&post).Where("id=? AND user_id=?", id, actualUser.ID).Updates(&input).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "message": "Update post failed!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Post updated."})
}

// DeletePost 		godoc
// @Security 			Bearer
// @Summary      	Delete post
// @Tags         	User
// @Accept       	json
// @Produce      	json
// @Param        	id path string true "Post ID"
// @Success      	200 "ok"
// @Router       	/post/{id} [delete]
func DeletePost(c *gin.Context) {
	user, _ := c.Get("user")
	actualUser, _ := user.(models.UserResponse)

	var post models.Post

	id := c.Param("id")

	if config.DB.Where("id=? AND user_id=?", id, actualUser.ID).Delete(&post).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": 0, "message": "Delete post failed!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Post deleted."})
}

// SearchPosts 		godoc
// @Security 			Bearer
// @Summary      	Search posts
// @Tags         	Post
// @Accept       	json
// @Produce      	json
// @Param        	query query string true "Search"
// @Success      	200 "ok"
// @Router       	/posts/search [get]
func SearchPosts(c *gin.Context) {
	query := c.Query("query")

	var post []models.Post

	if err := config.DB.Debug().
		Preload("User").
		Where("title ILIKE ? OR content ILIKE ?", "%"+query+"%", "%"+query+"%").
		Find(&post).Error; err != nil {
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
