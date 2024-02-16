package controllers

import (
	"fmt"
	"net/http"
	"project-mygram/database"
	"project-mygram/helpers"
	"project-mygram/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	newComment := models.Comment{}
	userId := uint(userData["id"].(float64))
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	if contentType == appJSON {
		c.ShouldBindJSON(&newComment)
	} else {
		c.ShouldBind(&newComment)
	}

	newComment.UserID = userId
	newComment.PhotoID = uint(photoId)
	err := db.Debug().Create(&newComment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "Bad Request",
		})
		return
	}

	c.JSON(http.StatusCreated, newComment)
}

func GetAllComments(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	userId := uint(userData["id"].(float64))
	var comments []models.Comment

	// Retrieve all comments associated with the user
	err := db.Preload("User").Preload("Photo").Where("user_id = ?", userId).Find(&comments).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed to get data",
			"error_message": "Something went wrong when trying to get data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}

func GetComment(c *gin.Context) {
	db := database.GetDB()

	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "Invalid comment ID",
		})
		return
	}

	comment := models.Comment{}
	err = db.Preload("User").Preload("Photo").First(&comment, commentID).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Comment not found",
			"error_message": fmt.Sprintf("Comment with ID %v not found", commentID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.First(&Comment, "id = ?", commentId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Delete(&Comment)

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("comment with id %v has been successfully deleted", commentId),
	})
}
