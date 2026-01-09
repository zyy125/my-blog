package handler

import (
	"context"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/internal/handler/dto"
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

// Create 创建文章（重写：支持标签）
// POST /api/admin/articles
func (h *ArticleHandler) Create(c *gin.Context) {
	ctx := context.Background()
	
	// 1. 绑定请求参数到 DTO
	var req dto.CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, "参数格式错误:  "+err.Error())
		return
	}
	
	// 2. DTO 转 Model
	article := &model. Article{
		Title:      req.Title,
		Content:    req.Content,
		Summary:    req.Summary,
		CoverImg:   req.CoverImg,
		CategoryID: req.CategoryID,
		Status:     req.Status,
		IsTop:      req.IsTop,
	}
	
	// 3. 调用 Service 创建（带标签）
	if err := h.service.CreateWithTags(ctx, article, req.TagIDs); err != nil {
		response. Error(c, err.Error())
		return
	}
	
	// 4. 返回成功响应
	response. SuccessWithMsg(c, article, "创建成功")
}

// Update 更新文章（重写：支持标签）
// PUT /api/admin/articles/:id
func (h *ArticleHandler) Update(c *gin.Context) {
	ctx := context.Background()
	
	// 1. 解析 ID
	id, err := strconv.ParseUint(c. Param("id"), 10, 32)
	if err != nil {
		response.Error(c, "ID格式错误")
		return
	}
	
	// 2. 绑定请求参数
	var req dto.UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, "参数格式错误: "+err.Error())
		return
	}
	
	// 3. DTO 转 Model
	article := &model.Article{
		ID:         uint(id),
		Title:      req.Title,
		Content:    req.Content,
		Summary:    req.Summary,
		CoverImg:   req.CoverImg,
		CategoryID:  req.CategoryID,
		Status:     req.Status,
		IsTop:      req.IsTop,
	}
	
	// 4. 调用 Service 更新（带标签）
	if err := h.service.UpdateWithTags(ctx, article, req. TagIDs); err != nil {
		response.Error(c, err.Error())
		return
	}
	
	// 5. 返回成功响应
	response.SuccessWithMsg(c, article, "更新成功")
}

// List 获取文章列表（重写：支持高级筛选）
// GET /api/articles? page=1&page_size=10&category_id=1&tag_id=2&keyword=Go
func (h *ArticleHandler) List(c *gin.Context) {
	ctx := context.Background()
	
	// 1. 绑定查询参数
	var query dto.ArticleListQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.Error(c, "参数格式错误: "+err.Error())
		return
	}
	
	// 2. 设置默认值
	if query.Page < 1 {
		query. Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 10
	}
	
	var articles []*model.Article
	var total int64
	var err error
	
	// 3. 根据不同条件查询
	if query.CategoryID != nil {
		// 按分类查询
		articles, total, err = h. service.ListByCategory(ctx, *query.CategoryID, query.Page, query.PageSize)
	} else if query.TagID != nil {
		// 按标签查询
		articles, total, err = h.service.ListByTag(ctx, *query. TagID, query.Page, query.PageSize)
	} else if query.Keyword != "" {
		// 搜索
		articles, total, err = h.service.Search(ctx, query.Keyword, query.Page, query.PageSize)
	} else {
		// 普通列表查询
		articles, total, err = h.service.List(ctx, query.Page, query.PageSize, query.Status)
	}
	
	if err != nil {
		response. ServerError(c, "查询失败:  "+err.Error())
		return
	}
	
	// 4. 返回分页响应
	response.PageSuccess(c, articles, total, query.Page, query.PageSize)
}

// GetByID 获取文章详情（保持不变，但使用关联查询）
// GET /api/articles/:id
func (h *ArticleHandler) GetByID(c *gin. Context) {
	ctx := context.Background()
	
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, "ID格式错误")
		return
	}
	
	article, err := h.service.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}
	
	response.Success(c, article)
}

// Delete 删除文章（保持不变）
// DELETE /api/admin/articles/:id
func (h *ArticleHandler) Delete(c *gin.Context) {
	ctx := context.Background()
	
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, "ID格式错误")
		return
	}
	
	if err := h.service.Delete(ctx, uint(id)); err != nil {
		response.Error(c, err.Error())
		return
	}
	
	response.SuccessWithMsg(c, nil, "删除成功")
}