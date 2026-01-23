package service

import (
	"context"
	"errors"
	"github.com/zyy125/my-blog/backend/internal/model"
	"github.com/zyy125/my-blog/backend/internal/repository"
	"gorm.io/gorm"
)

// CommentService 评论业务逻辑层
type CommentService struct {
	repo        *repository.CommentRepository
	articleRepo *repository.ArticleRepository
}

// NewCommentService 创建评论服务实例
func NewCommentService(
	repo *repository.CommentRepository,
	articleRepo *repository. ArticleRepository,
) *CommentService {
	return &CommentService{
		repo:        repo,
		articleRepo: articleRepo,
	}
}

// Create 创建评论
func (s *CommentService) Create(ctx context.Context, comment *model.Comment) error {
	// 1. 验证必填字段
	if comment.ArticleID == 0 {
		return errors.New("文章ID不能为空")
	}
	if comment.Nickname == "" {
		return errors.New("昵称不能为空")
	}
	// 邮箱不再是必填项，允许为空
	if comment.Content == "" {
		return errors.New("评论内容不能为空")
	}
	
	// 2. 验证文章是否存在
	_, err := s.articleRepo.GetByID(ctx, comment.ArticleID)
	if err != nil {
		if errors.Is(err, gorm. ErrRecordNotFound) {
			return errors.New("文章不存在")
		}
		return err
	}
	
	// 3. ✅ 修改：如果是回复，验证父评论是否存在
	if comment.ParentID != nil && *comment.ParentID != 0 {
		parent, err := s.repo.GetByID(ctx, *comment.ParentID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("父评论不存在")
			}
			return err
		}
		
		// 只允许二级评论，不允许三级
		if parent.ParentID != nil && *parent. ParentID != 0 {
			return errors.New("不支持三级评论")
		}
	}
	
	// 4. 设置默认状态（待审核）
	comment.Status = 0
	
	// 5. 创建评论
	return s.repo.Create(ctx, comment)
}

// ListByArticle 获取文章的评论列表
func (s *CommentService) ListByArticle(ctx context.Context, articleID uint) ([]*model.Comment, error) {
	// 验证文章是否存在
	_, err := s.articleRepo.GetByID(ctx, articleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		return nil, err
	}
	
	return s.repo.ListByArticle(ctx, articleID)
}

// ListAll 获取所有评论（管理后台）
func (s *CommentService) ListAll(ctx context.Context, page, pageSize int, status *int8) ([]*model.Comment, int64, error) {
	// 参数验证
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	
	return s.repo.ListAll(ctx, page, pageSize, status)
}

// Approve 审核通过评论
func (s *CommentService) Approve(ctx context.Context, id uint) error {
	// 检查评论是否存在
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("评论不存在")
		}
		return err
	}
	
	// 更新状态为已通过
	return s. repo.UpdateStatus(ctx, id, 1)
}

// Reject 拒绝评论
func (s *CommentService) Reject(ctx context. Context, id uint) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("评论不存在")
		}
		return err
	}
	
	return s.repo.UpdateStatus(ctx, id, 2)
}

// Delete 删除评论
func (s *CommentService) Delete(ctx context.Context, id uint) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("评论不存在")
		}
		return err
	}
	
	return s.repo.Delete(ctx, id)
}