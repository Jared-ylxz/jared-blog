package controllers

import (
	"encoding/json"
	"exchangeapp/global"
	"exchangeapp/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var allCacheKey string = "articles"
var oneCacheKey string = "articles:%d"

func CreateArticle(ctx *gin.Context) {
	var article models.Article
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if exists {
		article.AuthorID = userID.(uint)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err := global.DB.Create(&article).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"CreateArticle error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, article)

	if err := global.RDB.Del(ctx, allCacheKey).Err(); err != nil {
		fmt.Println("Redis delete error:", err)
	}
}

func GetArticles(ctx *gin.Context) {
	var articles []models.Article
	redisData, err := global.RDB.Get(ctx, allCacheKey).Result()
	if err == nil && redisData != "" {
		// redis 缓存命中
		fmt.Println("Redis get data!")
		// 将 JSON 字符串反序列化为文章列表
		err := json.Unmarshal([]byte(redisData), &articles)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal articles from cache"})
			return
		}
		ctx.JSON(http.StatusOK, articles)
		return
	} else if err != nil {
		// redis 缓存未命中, 从数据库获取数据并缓存
		fmt.Println("Redis not found:", err)
		result := global.DB.Find(&articles, "deleted_at IS NULL")
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Articles not found"})
			return
		}
		// 将文章列表序列化为 JSON 字符串
		jsonData, err := json.Marshal(articles)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal articles"})
			return
		}
		// 将 JSON 字符串存储到 Redis 中
		statusCmd := global.RDB.Set(ctx, allCacheKey, jsonData, time.Hour*24)
		if statusCmd.Err() != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache in Redis"})
			return
		}
		ctx.JSON(http.StatusOK, articles)
		return
	} else {
		// redis 报错
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}

func GetArticle(ctx *gin.Context) {
	var article models.Article
	id_str := ctx.Param("id")
	id_num, err := strconv.Atoi(id_str)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cacheKey := fmt.Sprintf(oneCacheKey, id_num)
	redisData, err := global.RDB.Get(ctx, cacheKey).Result()
	if err == nil && redisData != "" {
		// redis 缓存命中
		fmt.Println("Redis get data!")
		// 将 JSON 字符串反序列化为文章
		err := json.Unmarshal([]byte(redisData), &article)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal articles from cache"})
			return
		}
		ctx.JSON(http.StatusOK, article)
		return
	} else if err != nil {
		// redis 缓存未命中, 从数据库获取数据并缓存
		fmt.Println("Redis not found:", err.Error())
		result := global.DB.First(&article, "id = ?", id_num)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Articles not found"})
			return
		}
		// 将文章序列化为 JSON 字符串
		jsonData, err := json.Marshal(article)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal articles"})
			return
		}
		// 将 JSON 字符串存储到 Redis 中
		statusCmd := global.RDB.Set(ctx, fmt.Sprintf(oneCacheKey, article.ID), jsonData, time.Hour*24)
		if statusCmd.Err() != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set cache in Redis"})
			return
		}
		ctx.JSON(http.StatusOK, article)
		return
	} else {
		// redis 报错
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}

func DeleteArticle(ctx *gin.Context) {
	var article models.Article
	id := ctx.Param("id")
	result := global.DB.First(&article, "id = ?", id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}
	global.DB.Delete(&article) // 如果一个 model 有 DeletedAt 字段，则软删除。硬删除需要 db.Unscoped().Delete(&article)
	ctx.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})

	if err := global.RDB.Del(ctx, allCacheKey).Err(); err != nil {
		fmt.Println("Redis delete error:", err)
	}
}
