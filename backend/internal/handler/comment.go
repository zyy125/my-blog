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

// CommentHandler 评论控制器
type CommentHandler struct {
	service *service.CommentService
}

// NewCommentHandler 创建评论控制器实例
func NewCommentHandler(service *service.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

// Create 创建评论
// POST /api/comments
func (h *CommentHandler) Create(c *gin.Context) {
	ctx := context.Background()

	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, "参数格式错误:  "+err.Error())
		return
	}

	// DTO 转 Model
	comment := &model.Comment{
		ArticleID: req.ArticleID,
		ParentID:  req.ParentID, // 指针类型，可以为 nil
		Nickname:  req.Nickname,
		Email:     req.Email,
		Website:   req.Website,
		Content:   req.Content,
		IP:        c.ClientIP(),
	}

	if err := h.service.Create(ctx, comment); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.SuccessWithMsg(c, comment, "评论提交成功，等待审核")
}

// ListByArticle 获取文章的评论列表
// GET /api/articles/:id/comments
func (h *CommentHandler) ListByArticle(c *gin.Context) {
	ctx := context.Background()

	articleID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, "文章ID格式错误")
		return
	}

	comments, err := h.service.ListByArticle(ctx, uint(articleID))
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, comments)
}

// ListAll 获取所有评论（管理后台）
// GET /api/admin/comments? page=1&page_size=10&status=0
func (h *CommentHandler) ListAll(c *gin.Context) {
	ctx := context.Background()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	var status *int8
	if statusStr := c.Query("status"); statusStr != "" {
		statusInt, err := strconv.Atoi(statusStr)
		if err == nil {
			s := int8(statusInt)
			status = &s
		}
	}

	comments, total, err := h.service.ListAll(ctx, page, pageSize, status)
	if err != nil {
		response.ServerError(c, "查询失败:  "+err.Error())
		return
	}

	response.PageSuccess(c, comments, total, page, pageSize)
}

// Approve 审核通过评论
// PATCH /api/admin/comments/:id/approve
func (h *CommentHandler) Approve(c *gin.Context) {
	ctx := context.Background()

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, "ID格式错误")
		return
	}

	if err := h.service.Approve(ctx, uint(id)); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.SuccessWithMsg(c, nil, "审核通过")
}

// Reject 拒绝评论
// PATCH /api/admin/comments/: id/reject
func (h *CommentHandler) Reject(c *gin.Context) {
	ctx := context.Background()

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, "ID格式错误")
		return
	}

	if err := h.service.Reject(ctx, uint(id)); err != nil {
		response.Error(c, err.Error())
		return
	}

	response.SuccessWithMsg(c, nil, "已拒绝")
}

// Delete 删除评论
// DELETE /api/admin/comments/:id
func (h *CommentHandler) Delete(c *gin.Context) {
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
