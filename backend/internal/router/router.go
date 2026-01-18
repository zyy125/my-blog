package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/internal/handler"
	"github.com/zyy125/my-blog/backend/internal/middleware"
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
	categoryRepo := repository.NewCategoryRepository(database.DB)
	tagRepo := repository.NewTagRepository(database.DB)
	commentRepo := repository.NewCommentRepository(database.DB)

	// Service 层
	articleService := service.NewArticleService(articleRepo, tagRepo, categoryRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	tagService := service.NewTagService(tagRepo)
	commentService := service.NewCommentService(commentRepo, articleRepo)

	// Handler 层
	articleHandler := handler.NewArticleHandler(articleService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	tagHandler := handler.NewTagHandler(tagService)
	commentHandler := handler.NewCommentHandler(commentService)
	uploadHandler := handler.NewUploadHandler()
	statsHandler := handler.NewStatsHandler()
	authHandler := handler.NewAuthHandler()

	// ========== 静态文件服务 ==========
	r.Static("/uploads", "./uploads") // 提供上传文件访问

	// ========== 公开 API ==========
	api := r.Group("/api")
	{
		// 文章相关
		api.GET("/articles", articleHandler.List)
		api.GET("/articles/:id", articleHandler.GetByID)

		// 分类相关
		api.GET("/categories", categoryHandler.List)
		api.GET("/categories/stats", categoryHandler.ListWithCount)
		api.GET("/categories/:id", categoryHandler.GetByID)

		// 标签相关
		api.GET("/tags", tagHandler.List)
		api.GET("/tags/stats", tagHandler.ListWithCount)
		api.GET("/tags/:id", tagHandler.GetByID)

		// 评论相关（公开）
		api.GET("/articles/:id/comments", commentHandler.ListByArticle) // 查看评论
		api.POST("/comments", commentHandler.Create)                    // 提交评论
	}

	// ========== 管理认证 API (公开，无需 IP 验证) ==========
	adminAuth := r.Group("/api/admin/auth")
	{
		adminAuth.POST("/verify", authHandler.VerifyKey) // 验证密钥并获取token
	}

	// ========== 管理 API（Token认证）==========
	admin := r.Group("/api/admin")
	admin.Use(middleware.AdminAuth()) // 仅 Token 验证
	{
		// 文章管理
		admin.POST("/articles", articleHandler.Create)
		admin.PUT("/articles/:id", articleHandler.Update)
		admin.DELETE("/articles/:id", articleHandler.Delete)

		// 分类管理
		admin.POST("/categories", categoryHandler.Create)
		admin.PUT("/categories/:id", categoryHandler.Update)
		admin.DELETE("/categories/:id", categoryHandler.Delete)

		// 标签管理
		admin.POST("/tags", tagHandler.Create)
		admin.PUT("/tags/:id", tagHandler.Update)
		admin.DELETE("/tags/:id", tagHandler.Delete)

		// 评论管理
		admin.GET("/comments", commentHandler.ListAll)               // 所有评论
		admin.PATCH("/comments/:id/approve", commentHandler.Approve) // 审核通过
		admin.PATCH("/comments/:id/reject", commentHandler.Reject)   // 拒绝
		admin.DELETE("/comments/:id", commentHandler.Delete)         // 删除

		// 文件上传
		admin.POST("/upload/image", uploadHandler.UploadImage) // 上传图片

		// 统计数据
		admin.GET("/stats", statsHandler.GetDashboard) // 后台统计
	}

	return r
}
