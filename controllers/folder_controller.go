package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"note-kay/config"
	"note-kay/models"
)

func CreateFolder(c *gin.Context) {
    var input struct {
        Name string `json:"name" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetUint("userID") 
    folder := models.Folder{
        Name:   input.Name,
        UserID: userID,
    }

    if err := config.DB.Create(&folder).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, folder)
}

func GetFolders(c *gin.Context) {
    userID := c.GetUint("userID")
    var folders []models.Folder
    config.DB.Where("user_id = ?", userID).Find(&folders)
    c.JSON(http.StatusOK, folders)
}

func UpdateFolder(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetUint("userID")

    var folder models.Folder
    if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&folder).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Folder not found"})
        return
    }

    var input struct {
        Name string `json:"name" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    folder.Name = input.Name
    config.DB.Save(&folder)
    c.JSON(http.StatusOK, folder)
}

func DeleteFolder(c *gin.Context) {
    id := c.Param("id")
    userID := c.GetUint("userID")

    if err := config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Folder{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Folder deleted"})
}

