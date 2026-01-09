package dto

// CreateArticleRequest 创建文章请求
type CreateArticleRequest struct {
	Title      string  `json:"title" binding:"required"`        // 标题（必填）
	Content    string  `json:"content" binding:"required"`      // 内容（必填）
	Summary    string  `json:"summary"`                         // 摘要（可选）
	CoverImg   string  `json:"cover_img"`                       // 封面图（可选）
	CategoryID *uint   `json:"category_id"`                     // 分类ID（可选）
	TagIDs     []uint  `json:"tag_ids"`                         // 标签ID列表
	Status     int8    `json:"status"`                          // 状态：0草稿 1已发布
	IsTop      bool    `json:"is_top"`                          // 是否置顶
}

// UpdateArticleRequest 更新文章请求
type UpdateArticleRequest struct {
	Title      string  `json:"title" binding:"required"`
	Content    string  `json:"content" binding:"required"`
	Summary    string  `json:"summary"`
	CoverImg   string  `json:"cover_img"`
	CategoryID *uint   `json:"category_id"`
	TagIDs     []uint  `json:"tag_ids"`
	Status     int8    `json:"status"`
	IsTop      bool    `json:"is_top"`
}

// ArticleListQuery 文章列表查询参数
type ArticleListQuery struct {
	Page       int   `form:"page"`        // 页码
	PageSize   int   `form:"page_size"`   // 每页数量
	Status     *int8 `form:"status"`      // 状态筛选
	CategoryID *uint `form:"category_id"` // 分类筛选
	TagID      *uint `form:"tag_id"`      // 标签筛选
	Keyword    string `form:"keyword"`     // 关键词搜索
}