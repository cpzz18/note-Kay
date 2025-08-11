package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"note-kay/config"
	"note-kay/models"
)

type TagInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateTag(c *gin.Context) {
	var input TagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uint)
	tag := models.Tag{
		Name:   input.Name,
		UserID: userID,
	}

	if err := config.DB.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tag)
}

func GetTags(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var tags []models.Tag
	if err := config.DB.Where("user_id = ?", userID).Find(&tags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

func UpdateTag(c *gin.Context) {
	id := c.Param("id")
	userID := c.MustGet("user_id").(uint)

	var tag models.Tag
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&tag).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	var input TagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag.Name = input.Name
	config.DB.Save(&tag)
	c.JSON(http.StatusOK, tag)
}

func DeleteTag(c *gin.Context) {
	id := c.Param("id")
	userID := c.MustGet("user_id").(uint)

	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Tag{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted"})
}
