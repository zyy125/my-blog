package handler

import (
	"context"
	
	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/internal/model"
	"github.com/zyy125/my-blog/backend/internal/pkg/database"
	"github.com/zyy125/my-blog/backend/internal/pkg/response"
)

// StatsHandler 统计控制器
type StatsHandler struct{}

// NewStatsHandler 创建统计控制器实例
func NewStatsHandler() *StatsHandler {
	return &StatsHandler{}
}

// GetDashboard 获取后台首页统计数据
// GET /api/admin/stats
func (h *StatsHandler) GetDashboard(c *gin.Context) {
	ctx := context.Background()
	db := database.DB. WithContext(ctx)
	
	// 1. 统计文章数量
	var articleCount int64
	db. Model(&model.Article{}).Count(&articleCount)
	
	// 2. 统计已发布文章数量
	var publishedCount int64
	db.Model(&model.Article{}).Where("status = 1").Count(&publishedCount)
	
	// 3. 统计分类数量
	var categoryCount int64
	db.Model(&model.Category{}).Count(&categoryCount)
	
	// 4. 统计标签数量
	var tagCount int64
	db.Model(&model.Tag{}).Count(&tagCount)
	
	// 5. 统计评论数量
	var commentCount int64
	db.Model(&model. Comment{}).Count(&commentCount)
	
	// 6. 统计待审核评论数量
	var pendingCommentCount int64
	db. Model(&model.Comment{}).Where("status = 0").Count(&pendingCommentCount)
	
	// 7. 统计总浏览量
	var totalViews int64
	db.Model(&model.Article{}).Select("COALESCE(SUM(views), 0)").Scan(&totalViews)
	
	// 8. 返回统计数据
	response.Success(c, gin.H{
		"article_count":          articleCount,
		"published_count":        publishedCount,
		"category_count":         categoryCount,
		"tag_count":              tagCount,
		"comment_count":          commentCount,
		"pending_comment_count":   pendingCommentCount,
		"total_views":            totalViews,
	})
}