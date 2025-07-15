package controllers

import (
	"jaredBlog/global"
	"jaredBlog/models"
	"jaredBlog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	// model中，Password 的 json:"-" 会导致在 ctx.ShouldBindJSON(&user) 时不会绑定密码字段。
	// 这会导致注册时，user.Password 为空字符串，进而生成哈希的其实是空字符串的哈希。
	// 因此注册时不要用 models.User 直接绑定 JSON，而是定义单独的注册输入结构体，生成哈希后再赋值给 User 对象。
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: hashedPassword,
	}

	if err := global.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"token": token})
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
	if err := global.DB.Where("username = ?", input_user.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User or password is incorrect"})
		return
	}

	if !utils.VerifyPassword(user.Password, input_user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User or password is incorrect"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "role": user.Role})
}
