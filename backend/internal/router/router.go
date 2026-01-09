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
	categoryRepo := repository. NewCategoryRepository(database.DB)  // ✅ 新增
	tagRepo := repository.NewTagRepository(database.DB)            // ✅ 新增
	
	// Service 层
	articleService := service.NewArticleService(articleRepo)
	categoryService := service.NewCategoryService(categoryRepo)    // ✅ 新增
	tagService := service.NewTagService(tagRepo)                   // ✅ 新增
	
	// Handler 层
	articleHandler := handler.NewArticleHandler(articleService)
	categoryHandler := handler.NewCategoryHandler(categoryService) // ✅ 新增
	tagHandler := handler.NewTagHandler(tagService)                // ✅ 新增
	
	// ========== 公开 API ==========
	api := r.Group("/api")
	{
		// 文章相关
		api.GET("/articles", articleHandler.List)
		api.GET("/articles/: id", articleHandler.GetByID)
		
		// ✅ 分类相关（公开）
		api.GET("/categories", categoryHandler.List)
		api.GET("/categories/:id", categoryHandler.GetByID)
		
		// ✅ 标签相关（公开）
		api.GET("/tags", tagHandler.List)
		api.GET("/tags/:id", tagHandler.GetByID)
	}
	
	// ========== 管理 API ==========
	admin := r.Group("/api/admin")
	{
		// 文章管理
		admin.POST("/articles", articleHandler.Create)
		admin.PUT("/articles/:id", articleHandler. Update)
		admin.DELETE("/articles/:id", articleHandler.Delete)
		
		// ✅ 分类管理
		admin.POST("/categories", categoryHandler.Create)
		admin.PUT("/categories/:id", categoryHandler.Update)
		admin.DELETE("/categories/:id", categoryHandler. Delete)
		
		// ✅ 标签管理
		admin.POST("/tags", tagHandler.Create)
		admin.PUT("/tags/:id", tagHandler.Update)
		admin.DELETE("/tags/:id", tagHandler.Delete)
	}
	
	return r
}