package handler

import (
	"context"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/internal/model"
	"github.com/zyy125/my-blog/backend/internal/pkg/response"
	"github.com/zyy125/my-blog/backend/internal/service"
)

// TagHandler 标签控制器
type TagHandler struct {
	service *service.TagService
}

// NewTagHandler 创建标签控制器实例
func NewTagHandler(service *service.TagService) *TagHandler {
	return &TagHandler{service: service}
}

// Create 创建标签
// POST /api/admin/tags
func (h *TagHandler) Create(c *gin.Context) {
	ctx := context.Background()
	
	var tag model.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		response.Error(c, "参数格式错误:  "+err.Error())
		return
	}
	
	if err := h.service.Create(ctx, &tag); err != nil {
		response.Error(c, err.Error())
		return
	}
	
	response.SuccessWithMsg(c, tag, "创建成功")
}

// List 获取标签列表
// GET /api/tags
func (h *TagHandler) List(c *gin.Context) {
	ctx := context.Background()
	
	tags, err := h.service.List(ctx)
	if err != nil {
		response.ServerError(c, "查询失败: "+err.Error())
		return
	}
	
	response. Success(c, tags)
}

// GetByID 获取标签详情
// GET /api/tags/:id
func (h *TagHandler) GetByID(c *gin.Context) {
	ctx := context.Background()
	
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, "ID格式错误")
		return
	}
	
	tag, err := h.service.GetByID(ctx, uint(id))
	if err != nil {
		response.NotFound(c, err. Error())
		return
	}
	
	response.Success(c, tag)
}

// Update 更新标签
// PUT /api/admin/tags/: id
func (h *TagHandler) Update(c *gin.Context) {
	ctx := context.Background()
	
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, "ID格式错误")
		return
	}
	
	var tag model.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		response.Error(c, "参数格式错误: "+err.Error())
		return
	}
	
	tag.ID = uint(id)
	
	if err := h.service. Update(ctx, &tag); err != nil {
		response.Error(c, err.Error())
		return
	}
	
	response.SuccessWithMsg(c, tag, "更新成功")
}

// Delete 删除标签
// DELETE /api/admin/tags/:id
func (h *TagHandler) Delete(c *gin.Context) {
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