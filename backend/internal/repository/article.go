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

// GetByIDWithAssociations 根据ID查询文章（包含分类和标签）
func (r *ArticleRepository) GetByIDWithAssociations(ctx context.Context, id uint) (*model.Article, error) {
	var article model.Article
	err := r.db.WithContext(ctx).
		Preload("Category").  // 预加载分类
		Preload("Tags").      // 预加载标签
		First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// ListWithAssociations 查询文章列表（包含分类和标签）
func (r *ArticleRepository) ListWithAssociations(ctx context.Context, page, pageSize int, status *int8) ([]*model.Article, int64, error) {
	var articles []*model.Article
	var total int64
	
	query := r.db.WithContext(ctx).Model(&model.Article{})
	
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	offset := (page - 1) * pageSize
	err := query. 
		Preload("Category").  // ✅ 预加载分类
		Preload("Tags").      // ✅ 预加载标签
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error
	
	if err != nil {
		return nil, 0, err
	}
	
	return articles, total, nil
}

// UpdateWithTags 更新文章并关联标签
func (r *ArticleRepository) UpdateWithTags(ctx context.Context, article *model.Article, tagIDs []uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm. DB) error {
		// 1. 更新文章基本信息
		if err := tx.Save(article).Error; err != nil {
			return err
		}
		
		// 2. 清空现有标签关联
		if err := tx. Model(article).Association("Tags").Clear(); err != nil {
			return err
		}
		
		// 3. 如果有新标签，添加关联
		if len(tagIDs) > 0 {
			var tags []model.Tag
			if err := tx.Where("id IN ? ", tagIDs).Find(&tags).Error; err != nil {
				return err
			}
			if err := tx.Model(article).Association("Tags").Append(tags); err != nil {
				return err
			}
		}
		
		return nil
	})
}

// ListByCategory 根据分类查询文章
func (r *ArticleRepository) ListByCategory(ctx context.Context, categoryID uint, page, pageSize int) ([]*model.Article, int64, error) {
	var articles []*model.Article
	var total int64
	
	query := r.db.WithContext(ctx).
		Model(&model.Article{}).
		Where("category_id = ?  AND status = ? ", categoryID, 1) // 只查已发布的
	
	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	err := query. 
		Preload("Category").
		Preload("Tags").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error
	
	return articles, total, err
}

// ListByTag 根据标签查询文章
func (r *ArticleRepository) ListByTag(ctx context.Context, tagID uint, page, pageSize int) ([]*model.Article, int64, error) {
	var articles []*model.Article
	var total int64
	
	// 多对多关系的查询：需要 JOIN 中间表
	query := r. db.WithContext(ctx).
		Model(&model.Article{}).
		Joins("JOIN article_tags ON articles.id = article_tags.article_id").
		Where("article_tags. tag_id = ? AND articles.status = ?", tagID, 1)
	
	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	err := query. 
		Preload("Category").
		Preload("Tags").
		Order("articles.created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error
	
	return articles, total, err
}

// Search 搜索文章（标题或内容）
func (r *ArticleRepository) Search(ctx context.Context, keyword string, page, pageSize int) ([]*model.Article, int64, error) {
	var articles []*model.Article
	var total int64
	
	// LIKE 模糊查询
	searchPattern := "%" + keyword + "%"
	query := r.db.WithContext(ctx).
		Model(&model.Article{}).
		Where("(title LIKE ? OR content LIKE ?) AND status = ?", searchPattern, searchPattern, 1)
	
	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	err := query.
		Preload("Category").
		Preload("Tags").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&articles).Error
	
	return articles, total, err
}