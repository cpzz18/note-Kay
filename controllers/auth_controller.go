package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"note-kay/config"
	"note-kay/models"
	"note-kay/utils"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
    var input RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := utils.HashpPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    user := models.User{Email: input.Email, Password: hashedPassword}
    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
        return
    }

    token, _ := utils.GenerateJWT(user.ID)
    c.JSON(http.StatusOK, gin.H{
        "message": "User registered successfully",
        "user": gin.H{
            "id":    user.ID,
            "email": user.Email,
        },
        "token": token,
    })
}


func Login(c *gin.Context) {
    var input LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if !utils.CheckPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, _ := utils.GenerateJWT(user.ID)
    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "user": gin.H{
            "id":    user.ID,
            "email": user.Email,
        },
        "token": token,
    })
}

