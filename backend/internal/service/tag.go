package service

import (
	"context"
	"errors"
	"github.com/zyy125/my-blog/backend/internal/model"
	"github.com/zyy125/my-blog/backend/internal/repository"
	"gorm.io/gorm"
)

// TagService 标签业务逻辑层
type TagService struct {
	repo *repository. TagRepository
}

// NewTagService 创建标签服务实例
func NewTagService(repo *repository.TagRepository) *TagService {
	return &TagService{repo: repo}
}

// Create 创建标签
func (s *TagService) Create(ctx context.Context, tag *model.Tag) error {
	if tag.Name == "" {
		return errors.New("标签名称不能为空")
	}
	
	// 检查是否已存在同名标签
	existing, err := s.repo.GetByName(ctx, tag.Name)
	if err == nil && existing != nil {
		return errors.New("标签名称已存在")
	}
	
	return s.repo.Create(ctx, tag)
}

// GetByID 获取标签详情
func (s *TagService) GetByID(ctx context.Context, id uint) (*model.Tag, error) {
	tag, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors. New("标签不存在")
		}
		return nil, err
	}
	return tag, nil
}

// List 获取所有标签
func (s *TagService) List(ctx context.Context) ([]*model.Tag, error) {
	return s.repo.List(ctx)
}

// Update 更新标签
func (s *TagService) Update(ctx context. Context, tag *model.Tag) error {
	_, err := s.repo.GetByID(ctx, tag.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors. New("标签不存在")
		}
		return err
	}
	
	if tag.Name == "" {
		return errors.New("标签名称不能为空")
	}
	
	// 检查名称是否与其他标签重复
	existing, err := s.repo.GetByName(ctx, tag.Name)
	if err == nil && existing != nil && existing.ID != tag.ID {
		return errors.New("标签名称已存在")
	}
	
	return s.repo.Update(ctx, tag)
}

// Delete 删除标签
func (s *TagService) Delete(ctx context.Context, id uint) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("标签不存在")
		}
		return err
	}
	
	// 注意：删除标签不检查文章关联
	// 因为多对多关系，删除标签时会自动删除关联关系（CASCADE）
	return s.repo.Delete(ctx, id)
}

// ListWithCount 获取标签列表（带文章数量）
func (s *TagService) ListWithCount(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo. ListWithArticleCount(ctx)
}