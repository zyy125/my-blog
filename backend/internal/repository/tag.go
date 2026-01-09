package repository

import (
	"context"
	"github.com/zyy125/my-blog/backend/internal/model"
	"gorm.io/gorm"
)

// TagRepository 标签数据访问层
type TagRepository struct {
	db *gorm.DB
}

// NewTagRepository 创建标签仓库实例
func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

// Create 创建标签
func (r *TagRepository) Create(ctx context.Context, tag *model.Tag) error {
	return r.db.WithContext(ctx).Create(tag).Error
}

// GetByID 根据ID查询标签
func (r *TagRepository) GetByID(ctx context.Context, id uint) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.WithContext(ctx).First(&tag, id).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// GetByName 根据名称查询标签
func (r *TagRepository) GetByName(ctx context.Context, name string) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// GetByIDs 根据ID列表批量查询标签
func (r *TagRepository) GetByIDs(ctx context.Context, ids []uint) ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&tags).Error
	return tags, err
}

// List 查询所有标签
func (r *TagRepository) List(ctx context.Context) ([]*model.Tag, error) {
	var tags []*model.Tag
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&tags).Error
	return tags, err
}

// Update 更新标签
func (r *TagRepository) Update(ctx context.Context, tag *model. Tag) error {
	return r.db.WithContext(ctx).Save(tag).Error
}

// Delete 删除标签
func (r *TagRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Tag{}, id).Error
}