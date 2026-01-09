package model

import (
	"time"
)

// Article 文章模型
type Article struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Title      string    `gorm:"size:200;not null" json:"title"`              // 标题
	Content    string    `gorm:"type:longtext;not null" json:"content"`       // Markdown 内容
	Summary    string    `gorm:"size:500" json:"summary"`                     // 摘要
	CoverImg   string    `gorm:"size:500" json:"cover_img"`                   // 封面图
	CategoryID *uint     `gorm:"index" json:"category_id"`                    // 分类 ID（外键）
	Views      int       `gorm:"default:0" json:"views"`                      // 浏览量
	Status     int8      `gorm:"default:0;index" json:"status"`               // 0=草稿 1=已发布
	IsTop      bool      `gorm:"default:false" json:"is_top"`                 // 是否置顶
	CreatedAt  time.Time `json:"created_at"`                                  // 创建时间
	UpdatedAt  time. Time `json:"updated_at"`                                  // 更新时间
	// 所属分类（多对一）
	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	
	// 关联标签（多对多）
	Tags []Tag `gorm:"many2many:article_tags;" json:"tags,omitempty"`
}

// TableName 指定表名
func (Article) TableName() string {
	return "articles"
}