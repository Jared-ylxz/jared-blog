package router

import (
	"jaredBlog/controllers"
	"jaredBlog/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()                                           // create a gin router instance
	router.Use(favicon.New("../front-end-client/public/favicon.ico")) // 后端可以不用设置favicon

	// middleware to handle CORS
	router.Use(middlewares.CORSMiddleware())

	// user routes
	users := router.Group("/api/v1/users")
	users.POST("/login/", controllers.Login)
	users.POST("/register/", controllers.Register)

	// article routes
	articles := router.Group("/api/v1/articles")
	{
		articles.GET("/", controllers.GetArticles) // 如果前端报错CORS，可以尝试在Chrome打开无痕模式
		articles.GET("/:id/", controllers.GetArticleDetail)
		articles.POST("/", middlewares.AuthMiddleware(), controllers.CreateArticle)
		articles.PATCH("/:id/", middlewares.AuthMiddleware(), controllers.UpdateArticle)
		articles.DELETE("/:id/", middlewares.AuthMiddleware(), controllers.DeleteArticle)
	}

	return router
}
