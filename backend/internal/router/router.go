package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/internal/handler"
	"github.com/zyy125/my-blog/backend/internal/pkg/database"
	"github.com/zyy125/my-blog/backend/internal/repository"
	"github.com/zyy125/my-blog/backend/internal/service"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	// ========== 初始化依赖 ==========
	// 1. Repository 层
	articleRepo := repository.NewArticleRepository(database.DB)
	
	// 2. Service 层
	articleService := service.NewArticleService(articleRepo)
	
	// 3. Handler 层
	articleHandler := handler.NewArticleHandler(articleService)
	
	// ========== 公开 API ==========
	api := r.Group("/api")
	{
		// 文章相关（公开）
		api.GET("/articles", articleHandler.List)           // 文章列表
		api.GET("/articles/:id", articleHandler. GetByID)    // 文章详情
	}
	
	// ========== 管理 API ==========
	admin := r.Group("/api/admin")
	// TODO: 后续添加认证中间件
	// admin.Use(middleware.AdminAuth())
	{
		// 文章管理
		admin.POST("/articles", articleHandler.Create)        // 创建文章
		admin.PUT("/articles/:id", articleHandler. Update)     // 更新文章
		admin.DELETE("/articles/:id", articleHandler.Delete)  // 删除文章
	}
	
	return r
}