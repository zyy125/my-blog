package model

import "time"

// Comment 评论模型
type Comment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	ArticleID uint      `gorm:"not null;index" json:"article_id"`
	ParentID  *uint     `gorm:"index" json:"parent_id"`
	Nickname  string    `gorm:"size:50;not null" json:"nickname"`
	Email     string    `gorm:"size:100" json:"email"` // Removed not null constraint
	Content   string    `gorm:"type:text;not null" json:"content"`
	Status    int8      `gorm:"default:0;index" json:"status"`
	IP        string    `gorm:"size:50" json:"ip"`
	CreatedAt time.Time `json:"created_at"`

	// 关联：所属文章
	Article *Article `gorm:"foreignKey:ArticleID" json:"article,omitempty"`

	// ✅ 修改：禁用外键约束
	Parent *Comment `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE;" json:"parent,omitempty"`

	// ✅ 修改：禁用外键约束
	Replies []Comment `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE;" json:"replies,omitempty"`
}

// TableName 指定表名
func (Comment) TableName() string {
	return "comments"
}
