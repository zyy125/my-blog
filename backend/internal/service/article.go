package service

import (
	"context"
	"errors"
	"github.com/zyy125/my-blog/backend/internal/model"
	"github.com/zyy125/my-blog/backend/internal/repository"
	"gorm.io/gorm"
)

// ArticleService 文章业务逻辑层
type ArticleService struct {
	repo *repository.ArticleRepository
}

// NewArticleService 创建文章服务实例
func NewArticleService(repo *repository.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

// Create 创建文章
func (s *ArticleService) Create(ctx context.Context, article *model.Article) error {
	// 1. 业务验证
	if article.Title == "" {
		return errors.New("标题不能为空")
	}
	if article.Content == "" {
		return errors.New("内容不能为空")
	}
	
	// 2. 数据处理
	// 如果没有提供摘要，自动从内容截取
	if article.Summary == "" && len(article.Content) > 100 {
		article.Summary = article.Content[:100] + "..."
	}
	
	// 3. 调用 Repository 创建
	return s.repo.Create(ctx, article)
}

// GetByID 获取文章详情（并增加浏览量）
func (s *ArticleService) GetByID(ctx context.Context, id uint) (*model.Article, error) {
	// 1. 查询文章
	article, err := s. repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm. ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		return nil, err
	}
	
	// 2. 增加浏览量（异步，不影响返回）
	go func() {
		_ = s.repo.IncrementViews(context.Background(), id)
	}()
	
	return article, nil
}

// List 获取文章列表
func (s *ArticleService) List(ctx context.Context, page, pageSize int, status *int8) ([]*model.Article, int64, error) {
	// 参数验证
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	
	return s.repo.List(ctx, page, pageSize, status)
}

// Update 更新文章
func (s *ArticleService) Update(ctx context.Context, article *model.Article) error {
	// 1. 检查文章是否存在
	existing, err := s.repo.GetByID(ctx, article.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("文章不存在")
		}
		return err
	}
	
	// 2. 业务验证
	if article. Title == "" {
		return errors.New("标题不能为空")
	}
	
	// 3. 保留某些字段（如创建时间、浏览量）
	article.CreatedAt = existing.CreatedAt
	article.Views = existing. Views
	
	// 4. 执行更新
	return s. repo.Update(ctx, article)
}

// Delete 删除文章
func (s *ArticleService) Delete(ctx context. Context, id uint) error {
	// 1. 检查文章是否存在
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("文章不存在")
		}
		return err
	}
	
	// 2. 执行删除
	return s. repo.Delete(ctx, id)
}