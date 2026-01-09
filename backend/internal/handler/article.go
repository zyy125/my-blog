package handler

import (
	"context"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/internal/model"
	"github.com/zyy125/my-blog/backend/internal/pkg/response"
	"github.com/zyy125/my-blog/backend/internal/service"
)

// ArticleHandler 文章控制器
type ArticleHandler struct {
	service *service.ArticleService
}

// NewArticleHandler 创建文章控制器实例
func NewArticleHandler(service *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: service}
}

// Create 创建文章
// POST /api/admin/articles
func (h *ArticleHandler) Create(c *gin.Context) {
	ctx := context.Background()
	
	// 1. 绑定请求参数
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		response.Error(c, "参数格式错误:  "+err.Error())
		return
	}
	
	// 2. 调用 Service 创建
	if err := h.service.Create(ctx, &article); err != nil {
		response.Error(c, err.Error())
		return
	}
	
	// 3. 返回成功响应
	response. SuccessWithMsg(c, article, "创建成功")
}

// List 获取文章列表
// GET /api/articles? page=1&page_size=10&status=1
func (h *ArticleHandler) List(c *gin.Context) {
	ctx := context.Background()
	
	// 1. 解析查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	
	// 解析状态参数（可选）
	var status *int8
	if statusStr := c.Query("status"); statusStr != "" {
		statusInt, err := strconv.Atoi(statusStr)
		if err == nil {
			s := int8(statusInt)
			status = &s
		}
	}
	
	// 2. 调用 Service 查询
	articles, total, err := h.service.List(ctx, page, pageSize, status)
	if err != nil {
		response.ServerError(c, "查询失败:  "+err.Error())
		return
	}
	
	// 3. 返回分页响应
	response.PageSuccess(c, articles, total, page, pageSize)
}

// GetByID 获取文章详情
// GET /api/articles/:id
func (h *ArticleHandler) GetByID(c *gin.Context) {
	ctx := context.Background()
	
	// 1. 解析 ID 参数
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response. Error(c, "ID 格式错误")
		return
	}
	
	// 2. 调用 Service 查询
	article, err := h.service.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}
	
	// 3. 返回响应
	response.Success(c, article)
}

// Update 更新文章
// PUT /api/admin/articles/:id
func (h *ArticleHandler) Update(c *gin.Context) {
	ctx := context.Background()
	
	// 1. 解析 ID
	id, err := strconv. ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, "ID 格式错误")
		return
	}
	
	// 2. 绑定请求体
	var article model.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		response.Error(c, "参数格式错误: "+err.Error())
		return
	}
	
	// 3. 设置 ID
	article.ID = uint(id)
	
	// 4. 调用 Service 更新
	if err := h.service.Update(ctx, &article); err != nil {
		response.Error(c, err.Error())
		return
	}
	
	// 5. 返回响应
	response. SuccessWithMsg(c, article, "更新成功")
}

// Delete 删除文章
// DELETE /api/admin/articles/:id
func (h *ArticleHandler) Delete(c *gin.Context) {
	ctx := context.Background()
	
	// 1. 解析 ID
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response. Error(c, "ID 格式错误")
		return
	}
	
	// 2. 调用 Service 删除
	if err := h.service.Delete(ctx, uint(id)); err != nil {
		response.Error(c, err.Error())
		return
	}
	
	// 3. 返回响应
	response.SuccessWithMsg(c, nil, "删除成功")
}