package repository

import (
	"context"
	"github.com/zyy125/my-blog/backend/internal/model"
	"gorm.io/gorm"
)

// CategoryRepository 分类数据访问层
type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository 创建分类仓库实例
func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// Create 创建分类
func (r *CategoryRepository) Create(ctx context.Context, category *model. Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

// GetByID 根据ID查询分类
func (r *CategoryRepository) GetByID(ctx context.Context, id uint) (*model.Category, error) {
	var category model.Category
	err := r.db.WithContext(ctx).First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetByName 根据名称查询分类
func (r *CategoryRepository) GetByName(ctx context.Context, name string) (*model.Category, error) {
	var category model.Category
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// List 查询所有分类
func (r *CategoryRepository) List(ctx context.Context) ([]*model.Category, error) {
	var categories []*model. Category
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&categories).Error
	return categories, err
}

// Update 更新分类
func (r *CategoryRepository) Update(ctx context.Context, category *model.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}

// Delete 删除分类
func (r *CategoryRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Category{}, id).Error
}

// CountArticles 统计分类下的文章数量
func (r *CategoryRepository) CountArticles(ctx context.Context, categoryID uint) (int64, error) {
	var count int64
	err := r. db.WithContext(ctx).
		Model(&model.Article{}).
		Where("category_id = ?", categoryID).
		Count(&count).Error
	return count, err
}