package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"note-kay/config"
	"note-kay/models"
)

type NoteInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Folder  *uint  `json:"folder_id"`
}

func CreateNotes(c *gin.Context) {
    var input NoteInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.MustGet("user_id").(uint)
    
    folderID := uint(0)
    if input.Folder != nil {
        folderID = *input.Folder
    }

    note := models.Note{
        UserID:   userID,
        FolderID: folderID,
        Title:    input.Title,
        Content:  input.Content,
    }

    if err := config.DB.Create(&note).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, note)
}

func GetNotes(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var notes []models.Note

	config.DB.Where("user_id = ?", userID).Find(&notes)
	c.JSON(http.StatusOK, notes)
}

func GetNote(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	id := c.Param("id")

	var note models.Note
	if err := config.DB.Where("user_id = ? AND id = ?", userID, id).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		return
	}
	c.JSON(http.StatusOK, note)
}

func UpdateNote(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	id := c.Param("id")

	var note models.Note
	if err := config.DB.Where("user_id = ? AND id = ?", userID, id).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "note not found"})
		return
	}

	var input NoteInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.Title = input.Title
	note.Content = input.Content
	note.FolderID = *input.Folder

	config.DB.Save(&note)
	c.JSON(http.StatusOK, note)
}

func DeleteNote(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	id := c.Param("id")

	var note models.Note
	if err := config.DB.Where("user_id = ? AND id = ?", userID, id).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	config.DB.Delete(&note)
	c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
