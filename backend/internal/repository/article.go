package repository

import (
	"context"
	"github.com/zyy125/my-blog/backend/internal/model"
	"gorm.io/gorm"
)

// ArticleRepository 文章数据访问层
type ArticleRepository struct {
	db *gorm.DB
}

// NewArticleRepository 创建文章仓库实例
func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

// Create 创建文章
func (r *ArticleRepository) Create(ctx context.Context, article *model.Article) error {
	return r.db.WithContext(ctx).Create(article).Error
}

// GetByID 根据 ID 查询文章
func (r *ArticleRepository) GetByID(ctx context.Context, id uint) (*model.Article, error) {
	var article model.Article
	err := r.db.WithContext(ctx).First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// List 查询文章列表（带分页）
func (r *ArticleRepository) List(ctx context.Context, page, pageSize int, status *int8) ([]*model.Article, int64, error) {
	var articles []*model.Article
	var total int64
	
	// 构建查询
	query := r.db.WithContext(ctx).Model(&model.Article{})
	
	// 如果指定了状态，添加条件
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	
	// 查询总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error
	
	if err != nil {
		return nil, 0, err
	}
	
	return articles, total, nil
}

// Update 更新文章
func (r *ArticleRepository) Update(ctx context.Context, article *model.Article) error {
	return r.db.WithContext(ctx).Save(article).Error
}

// Delete 删除文章
func (r *ArticleRepository) Delete(ctx context. Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Article{}, id).Error
}

// IncrementViews 增加浏览量
func (r *ArticleRepository) IncrementViews(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&model.Article{}).
		Where("id = ?", id).
		UpdateColumn("views", gorm. Expr("views + 1")).Error
}