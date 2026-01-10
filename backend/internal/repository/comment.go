package repository

import (
	"context"
	"github.com/zyy125/my-blog/backend/internal/model"
	"gorm.io/gorm"
)

// CommentRepository 评论数据访问层
type CommentRepository struct {
	db *gorm.DB
}

// NewCommentRepository 创建评论仓库实例
func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

// Create 创建评论
func (r *CommentRepository) Create(ctx context.Context, comment *model.Comment) error {
	return r.db.WithContext(ctx).Create(comment).Error
}

// GetByID 根据ID查询评论
func (r *CommentRepository) GetByID(ctx context.Context, id uint) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.WithContext(ctx).First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// ListByArticle 查询文章的评论列表（带回复）
func (r *CommentRepository) ListByArticle(ctx context.Context, articleID uint) ([]*model.Comment, error) {
	var comments []*model.Comment
	
	// ✅ 修改条件：parent_id IS NULL
	err := r.db.WithContext(ctx).
		Where("article_id = ?  AND parent_id IS NULL AND status = 1", articleID).
		Preload("Replies", "status = 1").
		Order("created_at DESC").
		Find(&comments).Error
	
	return comments, err
}

// ListAll 查询所有评论（管理后台用）
func (r *CommentRepository) ListAll(ctx context. Context, page, pageSize int, status *int8) ([]*model.Comment, int64, error) {
	var comments []*model.Comment
	var total int64
	
	query := r.db.WithContext(ctx).Model(&model.Comment{})
	
	// 状态筛选
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	
	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	err := query. 
		Preload("Article", func(db *gorm. DB) *gorm.DB {
			return db.Select("id, title")  // 只加载文章的ID和标题
		}).
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&comments).Error
	
	return comments, total, err
}

// UpdateStatus 更新评论状态
func (r *CommentRepository) UpdateStatus(ctx context.Context, id uint, status int8) error {
	return r.db.WithContext(ctx).
		Model(&model.Comment{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// Delete 删除评论
func (r *CommentRepository) Delete(ctx context.Context, id uint) error {
	// 使用事务：删除评论时同时删除其所有回复
	return r.db.WithContext(ctx).Transaction(func(tx *gorm. DB) error {
		// 1. 删除所有回复
		if err := tx.Where("parent_id = ?", id).Delete(&model.Comment{}).Error; err != nil {
			return err
		}
		
		// 2. 删除评论本身
		if err := tx. Delete(&model.Comment{}, id).Error; err != nil {
			return err
		}
		
		return nil
	})
}

// CountByArticle 统计文章的评论数量
func (r *CommentRepository) CountByArticle(ctx context.Context, articleID uint) (int64, error) {
	var count int64
	err := r. db.WithContext(ctx).
		Model(&model.Comment{}).
		Where("article_id = ?  AND status = 1", articleID).
		Count(&count).Error
	return count, err
}