package model

import "time"

// Category 分类模型
type Category struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"size:50;not null;unique" json:"name"`        // 分类名称（唯一）
	Description string    `gorm:"size:200" json:"description"`                // 分类描述
	CreatedAt   time.Time `json:"created_at"`
	
	// 关联：一个分类有多篇文章
	Articles []Article `gorm:"foreignKey:CategoryID" json:"articles,omitempty"`
}

// TableName 指定表名
func (Category) TableName() string {
	return "categories"
}