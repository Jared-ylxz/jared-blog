package controllers

import (
	"jaredBlog/global"
	"jaredBlog/models"
	"jaredBlog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = hashedPassword

	if err := global.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user_count int64
	global.DB.Model(&models.User{}).Count(&user_count)
	if user_count == 1 || user.Username == "Jared" {
		user.Role = 9
	}
	if err := global.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change user role."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func Login(ctx *gin.Context) {
	var input_user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&input_user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := global.DB.Where("username = ?", input_user.Username).Find(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User or password is incorrect"})
		return
	}

	if !utils.VerifyPassword(user.Password, input_user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User or password is incorrect"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
