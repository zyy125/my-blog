package service

import (
	"context"
	"errors"
	"github.com/zyy125/my-blog/backend/internal/model"
	"github.com/zyy125/my-blog/backend/internal/repository"
	"gorm.io/gorm"
)

// CategoryService 分类业务逻辑层
type CategoryService struct {
	repo *repository.CategoryRepository
}

// NewCategoryService 创建分类服务实例
func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// Create 创建分类
func (s *CategoryService) Create(ctx context.Context, category *model.Category) error {
	// 1. 验证
	if category.Name == "" {
		return errors.New("分类名称不能为空")
	}
	
	// 2. 检查是否已存在同名分类
	existing, err := s.repo.GetByName(ctx, category.Name)
	if err == nil && existing != nil {
		return errors.New("分类名称已存在")
	}
	
	// 3. 创建
	return s.repo.Create(ctx, category)
}

// GetByID 获取分类详情
func (s *CategoryService) GetByID(ctx context. Context, id uint) (*model.Category, error) {
	category, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors. New("分类不存在")
		}
		return nil, err
	}
	return category, nil
}

// List 获取所有分类
func (s *CategoryService) List(ctx context.Context) ([]*model.Category, error) {
	return s.repo.List(ctx)
}

// Update 更新分类
func (s *CategoryService) Update(ctx context.Context, category *model.Category) error {
	// 1. 检查是否存在
	_, err := s.repo.GetByID(ctx, category.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors. New("分类不存在")
		}
		return err
	}
	
	// 2. 验证
	if category.Name == "" {
		return errors.New("分类名称不能为空")
	}
	
	// 3. 检查名称是否与其他分类重复
	existing, err := s.repo.GetByName(ctx, category.Name)
	if err == nil && existing != nil && existing.ID != category.ID {
		return errors.New("分类名称已存在")
	}
	
	// 4. 更新
	return s.repo.Update(ctx, category)
}

// Delete 删除分类
func (s *CategoryService) Delete(ctx context.Context, id uint) error {
	// 1. 检查是否存在
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("分类不存在")
		}
		return err
	}
	
	// 2. 检查是否有文章使用该分类
	count, err := s.repo.CountArticles(ctx, id)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors. New("该分类下有文章，无法删除")
	}
	
	// 3. 删除
	return s.repo.Delete(ctx, id)
}