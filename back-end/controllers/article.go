package controllers

import (
	"encoding/json"
	"fmt"
	"jaredBlog/global"
	"jaredBlog/models"
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
	if err == nil {
		// 如果缓存命中，则直接从缓存中获取数据，解析为文章列表并返回
		err := json.Unmarshal([]byte(redisData), &articles) // 将 JSON 字符串反序列化为文章列表
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal articles from cache"})
			return
		}

		ctx.JSON(http.StatusOK, articles)
		return
	} else {
		// 如果缓存未命中, 则从数据库获取数据并缓存
		result := global.DB.Select("id, title, content").Find(&articles, "deleted_at IS NULL") // Select 仅查询部分字段
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch articles"})
			return
		}

		// 将查询结果转为简单的结构体，仅返回 Title 和 Description
		var responseData []map[string]interface{}
		for _, article := range articles {
			responseData = append(responseData, map[string]interface{}{
				"title":       article.Title,
				"description": article.Description[:100],
			})
		}

		// 将结果存入 Redis，设置缓存过期时间
		jsonData, _ := json.Marshal(responseData)                // 将文章列表序列化为 JSON 字符串
		global.RDB.Set(ctx, allCacheKey, jsonData, time.Hour*24) // 将 JSON 字符串存储到 Redis 中
		ctx.JSON(http.StatusOK, responseData)
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
	if err == nil {
		// 如果缓存命中，则直接从缓存中获取数据，解析为文章列表并返回
		fmt.Println("Redis get data!")
		err := json.Unmarshal([]byte(redisData), &article) // 将 JSON 字符串反序列化为文章
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal articles from cache"})
			return
		}

		ctx.JSON(http.StatusOK, article)
		return
	} else {
		// 如果缓存未命中, 则从数据库获取数据并缓存
		fmt.Println("Redis not found!")
		result := global.DB.First(&article, "id = ?", id_num)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch article"})
			return
		}

		// 将文章序列化为 JSON 字符串
		jsonData, err := json.Marshal(article)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal article to JSON"})
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
