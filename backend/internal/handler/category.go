package handler

import (
	"context"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/internal/model"
	"github.com/zyy125/my-blog/backend/internal/pkg/response"
	"github.com/zyy125/my-blog/backend/internal/service"
)

// CategoryHandler 分类控制器
type CategoryHandler struct {
	service *service.CategoryService
}

// NewCategoryHandler 创建分类控制器实例
func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// Create 创建分类
// POST /api/admin/categories
func (h *CategoryHandler) Create(c *gin.Context) {
	ctx := context.Background()
	
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, "参数格式错误:  "+err.Error())
		return
	}
	
	if err := h.service.Create(ctx, &category); err != nil {
		response.Error(c, err.Error())
		return
	}
	
	response. SuccessWithMsg(c, category, "创建成功")
}

// List 获取分类列表
// GET /api/categories
func (h *CategoryHandler) List(c *gin.Context) {
	ctx := context.Background()
	
	categories, err := h.service.List(ctx)
	if err != nil {
		response.ServerError(c, "查询失败:  "+err.Error())
		return
	}
	
	response.Success(c, categories)
}

// GetByID 获取分类详情
// GET /api/categories/:id
func (h *CategoryHandler) GetByID(c *gin.Context) {
	ctx := context.Background()
	
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response. Error(c, "ID格式错误")
		return
	}
	
	category, err := h.service.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}
	
	response.Success(c, category)
}

// Update 更新分类
// PUT /api/admin/categories/:id
func (h *CategoryHandler) Update(c *gin.Context) {
	ctx := context.Background()
	
	id, err := strconv. ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, "ID格式错误")
		return
	}
	
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		response.Error(c, "参数格式错误: "+err.Error())
		return
	}
	
	category.ID = uint(id)
	
	if err := h.service. Update(ctx, &category); err != nil {
		response.Error(c, err.Error())
		return
	}
	
	response.SuccessWithMsg(c, category, "更新成功")
}

// Delete 删除分类
// DELETE /api/admin/categories/:id
func (h *CategoryHandler) Delete(c *gin.Context) {
	ctx := context.Background()
	
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response. Error(c, "ID格式错误")
		return
	}
	
	if err := h.service.Delete(ctx, uint(id)); err != nil {
		response.Error(c, err.Error())
		return
	}
	
	response. SuccessWithMsg(c, nil, "删除成功")
}