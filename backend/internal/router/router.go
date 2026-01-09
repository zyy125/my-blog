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
	// Repository 层
	articleRepo := repository.NewArticleRepository(database.DB)
	categoryRepo := repository. NewCategoryRepository(database.DB)
	tagRepo := repository. NewTagRepository(database.DB)
	
	// Service 层（注意：ArticleService 现在需要三个依赖）
	articleService := service.NewArticleService(articleRepo, tagRepo, categoryRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	tagService := service.NewTagService(tagRepo)
	
	// Handler 层
	articleHandler := handler.NewArticleHandler(articleService)
	categoryHandler := handler. NewCategoryHandler(categoryService)
	tagHandler := handler.NewTagHandler(tagService)
	
	// ========== 公开 API ==========
	api := r.Group("/api")
	{
		// 文章相关
		api.GET("/articles", articleHandler.List)           // 支持多种筛选
		api.GET("/articles/:id", articleHandler.GetByID)    // 详情（带关联）
		
		// 分类相关
		api.GET("/categories", categoryHandler.List)
		api.GET("/categories/stats", categoryHandler.ListWithCount) // ✅ 新增：带统计
		api.GET("/categories/: id", categoryHandler.GetByID)
		
		// 标签相关
		api.GET("/tags", tagHandler.List)
		api.GET("/tags/stats", tagHandler.ListWithCount)            // ✅ 新增：带统计
		api. GET("/tags/:id", tagHandler.GetByID)
	}
	
	// ========== 管理 API ==========
	admin := r.Group("/api/admin")
	{
		// 文章管理
		admin.POST("/articles", articleHandler.Create)        // 支持标签关联
		admin.PUT("/articles/:id", articleHandler.Update)     // 支持标签更新
		admin.DELETE("/articles/:id", articleHandler.Delete)
		
		// 分类管理
		admin.POST("/categories", categoryHandler.Create)
		admin.PUT("/categories/:id", categoryHandler.Update)
		admin.DELETE("/categories/:id", categoryHandler.Delete)
		
		// 标签管理
		admin.POST("/tags", tagHandler.Create)
		admin.PUT("/tags/:id", tagHandler.Update)
		admin.DELETE("/tags/:id", tagHandler. Delete)
	}
	
	return r
}