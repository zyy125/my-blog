package dto

// CreateCommentRequest 创建评论请求
type CreateCommentRequest struct {
	ArticleID uint   `json:"article_id" binding:"required"`
	ParentID  *uint  `json:"parent_id"` // 可选，指针类型
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Content   string `json:"content" binding:"required"`
}
