package model

import "time"

// Tag 标签模型
type Tag struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"size:50;not null;unique" json:"name"` // 标签名称（唯一）
	CreatedAt time.Time `json:"created_at"`
	
	// 关联：一个标签对应多篇文章（多对多）
	Articles []Article `gorm:"many2many:article_tags;" json:"articles,omitempty"`
}

// TableName 指定表名
func (Tag) TableName() string {
	return "tags"
}